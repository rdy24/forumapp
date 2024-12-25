package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/forumapp/internal/model/posts"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()
	var req posts.CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIdStr := c.Param("post_id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid post_id"),
		})

		return
	}

	userId := c.GetInt64("userId")

	err = h.postSvc.CreateComment(ctx, postId, userId, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "comment created"})
}
