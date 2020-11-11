package main

import (
	"fmt"
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
	data := make(map[string]string)
	d, err := ReadFile("input.txt", ReadModeMultiline)
	if err != nil {
	    return err
	}

	for _, s := range d {
		p := strings.Replace(s, "\r", "", -1)
		if len(p) < 3 {
			continue
		}

		n := strings.Index(p, ")")

		a := p[:n]
		b := p[n+1:]

		data[b] = a
	}

	total := 0
	for _, parent := range data {
		for parent != "COM" {
			parent = data[parent]
			total++
		}
		total++
	}
	
	fmt.Printf("Part1: %v\n", total)

	var sans []string
	var yous []string
	parent := data["YOU"]
	for parent != "COM" {
		yous = append(yous, parent)
		parent = data[parent]
	}
	parent = data["SAN"]
	for parent != "COM" {
		sans = append(sans, parent)
		parent = data[parent]
	}


	var n1 []string
	var n2 []string
	if len(sans) > len(yous) {
		n1 = sans
		n2 = yous
	} else {
		n1 = yous
		n2 = sans
	}

	inte := ""
	di := len(n1) - len(n2)
	for i := len(n1) - 1; i >= 0; i-- {
		if n1[i] != n2[i-di] {
			inte = n1[i+1]
			break
		}
	}

	jumps1 := 0
	jumps2 := 0
	dat := "YOU"
	for data[dat] != inte {
		dat = data[dat]
		jumps1++
	}
	dat = "SAN"
	for data[dat] != inte {
		dat = data[dat]
		jumps2++
	}

	fmt.Printf("Part2: %v\n", jumps1+jumps2)


	return nil
}

func insert(a string, b string) {

}


