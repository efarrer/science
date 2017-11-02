package science

import (
	"fmt"
	"net/http"
)

const VERSION = "0.1"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thanks for being willing to help with my science project by taking this survey! (%s)", VERSION)
}
