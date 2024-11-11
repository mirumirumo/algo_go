package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mirusort "github.com/mirumirumo/algo_go/sort"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var nums []int
	for _, num := range strings.Split(sc.Text(), " ") {
		num, _ := strconv.Atoi(num)
		nums = append(nums, num)
	}
	mirusort.InsertionSort(nums)
	fmt.Printf("%v\n", nums)
}
