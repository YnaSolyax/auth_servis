package main

import (
	"fmt"
	"net/http"

	"rest_api/access"
)

// - migrate -source=file:///home/pentagon/telegram_bot/db -database=postgres://postgres:1@localhost:5432/bot up

func main() {
	token, err := access.AccessToken("123")
	if err != nil {
		fmt.Println("Ошибка генерации токена:", err)
		return
	}

	fmt.Println("Starting server at port 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Authorization", token)
	})
	//http.HandleFunc("/access,")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
