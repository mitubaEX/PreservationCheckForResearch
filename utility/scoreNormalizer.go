package utility

import (
	"math"
	"strconv"
	"strings"
)

// resultを受け取って，そいつを正規化して結果を返す
func ScoreNormalization(result string) string {
	splitString := strings.Split(result, "\n")
	var maxScore = 0.0

	// MaxScoreを求める
	for _, v := range splitString {
		row := strings.Split(v, ",")
		if len(row) <= 4 || row[1] == "score" {
			continue
		}
		floatVal, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			panic(err)
		}
		maxScore = math.Max(maxScore, floatVal)
	}

	// 正規化した検索結果を取得する
	var newResults = []string{}
	for _, v := range splitString {
		row := strings.SplitN(v, ",", 5)
		if len(row) <= 4 || row[1] == "score" {
			continue
		}

		floatVal, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			panic(err)
		}
		newResults = append(newResults, row[0]+","+
			strconv.FormatFloat(floatVal/maxScore, 'f', -1, 64)+","+
			row[2]+","+row[4])
	}
	return strings.Join(newResults, "\n")
}
