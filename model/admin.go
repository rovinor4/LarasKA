package model

type Admin struct {
	Name     string
	Username string
	Pass     string
}

var ListAdmin = []Admin{
	{Name: "Admin", Username: "admin", Pass: "admin"},
	{Name: "Bording Pass", Username: "bording", Pass: "bording"},
}
