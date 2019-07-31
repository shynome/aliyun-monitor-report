package aliyun

import (
	"os"
)

// Aliyun instance
type Aliyun struct {
	AccessKey       string
	AccessKeySecret string
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
