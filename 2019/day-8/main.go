package main

import (
	"fmt"
	"strings"
)

const width = 25
const height = 6
const combine = width*height

func main() {
	err := run()
	if err != nil {
	    fmt.Printf("err: %v", err)
	    return
	}
}

func run() error {
	raw, err := ReadFile("input.txt", ReadModeSingleCollapsed)
	if err != nil {
	  return err
	}

	data := raw[0]
	temp := data[:combine]
	i := 0
	min := strings.Count(temp, "0")
	minl := 0
	for temp != "" {
		if strings.Count(temp, "0") < min {
			minl = i
			min = strings.Count(temp, "0")
		}
		i++
		if len(data) >= combine*(i+1) {
			temp = data[combine*i:combine*(i+1)]
		} else {
			temp = ""
		}
	}

	slice := data[combine*minl:combine*(minl+1)]
	out := strings.Count(slice, "1") * strings.Count(slice, "2")
	fmt.Printf("Part 1: %v\n", out)

	var out2 []string
	for i := 0; i < combine; i++ {
		s := string(data[i])
		l:=1
		for s == "2" {
			s = string(data[combine*l+i])
			l++
		}
		out2 = append(out2, s)
	}

	fmt.Println("Part 2:")
	for i := 0; i < len(out2)/width; i++ {
		fmt.Println(strings.Replace(fmt.Sprintf("%v", out2[i*width:(i+1)*width]), "0", " ", -1))
	}

	return nil
}

