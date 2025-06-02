package User

import (
	"bufio"
	"fmt"
	"laraska/model"
	"laraska/utils"
	"os"
	"strconv"
	"strings"
)

func ShowHistoryTiket(authUser model.User) {
	reader := bufio.NewReader(os.Stdin)
	var mappedTiket []map[string]string

	if len(model.ListTiket) == 0 {
		utils.ClearScreen()
		utils.PrintMessage("Belum ada tiket yang dipesan!", "error")
		MenuAwalUser(authUser)
	}

	for i, tiket := range model.ListTiket {
		var penumpangList strings.Builder
		for _, p := range tiket.Penumpang {
			penumpangList.WriteString(fmt.Sprint(p.Nama))
		}
		penumpangStr := strings.TrimSuffix(penumpangList.String(), ", ")

		mappedTiket = append(mappedTiket, map[string]string{
			"no":       strconv.Itoa(i + 1),
			"kode":     tiket.Kode,
			"atasNama": penumpangStr,
			"rute":     fmt.Sprintf("%s - %s", tiket.Rute.StasiunAwal.Nama, tiket.Rute.StasiunAkhir.Nama),
			"jadwal":   tiket.Rute.RuteBerhenti[0].Berangkat.Format("02/01/2006 15:04"),
		})
	}

	utils.PrintTable(
		[]string{"No.", "Kode", "Atas Nama", "Dari/Ke", "Jadwal"},
		mappedTiket,
		[]string{"no", "kode", "atasNama", "rute", "jadwal"},
		2,
		"History Pemesanan Tiket",
	)

	fmt.Println("[1] Detail tiket")
	fmt.Println("[2] Kembali")

	input := utils.Input("Pilih Menu: ", func(value string) (bool, string) {
		if value == "" {
			return false, ""
		}

		if !utils.IsIn(value, []string{"1", "2"}) {
			return false, "Pilihan tidak tersedia"
		}

		return true, ""

	})

	switch input {
	case "1":
		detailTiket(authUser, *reader)
	case "2":
		utils.ClearScreen()
		MenuAwalUser(authUser)
	default:
		utils.ClearScreen()
		if input != "" {
			s := fmt.Sprintf("Tidak ada pilihan \"%s\"", input)
			utils.PrintMessage(s, "error")
		}
		ShowHistoryTiket(authUser)
	}
}

// TODO: Tiket layout
func detailTiket(authUser model.User, reader bufio.Reader) {
	inputNomor := utils.Input("Pilih nomor tiket: ", func(value string) (bool, string) {
		if value == "" {
			return false, ""
		}

		return true, ""

	})

	nomor, err := strconv.Atoi(inputNomor)
	if err != nil {
		errFmt := fmt.Sprintf("Invalid input \"%s\"", inputNomor)
		utils.PrintMessage(errFmt, "error")
		detailTiket(authUser, reader)
	} else if nomor < 1 || nomor > len(model.ListTiket) {
		errFmt := fmt.Sprintf("Error: Out of range \"%s\"", strconv.Itoa(nomor))
		utils.PrintMessage(errFmt, "error")
		detailTiket(authUser, reader)
	}

	selectedTiket := model.ListTiket[nomor-1]
	ruteBerhenti := selectedTiket.Rute.RuteBerhenti

	bannerFmt := fmt.Sprintf("%s - %s | %s", selectedTiket.Rute.StasiunAwal.Kota, selectedTiket.Rute.StasiunAkhir.Kota, selectedTiket.Rute.Nama)
	utils.PrintHead(bannerFmt)
	utils.PrintBoxWithText(60, []string{
		"\n",
		fmt.Sprintf("%s - %s", selectedTiket.Rute.StasiunAwal.Nama, selectedTiket.Rute.StasiunAkhir.Nama),
		fmt.Sprintf("Atas Nama: %s", selectedTiket.Penumpang[0].Nama),
		fmt.Sprintf("Jadwal Keberangkatan : %s", selectedTiket.Rute.RuteBerhenti[0].Berangkat.Format("2 January 2006 15:04")),
		fmt.Sprintf("Kode Tiket: %s", selectedTiket.Kode),
		fmt.Sprintf("Kereta: %s (%s)", selectedTiket.Rute.Kereta.Nama, selectedTiket.Rute.Kereta.Kelas),
		fmt.Sprintf("Gerbong: %d", selectedTiket.Penumpang[0].Gerbong),
		fmt.Sprintf("Tempat Duduk: %s", selectedTiket.Penumpang[0].TempatDuduk),
		fmt.Sprintf("Perkiraan Tiba: %s", ruteBerhenti[len(ruteBerhenti)-1].Tiba.Format("2 January 2006 15:04")),
		fmt.Sprintf("Harga: %s", utils.RupiahFormat(selectedTiket.Rute.Harga)),
		"\n",
	})

	fmt.Print("Press ENTER to continue")
	reader.ReadByte()
	fmt.Println()
	utils.ClearScreen()
	ShowHistoryTiket(authUser)
}
