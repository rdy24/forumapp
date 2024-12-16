package memberships

import (
	"context"
	"errors"

	"github.com/rdy24/forumapp/internal/model/memberships"
	"github.com/rdy24/forumapp/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req *memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")

	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return "", errors.New("email or password is incorrect")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}

	return token, nil

}
