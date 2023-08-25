package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello world...")

}
