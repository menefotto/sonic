package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		total float64
		nums  float64
	)

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {

		splitted := strings.Split(scan.Text(), " ")
		if len(splitted) > 4 {
			percent := splitted[3:4]
			num, err := strconv.ParseFloat(percent[0][:3], 64)
			if err != nil {
				return
			}

			total += num
			nums++
		}
	}

	fmt.Printf("Total project test coverage is %.2f%s\n", total/nums, "%")
}
