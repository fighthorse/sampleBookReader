package log

type ConfigLog struct {
	LogType string `mapstructure:"type"`
	App     struct {
		FilePath string `mapstructure:"file_path"`
		Level    string `mapstructure:"level"`
		Enable   bool   `mapstructure:"enable"`
	} `mapstructure:"app"`

	Access struct {
		FilePath string `mapstructure:"file_path"`
		Enable   bool   `mapstructure:"enable"`
	} `mapstructure:"access"`
}
