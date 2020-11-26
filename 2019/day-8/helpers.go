package main

import (
	"io/ioutil"
	"strings"
)

/*
 * Default: ReadModeSingle (Single is stored in the 0th element)
 * ReadModeSingle returns a single string, including all \n
 * ReadModeSingleCollapsed returns a single string, with all \n stripped away
 * ReadModeMultiline returns an array of string, split on \n (each line)
 */
const (
	ReadModeSingle = iota
	ReadModeSingleCollapsed
	ReadModeMultiline
)
func ReadFile(file string, method... int) ([]string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if method == nil {
		method = []int { ReadModeSingle }
	}

	switch method[0] {
	case ReadModeSingle:
		return []string { string(data) }, nil
	case ReadModeSingleCollapsed:
		return []string { strings.Replace(string(data), "\n", "", -1) }, nil
	case ReadModeMultiline:
		return strings.Split(string(data), "\n"), nil
	}

	return []string { string(data) }, nil
}