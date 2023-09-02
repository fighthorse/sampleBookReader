package trace_redis

type Redis struct {
	Name         string  `mapstructure:"name"`
	Addr         string  `mapstructure:"addr"`
	Pwd          string  `mapstructure:"pwd"`
	Db           float64 `mapstructure:"db"`
	DialTimeout  float64 `mapstructure:"dial_timeout"`
	ReadTimeout  float64 `mapstructure:"read_timeout"`
	WriteTimeout float64 `mapstructure:"write_timeout"`
	PoolSize     float64 `mapstructure:"pool_size"`
	MinIdleConns float64 `mapstructure:"min_idle_conns"`
	MaxRetries   float64 `mapstructure:"max_retries"`
}
