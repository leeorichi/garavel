package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Respond struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj := Respond{
		Status:  true,
		Message: "Hello from Garavel",
	}

	fmt.Printf("got /hello request\n")
	// io.WriteString(w, message)
	w.Write([]byte{})

	respBytes, _ := json.Marshal(obj)
	w.Write(respBytes)

}
