package main

import (
	"fmt"
	"os"
)

func cek() {
	data := map[string]string{
		"1":  `{"id": "1", "nama": "Anwar Sahid", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"mengisi waktu luang"}`,
		"2":  `{"id": "2", "nama": "Dian", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"mengisi waktu sempit"}`,
		"3":  `{"id": "3", "nama": "Doni", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"naganggur"}`,
		"4":  `{"id": "4", "nama": "Marwan", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"pengen sibuk"}`,
		"5":  `{"id": "5", "nama": "Faro", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"yntkts"}`,
		"6":  `{"id": "6", "nama": "Master Wah", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"pengen beli mobil"}`,
		"7":  `{"id": "7", "nama": "Marc Andreas", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"pengen jadi pro golang"}`,
		"8":  `{"id": "8", "nama": "Yuniarti", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"aku tak tau"}`,
		"9":  `{"id": "9", "nama": "Kalianda pratama", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"senang golang"}`,
		"10": `{"id": "10", "nama": "jiah xiaw", "Alamat": "sinar Jaya", "Pekerjaan": "developer", "alasan" :"fresh graduate aku"}`,
	}

	id := os.Args[1]
	fmt.Println(data[id])
	if id > "10" {
		fmt.Println("data tidak boleh lebih dari 10 atau harus angka")
	} else if id < "1" {
		fmt.Println("data tidak boleh kurang dari 1")
	}
}

func main() {
	cek()
}
