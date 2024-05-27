package conduit

import (
	"github.com/gin-gonic/gin"
	"github.com/wupyweb/realworld-kit/internal/middleware"
)

type ConduitService struct {
	store  *Storage
	router *gin.RouterGroup
}

func NewService(store *Storage, router *gin.RouterGroup) ConduitService {

	return ConduitService{
		router: router,
		store:  store,
	}
}

// 把路由注册到gin.engine
func (s ConduitService) RegisterRoutes() {

	// 不需要鉴权的路由
	auth := s.router
	{
		auth.POST("/users/login", s.Authentication)
		auth.POST("/users", s.Registration)
	}

	// 需要鉴权的路由
	v1 := s.router
	v1.Use(middleware.AuthMiddleware(s.store.db))
	{
		// ...
		v1.GET("/user", s.GetCurrentUser)
		v1.PUT("/user", s.UpdateCurrentUser)
		v1.GET("/profiles/:username", s.GetProfile)
		v1.POST("/profiles/:username/follow", s.FollowUser)
		v1.DELETE("/profiles/:username/follow", s.UnFollowUser)
		v1.GET("/articles", s.ListArticles)
		v1.GET("/articles/feed", s.FeedArticles)
		v1.GET("/articles/:slug", s.GetArticle)
		v1.POST("/articles", s.CreateArticle)
		v1.PUT("/articles/:slug", s.UpdateArticle)
		v1.DELETE("/articles/:slug", s.DeleteArticle)
		v1.POST("/articles/:slug/comments", s.AddComment)
		v1.GET("/articles/:slug/comments", s.GetComments)
		v1.DELETE("/articles/:slug/comments/:id", s.DeleteComment)
		v1.POST("/articles/:slug/favorite", s.FavoriteArticle)
		v1.DELETE("/articles/:slug/favorite", s.UnfavoriteArticle)
		v1.GET("/tags", s.GetTags)
	}
}
