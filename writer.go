package main

import (
	"fmt"
	"log"
	"os"
	"math/rand"
	"time"
)

func write(data, tempDir string) {

	// Making directory for paste
	os.Mkdir("/var/www/html/paste/" + tempDir, 0777)
	file, err := os.Create("/var/www/html/paste/" + tempDir + "/index.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	// Write data to file
	fmt.Fprintf(file, data)
}

// Chars allowed for the random string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// Generates random string of 'n' length
func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}