package health

import (
	"fmt"
	"net/http"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server running fine!")
}
