package main

import (
	"net/http"
	"encoding/json"
	"time"
	"fmt"
	"strconv"
)

type Response struct {
	Desc string	`json:"desc"`
	Nano int64	`json:"nano"`
	Time string	`json:"time"`
}

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now().UTC()
		_str1 := "UTC time: " + now.String()
		_str2 := "UTC time in nano second precision: " + strconv.FormatInt(now.UnixNano(), 10)
		_str3 := _str1 + "\n" + _str2
		fmt.Println(_str3)

		res := &Response{
			Desc:   string("UTC time"),
			Nano:	now.UnixNano(),
			Time:	string(now.String()),
			}
		resJson, _ := json.Marshal(res)
		writer.Write([]byte(string(resJson)))
	})
	http.ListenAndServe(":8080", nil)
}