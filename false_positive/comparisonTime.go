package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func EachCOmparisonTime(threshold string) {
	searchResultFiles, err := ioutil.ReadDir("../each_compare_result/" + threshold)
	if err != nil {
		panic(err)
	}
	var searchTime = 0.0
	var compareTime = 0
	for _, file := range searchResultFiles {
		f, i := CheckComparisonTime("../each_compare_result/" + threshold + "/" + file.Name())
		searchTime += f
		compareTime += i
	}
	fmt.Println(searchTime, compareTime)
}

// 検索ファイル，比較ファイルをもらい検索時間，比較時間を返す
func CheckComparisonTime(filename string) (float64, int) {
	// input
	input, err := os.Open("../search_result/" + filename)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		// 検索時間
		if strings.Contains(scanner.Text(), "sec") {
			val, err := strconv.ParseFloat(
				strings.Replace(
					strings.Replace(scanner.Text(), "sec", "", -1),
					" ns", "", -1), 64)
			if err != nil {
				continue
				//panic(err)
			}
			return val, 0
		} else if strings.Contains(scanner.Text(), " ns") { // 比較時間
			val, err := strconv.Atoi(
				strings.Replace(
					strings.Replace(scanner.Text(), "sec", "", -1),
					" ns", "", -1))
			if err != nil {
				continue
			}
			return 0.0, val
		}
	}
	return 0.0, 0
}
