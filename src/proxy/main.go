package main

import (
	"fmt"
	"github.com/kabukky/httpscerts"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request.
	host := strings.Split(r.Host, ":")
	http.Redirect(w, r, "https://"+host[0]+r.RequestURI, http.StatusMovedPermanently)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested URI: " + r.RequestURI)
	response, err := http.Get("http://ec2-54-166-49-196.compute-1.amazonaws.com")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(contents)
	}
}

func main() {
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "proxy")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}
	http.HandleFunc("/", handler)
	go http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
	// Start the HTTP server and redirect all incoming connections to HTTPS
	http.ListenAndServe(":8080", http.HandlerFunc(redirectToHttps))
}