package main
type config struct {
	API apiconfig `yaml:"api"`
}
type apiconfig struct{
	Port int `yaml:"port"`
}
