package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	idNumber int
	rstring  string
	webPrint string
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func main() {
	web()
}

func random()  {
	rand.Seed(time.Now().UnixNano())
	idNumber = rand.Intn(1000)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func jsonobject()  {
	x := map[string]interface{}{"IDNumber": idNumber, "Feels": rstring}
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	webPrint = string(b)
}

func web()  {
	fmt.Println("RANDOM GENERATOR \nhttp://localhost/")
	for  {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			rstring = RandStringBytes(10)
			fmt.Fprint(w, webPrint)
			jsonobject()
			random()
		})
		http.ListenAndServe(":80", nil)
	}
}