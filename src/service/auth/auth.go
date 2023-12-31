package auth

import (
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/service/user"
	"github.com/mensylisir/kmpp-middleware/src/util/captcha"
	"github.com/mensylisir/kmpp-middleware/src/util/jwt"
	"net/http"
)

type AuthService interface {
	Login(aul entity.LoginProfile) (*entity.Profile, error)
	Refresh(r *http.Request) (*entity.Profile, error)
	Logout(r *http.Request) error
	Register(usr entity.User) (*entity.Profile, error)
}

type authService struct {
	UserService user.UserService
}

func NewAuthService() AuthService {
	return &authService{
		UserService: user.NewUserService(),
	}
}

func (as authService) Login(aul entity.LoginProfile) (*entity.Profile, error) {
	if aul.CaptchaId != "" {
		err := captcha.VerifyCode(aul.CaptchaId, aul.Code)
		if err != nil {
			logger.Log.Errorf("Login failed: %s", err.Error())
			return nil, err
		}
	}
	profile := entity.Profile{}
	usr, err := as.UserService.UserAuth(aul.Username, aul.Password)
	if err != nil {
		logger.Log.Errorf("User auth failed: %s", err.Error())
		return nil, err
	}
	profile.User = *usr
	token, err := jwt.CreateToken(*usr)
	if err != nil {
		logger.Log.Errorf("Create token failed: %s", err.Error())
		return nil, err
	}
	profile.AccessToken = token
	refreshToken, err := jwt.CreateRefreshToken(*usr)
	if err != nil {
		logger.Log.Errorf("Create refresh token failed: %s", err.Error())
		return nil, err
	}
	profile.RefreshToken = refreshToken
	return &profile, nil
}

func (as authService) Refresh(r *http.Request) (*entity.Profile, error) {
	usr, err := jwt.ExtractTokenMetadata(r)
	if err != nil {
		logger.Log.Errorf("Extract token failed: %s", err.Error())
		return nil, err
	}
	if usr == nil {
		logger.Log.Errorf("User not found: %s", err.Error())
		return nil, constant.ErrUserNotFound
	}
	profile := entity.Profile{}
	profile.User = *usr
	token, err := jwt.CreateToken(*usr)
	if err != nil {
		logger.Log.Errorf("Create token failed: %s", err.Error())
		return nil, err
	}
	profile.AccessToken = token
	refreshToken, err := jwt.CreateRefreshToken(*usr)
	if err != nil {
		logger.Log.Errorf("Create refresh token failed: %s", err.Error())
		return nil, err
	}
	profile.RefreshToken = refreshToken
	return &profile, nil
}

func (as authService) Logout(r *http.Request) error {
	usr, err := jwt.ExtractTokenMetadata(r)
	if err != nil {
		logger.Log.Errorf("Extract token failed: %s", err.Error())
		return err
	}
	if usr == nil {
		logger.Log.Errorf("User not found: %s", err.Error())
		return constant.ErrUserNotFound
	}
	return nil
}

func (as authService) Register(usr entity.User) (*entity.Profile, error) {
	usrEntity, err := as.UserService.Create(usr)
	if err != nil {
		logger.Log.Errorf("Create User failed: %s", err.Error())
		return nil, err
	}
	profile := entity.Profile{}
	usrEntity1, err := as.UserService.UserAuth(usrEntity.Name, usrEntity.Password)
	if err != nil {
		logger.Log.Errorf("User %s login failed: %s", usrEntity.Name, err.Error())
		return nil, err
	}
	profile.User = *usrEntity1
	token, err := jwt.CreateToken(*usrEntity1)
	if err != nil {
		logger.Log.Errorf("Create token failed: %s", err.Error())
		return nil, err
	}
	profile.AccessToken = token
	refreshToken, err := jwt.CreateRefreshToken(*usrEntity1)
	if err != nil {
		logger.Log.Errorf("Create refresh token failed: %s", err.Error())
		return nil, err
	}
	profile.RefreshToken = refreshToken
	return &profile, nil
}
