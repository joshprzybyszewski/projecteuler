package primes

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	primesCacheFileName    = `primes/cache.txt`
	primesBigCacheFileName = `primes/big_cache.txt`
)

func InitCache() {
	dat, err := ioutil.ReadFile(primesBigCacheFileName)
	if err != nil {
		dat, err = ioutil.ReadFile(primesCacheFileName)
		if err != nil {
			log.Printf("primes cache cannot read file: %v", err)
			return
		}
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

	cacheFile, err := os.Create(primesBigCacheFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer cacheFile.Close()
	w := bufio.NewWriter(cacheFile)
	_, err = w.WriteString(fileString)
	if err != nil {
		log.Fatal(err)
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
