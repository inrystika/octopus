package config

type Config struct {
	Harbor Harbor `yaml:"harbor"`
}

type Harbor struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SSL      bool   `yaml:"ssl"`
}
