package postgresql

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/jinzhu/gorm"
	"github.com/mensylisir/kmpp-middleware/src/constant"
	"github.com/mensylisir/kmpp-middleware/src/db"
	"github.com/mensylisir/kmpp-middleware/src/entity"
	"github.com/mensylisir/kmpp-middleware/src/logger"
	"github.com/mensylisir/kmpp-middleware/src/model"
	"github.com/mensylisir/kmpp-middleware/src/repository"
	"github.com/mensylisir/kmpp-middleware/src/service/cluster"
	"github.com/mensylisir/kmpp-middleware/src/service/pod"
	"github.com/mensylisir/kmpp-middleware/src/service/secret"
	"github.com/mensylisir/kmpp-middleware/src/service/templates"
	"github.com/mensylisir/kmpp-middleware/src/util/kubernetes"
	"github.com/sirupsen/logrus"
	v1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	"regexp"
)

type PostgresService interface {
	Get(name string) (entity.Instance, error)
	GetById(ID string) (entity.Instance, error)
	GetStatusById(ID string) (string, error)
	GetPostgresOperatorPodName(clusterId string) (string, error)
	Page(num, size int, userId string) (*entity.InstancePage, error)
	Save(instance entity.Instance) error
	Create(postgres entity.Postgres) (string, error)
	Sync(name string) (entity.Instance, error)
	Delete(name string) error
	Update(instance entity.Instance) error
	Edit(instance entity.Instance) error
	GetPostgresLog(instance entity.Instance, log chan string)
	GetUsernameAndPassword(instance entity.Instance) ([]entity.SecretInfo, error)
}

type postgresService struct {
	instanceRepo     repository.InstanceRepository
	userInstanceRepo repository.UserInstanceRepository
	userRepo         repository.UserRepository
	clusterService   cluster.ClusterService
	templateService  templates.TemplatesService
	podService       pod.PodService
	secretService    secret.SecretService
}

func NewPostgresService() PostgresService {
	return &postgresService{
		instanceRepo:     repository.NewInstanceRepository(),
		userInstanceRepo: repository.NewUserInstanceRepository(),
		userRepo:         repository.NewUserRepository(),
		clusterService:   cluster.NewClusterService(),
		templateService:  templates.NewTemplatesService(),
		podService:       pod.NewPodService(),
		secretService:    secret.NewSecretService(),
	}
}

func (c postgresService) Get(name string) (entity.Instance, error) {
	var instanceDTO entity.Instance
	mo, err := c.instanceRepo.Get(name)
	if err != nil {
		logger.Log.Errorf("Faile to get instance info by name[%s]: %s", name, err.Error())
		return instanceDTO, err
	}
	instanceDTO.Instance = mo

	clusterModel, err := c.clusterService.GetByID(mo.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", mo.ClusterID, err.Error())
		return instanceDTO, err
	}
	instanceDTO.Instance.Cluster = clusterModel.Cluster

	template, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		logger.Log.Errorf("Faile to get template info by name[%s]: %s", constant.POSTGRESQL, err.Error())
		return instanceDTO, err
	}
	instanceDTO.Instance.Template = template.Templates

	if err := kubernetes.GatherPostgresInfo(&instanceDTO); err != nil {
		logger.Log.Errorf("Faile to get postgres info: %s", err.Error())
		return instanceDTO, err
	}
	if err := kubernetes.GatherPostgresStatus(&instanceDTO); err != nil {
		logger.Log.Errorf("Faile to get postgres status: %s", err.Error())
		return instanceDTO, err
	}
	return instanceDTO, nil
}

func (c postgresService) GetById(ID string) (entity.Instance, error) {
	var instanceDTO entity.Instance
	mo, err := c.instanceRepo.GetByID(ID)
	if err != nil {
		logger.Log.Errorf("Faile to get instance info by id[%s]: %s", ID, err.Error())
		return instanceDTO, err
	}
	instanceDTO.Instance = mo
	clusterModel, err := c.clusterService.GetByID(mo.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", mo.ClusterID, err.Error())
		return instanceDTO, err
	}
	instanceDTO.Instance.Cluster = clusterModel.Cluster

	template, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		logger.Log.Errorf("Faile to get template info by name[%s]: %s", constant.POSTGRESQL, err.Error())
		return instanceDTO, err
	}
	instanceDTO.Instance.Template = template.Templates

	if err := kubernetes.GatherPostgresInfo(&instanceDTO); err != nil {
		logger.Log.Errorf("Faile to get postgres info: %s", err.Error())
		return instanceDTO, err
	}
	if err := kubernetes.GatherPostgresStatus(&instanceDTO); err != nil {
		logger.Log.Errorf("Faile to get postgres status: %s", err.Error())
		return instanceDTO, err
	}
	return instanceDTO, nil
}

func (c postgresService) GetStatusById(ID string) (string, error) {
	mo, err := c.instanceRepo.GetByID(ID)
	if err != nil {
		logger.Log.Errorf("Faile to get instance info by id[%s]: %s", ID, err.Error())
		return "", err
	}
	instanceEntity := entity.Instance{}
	instanceEntity.Instance = mo
	status, err := kubernetes.PostgresStatus(&instanceEntity)
	if err != nil {
		logger.Log.Errorf("Faile to get postgres status: %s", err.Error())
		return "", err
	}
	return status, nil
}

func (c postgresService) GetPostgresLog(instance entity.Instance, log chan string) {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by clusterid[%s]: %s", instance.ClusterID, err.Error())
		log <- err.Error()
	}
	instance.Cluster = clusterObj.Cluster
	instance.Namespace = constant.PostgresOperatorNamespace
	name, err := c.GetPostgresOperatorPodName(instance.ClusterID)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get postgres-operator pod name:%s", err.Error())
		logger.Log.Errorf(errMsg)
		log <- errMsg
	}
	instance.Name = name
	go kubernetes.GetLogs(&instance, log)
}

func (c postgresService) GetPostgresOperatorPodName(clusterId string) (string, error) {
	clusterObj, err := c.clusterService.GetByID(clusterId)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", clusterId, err.Error())
		return "", err
	}
	instance := entity.Instance{}
	instance.Cluster = clusterObj.Cluster
	instance.ClusterID = clusterObj.ID
	instance.Namespace = constant.PostgresOperatorNamespace
	podNames, err := c.podService.GetPods(instance)

	if err != nil {
		logger.Log.Errorf("Faile to get pod list: %s", err.Error())
		return "", err
	}
	regexString := fmt.Sprintf("%s-[a-z0-9-].*?", constant.PostgresOperatorPrefix)
	regex := regexp.MustCompile(regexString)
	for _, podName := range podNames {
		res := regex.MatchString(podName)
		if res {
			return podName, nil
		}
	}
	return "", constant.ErrResourceNotExist
}

func (c postgresService) Page(num, size int, userID string) (*entity.InstancePage, error) {
	user, err := c.userRepo.GetByID(userID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by userid[%s]: %s", userID, err.Error())
		return nil, err
	}
	if user.Role == 0 {
		return c.getPageForAdmin(num, size)
	} else {
		return c.getPageForUser(num, size, userID)
	}
}

func (c postgresService) Save(instance entity.Instance) error {
	loginfo, _ := json.Marshal(instance)
	logger.Log.WithFields(logrus.Fields{"instance_info": string(loginfo)}).Debugf("start to add the instance %s", instance.Name)
	//tx := db.DB.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	modelInstance := model.Instance{
		Name:          instance.Name,
		Type:          constant.POSTGRESQL,
		ClusterID:     instance.ClusterID,
		Namespace:     instance.Namespace,
		Count:         instance.Count,
		RequestCpu:    instance.RequestCpu,
		RequestMemory: instance.RequestMemory,
		LimitCpu:      instance.LimitCpu,
		LimitMemory:   instance.LimitMemory,
		Volume:        instance.Volume,
	}

	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", instance.ClusterID, err.Error())
		return err
	}
	modelInstance.Cluster = clusterObj.Cluster

	template1, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		logger.Log.Errorf("Faile to get template info by name[%s]: %s", constant.POSTGRESQL, err.Error())
		return err
	}
	modelInstance.TemplateID = template1.Templates.ID
	modelInstance.Template = template1.Templates

	var inst entity.Instance
	inst.Instance = modelInstance

	res, err := kubernetes.CreatePostgres(&inst)
	if err != nil {
		logger.Log.Errorf("Faile to create postgres instance: %s", err.Error())
		return err
	}

	inst.Instance.Status = res.Status.PostgresClusterStatus
	if err := c.instanceRepo.Save(&inst.Instance); err != nil {
		//if err := tx.Create(&inst.Instance).Error; err != nil {
		//	tx.Rollback()
		logger.Log.Errorf("Faile to insert postgres instance info to database: %s", err.Error())
		return fmt.Errorf("can not create postgres %s", err.Error())
	}

	user, err := c.userRepo.GetByID(instance.UserId)
	if err != nil {
		logger.Log.Errorf("Faile to get user info by id[%s]: %s", instance.UserId, err.Error())
		return err
	}
	userInstance := model.UserInstance{
		UserID:     user.ID,
		InstanceID: inst.Instance.ID,
	}
	if err := c.userInstanceRepo.Save(&userInstance); err != nil {
		//if err := tx.Create(&userInstance).Error; err != nil {
		//	tx.Rollback()
		logger.Log.Errorf("Faile to create user-instance info: %s", err.Error())
		return fmt.Errorf("can not create postgres%s", err.Error())
	}
	return nil
}

func (c postgresService) Create(postgres entity.Postgres) (string, error) {
	template, err := c.templateService.Get(constant.POSTGRESQL)
	if err != nil {
		logger.Log.Errorf("Faile to get template info by name[%s]: %s", constant.POSTGRESQL, err.Error())
		return "", err
	}
	clusterObj, err := c.clusterService.GetByID(postgres.ClusterId)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", postgres.ClusterId, err.Error())
		return "", err
	}
	postgres.Cluster = clusterObj
	var postgresql *v1.Postgresql
	if postgres.Type == constant.BASIC {
		err = yaml.Unmarshal([]byte(template.BaseTemplate), &postgresql)
		if err != nil {
			logger.Log.Errorf("Faile to unmarshal posrgres yaml: %s", err.Error())
			return "", err
		}
	} else if postgres.Type == constant.ADVANCE {
		err = yaml.Unmarshal([]byte(template.AdvanceTemplate), &postgresql)
		if err != nil {
			logger.Log.Errorf("Faile to unmarshal posrgres yaml: %s", err.Error())
			return "", err
		}
	}
	pInst, err := kubernetes.CreatePostgresFromTemplate(&postgres, postgresql)
	if err != nil {
		logger.Log.Errorf("Faile to create posrgres instance: %s", err.Error())
		return "", err
	}
	//tx := db.DB.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	modelInstance := model.Instance{
		Name:          postgres.Name,
		Type:          constant.POSTGRESQL,
		ClusterID:     postgres.ClusterId,
		Namespace:     postgres.Namespace,
		Count:         pInst.Spec.NumberOfInstances,
		RequestCpu:    pInst.Spec.Resources.ResourceRequests.CPU,
		RequestMemory: pInst.Spec.Resources.ResourceRequests.Memory,
		LimitCpu:      pInst.Spec.Resources.ResourceLimits.CPU,
		LimitMemory:   pInst.Spec.Resources.ResourceLimits.Memory,
		Volume:        pInst.Spec.Volume.Size,
		Status:        pInst.Status.PostgresClusterStatus,
	}
	modelInstance.TemplateID = template.Templates.ID
	modelInstance.Template = template.Templates
	if err := c.instanceRepo.Save(&modelInstance); err != nil {
		//if err := tx.Create(&modelInstance).Error; err != nil {
		//	tx.Rollback()
		logger.Log.Errorf("Faile to insert postgres instance info to database: %s", err.Error())
		return "", fmt.Errorf("can not create postgres %s", err.Error())
	}

	user, err := c.userRepo.GetByID(postgres.UserId)
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			logger.Log.Errorf("Faile to get user info by userid[%s]: %s", postgres.UserId, err.Error())
			return "", fmt.Errorf("Can't find user %s", err.Error())
		}
		logger.Log.Errorf("Faile to get user info by userid[%s]: %s", postgres.UserId, err.Error())
		return "", fmt.Errorf("Can't find user %s", err.Error())
	}
	userInstance := model.UserInstance{
		UserID:     user.ID,
		InstanceID: modelInstance.ID,
	}
	if err := c.userInstanceRepo.Save(&userInstance); err != nil {
		//if err := tx.Create(&userInstance).Error; err != nil {
		//	tx.Rollback()
		logger.Log.Errorf("Faile to create user-instance info: %s", err.Error())
		return "", fmt.Errorf("can not create postgres%s", err.Error())
	}
	return modelInstance.ID, nil
}

func (c postgresService) Sync(name string) (entity.Instance, error) {
	instance, err := c.instanceRepo.Get(name)
	if err != nil {
		logger.Log.Errorf("instance of %s not found error: %s", name, err.Error())
		return entity.Instance{Instance: instance}, err
	}
	var inst entity.Instance
	inst.Instance = instance

	clusterModel, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("instance of %s not found error: %s", name, err.Error())
		return entity.Instance{Instance: instance}, err
	}
	inst.Instance.Cluster = clusterModel.Cluster
	if err := kubernetes.GatherPostgresStatus(&inst); err != nil {
		logger.Log.Errorf("Faile to get postgres status: %s", err.Error())
		return inst, err
	}
	logger.Log.Infof("instance info:%v", inst)
	err = c.instanceRepo.Update(instance.ID, map[string]interface{}{
		"status": inst.Status,
	})
	if err != nil {
		logger.Log.Errorf("Faile to update postgres status: %s", err.Error())
		return inst, err
	}
	return inst, nil
}

func (c postgresService) Delete(name string) error {
	instance, err := c.instanceRepo.Get(name)
	if err != nil {
		logger.Log.Errorf("instance of %s not found error: %s", name, err.Error())
		return err
	}
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("instance of id not found error: %s", instance.ClusterID, err.Error())
		return err
	}
	instance.Cluster = clusterObj.Cluster
	instanceEntity := entity.Instance{}
	instanceEntity.Instance = instance
	err = kubernetes.DeletePostgres(&instanceEntity)
	if err != nil {
		logger.Log.Errorf("Faile to delete postgres instance: %s", err.Error())
		return err
	}
	err = c.instanceRepo.Delete(name)
	if err != nil {
		logger.Log.Errorf("Faile to delete postgres instance info from database: %s", err.Error())
		return err
	}
	return c.userInstanceRepo.DeleteByInstanceId(instance.ID)
}

func (c postgresService) Update(instance entity.Instance) error {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", instance.ClusterID, err.Error())
		return err
	}
	instance.Cluster = clusterObj.Cluster
	_, err = kubernetes.UpdatePostgres(&instance)
	if err != nil {
		logger.Log.Errorf("Faile to update postgres info: %s", err.Error())
		return err
	}
	jsonInstance, err := json.Marshal(instance)
	if err != nil {
		logger.Log.Errorf("Faile to marshal instance info: %s", err.Error())
		return err
	}
	mapInstance := make(map[string]interface{})
	err = json.Unmarshal(jsonInstance, &mapInstance)
	if err != nil {
		logger.Log.Errorf("Faile to unmarshal instance info: %s", err.Error())
		return err
	}
	return c.instanceRepo.Update(instance.ID, mapInstance)
}

func (c postgresService) Edit(instance entity.Instance) error {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by id[%s]: %s", instance.ClusterID, err.Error())
		return err
	}
	instance.Cluster = clusterObj.Cluster
	_, err = kubernetes.EditPostgres(&instance)
	if err != nil {
		logger.Log.Errorf("Faile to edit postgres info: %s", err.Error())
		return err
	}
	jsonInstance, err := json.Marshal(instance)
	if err != nil {
		logger.Log.Errorf("Faile to marshal instance info: %s", err.Error())
		return err
	}
	mapInstance := make(map[string]interface{})
	err = json.Unmarshal(jsonInstance, &mapInstance)
	if err != nil {
		logger.Log.Errorf("Faile to unmarshal instance info: %s", err.Error())
		return err
	}
	return c.instanceRepo.Update(instance.ID, mapInstance)
}

func (c postgresService) getPageForUser(num, size int, userID string) (*entity.InstancePage, error) {
	var (
		page        entity.InstancePage
		instances   []model.Instance
		instanceIds []string
	)
	userInstance, err := c.userInstanceRepo.Get(userID)
	if err != nil {
		logger.Log.Errorf("Faile to get instance infoby userid[%s]: %s", userID, err.Error())
		return nil, err
	}

	for _, instance := range userInstance {
		instanceIds = append(instanceIds, instance.InstanceID)
	}

	if err := db.DB.Model(&model.Instance{}).
		Where("id in (?) and type = ?", instanceIds, constant.POSTGRESQL).
		Preload("Cluster").
		Preload("Template").
		Count(&page.Total).
		Offset((num - 1) * size).
		Limit(size).
		Order("created_at ASC").
		Find(&instances).Error; err != nil {
		logger.Log.Errorf("Faile to get instance list info by instanceids[%s]: %s", instanceIds, err.Error())
		return nil, err
	}

	for _, mo := range instances {
		instanceDTO := entity.Instance{
			Instance: mo,
		}
		page.Items = append(page.Items, instanceDTO)
	}
	return &page, nil
}

func (c postgresService) getPageForAdmin(num, size int) (*entity.InstancePage, error) {
	var (
		page      entity.InstancePage
		instances []model.Instance
	)
	if err := db.DB.Model(&model.Instance{}).
		Where("type = ?", constant.POSTGRESQL).
		Preload("Cluster").
		Preload("Template").
		Count(&page.Total).
		Offset((num - 1) * size).
		Limit(size).
		Order("created_at ASC").
		Find(&instances).Error; err != nil {
		logger.Log.Errorf("Faile to get instance list info by type[%s]: %s", constant.POSTGRESQL, err.Error())
		return nil, err
	}

	for _, mo := range instances {
		instanceDTO := entity.Instance{
			Instance: mo,
		}
		page.Items = append(page.Items, instanceDTO)
	}
	return &page, nil
}

func (c postgresService) GetUsernameAndPassword(instance entity.Instance) ([]entity.SecretInfo, error) {
	clusterObj, err := c.clusterService.GetByID(instance.ClusterID)
	if err != nil {
		logger.Log.Errorf("Faile to get cluster info by clusterid[%s]: %s", instance.ClusterID, err.Error())
		return nil, err
	}
	instance.Cluster = clusterObj.Cluster
	secretInfos, err := c.secretService.GetSecrets(instance)
	if err != nil {
		logger.Log.Errorf("Failed to get postgres secrets:%s", err.Error())
		return nil, err
	}
	regexString := fmt.Sprintf(".*?\\.%s.credentials.postgresql.acid.zalan.do", instance.Name)
	regex := regexp.MustCompile(regexString)
	secrets := []entity.SecretInfo{}
	for _, secret := range secretInfos {
		res := regex.MatchString(secret.Name)
		if res {
			secrets = append(secrets, secret)
		}
	}
	return secrets, nil
}
