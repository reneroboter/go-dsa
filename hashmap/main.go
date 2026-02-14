package main

import "fmt"

func main() {
	s := "abc"
	hash := hash(s)
	fmt.Println(hash)
}

func hash(s string) uint32 {
	var hash uint32 = 0
	for i := 0; i < len(s); i++ {
		hash = hash*31 + uint32(s[i])
	}
	return hash
}
