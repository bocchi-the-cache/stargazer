package model

type Target struct {
	Name        string      `json:"name"`
	Enabled     bool        `json:"enabled"`
	Type        string      `json:"type"`
	Destination string      `json:"destination"`
	Interval    string      `json:"interval"`
	Timeout     string      `json:"timeout"`
	Thresholds  []Threshold `json:"thresholds"`
}

type Threshold struct {
	Level    string `json:"level"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}
