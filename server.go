package main

import (
	"fmt"
	"github.com/frankrafael/predictive-maintenance-go/src/model"
)

func main() {

	var frank *model.User
	frank = new(model.User)
	
	fmt.Print(frank)
  /*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8083", nil))*/



}