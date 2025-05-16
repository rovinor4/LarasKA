package controller

import (
	"fmt"
	"laraska/model"
)

func AuthController() {

	PrintBoxWithText(60, []string{
		"Login",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

	var username, pass string
	fmt.Print("username : ")
	fmt.Scan(&username)
	fmt.Print("password : ")
	fmt.Scan(&pass)

	_, found := AuthCheck(model.ListAdmin, username, pass)
	if found {
		ClearScreen()
		MenuAwal()
	} else {
		PrintError("Username atau password salah")
		AuthController()
	}

}

func AuthCheck(admins []model.Admin, username, pass string) (model.Admin, bool) {
	for _, admin := range admins {
		if admin.Username == username && admin.Pass == pass {
			return admin, true
		}
	}
	return model.Admin{}, false
}
