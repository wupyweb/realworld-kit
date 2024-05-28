package conduit

import (
	"context"

	"github.com/wupyweb/realworld-kit/ent"
	"github.com/wupyweb/realworld-kit/ent/article"
	"github.com/wupyweb/realworld-kit/ent/comment"
	"github.com/wupyweb/realworld-kit/ent/tag"
	"github.com/wupyweb/realworld-kit/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type Storage struct {
	db  *ent.Client
	ctx context.Context
}

func NewStorage(db *ent.Client) *Storage {
	ctx := context.Background()
	return &Storage{
		db:  db,
		ctx: ctx,
	}
}

func HashPassword(password string) (string, error) {
	str, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(str), err
}

func CheckPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func (s *Storage) GetUserByID(id int) (*ent.User, error) {
	return s.db.User.Get(s.ctx, id)
}

func (s *Storage) GetUserByEmail(email string) (*ent.User, error) {
	return s.db.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(s.ctx)
}

func (s *Storage) GetUserByName(username string) (*ent.User, error) {
	return s.db.User.
		Query().
		Where(user.UsernameEQ(username)).
		Only(s.ctx)
}

func (s *Storage) CreateUser(req RegisterRequest) (*ent.User, error) {
	hashedPassword, _ := HashPassword(req.User.Password)

	return s.db.User.Create().
		SetUsername(req.User.Username).
		SetEmail(req.User.Email).
		SetPassword(hashedPassword).
		Save(s.ctx)
}

func (s *Storage) UpdateUser(user_id int, req UserUpdate) (*ent.User, error) {
	t := s.db.User.UpdateOneID(user_id)

	if req.User.Email != "" {
		t.SetEmail(req.User.Email)
	}
	if req.User.Username != "" {
		t.SetUsername(req.User.Username)
	}
	if req.User.Password != "" {
		hashedPassword, _ := HashPassword(req.User.Password)
		t.SetPassword(hashedPassword)
	}
	if req.User.Bio != "" {
		t.SetBio(req.User.Bio)
	}
	if req.User.Image != "" {
		t.SetImage(req.User.Image)
	}

	return t.Save(s.ctx)
}

func (s *Storage) GetProfile(username string) (*ent.User, error) {

	return s.db.User.Query().
		WithFollowers().
		Where(
			user.UsernameEQ(username),
		).
		Only(s.ctx)
}

func (s *Storage) IsFollowing(follower_id, followee_id int) (bool, error) {
	return s.db.User.Query().
		Where(user.ID(follower_id)).
		QueryFollowing().
		Where(user.ID(followee_id)).
		Exist(s.ctx)
}

func (s *Storage) Follow(follower_id, followee_id int) (*ent.User, error) {
	return s.db.User.UpdateOneID(follower_id).
		AddFollowingIDs(followee_id).
		Save(s.ctx)
}

func (s *Storage) FollowUser(username string, follower_id int) (*ent.User, error) {
	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return nil, err
	}
	// 查询用户
	u, err := tx.User.Query().
		Where(user.UsernameEQ(username)).
		Only(s.ctx)
	if err != nil {
		return nil, err
	}
	// 添加关注者
	err = tx.User.UpdateOne(u).
		AddFollowerIDs(follower_id).
		Exec(s.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	uu, err := tx.User.Query().
		WithFollowers().
		Where(user.IDEQ(u.ID)).
		Only(s.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 提交事务
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return uu, nil
}

func (s *Storage) Unfollow(follower_id, followee_id int) (*ent.User, error) {
	return s.db.User.UpdateOneID(follower_id).
		RemoveFollowingIDs(followee_id).
		Save(s.ctx)
}

func (s *Storage) UnFollowUser(username string, follower_id int) (*ent.User, error) {
	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return nil, err
	}
	// 查询用户
	u, err := tx.User.Query().
		Where(user.UsernameEQ(username)).
		Only(s.ctx)
	if err != nil {
		return nil, err
	}
	// 移除关注者
	err = tx.User.UpdateOne(u).
		RemoveFollowerIDs(follower_id).
		Exec(s.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	uu, err := tx.User.Query().
		WithFollowers().
		Where(user.IDEQ(u.ID)).
		Only(s.ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 提交事务
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return uu, nil
}

func (s *Storage) GetArticles(query *ArticleQuery) ([]*ent.Article, error) {
	articles := s.db.Article.Query().
		WithOwner().
		WithTags().
		WithLikedUsers()

	if query.Tag != "" {
		articles = articles.
			Where(article.HasTagsWith(tag.Name(query.Tag)))
	}
	if query.Author != "" {
		articles = articles.
			Where(article.HasOwnerWith(user.Username(query.Author)))
	}
	if query.Favorited != "" {
		articles = articles.
			Where(article.HasLikedUsersWith(user.Username(query.Favorited)))
	}

	return articles.
		Order(ent.Desc("created_at")).
		Limit(query.Limit).
		Offset(query.Offset).
		All(s.ctx)
}

func (s *Storage) GetFeed(user_id int, limit, offset int) ([]*ent.Article, error) {

	return s.db.Article.Query().
		WithOwner().
		WithTags().
		WithLikedUsers().
		Where(
			article.HasOwnerWith(
				user.HasFollowersWith(
					user.IDEQ(user_id),
				),
			),
		).
		Order(ent.Desc("created_at")).
		Limit(limit).
		Offset(offset).
		All(s.ctx)
}

func (s *Storage) GetArticleBySlug(slug string) (*ent.Article, error) {

	return s.db.Article.Query().
		WithOwner().
		WithTags().
		WithLikedUsers().
		Where(article.Slug(slug)).
		First(s.ctx)
}

func (s *Storage) CreateArticle(user_id int, req *ArticleRequest) (*ent.Article, error) {
	t := s.db.Article.Create()

	if req.Article.Title != "" {
		t.SetTitle(req.Article.Title)
		t.SetSlug(req.Article.Title)
	}
	if req.Article.Description != "" {
		t.SetDescription(req.Article.Description)
	}
	if req.Article.Body != "" {
		t.SetBody(req.Article.Body)
	}
	if len(req.Article.TagList) > 0 {
		for _, tag := range req.Article.TagList {
			tmp, _ := s.GetOrCreateTag(tag)
			t.AddTags(tmp)
		}
	}

	return t.SetOwnerID(user_id).
		Save(s.ctx)
}

func (s *Storage) UpdateArticle(user_id int, slug string, req *ArticleRequest) (*ent.Article, error) {
	article, _ := s.GetArticleBySlug(slug)
	t := article.Update()

	if req.Article.Title != "" {
		t.SetTitle(req.Article.Title)
		t.SetSlug(req.Article.Title)
	}
	if req.Article.Description != "" {
		t.SetDescription(req.Article.Description)
	}
	if req.Article.Body != "" {
		t.SetBody(req.Article.Body)
	}
	if len(req.Article.TagList) > 0 {
		for _, tag := range req.Article.TagList {
			tmp, _ := s.GetOrCreateTag(tag)
			t.AddTags(tmp)
		}
	}

	return t.Save(s.ctx)
}

func (s *Storage) DeleteArticle(slug string) error {

	// ent: constraint failed: FOREIGN KEY constraint failed
	// 需要先接触关联关系，才能删除实例

	// 开启事务
	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return err
	}

	// 查询文章
	a, err := tx.Article.Query().
		Where(article.SlugEQ(slug)).
		Only(s.ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 接触文章的所有关联关系
	err = a.Update().
		ClearOwner().
		ClearLikedUsers().
		ClearComments().
		ClearTags().
		Exec(s.ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 删除文章
	if err := tx.Article.DeleteOne(a).Exec(s.ctx); err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddComment(user_id int, slug string, req *CommentRequest) (*ent.Comment, error) {
	article, _ := s.GetArticleBySlug(slug)

	c, _ := s.db.Comment.Create().
		SetOwnerID(user_id).
		SetArticleID(article.ID).
		SetBody(req.Comment.Body).
		Save(s.ctx)
	return s.db.Comment.Query().WithOwner().Where(comment.IDEQ(c.ID)).Only(s.ctx)
}

func (s *Storage) GetComments(slug string) ([]*ent.Comment, error) {

	return s.db.Comment.Query().
		Where(comment.HasArticleWith(article.Slug(slug))).
		WithOwner().
		All(s.ctx)
}

func (s *Storage) DeleteComment(id int, slug string, comment_id int) error {

	return s.db.Comment.
		DeleteOneID(comment_id).
		Where(
			comment.HasOwnerWith(user.ID(id)),
			comment.HasArticleWith(article.Slug(slug)),
		).
		Exec(s.ctx)
}

func (s *Storage) IsFavorited(user_id int, slug string) (bool, error) {

	return s.db.Article.Query().
		Where(article.SlugEQ(slug), article.HasLikedUsersWith(user.ID(user_id))).
		Exist(s.ctx)
}

func (s *Storage) FavoriteArticle(user_id int, slug string) (*ent.Article, error) {
	article_id, err := s.db.Article.
		Update().
		AddLikedUserIDs(user_id).
		Where(
			article.SlugEQ(slug),
		).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}

	return s.db.Article.Query().
		WithOwner().
		WithTags().
		WithLikedUsers().
		Where(article.ID(article_id)).
		First(s.ctx)
}

func (s *Storage) UnfavoriteArticle(user_id int, slug string) (*ent.Article, error) {
	article_id, err := s.db.Article.Update().
		RemoveLikedUserIDs(user_id).
		Where(
			article.Slug(slug),
		).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}

	return s.db.Article.Query().
		WithOwner().
		WithTags().
		WithLikedUsers().
		Where(article.ID(article_id)).
		First(s.ctx)
}

func (s *Storage) GetTags() ([]*ent.Tag, error) {
	return s.db.Tag.Query().All(s.ctx)
}

func (s *Storage) GetOrCreateTag(name string) (*ent.Tag, error) {
	// 开启事务
	tx, _ := s.db.Tx(s.ctx)
	// 事务回滚
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	tag, err := tx.Tag.Query().
		Where(tag.Name(name)).
		Only(s.ctx)
	if err != nil {
		// tag不存在
		tag, err = tx.Tag.Create().
			SetName(name).
			Save(s.ctx)
		if err != nil {
			return nil, err
		}
	}
	// 提交事务
	tx.Commit()

	return tag, nil
}
