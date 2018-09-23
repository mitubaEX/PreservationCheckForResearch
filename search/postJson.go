package search

import (
	"bytes"
	"github.com/mitubaEX/ToolForResearch/settings"
	"io/ioutil"
	"log"
	"net/http"
)

func postJson(argument settings.Argument, encodedString string) (string, error) {
	jsonStr := `{query:'encode_data:` + encodedString + `'}`

	// log
	log.Println("http://" + argument.Host + ":" + argument.Port + "/solr/" + argument.Birthmark + "/query?" +
		"fl=filename,score,place,birthmark,data&rows=" + argument.Rows + "&sort=score%20desc&wt=csv")
	log.Println(jsonStr)

	// request設定
	req, err := http.NewRequest(
		"POST",
		"http://"+argument.Host+":"+argument.Port+"/solr/"+argument.Birthmark+"/query?"+
			"fl=filename,score,place,birthmark,data&rows="+argument.Rows+"&sort=score%20desc&wt=csv",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// post execute
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 検索結果を文字列を取得する
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return bodyString, nil
	}
	err = nil
	return "", err
}
