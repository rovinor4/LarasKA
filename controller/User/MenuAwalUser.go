package User

import (
	"bufio"
	"fmt"
	"laraska/controller"
	"laraska/model"
	"laraska/utils"
	"os"
	"strconv"
	"strings"
)

func Login() {
	render := bufio.NewReader(os.Stdin)

	utils.PrintBoxWithText(60, []string{
		"Login Sebagai Pengguna",
		"LarasKA (Layanan Reservasi Kereta Api)",
	})

FormUsername:
	fmt.Print("Masukan username (email/nomor hp) : ")
	username, err := render.ReadString('\n')
	username = strings.TrimSpace(username)
	if username == "" || err != nil {
		utils.PrintMessage("Username wajib di isi", "error")
		goto FormUsername
	}
FormPassword:
	fmt.Print("Masukan password : ")
	password, err := render.ReadString('\n')
	password = strings.TrimSpace(password)
	if password == "" || err != nil {
		utils.PrintMessage("Password wajib di isi", "error")
		goto FormPassword
	}

	isNull := true
	for _, usr := range model.ListUser {
		if (usr.Email == username || usr.NoHP == username) && usr.Password == password {
			controller.AuthData = controller.AuthStruct{
				User: usr,
			}
			isNull = false
			utils.ClearScreen()
			MenuAwalUser(controller.AuthData.User)
			return
		}
	}

	if isNull {
		utils.PrintMessage("Username dan password salah", "error")
		utils.Divider("-")
		goto FormUsername
	}

}

func MenuAwalUser(authUser model.User) {

	utils.PrintBoxWithText(60, []string{
		"LarasKA (Layanan Reservasi Kereta Api)",
		fmt.Sprintf("Hai, %s", authUser.NamaLengkap),
	})

	fmt.Println("[1] Pesan Tiket")
	fmt.Println("[2] History Pemesanan Tiket")
	fmt.Println("[3] Edit Akun")
	fmt.Println("[4] Log out")

	utils.Divider("-")

	InputSelect := utils.Input("Pilih menu: ", func(input string) (bool, string) {
		if input == "" {
			return false, "Input tidak boleh kosong"
		}
		if !utils.IsNumeric(input) {
			return false, "Input harus berupa angka"
		}

		if !utils.IsIn(input, []string{"1", "2", "3", "4"}) {
			return false, "Pilihan menu tidak tersedia"
		}

		return true, ""
	})

	input, _ := strconv.Atoi(InputSelect)

	switch input {
	case 1:
		utils.ClearScreen()
		PesanTiket(authUser)
	case 2:
		utils.ClearScreen()
		ShowHistoryTiket(authUser)
	case 3:
		utils.ClearScreen()
		EditAkunUser(&authUser)
	case 4:
		authUser = model.User{}
		utils.ClearScreen()
		Login()
	}

}
