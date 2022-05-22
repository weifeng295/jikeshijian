package config

// MysqlConfig mysql 配置 接口
type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

// defaultMysqlConfig mysql 配置
type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

// GetURL URL mysql 连接
func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

// GetEnabled Enabled 激活
func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enable
}

// GetMaxIdleConnection 闲置连接数
func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// GetMaxOpenConnection 打开连接数
func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}
