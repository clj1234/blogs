package config

type Server struct {
	Mysql      Mysql      `yaml:"mysql"`
	GormConfig GormConfig `yaml:"gorm"`
	System     System     `yaml:"system"`
}
