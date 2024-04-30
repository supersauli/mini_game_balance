package configs

import "time"

type System struct {
	OutsideHttpAddr string        `json:"outside-http-addr" yaml:"outside-http-addr" mapstructure:"outside-http-addr"`
	InsideHttpAddr  string        `json:"inside-http-addr" yaml:"inside-http-addr" mapstructure:"inside-http-addr"`
	WebHttpAddr     string        `json:"web-http-addr" yaml:"web-http-addr" mapstructure:"web-http-addr"`
	TimeoutSecond   time.Duration `json:"timeout-second" yaml:"timeout-second" mapstructure:"timeout-second"`
}
