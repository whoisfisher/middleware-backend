package encrypt

import (
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"golang.org/x/crypto/bcrypt"
)

func encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func isEncrypted(password string) bool {
	cost, _ := bcrypt.Cost([]byte(password))
	return cost > 0
}

func ComparePassword(hashedPassword, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPassword))
	if err != nil {
		logger.Log.Errorf(err.Error())
		return false
	}
	return true
}

func EncryptPassword(originPassword string) (string, error) {
	if originPassword != "" && !isEncrypted(originPassword) {
		password, err := encrypt(originPassword)
		if err != nil {
			logger.Log.Errorf(err.Error())
			return "", err
		}
		return password, nil
	}
	return originPassword, nil
}
