package user

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateKodePenumpang() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return "P-" + time.Now().Format("20060102") + "-" + strconv.Itoa(rand.Intn(8000)) // YYYYMMDD
}

func GenerateKodeTiket() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return "T-" + time.Now().Format("020106") + "-" + strconv.Itoa(rand.Intn(10000)) // DDMMYY
}
