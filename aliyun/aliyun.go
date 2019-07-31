package aliyun

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// CommonParams type
type CommonParams struct {
	RegionID string `json:"RegionId"`
}

// Aliyun instance
type Aliyun struct {
	AccessKey       string
	AccessKeySecret string
}

// GetClient reurn a cms client
func (aliyun *Aliyun) GetClient(regionID string) (client *cms.Client, err error) {
	client, err = cms.NewClientWithAccessKey(regionID, aliyun.AccessKey, aliyun.AccessKeySecret)
	return
}

// NewWithEnv new Aliyun with env
func NewWithEnv() *Aliyun {
	var key, secret string
	if key = os.Getenv("AccessKey"); key == "" {
		panic("env AccessKey is required")
	}
	if secret = os.Getenv("AccessKeySecret"); secret == "" {
		panic("env AccessKeySecret is required")
	}
	return &Aliyun{
		AccessKey:       key,
		AccessKeySecret: secret,
	}
}
