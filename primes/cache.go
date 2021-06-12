package primes

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	primesCacheFileName = `primes/cache.txt`
)

func InitCache() {
	dat, err := ioutil.ReadFile(primesCacheFileName)
	if err != nil {
		log.Fatalf("primes cache cannot read file: %v", err)
	}
	lines := strings.Split(string(dat), "\n")
	if len(lines) == 0 {
		return
	}

	cache = newSliceCacheFromFile(lines)
}

func SaveCache() {
	lines := cache.knownToString()
	fileString := strings.Join(lines, "\n")

	mainFileWriter, err := os.Create(primesCacheFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer mainFileWriter.Close()
	w := bufio.NewWriter(mainFileWriter)
	w.WriteString(fileString)

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
