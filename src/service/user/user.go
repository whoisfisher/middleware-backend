package user

import (
	"encoding/json"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/util/encrypt"
)

type UserService interface {
	Get(name string) (*entity.User, error)
	GetUserById(id string) (*entity.User, error)
	List() ([]entity.User, error)
	Page(num, size int) (*entity.Page, error)
	Create(usr entity.User) (*entity.User, error)
	Update(usr entity.User) (*entity.User, error)
	Delete(name string) error
	Batch(usr entity.OperateUser) error
	ChangePassword(usr entity.UserChangePassword) error
	UserAuth(name string, password string) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService() UserService {
	return &userService{
		userRepo: repository.NewUserRepository(),
	}
}

func (u *userService) Get(name string) (*entity.User, error) {
	user, err := u.userRepo.Get(name)
	if err != nil {
		logger.Log.Errorf("Faile to get user info by name[%s]: %s", name, err.Error())
		return nil, err
	}

	var userEntity *entity.User
	userEntity.User = user
	return userEntity, nil
}

func (u *userService) GetUserById(id string) (*entity.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		logger.Log.Errorf("Faile to get user info by id[%s]: %s", id, err.Error())
		return nil, err
	}

	var userEntity *entity.User
	userEntity.User = user
	return userEntity, nil
}

func (u *userService) List() ([]entity.User, error) {
	users, err := u.userRepo.List()
	if err != nil {
		logger.Log.Errorf("Faile to get user list info: %s", err.Error())
		return nil, err
	}
	var usersEntity []entity.User
	for _, user := range users {
		var userEntity entity.User
		userEntity.User = user
		usersEntity = append(usersEntity, userEntity)
	}
	return usersEntity, nil
}

func (u *userService) Page(num, size int) (*entity.Page, error) {
	var (
		page        entity.Page
		usersEntity []entity.User
	)
	total, users, err := u.userRepo.Page(num, size)
	if err != nil {
		logger.Log.Errorf("Faile to get user list info: %s", err.Error())
		return nil, err
	}
	for _, user := range users {
		var userEntity entity.User
		userEntity.User = user
		usersEntity = append(usersEntity, userEntity)
	}
	page.Total = total
	page.Items = usersEntity
	return &page, nil
}

func (u *userService) Create(usr entity.User) (*entity.User, error) {
	if usr.Name == usr.Password {
		return nil, constant.ErrNamePwdFailed
	}
	user, err := u.Get(usr.Name)
	if err != nil {
		logger.Log.Errorf("Faile to get user info by name[%s]: %s", usr.Name, err.Error())
		return nil, err
	}
	if user != nil {
		logger.Log.Errorf("Faile to get user info, can't not find user: %s", err.Error())
		return nil, constant.ErrUserNameExist
	}

	password, err := encrypt.StringEncrypt(usr.Password)
	if err != nil {
		return nil, err
	}
	userModel := model.User{
		Name:     usr.Name,
		Password: password,
		IsAdmin:  usr.IsAdmin,
		Role:     usr.Role,
		IsActive: true,
		Type:     usr.Type,
	}
	err = u.userRepo.Save(&userModel)
	if err != nil {
		logger.Log.Errorf("Faile to save user info to database: %s", err.Error())
		return nil, err
	}

	var userEntity *entity.User
	userEntity.User = userModel
	return userEntity, nil
}

func (u *userService) Update(usr entity.User) (*entity.User, error) {
	mapUser := make(map[string]interface{})
	jsonUser, err := json.Marshal(usr)
	if err != nil {
		logger.Log.Errorf("Faile to marshal user info: %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(jsonUser, &mapUser)
	if err != nil {
		logger.Log.Errorf("Faile to unmarshal user info: %s", err.Error())
		return nil, err
	}
	err = u.userRepo.Update(usr.ID, mapUser)
	if err != nil {
		logger.Log.Errorf("Faile to update user info: %s", err.Error())
		return nil, err
	}
	var userEntity *entity.User
	userEntity.User, err = u.userRepo.GetByID(usr.ID)
	if err != nil {
		logger.Log.Errorf("Faile to get user info: %s", err.Error())
		return nil, err
	}
	return userEntity, nil
}

func (u *userService) Delete(name string) error {
	return u.userRepo.Delete(name)
}

func (u *userService) Batch(usr entity.OperateUser) error {
	return u.userRepo.Batch(usr.Operation, usr.Items)
}

func (u *userService) ChangePassword(usr entity.UserChangePassword) error {
	success, err := u.ValidateOldPassword(usr)
	if err != nil {
		logger.Log.Errorf("Faile to validate old password: %s", err.Error())
		return err
	}
	if !success {
		return constant.ErrOriginalNotMatch
	}
	if usr.Name == usr.Password {
		return constant.ErrNamePwdFailed
	}
	user, err := u.userRepo.Get(usr.Name)
	if err != nil {
		logger.Log.Errorf("Faile to get user info by name[%s]: %s", usr.Name, err.Error())
		return err
	}
	user.Password, err = encrypt.StringEncrypt(usr.Password)
	if err != nil {
		logger.Log.Errorf("Faile to encrypt password: %s", err.Error())
		return err
	}
	err = u.userRepo.Save(&user)
	if err != nil {
		logger.Log.Errorf("Faile to save user info: %s", err.Error())
		return err
	}
	return nil
}

func (u *userService) UserAuth(name string, password string) (*entity.User, error) {
	var userEntity entity.User
	user, err := u.userRepo.Get(name)
	if err != nil {
		logger.Log.Errorf("Faile to get user info by name[%s]: %s", name, err.Error())
		return nil, err
	}
	if !user.IsActive {
		return nil, constant.ErrUserIsNotActive
	}
	if user.Type == constant.Local {
		uPassword, err := encrypt.StringDecrypt(user.Password)
		if err != nil {
			logger.Log.Errorf("Faile to decrypt password: %s", err.Error())
			return nil, err
		}
		if uPassword != password {
			return nil, constant.ErrNameOrPasswordErr
		}
	}
	userEntity.User = user
	return &userEntity, nil
}

func (u *userService) ResetPassword(forget entity.UserPasswordForget) error {
	_, err := u.userRepo.GetByID(forget.ID)
	if err != nil {
		logger.Log.Errorf("Faile to get user info: %s", err.Error())
		return err
	}
	return nil
}

func (u *userService) ValidateOldPassword(usr entity.UserChangePassword) (bool, error) {
	user, err := u.userRepo.Get(usr.Name)
	if err != nil {
		logger.Log.Errorf("Faile to get user info by name[%s]: %s", usr.Name, err.Error())
		return false, err
	}
	oldPassword, err := encrypt.StringDecrypt(user.Password)
	if err != nil {
		logger.Log.Errorf("Faile to decrypt password: %s", err.Error())
		return false, err
	}
	if oldPassword != usr.Original {
		logger.Log.Errorf("Faile to validata oldpassword and password: %s", err.Error())
		return false, err
	}
	return true, err
}
