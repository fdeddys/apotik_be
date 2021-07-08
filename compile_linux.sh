#!/usr/bin/env bash

# env GOOS=linux GOARCH=amd64 go build -o oasis_be

# scp oasis_be deddy@13.229.85.120:/home/deddy

env GOOS=windows GOARCH=amd64 go build -o oasis_be

