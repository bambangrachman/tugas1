package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Temans struct {
	Temans []Teman `json:"teman"`
}

type Teman struct {
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Alasan    string `json:"alasan"`
}

func main() {
	fmt.Println("halo")
	dataJson, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	defer dataJson.Close()

	byteValue, _ := io.ReadAll(dataJson)

	var temans Temans

	json.Unmarshal(byteValue, &temans)

	args := os.Args

	if len(args) == 0 {
		fmt.Println("tolong masukan inputan")
	} else {
		i, _ := strconv.Atoi(args[1])
		fmt.Println(i)

		fmt.Println("Nama:", temans.Temans[i-1].Nama)
		fmt.Println("Alamat:", temans.Temans[i-1].Alamat)
		fmt.Println("Pekerjaan:", temans.Temans[i-1].Pekerjaan)
		fmt.Println("Alasan Mempelajari Golang:", temans.Temans[i-1].Alasan)
	}

}
