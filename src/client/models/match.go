package models

type Lance struct {
	Tipo      string `json:"tipo"`
	Minuto    int    `json:"minuto"`
	Descricao string `json:"descricao"`
	Autor     string `json:"autor"`
}
