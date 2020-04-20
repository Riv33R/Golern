package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	fmt.Println("RANDOM GENERATOR \nhttp://localhost/")
	for {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, jsonobject())
		})
		http.ListenAndServe(":80", nil)
	}
}

func random() {
	rand.Seed(time.Now().UnixNano())
}

//RandomString ...
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = characterSet[rand.Int63()%int64(len(characterSet))]
	}
	return string(b)
}

func jsonobject() string {
	x := map[string]interface{}{"IDNumber": rand.Intn(1000), "Feels": RandomString(15)}
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}
