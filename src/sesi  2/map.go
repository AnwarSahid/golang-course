package main

import "fmt"

func main() {
	var nameObject = []map[string]string{

		{"name": "anwar Sahid", "job": "on go"},
		{"name": "anwar Sahid2", "job": "on go2"},
		{"name": "anwar Sahid3", "job": "on go3"},
	}
	for _, nameObject := range nameObject {
		fmt.Println(nameObject["name"])
	}

	var name = map[string]string{
		"name": "cek",
		"s":    "cek",
		"c":    "cek",
	}

	fmt.Println(name["name"])
}
