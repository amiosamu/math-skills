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
	data := make([]float64, 0)
	for _, val := range res {
		key, _ := strconv.Atoi(val)
		data = append(data, float64(key))
	}

	fmt.Println("Average:", math.Round(Average(data)))
	fmt.Println("Median:", math.Round(Median(data)))
	fmt.Println("Variance:", int(math.Round((Variance(data)))))
	fmt.Println("Standard Deviation:", math.Round(StandardDeviation(data)))
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

func StandardDeviation(data []float64) (stddev float64) {
	return StandardDeviationPopulation(data)
}

func StandardDeviationPopulation(data []float64) (stddev float64) {
	vp := PopulationVariance(data)
	return math.Sqrt(vp)
}

func PopulationVariance(data []float64) (pvar float64) {
	v := _variance(data, 0)
	return v
}

func _variance(data []float64, sample int) (variance float64) {
	m := Average(data)
	for _, n := range data {
		variance += (n - m) * (n - m)
	}
	return variance / float64((len(data) - (1 * sample)))
}

func Variance(data []float64) (stddev float64) {
	return PopulationVariance(data)
}
