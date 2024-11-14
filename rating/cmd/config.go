package main
type config struct {
	API apiconfig `yaml:"api"`
}
type apiconfig struct{
	port int `yaml:"port"`
}