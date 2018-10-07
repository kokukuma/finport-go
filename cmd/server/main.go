package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type authData struct {
	UserID uint32 `json:"user_id"`
	UUID   string `json:"uuid"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Print(r)

	auth := authData{
		UserID: 1,
		UUID:   "12345678",
	}

	// jsonエンコード
	outputJSON, err := json.Marshal(&auth)
	if err != nil {
		fmt.Fprintf(w, "HelloWorld")
		return
	}

	// jsonヘッダーを出力
	w.Header().Set("Content-Type", "application/json")

	// jsonデータを出力
	fmt.Fprint(w, string(outputJSON))
}

func main() {
	http.HandleFunc("/", helloWorld)

	//
	keyFile := os.Getenv("KEY_PATH")
	crtFile := os.Getenv("CRT_PATH")

	//
	if keyFile == "" || crtFile == "" {
		log.Print("Start http")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		log.Print("Start https")
		log.Fatal(http.ListenAndServeTLS(":8080", crtFile, keyFile, nil))
	}
}
