package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	args := os.Args[1:]
	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	ints := make([]int, 0)
	for _, s := range lines {
		temp, _ := strconv.Atoi(s)
		ints = append(ints, temp)
	}
	ints = ints[:len(ints)-1]
	average := 0
	// median := 0
	variance := 0
	// standardDeviation := 0
	sum := 0
	var mean int

	for i := 0; i < len(ints); i++ {
		sum += ints[i]
	}
	average = mean
	mean = sum / len(ints)
	mean = mean + 1
	// 15 =avg, (1-15)^2 + 2-15^2 + 3-15^2 + 4-15^2 + 5-15^2, 196+169+144+121+100 =
	fmt.Println(ints[0] - mean)
	for i := 0; i < len(ints); i++ {
		fmt.Println(quadrat(ints[i] - mean))
	}
	// variance = variance / len(ints)
	// standardDeviation = int(math.Sqrt(float64(variance)))
	fmt.Println(variance)
	fmt.Println(sum)
	fmt.Println(mean)
	fmt.Println(ints)
}

func quadrat(n int) int {
	return n * n
}
