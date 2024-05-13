package configs

type Balance struct {
	Formula         string `mapstructure:"formula" json:"formula" yaml:"formula"`                                  // 计算分数公式
	CheckTimeOutSec int    `mapstructure:"check_time_out_sec" json:"check_time_out_sec" yaml:"check_time_out_sec"` // 检查超时时间
	MaxTimeOutSec   int    `mapstructure:"max_time_out_sec" json:"max_time_out_sec" yaml:"max_time_out_sec"`       // 最大超时时间
}
