package config

type Mongodb struct {
	Uri string `mapstructure:"uri" json:"uri" yaml:"uri"` //mongodb配置连接
}
