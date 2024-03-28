package main

import (
	"fmt"
	"net/http"

	"github.com/ohayouarmaan/learning_go/internal/routers"
)

func main() {
	fmt.Println("Hello, world!")
	r := router.GetRouter()
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:3000", r)
}
