// could run through this again a little later using loops but wanted to learn regex

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Advent!!")
	cd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := filepath.Dir(cd)
	fmt.Println(p)

	fp := filepath.Join(p, "input.txt")
	fmt.Println(fp)

	file, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	r, _ := regexp.Compile(`\d`)
	cumVal := 0
	for scanner.Scan() {
		match := r.FindAllString(scanner.Text(), -1)
		fmt.Println(match)
		calVal, err := strconv.Atoi(match[0] + match[len(match)-1]) // assumes I can duplicate values where len of match == 1
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cumVal += calVal
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading text file", err)
	}
	fmt.Println("Calibration values", cumVal)
}
