package chat

import (
	"fmt"
	"net/http"
)

func HandleMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Body)
}
