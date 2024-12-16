package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/forumapp/internal/configs"
	"github.com/rdy24/forumapp/internal/handlers/memberships"
	"github.com/rdy24/forumapp/internal/handlers/posts"
	membershipRepo "github.com/rdy24/forumapp/internal/repository/memberships"
	postRepo "github.com/rdy24/forumapp/internal/repository/posts"
	membershipSvc "github.com/rdy24/forumapp/internal/service/memberships"
	postSvc "github.com/rdy24/forumapp/internal/service/posts"
	"github.com/rdy24/forumapp/pkg/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	serviceMembership := membershipSvc.NewService(cfg, membershipRepo)
	servicePost := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, serviceMembership)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, servicePost)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
