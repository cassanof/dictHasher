package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

func main() {
	createHashDict("test.txt", readLines("dicts/rockyou.txt"))
}

func createHashDict(fileName string, plaintextDict []string) {
	fmt.Printf("Creating hashed dictionary\n")
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	defer file.Close()

	for i := 0; i < len(plaintextDict)-1; i++ {
		// hash values are often printed in hex, for example
		// Use the `%x` format verb to convert
		// a hash results to a hex string.
		n, err := file.WriteString(hasher(plaintextDict[i]))

		if err != nil {
			log.Fatalf("Failed writing to file: %s %d", err, n)
		}
	}

	fmt.Printf("\nSuccessfully created hashed dictionary!")
	fmt.Printf("\nFile Name: %s", file.Name())
}

func readLines(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	// for _, eachline := range txtlines {
	// fmt.Println(eachline)
	// }

	return txtlines
}

func hasher(plaintext string) string {
	h := md5.New()

	h.Write([]byte(plaintext))

	ciphertext := h.Sum(nil)
	return fmt.Sprintf("%x\n", string(ciphertext))
}
