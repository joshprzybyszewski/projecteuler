package easy

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

const (
	problem59CipherTextFile = `easy/p059_cipher.txt`
)

func SolveProblem59() string {
	/*
		decrypt the message and find the sum of the ASCII values
		in the original text.
	*/
	cipher := getProbelm59Ascii()
	// by running them all, then looking through them I found the expected key.
	// attemptAllProblem59Decodes(cipher)
	output := decipher(cipher, []byte(`exp`))
	outputInts := make([]int, len(output))
	for i := range output {
		outputInts[i] = int(output[i])
	}
	ans := mathUtils.Sum(outputInts)
	return fmt.Sprintf("%v", ans)
}

func getProbelm59Ascii() []byte {
	dat, err := ioutil.ReadFile(problem59CipherTextFile)
	if err != nil {
		log.Fatalf("problem 59 cannot read file: %v", err)
	}
	numbers := strings.Split(string(dat), ",")

	bs := make([]byte, len(numbers))
	for i, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatalf("problem 59 unexpected number: %v", err)
		}
		bs[i] = byte(n)
	}

	return bs
}

func decipher(cipher, key []byte) []byte {
	output := make([]byte, len(cipher))
	copy(output, cipher)

	for i, o := range output {
		output[i] = key[i%len(key)] ^ o
	}

	return output
}

func attemptAllProblem59Decodes(cipher []byte) {

	commonEnglishWords := []string{
		`and`,
		`of`,
		`in`,
		`on`,
		`a`,
		`hello`,
	}

	min := byte('a')
	max := byte('z')
	for k1 := min; k1 <= max; k1++ {
		for k2 := min; k2 <= max; k2++ {
			for k3 := min; k3 <= max; k3++ {
				key := []byte{k1, k2, k3}
				d := decipher(cipher, key)
				msg := string(d)
				if msgContainsAnyWords(msg, commonEnglishWords) {
					log.Printf("key %s => %q\n\n", string(key), msg)
				}
			}
		}
	}
}

func msgContainsAnyWords(msg string, words []string) bool {
	for _, word := range words {
		if strings.Contains(msg, ` `+word+` `) {
			return true
		}
	}
	return false
}
