package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	log.Println("Reading hashtags...")
	b, err := ioutil.ReadFile("hashtags.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// Convert the bytes into readable string
	str := string(b)

	// Remove whitespaces or line breaks
	fields := strings.Fields(str)

	// Sort from less..
	log.Println("Sorting your hashtags...")
	sort.Sort(byLength(fields))

	// Create a file in the local directory
	file, err := os.Create("hashtags-sorted.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	defer file.Close()

	// Just if you need a line break like us
	for index, hashtag := range fields {

		//		fmt.Println(hashtag)
		if index%3 == 0 && index != 0 {
			fmt.Fprint(file, "\n")
		}
		fmt.Fprint(file, hashtag+" ")
	}

	fmt.Println("Your hashtags has been sorted.")
}

// Sort functions
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
