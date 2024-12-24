package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()

	postIdStr := c.Param("post_id")

	postId, err := strconv.ParseInt(postIdStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("invalid post id").Error(),
		})
		return
	}

	post, err := h.postSvc.GetPostByID(ctx, postId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    post,
	})

}
