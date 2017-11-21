package main

import (
    "fmt"
    "net/http"
	"os"
	"time"
	"strconv"
    "crypto/tls"
    )

func main() {
	//http://ec2-54-166-49-196.compute-1.amazonaws.com // Website
	//https://ec2-54-167-207-16.compute-1.amazonaws.com // Proxy

	before := time.Now().UTC().UnixNano()
	response, err := http.Get("http://our-website.com")
	after := time.Now().UTC().UnixNano()
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()        
        fmt.Printf("Time taken for request without Proxy: %s\n", strconv.FormatInt(after - before, 10))
	}
	
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	before = time.Now().UTC().UnixNano()
	_, err = client.Get("https://our-proxy.com")
	after = time.Now().UTC().UnixNano()
	if err != nil {
		fmt.Println(err)
	}else {			
		fmt.Printf("Time taken for request with Proxy: %s\n", strconv.FormatInt(after - before, 10))
	}	
}