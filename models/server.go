package models

type Server struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    string `json:"port"`
	Status  string `json:"status"`
}
