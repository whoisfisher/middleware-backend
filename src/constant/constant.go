package constant

import "errors"

const (
	DefaultPassword = "Def@u1tpwd"
)

const (
	BatchOperationUpdate = "update"
	BatchOperationCreate = "create"
	BatchOperationDelete = "delete"
)

var (
	NotSupportedBatchOperation = errors.New("not supported operation")
)

const (
	CenterCluster = "center"
)

const (
	ClusterNormal   = "Normal"
	ClusterInnormal = "Innormal"
)

const (
	BASIC      = "basic"
	ADVANCE    = "advance"
	POSTGRESQL = "postgres"
)

const (
	PARAM_MISSING = "parameter %v is missing, please check your requests parameters"
	TOKEN_INVALID = "token is invalid"
	PARAM_EMPTY   = "param_empty"
	NOT_SUPPORT   = "not_support"
)

const (
	Azure = "AZURE"
	S3    = "S3"
	OSS   = "OSS"
	Sftp  = "SFTP"
	Minio = "MINIO"
)

var (
	ErrOriginalNotMatch  = errors.New("ORIGINAL_NOT_MATCH")
	ErrUserNotFound      = errors.New("USER_NOT_FOUND")
	ErrUserIsNotActive   = errors.New("USER_IS_NOT_ACTIVE")
	ErrUserNameExist     = errors.New("NAME_EXISTS")
	ErrLdapDisable       = errors.New("LDAP_DISABLE")
	ErrEmailExist        = errors.New("EMAIL_EXIST")
	ErrNamePwdFailed     = errors.New("NAME_PASSWORD_SAME_FAILED")
	ErrEmailDisable      = errors.New("EMAIL_DISABLE")
	ErrEmailNotMatch     = errors.New("EMAIL_NOT_MATCH")
	ErrNameOrPasswordErr = errors.New("NAME_PASSWORD_ERROR")
	ErrResourceExist     = errors.New("RESOURCE_EXISTS")
	ErrResourceNotExist  = errors.New("RESOURCE_NOT_EXISTS")
)

const (
	SystemRoleSuperAdmin = 0
)

const (
	Local = "LOCAL"
	Ldap  = "LDAP"
)

const (
	ContainersReady string = "ContainersReady"
	PodInitialized  string = "Initialized"
	PodReady        string = "Ready"
	PodScheduled    string = "PodScheduled"
)

const (
	ConditionTrue    string = "True"
	ConditionFalse   string = "False"
	ConditionUnknown string = "Unknown"
)

const (
	PostgresOperatorNamespace = "postgres-operator"
	PostgresOperatorPrefix    = "postgres-operator"
)

const (
	PLATFORM_ADMIN             = "platform-admin"            // 管理 KubeSphere 平台上的所有资源。
	PLATFORM_REGULAR           = "platform-regular"          // 被邀请加入租户空间之前无法访问任何资源。
	PLATFORM_SELF_PROVISIONER  = "platform-self-provisioner" // 创建租户空间并成为所创建的租户空间的管理员。
	WORKSPACE_ADMIN            = "admin"                     // 管理租户空间中的所有资源。
	WORKSPAC_REGULAR           = "regular"                   // 查看租户空间设置。
	WORKSPACE_SELF_PROVISIONER = "self-provisioner"          // 查看企业设置、管理应用模板、创建环境和 DevOps 环境。
	WORKSPACE_VIEWER           = "viewer"                    // 查看租户空间中的所有资源。
	PROJECT_ADMIN              = "admin"                     // 管理环境中的所有资源。
	PROJECT_OPERATOR           = "operator"                  // 管理环境中除用户和角色之外的资源。
	PROJECT_VIEWER             = "viewer"                    // 查看环境中的所有资源。
)
