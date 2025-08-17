package config

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Path     string `yaml:"path"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"dbName"`
	Config   string `yaml:"config"`
}

func (m *Mysql) GetDSN() string {
	result := m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName
	if m.Config != "" {
		return result + "?" + m.Config
	}
	return result
}
