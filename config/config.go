package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
}
