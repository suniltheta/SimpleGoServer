# SimpleGoServer
A simple web server in GoLang which replies time which can be deployed as docker container.

<br>

## Steps to run
Install GoLang in Ubuntu <br>
` $ sudo apt-get install golang-go`

Set GOBIN env var <br>
`$ export  GOBIN=/home/{USER}/bin`

Git clone this repo <br>
`$ git clone https://github.com/suniltheta/SimpleGoServer`

Install src/time as executable into GOBIN <br>
`$ go install SimpleGoServer/src/time/main.go`

Run the application <br>
`$ /home/{USER}/bin/main`

## How to use?
Go to the home page of URL and you can see the string response as <br>
`{"desc":"UTC time","nano":1510968028319270900,"time":"2017-11-18 01:20:28.3192709 +0000 UTC"}`
