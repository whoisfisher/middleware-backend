package repository

import (
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
)

type HelmRepository interface {
	Page(num, size int) (int, []model.HelmRepository, error)
	List() ([]model.HelmRepository, error)
	Get(name string) (model.HelmRepository, error)
	GetByID(id string) (model.HelmRepository, error)
	Save(item *model.HelmRepository) error
	Delete(name string) error
	Batch(operation string, items []model.HelmRepository) error
	Update(ID string, values map[string]interface{}) error
}

type helmRepository struct {
}

func NewHelmRepository() HelmRepository {
	return &helmRepository{}
}

func (u helmRepository) Page(num, size int) (int, []model.HelmRepository, error) {
	var total int
	var helmRepositories []model.HelmRepository
	err := db.DB.Model(&model.HelmRepository{}).Count(&total).Order("name").Offset((num - 1) * size).Limit(size).Find(&helmRepositories).Error
	if err != nil {
		logger.Log.Errorf("Failed to get Helm repository list info: %s", err.Error())
		return -1, nil, err
	}
	return total, helmRepositories, nil
}

func (u helmRepository) List() ([]model.HelmRepository, error) {
	var helmRepositories []model.HelmRepository
	err := db.DB.Order("name").Find(&helmRepositories).Error
	if err != nil {
		logger.Log.Errorf("Failed to get Helm repository list info: %s", err.Error())
		return nil, err
	}
	return helmRepositories, nil
}

func (u helmRepository) Get(name string) (model.HelmRepository, error) {
	var helmRepo model.HelmRepository
	helmRepo.Name = name
	if err := db.DB.Where("name = ?", name).First(&helmRepo).Error; err != nil {
		logger.Log.Errorf("Failed to get Helm repository info[%s]: %s", name, err.Error())
		return helmRepo, err
	}
	return helmRepo, nil
}

func (u helmRepository) GetByID(id string) (model.HelmRepository, error) {
	var helmRepo model.HelmRepository
	if err := db.DB.Where("id = ?", id).First(&helmRepo).Error; err != nil {
		logger.Log.Errorf("Failed to get Helm repository info[%s]: %s", id, err.Error())
		return helmRepo, err
	}
	return helmRepo, nil
}

func (u helmRepository) Save(item *model.HelmRepository) error {
	if db.DB.NewRecord(item) {
		return db.DB.Create(&item).Error
	} else {
		return db.DB.Save(&item).Error
	}
}

func (u helmRepository) Delete(name string) error {
	helmRepo, err := u.Get(name)
	if err != nil {
		logger.Log.Errorf("Failed to get Helm repository info: %s", err.Error())
		return err
	}
	return db.DB.Delete(&helmRepo).Error
}

func (u helmRepository) Batch(operation string, items []model.HelmRepository) error {
	switch operation {
	case constant.BatchOperationDelete:
		tx := db.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		for _, item := range items {
			err := tx.Delete(&item).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	default:
		return constant.NotSupportedBatchOperation
	}
	return nil
}

func (u helmRepository) Update(ID string, values map[string]interface{}) error {
	err := db.DB.Model(model.HelmRepository{}).Where("id = ?", ID).Updates(values).Error
	if err != nil {
		logger.Log.Errorf("Failed to update Helm repository info: %s", err.Error())
		return err
	}
	return nil
}
