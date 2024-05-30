package alltypes

type LangConf struct {
	Active   string
	Path     string
	FileName string `yaml:"fileName"`
}

type DbConf struct {
	Dsn          string
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxIdleTime  string `yaml:"maxIdleTime"`
	Dialect      string `yaml:"dialect"`
}

type MsConf struct {
	Dsn  string
	Host string
	Port int
}

type HttpConf struct {
	Host string
	Port int
}

type RateLimiterConf struct {
	Enabled bool
	Rps     float64
	Burst   int
}
