package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/forumapp/internal/middleware"
	"github.com/rdy24/forumapp/internal/model/posts"
)

type postService interface {
	CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, postId, userId int64, req posts.CreateCommentRequest) error

	UpsertUserActivity(ctx context.Context, postId, userId int64, req posts.UserActivityRequest) error
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:post_id", h.CreateComment)
	route.POST("/activity/:post_id", h.UpsertUserActivity)
}
