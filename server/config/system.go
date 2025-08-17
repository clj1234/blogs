package config

type System struct {
	Addr       string `yaml:"addr"`
	Port       string `yaml:"port"`
	SigningKey string `yaml:"signingKey"`
}
