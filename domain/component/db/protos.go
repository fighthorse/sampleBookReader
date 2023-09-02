package db

type Mysql struct {
	Name   string      `mapstructure:"name"`
	Master MysqlConfig `mapstructure:"master"`
	Slave  MysqlConfig `mapstructure:"slave"`
}

type MysqlConfig struct {
	Driver         string  `mapstructure:"driver"`
	DSN            string  `mapstructure:"dsn"`
	MaxOpenConns   int32   `mapstructure:"max_open_conns"`
	MaxIdleConns   int32   `mapstructure:"max_idle_conns"`
	MaxLifeTimeout float64 `mapstructure:"max_life_timeout"`
}
