package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type solution struct {
	answer   string
	solved   bool
	num      int
	duration time.Duration
}

func buildOutput(
	solutions []solution,
) (int, string) {
	numSolved := 0
	var sb strings.Builder
	sb.WriteString("|Problem #|Answer|Duration|\n")
	sb.WriteString("|-:|-|-:|\n")
	for _, s := range solutions {
		if !s.solved {
			sb.WriteString(fmt.Sprintf("|%d|Skipped|N/A|\n", s.num))
			continue
		}
		numSolved++
		sb.WriteString(fmt.Sprintf("|%d|%s|%s|\n", s.num, s.answer, s.duration))
	}
	return numSolved, sb.String()
}

func writeToFile(
	fileName, fileString string,
) {
	answerFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer answerFile.Close()
	w := bufio.NewWriter(answerFile)
	_, err = w.WriteString(fileString)
	if err != nil {
		log.Fatal(err)
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
