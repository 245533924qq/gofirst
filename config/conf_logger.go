package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       int    `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`      // 行号
	LogInConsole bool   `yaml:"log_in_console"` //显示路径
}
