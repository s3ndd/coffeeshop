#!/bin/sh

CGO_ENABLED=0 go build -o coffeeshop ./cmd/main.go

./coffeeshop