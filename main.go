package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

var CORES int
var HASH_METHOD, TARGET, DESTINATION string

func main() {
	flag.StringVar(&TARGET, "t", "none", "REQUIRED. Specify the target (plaintext) dictionary to hash")
	flag.StringVar(&DESTINATION, "d", "none", "REQUIRED. Specify the destination of the hashed dictionary")
	flag.StringVar(&HASH_METHOD, "hash", "md5", "Specify the hash method. \nmd5, sha1, sha256, sha512")
	flag.IntVar(&CORES, "cores", 1, "Number of cores the program should use")
	flag.Parse()

	if TARGET == "none" || DESTINATION == "none" {
		fmt.Print("You must specify -t and -d.\nRead the help with the flag -h for more information.")
		os.Exit(-1)
	}

	runtime.GOMAXPROCS(CORES)

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go createHashDict(DESTINATION, readLines(TARGET), &waitgroup)
	waitgroup.Wait()

	fmt.Printf("\nProgram exited successfully\n")
}

func writeDict(plaintextDict []string, file *os.File) {
	if HASH_METHOD == "md5" {
		for i := 0; i < len(plaintextDict)-1; i++ {
			_, err := file.WriteString(hasherMD5(plaintextDict[i]))

			if err != nil {
				log.Fatalf("Failed writing to file: %s", err)
			}
		}
	} else if HASH_METHOD == "sha1" {
		for i := 0; i < len(plaintextDict)-1; i++ {
			_, err := file.WriteString(hasherSHA1(plaintextDict[i]))

			if err != nil {
				log.Fatalf("Failed writing to file: %s", err)
			}
		}
	} else if HASH_METHOD == "sha256" {
		for i := 0; i < len(plaintextDict)-1; i++ {
			_, err := file.WriteString(hasherSHA256(plaintextDict[i]))

			if err != nil {
				log.Fatalf("Failed writing to file: %s", err)
			}
		}
	} else if HASH_METHOD == "sha512" {
		for i := 0; i < len(plaintextDict)-1; i++ {
			_, err := file.WriteString(hasherSHA512(plaintextDict[i]))

			if err != nil {
				log.Fatalf("Failed writing to file: %s", err)
			}
		}
	} else {
		fmt.Printf("\nThe hash method specified does not exist. Stopping the program.\n")
		os.Exit(-1)
	}
}

func createHashDict(hashedDict string, plaintextDict []string, waitgroup *sync.WaitGroup) {
	fmt.Printf("Creating hashed dictionary\n")
	file, err := os.Create(hashedDict)

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	writeDict(plaintextDict, file)

	defer file.Close()

	fmt.Printf("\nSuccessfully created hashed dictionary!")
	fmt.Printf("\nFile Name: %s", file.Name())
	waitgroup.Done()
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

func hasherMD5(plaintext string) string {
	h := md5.New()

	h.Write([]byte(plaintext))

	hashed := h.Sum(nil)

	// hash values are often printed in hex, for example
	// Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x\n", string(hashed))
}

func hasherSHA1(plaintext string) string {
	h := sha1.New()

	h.Write([]byte(plaintext))

	hashed := h.Sum(nil)

	return fmt.Sprintf("%x\n", string(hashed))
}

func hasherSHA256(plaintext string) string {
	h := sha256.New()

	h.Write([]byte(plaintext))

	hashed := h.Sum(nil)

	return fmt.Sprintf("%x\n", string(hashed))
}

func hasherSHA512(plaintext string) string {
	h := sha512.New()

	h.Write([]byte(plaintext))

	hashed := h.Sum(nil)

	return fmt.Sprintf("%x\n", string(hashed))
}
