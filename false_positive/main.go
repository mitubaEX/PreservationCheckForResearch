package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// 検索結果と比較結果を読み取りzipで回し，類似度を比較する
func main() {
	searchResultFiles, err := ioutil.ReadDir("../search_result")
	if err != nil {
		panic(err)
	}

	compareResultFiles, err := ioutil.ReadDir("../compare_result")
	if err != nil {
		panic(err)
	}
	for index, file := range searchResultFiles {
		// input
		input, err := os.Open("../search_result/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer input.Close()

		scanner := bufio.NewScanner(input)

		compareInput, err := os.Open("../compare_result/" + compareResultFiles[index].Name())
		if err != nil {
			panic(err)
		}
		defer compareInput.Close()

		compareScanner := bufio.NewScanner(compareInput)

		// each count
		var count075, count05, count025 = 0, 0, 0
		var compareCount075, compareCount05, compareCount025 = 0, 0, 0
		scanner.Scan()
		// 検索バースマークと比較バースマークを取得する
		for scanner.Scan() {
			compareScanner.Scan()

			searchSplitString := strings.Split(scanner.Text(), ",")
			compareSplitString := strings.Split(scanner.Text(), ",")
			if len(searchSplitString) <= 3 || len(compareSplitString) <= 3 {
				continue
			}

			searchScore, err := strconv.ParseFloat(strings.Split(scanner.Text(), ",")[1], 64)
			compareScore, err := strconv.ParseFloat(strings.Split(compareScanner.Text(), ",")[2], 64)
			if err != nil {
				panic(err)
			}
			if searchScore >= 0.75 {
				count075++
				if compareScore >= 0.75 {
					compareCount075++
				}
			}
			if searchScore >= 0.5 {
				count05++
				if compareScore >= 0.75 {
					compareCount05++
				}
			}
			if searchScore >= 0.25 {
				count025++
				if compareScore >= 0.75 {
					compareCount025++
				}
			}
		}
		fmt.Printf("%d,%d,%d,%d,%d,%d\n",
			count075, count05, count025, compareCount075, compareCount05, compareCount025)
	}
}
