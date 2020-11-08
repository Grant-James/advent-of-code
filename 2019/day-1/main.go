package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	err := run()
	if err != nil {
	    fmt.Printf("err: %v", err)
	    return
	}
}

func run() error {
	data, err := ReadFile("input.txt", ReadModeMultiline)
	if err != nil {
	    return err
	}

	var total int
	for _, datum := range data {
		datum2 := strings.Replace(datum, "\r", "", -1)

		n, err := strconv.Atoi(datum2)
		if err != nil {
			continue
		}

		total += fuelNeeded(n)
	}
	
	fmt.Printf("Part1: %v\n", total)

	var total2 int
	for _, datum := range data {
		datum2 := strings.Replace(datum, "\r", "", -1)

		mass, err := strconv.Atoi(datum2)
		if err != nil {
			continue
		}

		for mass > 0 {
			fuel := fuelNeeded(mass)
			mass = fuel
			if mass > 0 {
				total2 += mass
			}
		}
	}

	fmt.Printf("Part2: %v\n", total2)

	return nil
}

func fuelNeeded(mass int) int {
	return mass/3-2
}
