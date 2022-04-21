package config

var Version string = "v0.0.1"

type TLSConfig struct {
	Enabled bool   `yaml:"enabled"`
	Cert    string `yaml:"cert"`
	Key     string `yaml:"key"`
}

type ServerConfig struct {
	Addr string    `yaml:"addr"`
	TLS  TLSConfig `yaml:"tls"`
}

type AuthConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type DirConfig struct {
	Path     string     `yaml:"path"`
	Readonly bool       `yaml:"readonly"`
	Auth     AuthConfig `yaml:"auth"`
}

type WebdavdConfig struct {
	Server   ServerConfig `yaml:"server"`
	RootPath string       `yaml:"rootPath"`
	Dir      DirConfig    `yaml:"dir"`
}
