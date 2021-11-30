#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o apotik_be

# sshpass -p "Lh-3*LxDmSz32p4$" scp -P 8288 apotik_be root@103.82.242.11:/root

# env GOOS=windows GOARCH=amd64 go build -o oasis_be

