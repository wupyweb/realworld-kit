package conduit

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wupyweb/realworld-kit/internal/utils"
)

// POST /api/users/login
// 用户登录
func (m *ConduitService) Authentication(c *gin.Context) {
	var req LoginRequest

	// gin根据Content-Type进行数据绑定
	//
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 对username和password校验
	// ...
	user, err := m.store.GetUserByEmail(req.User.Email)
	// 用户不存在
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 密码不正确
	if !CheckPassword(user.Password, req.User.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pass"})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := LoginResponse{}
	res.User.Email = user.Email
	res.User.Token = token
	res.User.Username = user.Username
	res.User.Bio = user.Bio
	res.User.Image = user.Image

	c.JSON(http.StatusOK, res)
}

// POST /api/users
// 用户注册
func (m *ConduitService) Registration(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证用户名和密码
	// ...
	// 创建用户
	user, err := m.store.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := LoginResponse{}
	res.User.Email = user.Email
	res.User.Token = token
	res.User.Username = user.Username
	res.User.Bio = user.Bio
	res.User.Image = user.Image

	c.JSON(http.StatusOK, res)
}

// GET /api/user
// 获取当前登录用户
func (m *ConduitService) GetCurrentUser(c *gin.Context) {

	// 从上下文拿到鉴权用户id
	userid, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	u, err := m.store.GetUserByID(userid.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := map[string]UserBase{
		"user": {
			Email:    u.Email,
			Username: u.Username,
			Bio:      u.Bio,
			Image:    u.Image,
		},
	}

	c.JSON(http.StatusOK, res)
}

// PUT /api/user
// 更新用户
func (m *ConduitService) UpdateCurrentUser(c *gin.Context) {
	var req UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID, _ := c.Get("user_id")
	// 更新用户信息
	user, err := m.store.UpdateUser(userID.(int), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := UserBase{
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
	}
	c.JSON(http.StatusOK, res)
}

// GET /api/profiles/:username
// 获取用户profile
func (m *ConduitService) GetProfile(c *gin.Context) {
	// 从上下文中取出鉴权后的当前用户id
	userid, _ := c.Get("user_id")
	// 要获取的用户profile
	username := c.Param("username")
	user, err := m.store.GetProfile(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := toProfile(user, userid.(int))
	c.JSON(http.StatusOK, res)
}

// POST /api/profiles/:username/follow
// 关注用户
func (m *ConduitService) FollowUser(c *gin.Context) {
	followerID, _ := c.Get("user_id")
	username := c.Param("username")

	// TODO 先判断是否已关注
	user, err := m.store.FollowUser(username, followerID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := toProfile(user, followerID.(int))
	c.JSON(http.StatusOK, res)
}

// DELETE /api/profiles/:username/follow
// 取消关注
func (m *ConduitService) UnFollowUser(c *gin.Context) {
	followerID, _ := c.Get("user_id")

	username := c.Param("username")
	user, err := m.store.UnFollowUser(username, followerID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := toProfile(user, followerID.(int))
	c.JSON(http.StatusOK, res)
}

// GET /api/articles
// 获取文章
// Query Parameters:
// Filter by tag:
// ?tag=AngularJS
// Filter by author:
// ?author=jake
// Favorited by user:
// ?favorited=jake
// Limit number of articles (default is 20):
// ?limit=20
// Offset/skip number of articles (default is 0):
// ?offset=0
func (m *ConduitService) ListArticles(c *gin.Context) {
	var query ArticleQuery
	current_user_id, _ := c.Get("user_id")

	if err := c.BindQuery(&query); err != nil {
		log.Print(1, query, err)
		return
	}

	// 赋予limit和offset默认值
	if query.Limit == 0 {
		query.Limit = 20
	}
	if query.Offset == 0 {
		query.Offset = 0
	}

	articles, err := m.store.GetArticles(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var articleList []ArticleBase
	for _, article := range articles {
		articleList = append(articleList, toArticle(article, current_user_id.(int)))
	}

	res := MultipleArticles{
		Articles:      articleList,
		ArticlesCount: len(articleList),
	}
	c.JSON(http.StatusOK, res)
}

// GET /api/articles/feed
// 获取关注用户的文章
func (m *ConduitService) FeedArticles(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	var query ArticleQuery

	if err := c.BindQuery(&query); err != nil {
		log.Print(1, query, err)
		return
	}

	// 赋予limit和offset默认值
	if query.Limit == 0 {
		query.Limit = 20
	}
	if query.Offset == 0 {
		query.Offset = 0
	}
	articles, err := m.store.GetFeed(current_user_id.(int), query.Limit, query.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var articleList []ArticleBase
	for _, article := range articles {
		articleList = append(articleList, toArticle(article, current_user_id.(int)))
	}

	res := MultipleArticles{
		Articles:      articleList,
		ArticlesCount: len(articleList),
	}
	c.JSON(http.StatusOK, res)
}

// GET /api/articles/:slug
// 获取文章详情，无需鉴权
func (m *ConduitService) GetArticle(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")

	slug := c.Param("slug")
	article, err := m.store.GetArticleBySlug(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := SingleArticle{
		Article: toArticle(article, current_user_id.(int)),
	}
	c.JSON(http.StatusOK, res)
}

// POST /api/articles
// 创建文章，需要鉴权
func (m *ConduitService) CreateArticle(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	var req *ArticleRequest

	// 绑定请求参数
	if err := c.BindQuery(&req); err != nil {
		log.Print(1, req, err)
		return
	}

	// 创建文章
	article, err := m.store.CreateArticle(current_user_id.(int), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	a, _ := m.store.GetArticleBySlug(article.Slug)
	res := SingleArticle{
		Article: toArticle(a, current_user_id.(int)),
	}

	c.JSON(http.StatusOK, res)
}

// PUT /api/articles/:slug
// 修改文章，需要鉴权
func (m *ConduitService) UpdateArticle(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	slug := c.Param("slug")
	var req *ArticleRequest

	// 绑定请求参数
	if err := c.BindJSON(&req); err != nil {
		log.Print(1, req, err)
		return
	}

	// 修改文章
	article, err := m.store.UpdateArticle(current_user_id.(int), slug, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	a, _ := m.store.GetArticleBySlug(article.Slug)
	res := SingleArticle{
		Article: toArticle(a, current_user_id.(int)),
	}

	c.JSON(http.StatusOK, res)
}

// DELETE /api/articles/:slug
// 删除文章，需要鉴权
func (m *ConduitService) DeleteArticle(c *gin.Context) {
	slug := c.Param("slug")

	err := m.store.DeleteArticle(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// POST /api/articles/:slug/comments
// 给文章添加评论，需要鉴权
func (m *ConduitService) AddComment(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	slug := c.Param("slug")
	var req *CommentRequest

	// 绑定请求参数
	if err := c.BindJSON(&req); err != nil {
		log.Print(1, req, err)
		return
	}
	// 添加评论
	comment, err := m.store.AddComment(current_user_id.(int), slug, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := SingleComment{
		Comment: toComment(comment, current_user_id.(int)),
	}
	c.JSON(http.StatusOK, res)
}

// GET /api/articles/:slug/comments
// 获取文章所有评论，鉴权可选
func (m *ConduitService) GetComments(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	slug := c.Param("slug")
	comments, err := m.store.GetComments(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var commentList []CommentBase
	for _, comment := range comments {
		commentList = append(commentList, toComment(comment, current_user_id.(int)))
	}
	res := MultipleComments{
		Comments: commentList,
	}
	c.JSON(http.StatusOK, res)
}

// DELETE /api/articles/:slug/comments/:id
// 删除评论，需要鉴权
func (m *ConduitService) DeleteComment(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	slug := c.Param("slug")
	commentID := c.Param("id")

	comment_id, err := strconv.Atoi(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = m.store.DeleteComment(current_user_id.(int), slug, comment_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// POST /api/articles/:slug/favorite
// 喜欢文章，需要鉴权
func (m *ConduitService) FavoriteArticle(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	slug := c.Param("slug")

	// TODO 先判断文章是否被喜欢
	article, err := m.store.FavoriteArticle(current_user_id.(int), slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := SingleArticle{
		Article: toArticle(article, current_user_id.(int)),
	}

	c.JSON(http.StatusOK, res)
}

// DELETE /api/articles/:slug/favorite
// unfavorite
func (m *ConduitService) UnfavoriteArticle(c *gin.Context) {
	current_user_id, _ := c.Get("user_id")
	slug := c.Param("slug")

	article, err := m.store.UnfavoriteArticle(current_user_id.(int), slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := SingleArticle{
		Article: toArticle(article, current_user_id.(int)),
	}

	c.JSON(http.StatusOK, res)
}

// GET /api/tags
// 无需鉴权
func (m *ConduitService) GetTags(c *gin.Context) {
	tags, _ := m.store.GetTags()

	var tagList []string
	for _, tag := range tags {
		tagList = append(tagList, tag.Name)
	}

	c.JSON(http.StatusOK, gin.H{"tags": tagList})
}
