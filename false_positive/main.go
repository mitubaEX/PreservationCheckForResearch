package main

func main() {
	var outputDirectoryName075 = "../search_result/075"
	var outputDirectoryName05 = "../search_result/05"
	var outputDirectoryName025 = "../search_result/025"
	CheckFalsePositive(outputDirectoryName075, outputDirectoryName05, outputDirectoryName025)

	// 各閾値の比較時間を出す
	for _, v := range []string{"075", "05", "025"} {
		CompareFromSearchResult(v)
		EachCOmparisonTime(v)
	}
}
