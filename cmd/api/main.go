package main

import "fmt"

/*
go run :- build + executable for devlopment
go build -o bin/main ./cmd/api/main.go :- build for production bin folder main file and build main.go
./bin/main :- this is command for execute build file.
Makefile :- used as a shortcut
Makefile :- running like command make build, make run, make clean

*/

func main() {
	fmt.Println("Olx-Api server is running!!!")
}