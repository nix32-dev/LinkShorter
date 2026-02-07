package main

import (
	"fmt"
	"mainMod/shorcut"
	"net/http"
)

func main() {
	http.HandleFunc("/", shorcut.Redirect)
	http.HandleFunc("/create", shorcut.CreateLink)
	http.HandleFunc("/get", shorcut.GetLink)

	fmt.Println("Старт!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка: ", err)
	}
}
