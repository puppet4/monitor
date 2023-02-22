package config

type GenrealDB struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
}
