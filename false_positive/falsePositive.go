package main

import (
	"../utility"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// 構造体を受け取り，false_positiveを調べる
func CheckFalsePositive(outputDirectoryName075 string, outputDirectoryName05 string, outputDirectoryName025 string) {
	var searchTime = 0.0
	var compareTime = 0

	utility.Mkdir(outputDirectoryName075)
	utility.Mkdir(outputDirectoryName05)
	utility.Mkdir(outputDirectoryName025)

	searchResultFiles, err := ioutil.ReadDir("../search_result")
	if err != nil {
		panic(err)
	}

	compareResultFiles, err := ioutil.ReadDir("../compare_result")
	if err != nil {
		panic(err)
	}

	var index = 0
	for _, file := range searchResultFiles {
		if file.IsDir() {
			continue
		}
		// 検索時間を求める
		for _, target := range []string{"../search_result/" + file.Name(), "../compare_result/" + compareResultFiles[index].Name()} {
			f, i := CheckComparisonTime(target)
			searchTime += f
			compareTime += i
		}

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

		output075, err := os.Create(outputDirectoryName075 + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer output075.Close()
		output05, err := os.Create(outputDirectoryName05 + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer output075.Close()
		output025, err := os.Create(outputDirectoryName025 + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer output075.Close()

		// each count
		var count075, count05, count025 = 0, 0, 0
		var compareCount075, compareCount05, compareCount025 = 0, 0, 0
		scanner.Scan()

		output075.WriteString(scanner.Text() + "\n")
		output05.WriteString(scanner.Text() + "\n")
		output025.WriteString(scanner.Text() + "\n")

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
				output075.WriteString(scanner.Text() + "\n")
				if compareScore >= 0.75 {
					compareCount075++
				}
			}
			if searchScore >= 0.5 {
				count05++
				output05.WriteString(scanner.Text() + "\n")
				if compareScore >= 0.75 {
					compareCount05++
				}
			}
			if searchScore >= 0.25 {
				count025++
				output025.WriteString(scanner.Text() + "\n")
				if compareScore >= 0.75 {
					compareCount025++
				}
			}
		}
		fmt.Printf("%d,%d,%d,%d,%d,%d\n",
			count075, count05, count025, compareCount075, compareCount05, compareCount025)
		index++
	}
	// 全体の比較時間
	fmt.Printf("%fsec,%dns\n", searchTime, compareTime)
}
