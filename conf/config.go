package conf

type Config struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
}

type System struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Dir  string `mapstructure:"dir" json:"dir" yaml:"dir"`
}
