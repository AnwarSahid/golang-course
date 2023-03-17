package main

import (
	"fmt"
	"os"
	"strconv"
)

type Siswa struct {
	Id, Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	data := []Siswa{
		{Id: "0", Nama: "Anwar Sahid", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "mengisi waktu luang"},
		{Id: "1", Nama: "Dian", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "mengisi waktu sempit"},
		{Id: "2", Nama: "Doni", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "naganggur"},
		{Id: "3", Nama: "Marwan", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "pengen sibuk"},
		{Id: "4", Nama: "Faro", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "yntkts"},
		{Id: "5", Nama: "Master Wah", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "pengen beli mobil"},
		{Id: "6", Nama: "Marc Andreas", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "pengen jadi pro golang"},
		{Id: "7", Nama: "Yuniarti", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "aku tak tau"},
		{Id: "8", Nama: "Kalianda pratama", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "senang golang"},
		{Id: "9", Nama: "jiah xiaw", Alamat: "sinar Jaya", Pekerjaan: "developer", Alasan: "fresh graduate aku"},
	}

	indexStr := os.Args[1]
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(data[index])

	if index < 0 || index >= len(data) {
		fmt.Printf("Error: Data tidak sesuai\n")
	}

}
