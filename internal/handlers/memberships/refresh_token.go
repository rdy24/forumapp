package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/forumapp/internal/model/memberships"
)

func (h *Handler) Refresh(c *gin.Context) {

	ctx := c.Request.Context()

	var req memberships.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetInt64("userId")

	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userId, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, memberships.RefreshTokenResponse{
		AccessToken: accessToken,
	})
}
