package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/forumapp/internal/model/posts"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var req posts.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetInt64("userId")

	err := h.postSvc.CreatePost(ctx, userId, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post created"})
}
