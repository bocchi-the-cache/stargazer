package conf

type AllConfig struct {
	Service ServiceConfig `yaml:"service"`
	Data    DataConfig    `yaml:"data"`
	Swagger SwaggerConfig `yaml:"swagger"`
}

type ServiceConfig struct {
	Label          string     `yaml:"label"`
	Http           HttpConfig `yaml:"http"`
	LogLevel       string     `yaml:"log_level"`
	LogCleanBefore int        `yaml:"log_clean_before"`
}

type HttpConfig struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

type DataConfig struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Driver     string `yaml:"driver"`
	Connection string `yaml:"connection"`
}

type SwaggerConfig struct {
	Enabled bool   `yaml:"enabled"`
	Addr    string `yaml:"addr"`
	Port    int    `yaml:"port"`
}
