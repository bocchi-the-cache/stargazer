package model

type ServerConfig struct {
	Label          string `json:"label" mapstructure:"label"`
	PanelPort      int    `json:"panel_port" mapstructure:"panel_port"`
	PrometheusPort int    `json:"prometheus_port" mapstructure:"prometheus_port"`
	LogLevel       string `json:"log_level" mapstructure:"log_level"`
}
