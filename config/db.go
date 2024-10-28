package config

type SysMysql struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func (m *SysMysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Database + "?" +
		"charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai&timeout=10s&readTimeout=30s&writeTimeout=30s"
}
