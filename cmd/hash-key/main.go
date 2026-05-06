package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	scope := flag.String("scope", "read", "API key scope: read or crud")
	cost := flag.Int("cost", bcrypt.DefaultCost, "bcrypt cost")
	flag.Parse()

	if *scope != "read" && *scope != "crud" {
		log.Fatal("scope must be read or crud")
	}
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: go run ./cmd/hash-key --scope read <api-key>\n")
		os.Exit(2)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(flag.Arg(0)), *cost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s:%s\n", hash, *scope)
}
