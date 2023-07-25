package model

import (
	"github.com/mensylisir/kmpp-middleware/src/model/common"
	uuid "github.com/satori/go.uuid"
)

type HelmRepository struct {
	common.BaseModel
	ID                    string `json:"id"`
	Name                  string `json:"name" gorm:"not null;unique"`
	Url                   string `json:"url" gorm:"not null;unique"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	IsActive              bool   `json:"is_active" gorm:"type:boolean;default:true"`
	CertFile              string `json:"cert_file"`
	KeyFile               string `json:"key_file"`
	CAFile                string `json:"ca_file"`
	InsecureSkipTlsVerify bool   `json:"insecure_skip_tls_verify"`
}

func (hr *HelmRepository) BeforeCreate() error {
	hr.ID = uuid.NewV4().String()
	return nil
}
