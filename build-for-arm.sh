#!/bin/bash

GOOS=linux GOARCH=arm GOARM=7 go build mycrawler/main.go

exit 0

