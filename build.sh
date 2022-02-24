#!/bin/bash

mkdir -p ../hashcode-2022-eval

env GOOS=linux GOARCH=amd64 go build -o ../hashcode-2022-eval/eval .
env GOOS=windows GOARCH=amd64 go build -o ../hashcode-2022-eval/eval.exe .
