package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func help() {
	fmt.Print(`
TOTP code generator.

This tool accepts a single TOTP seed as an argument and prints the current code to stout.

Usage: totpgen [seed]

`)
}

func stringinarray(array []string, string string) bool {
	for _, value := range array {
		if value == string {
			return true
		}
	}
	return false
}

var seed string

func main() {
	for _, arg := range os.Args[1:] {
		if stringinarray([]string{"-h", "--help", "help"}, arg) {
			help()
			os.Exit(0)
		}
	}
	if len(os.Args[1:]) != 1 {
		log.Printf("error: wrong number of arguments\nargs: %s", os.Args[1:])
		help()
		os.Exit(1)
	}
	seed = os.Args[1]
	code, _ := totp.GenerateCode(seed, time.Now())
	fmt.Println(code)
}
