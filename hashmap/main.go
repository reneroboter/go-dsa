package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	readLinesFromFileAndHash()
}

func readLinesFromFileAndHash() error {
	m := make(map[int]int)
	fileHandle, err := os.Open("wordlist.txt")
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	scanner := bufio.NewReader(fileHandle)
	for {
		textLine, err := scanner.ReadString('\n')
		if err == io.EOF {
			if len(textLine) != 0 {
				hashedTextline := hash(textLine)
				m[int(hashedTextline)] = m[int(hashedTextline)] + 1
			}
			break
		}
		if err != nil {
			return fmt.Errorf("error reading from file: %w", err)
		}
		hashedTextline := hash(textLine)
		m[int(hashedTextline)] = m[int(hashedTextline)] + 1
	}

	displayResults(m)

	return nil
}

func displayResults(m map[int]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func hash(s string) uint32 {
	var hash uint32 = 0
	for i := 0; i < len(s); i++ {
		hash = hash*31 + uint32(s[i])
	}
	return hash
}
