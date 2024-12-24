package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()
	pageSize, pageIndex := c.Query("page_size"), c.Query("page_index")

	if pageSize == "" || pageIndex == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page_size and page_index query params are required"})
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page_size query param must be an integer"})
		return
	}

	pageIndexInt, err := strconv.Atoi(pageIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page_index query param must be an integer"})
		return
	}

	response, err := h.postSvc.GetAllPost(ctx, pageSizeInt, pageIndexInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
