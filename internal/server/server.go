package server

// import "github.com/gin-gonic/gin"
import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wupyweb/realworld-kit/ent"
	"github.com/wupyweb/realworld-kit/internal/apis/conduit"
	"github.com/wupyweb/realworld-kit/internal/config"
	"github.com/wupyweb/realworld-kit/internal/database"
)

type APIServer struct {
	Config *config.Config
	Router *gin.Engine
	DB     *ent.Client
	// ...
}

func NewServer(cfg *config.Config) (*APIServer, error) {
	// 初始化gin, 注册中间件
	e := gin.Default()
	e.Use(
	// ...
	)

	// 连接数据库，实例化services
	db, err := database.NewSqlite(cfg)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	conduitStore := conduit.NewStorage(db)
	conduitService := conduit.NewService(conduitStore, e.Group("/api"))
	conduitService.RegisterRoutes()

	return &APIServer{
		Config: cfg,
		Router: e,
		DB:     db,
		// ...
    }, nil
}

func (s *APIServer) Serve() {
	// 启动gin
	s.Router.Run(":8080")
}

func (s *APIServer) Shutdown() {

}
