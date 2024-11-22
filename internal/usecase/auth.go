package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/bcrypt"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/Ndraaa15/diabetix-server/pkg/gomail"
	"github.com/Ndraaa15/diabetix-server/pkg/jwt"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/allegro/bigcache"
	"github.com/bwmarrin/snowflake"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type IAuthUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (string, error)
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
	Verify(ctx context.Context, req dto.VerificationRequest) error
}

type AuthUsecase struct {
	authStore store.IAuthStore
	zap       *zap.Logger
	snowFlake *snowflake.Node
	bigCache  *bigcache.BigCache
	gomail    *gomail.Gomail
}

func NewAuthUsecase(authStore store.IAuthStore, zap *zap.Logger, snowFlake *snowflake.Node, bigCache *bigcache.BigCache, gomail *gomail.Gomail) IAuthUsecase {
	return &AuthUsecase{
		authStore: authStore,
		zap:       zap,
		snowFlake: snowFlake,
		bigCache:  bigCache,
		gomail:    gomail,
	}
}

func (uc *AuthUsecase) Register(ctx context.Context, req dto.RegisterRequest) (string, error) {
	if req.Password != req.ConfirmPassword {
		return "", errx.New().
			WithCode(iris.StatusUnprocessableEntity).
			WithMessage("password and confirm password must be the same").
			WithError(errors.New("password and confirm password must be the same"))
	}

	birth, err := time.Parse("02-01-2006", req.Birth)
	if err != nil {
		return "", errx.New().WithCode(iris.StatusUnprocessableEntity).WithMessage("invalid birth date format").WithError(err)
	}

	hashedPassword, err := bcrypt.EncryptPassword(req.Password)
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed to encrypt password").
			WithError(err)
	}

	data := domain.User{
		ID:       uc.snowFlake.Generate().String(),
		Name:     req.Name,
		Email:    req.Email,
		Birth:    birth,
		IsActive: false,
		Password: hashedPassword,
		LevelID:  1,
	}

	if err := uc.authStore.CreateUser(ctx, data); err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed to create user").
			WithError(err)
	}

	code, err := util.GenerateCode(5)
	err = uc.bigCache.Set(data.ID, []byte(code))
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed to set cache").
			WithError(err)
	}

	uc.gomail.SetSubject("Verification Code")
	uc.gomail.SetReciever(data.Email)
	uc.gomail.SetSender("fuwafu212@gmail.com")
	err = uc.gomail.SetBodyHTML("verification_code.html", struct{ Code string }{Code: code})
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed to set body").
			WithError(err)
	}

	err = uc.gomail.Send()
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed to send email").
			WithError(err)

	}

	return data.ID, nil
}

func (uc *AuthUsecase) Verify(ctx context.Context, req dto.VerificationRequest) error {
	code, err := uc.bigCache.Get(req.ID)
	if err != nil {
		return errx.New().
			WithCode(iris.StatusNotFound).
			WithMessage("verification code not found").
			WithError(err)
	}

	if req.Code != string(code) {
		return errx.New().
			WithCode(iris.StatusUnprocessableEntity).
			WithMessage("invalid verification code").
			WithError(errors.New("invalid verification code"))
	}

	user, err := uc.authStore.GetUserByID(ctx, req.ID)
	if err != nil {
		return errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed get user by id").
			WithError(err)
	}

	if user.IsActive {
		return errx.New().
			WithCode(iris.StatusUnprocessableEntity).
			WithMessage("user already verified").
			WithError(errors.New("user already verified"))
	}

	user.IsActive = true

	if err := uc.authStore.UpdateUser(ctx, user); err != nil {
		return errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed to update user").
			WithError(err)
	}

	return nil
}

func (uc *AuthUsecase) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := uc.authStore.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed get user with email").
			WithError(err)
	}

	if !user.IsActive {
		return "", errx.New().
			WithCode(iris.StatusUnprocessableEntity).
			WithMessage("user not verified").
			WithError(errors.New("user not verified"))
	}

	if err := bcrypt.ComparePassword(user.Password, req.Password); err != nil {
		return "", errx.New().
			WithCode(iris.StatusUnprocessableEntity).
			WithMessage("invalid password").
			WithError(err)
	}

	token, err := jwt.EncodeToken(&user, 24*time.Hour)
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("failed to encode token").
			WithError(err)
	}

	return token, nil
}
