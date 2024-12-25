package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rdy24/forumapp/internal/model/memberships"
	"github.com/rdy24/forumapp/pkg/jwt"
	tokenUtil "github.com/rdy24/forumapp/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req *memberships.LoginRequest) (string, string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)

	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", "", err
	}

	if user == nil {
		return "", "", errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", "", errors.New("email or password is incorrect")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", "", err
	}

	existingToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())

	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", "", err
	}

	if existingToken != nil {
		return token, existingToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()

	if refreshToken == "" {
		log.Error().Msg("failed to generate refresh token")
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefershToken(ctx, memberships.RefreshTokenModel{
		UserId:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})

	if err != nil {
		log.Error().Err(err).Msg("failed to insert refresh token")
		return token, refreshToken, err
	}

	return token, refreshToken, nil

}
