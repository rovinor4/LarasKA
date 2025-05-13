package main

import (
	"fmt"
	"laraska/model"
	"laraska/view"
)

func main() {
	var DataStasiun model.Stasiun
	fmt.Scan(&DataStasiun.Nama)
	view.FormStasiun(DataStasiun)
}
