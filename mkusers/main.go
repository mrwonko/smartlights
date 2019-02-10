package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userData struct {
	Name         string
	PasswordHash []byte // bcrypt
}

func main() {
	fmt.Println("enter rows of key:password:")
	reader := bufio.NewReader(os.Stdin)
	var users []userData
readLoop:
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			pieces := strings.SplitN(line[:len(line)-1], ":", 2)
			if len(pieces) != 2 {
				log.Fatal("invalid input contained no ':'")
			}
			hash, err := bcrypt.GenerateFromPassword([]byte(pieces[1]), bcrypt.DefaultCost)
			if err != nil {
				log.Fatalf("error hashing password: %s", err)
			}
			users = append(users, userData{
				Name:         pieces[0],
				PasswordHash: hash,
			})
		}
		switch err {
		case nil:
		case io.EOF:
			break readLoop
		default:
			log.Fatalf("failed to read from stdin: %s", err)
		}
	}
	if err := json.NewEncoder(os.Stdout).Encode(users); err != nil {
		log.Fatalf("error encoding JSON: %s", err)
	}
}
