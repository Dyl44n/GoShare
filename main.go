package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "Dylan" {
		return "$1$c.FyXnmG$6rDIIU/XA0qZSA4wcmaT20"
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta> ")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	porta := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("Rodando Server na porta: %s ...", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil))
}
