package aliyun

// Dimension type
type Dimension struct {
	InstanceID string `json:"instanceId"`
}

// SLBDimension type
type SLBDimension struct {
	Dimension
	Port string `json:"port"`
}