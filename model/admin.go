package model

type Admin struct {
	Name     string
	Username string
	Pass     string
	Role     int // 1 = admin, 2 = bording pass
}

var ListAdmin = []Admin{
	{Name: "Admin", Username: "admin", Pass: "admin", Role: 1},
	{Name: "Bording Pass", Username: "bording", Pass: "bording", Role: 2},
}
