#!/bin/bash

rm -f ./bin/createRSVP
GOOS=linux go build -ldflags="-s -w" -o bin/createRSVP rest/rsvp/main.go
