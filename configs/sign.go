package configs

// Sign 签名
type Sign struct {
	Token string `json:"token" yaml:"token"`
	Use   bool   `json:"use" yaml:"use"`
}
