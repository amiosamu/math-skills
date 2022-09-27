package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Println(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("data.txt")
	if err != nil {
		log.Fatalf("readlines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
	if err := writeLines(lines, "data.new.txt"); err != nil {
		log.Fatalf("writelines: %s", err)
	}
}
*/

func main() {
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	ints := make([]int, len(lines))
	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}
	fmt.Println(ints[0] + ints[1])
}
