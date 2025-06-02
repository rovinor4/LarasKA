package User

import (
	"fmt"
	"laraska/model"
	"laraska/utils"
	"regexp"
	"strings"
)

func EditAkunUser(authUser *model.User) {
	utils.PrintHead("Edit User")

	utils.PrintBoxWithText(60, []string{
		fmt.Sprintf("NIK: %s", authUser.NIK),
		fmt.Sprintf("Nama: %s", authUser.NamaLengkap),
		fmt.Sprintf("Pekerjaan: %s", authUser.Pekerjaan),
	})

	fmt.Println("Biarkan kosong jika tidak ingin mengeditnya . . .")

	nik := utils.Input("NIK: ", func(nik string) (bool, string) {
		if !regexp.MustCompile(`^\d{16}$`).MatchString(nik) && nik != "" {
			return false, "Format NIK salah"
		}

		return true, ""
	})

	nama := utils.Input("Nama: ", func(nama string) (bool, string) {
		return true, ""
	})

	pekerjaan := utils.Input("Pekerjaan: ", func(nama string) (bool, string) {
		return true, ""
	})

	if nik != "" && strings.TrimSpace(authUser.NIK) != nik {
		nikFmt := fmt.Sprintf("NIK: %s -> %s", authUser.NIK, nik)
		utils.PrintMessage(nikFmt, "success")
	} else {
		fmt.Printf("NIK: %s (Tidak Berubah)\n", authUser.NIK)
	}

	if nama != "" && strings.TrimSpace(authUser.NamaLengkap) != nama {
		namaFmt := fmt.Sprintf("Nama: %s -> %s", authUser.NamaLengkap, nama)
		utils.PrintMessage(namaFmt, "success")
	} else {
		fmt.Printf("Nama: %s (Tidak Berubah)\n", authUser.NamaLengkap)
	}

	if pekerjaan != "" && strings.TrimSpace(authUser.Pekerjaan) != pekerjaan {
		pekerjaanFmt := fmt.Sprintf("Pekerjaan: %s -> %s", authUser.Pekerjaan, pekerjaan)
		utils.PrintMessage(pekerjaanFmt, "success")
	} else {
		fmt.Printf("Pekerjaan: %s (Tidak Berubah)\n", authUser.Pekerjaan)
	}

	yesno := utils.Input("Apakah anda yakin akan melakukan perubahan ini (y/n): ", func(yesno string) (bool, string) {
		if yesno == "" {
			return false, "Mohon konfirmasi (y/n)"
		}

		if strings.ToLower(yesno) == "y" || strings.ToLower(yesno) == "n" {
			return true, ""
		}

		return false, "Pilihan tidak tersedia"
	})

	if strings.ToLower(yesno) == "y" {
		for i, usr := range model.ListUser {
			if usr.NIK == authUser.NIK {
				if nik != "" {
					authUser.NIK = nik
					model.ListUser[i].NIK = nik
				}

				if nama != "" {
					authUser.NamaLengkap = nama
					model.ListUser[i].NamaLengkap = nama
				}

				if pekerjaan != "" {
					authUser.Pekerjaan = pekerjaan
					model.ListUser[i].Pekerjaan = pekerjaan
				}
			}
		}
		utils.ClearScreen()
	} else {
		utils.ClearScreen()
		utils.PrintMessage("Edit User Dibatalkan", "warning")
	}
	MenuAwalUser(*authUser)

}
