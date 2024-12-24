package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/rdy24/forumapp/internal/model/memberships"
	"github.com/rdy24/forumapp/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userId int64, request memberships.RefreshTokenRequest) (string, error) {
	existingToken, err := s.membershipRepo.GetRefreshToken(ctx, userId, time.Now())

	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", err
	}

	if existingToken == nil {
		return "", errors.New("refresh token not found or expired")
	}

	if existingToken.RefreshToken != request.Token {
		return "", errors.New("invalid refresh token")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userId)

	token, err := jwt.CreateToken(userId, existingToken.UserId, s.cfg.Service.SecretJWT)
}
