package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"01.alem.school/git/dias_amitullayev/math-skills/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	file, _ := ioutil.ReadFile(args[0])
	str := string(file)
	res := strings.Fields(str)
	keys := make([]float64, 0)
	for _, val := range res {
		key, _ := strconv.Atoi(val)
		keys = append(keys, float64(key))
	}
	/*
		out_avg := fmt.Sprintf("%s: %d", "Average", Average(ints))
		io.WriteString(os.Stdout, out_avg)
		fmt.Println()

		out_median := fmt.Sprintf("%s: %d", "Median", Median(ints))
		io.WriteString(os.Stdout, out_median)
		fmt.Println()

		out_variance := fmt.Sprintf("%s: %d", "Variance", Variance(ints))
		io.WriteString(os.Stdout, out_variance)
		fmt.Println()

		out_sd := fmt.Sprintf("%s: %d", "Standard deviation", StandardDeviation(ints))
		io.WriteString(os.Stdout, out_sd)
		fmt.Println()
	*/
	average := Average(keys)
	median := Median(keys)
	fmt.Println("Average:", math.Round(average))
	fmt.Println("Median:", math.Round(median))
}

func Average(data []float64) float64 {
	sum := Sum(data)
	return sum / float64(len(data))
}

func Sum(data []float64) (sum float64) {
	for _, n := range data {
		sum += n
	}
	return sum
}

func Median(data []float64) (median float64) {
	c := utils.SortedCopy(data)
	l := len(c)
	if l%2 == 0 {
		median = Average(c[l/2-1 : l/2+1])
	} else {
		median = c[l/2]
	}
	return median
}

func Variance(data []float64)
