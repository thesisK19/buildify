package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/api"
	"github.com/thesisK19/buildify/app/user/internal/model"
	context_lib "github.com/thesisK19/buildify/library/context"
	errors_lib "github.com/thesisK19/buildify/library/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(ctx context.Context, in *api.SignUpRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "SignUp")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.WithError(err).Error("Failed to hashPassword")
		return nil, err
	}

	createParams := model.CreateUserParams{
		Name:     in.Name,
		Username: in.Username,
		Password: string(hashedPassword),
	}
	_, err = s.repository.CreateUser(ctx, createParams)
	if err != nil {
		logger.WithError(err).Error("Failed to repo.CreateUser")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}

func (s *Service) SignIn(ctx context.Context, in *api.SignInRequest) (*api.SignInResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "SignIn")

	user, err := s.repository.GetUserByUsername(ctx, in.Username)
	if err != nil {
		logger.WithError(err).Error("Failed to GetUserByUsername")
		return nil, err
	}
	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		logger.WithError(err).Error("Incorrect Password")
		return nil, errors_lib.ToUnauthenticatedError(err) // TODO: not return err directly
	}

	// If the password matches, generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.Claims{
		Username: in.Username,
		StandardClaims: jwt.StandardClaims{
			// Set additional standard claims as needed
			ExpiresAt: time.Now().Add(time.Hour * 100000).Unix(), // Token expiration time
		},
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		logger.WithError(err).Error("Failed to SignedString token")
		return nil, err
	}

	return &api.SignInResponse{
		Token: tokenString,
	}, nil
}

func (s *Service) GetUser(ctx context.Context, in *api.EmptyRequest) (*api.GetUserResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetUser")

	username := ctx.Value(context_lib.USERNAME).(string)
	user, err := s.repository.GetUserByUsername(ctx, username)

	if err != nil {
		logger.WithError(err).Error("Failed to repository.GetUserByUsername")
		return nil, err
	}

	return &api.GetUserResponse{
		User: &api.User{
			Username: user.Username,
			Name:     user.Name,
		},
	}, nil
}

func (s *Service) UpdateUser(ctx context.Context, in *api.UpdateUserRequest) (*api.EmptyResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "SignUp")

	username := ctx.Value(context_lib.USERNAME).(string)

	var (
		err            error
		hashedPassword []byte
	)

	// Hash the password
	if in.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			logger.WithError(err).Error("Failed to hashPassword")
			return nil, err
		}
	}

	updateParams := model.UpdateUserParams{
		Name:     in.Name,
		Password: string(hashedPassword),
	}

	err = s.repository.UpdateUserByUsername(ctx, username, updateParams)

	if err != nil {
		logger.WithError(err).Error("Failed to repository.UpdateUserByUsername")
		return nil, err
	}

	return &api.EmptyResponse{}, nil
}
