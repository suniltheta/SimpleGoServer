# Using docker with GoLang latest
FROM golang:latest

# Maintainer information
MAINTAINER "Sunil Narasimhamurthy suniltheta.com" <suniltheta@gmail.com>

# Expose port so that user of this file knows
EXPOSE 8080

# Set label
LABEL version="1.0"

# Set working dir and copy application contents
WORKDIR /go/src/app
COPY app/* .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

# This command runs when docker starts
ENTRYPOINT ["go-wrapper","run"] # ["app"]