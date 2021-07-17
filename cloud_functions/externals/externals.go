package externals

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func ExecuteRequest(urls string, method string, requests interface{}, headers map[string]string) ([]byte, error) {

	reqJSON, err := json.Marshal(requests)

	req, _ := http.NewRequest(method, urls, strings.NewReader(string(reqJSON)))

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	cli := &http.Client{Transport: http.DefaultTransport}

	response, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}

func SlackWebhook(req interface{}) ([]byte, error) {
	slackUrl := "https://slack.com/api/chat.postMessage"

	token := Token

	bearer := "Bearer " + token

	headers := map[string]string{
		"Authorization": bearer,
		"Content-Type":  "application/json",
	}

	request := req.(map[string]interface{})

	msg := request["text"]
	channel := request["channel"]
	ts := request["ts"]
	user := request["user"]

	return ExecuteRequest(slackUrl, "POST", map[string]interface{}{
		"text":      "<@" + user.(string) + ">, what do you mean " + msg.(string) + "? :rage:",
		"channel":   channel.(string),
		"thread_ts": ts.(string),
		"username":  "AngryBots",
		"icon_url":  "https://image.flaticon.com/icons/png/512/528/528076.png",
	}, headers)
}

func CoinGecko(req interface{}) ([]byte, error) {
	// for slack request
	slackUrl := "https://slack.com/api/chat.postMessage"
	request := req.(map[string]interface{})
	msg := request["text"].(string)
	channel := request["channel"]
	ts := request["ts"]
	user := request["user"]

	allCrytpoCurrencies := SliceAContainsAnyStringInSliceB(GetCryptoList(), msg)

	allCurrencies := SliceAContainsAnyStringInSliceB(GetCurrencyList(), msg)

	coinGeckoUrl := "https://api.coingecko.com/api/v3/simple/price"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	params := url.Values{}
	idsParam := strings.Join(allCrytpoCurrencies[:], ",")
	vsCurrenciesParam := strings.Join(allCurrencies[:], ",")

	params.Add("ids", idsParam)
	params.Add("vs_currencies", vsCurrenciesParam)

	coinGeckoUrl = fmt.Sprintf("%s?%s",coinGeckoUrl, params.Encode())

	log.Print("coinGeckoUrl: ", coinGeckoUrl)

	//body, err:=ExecuteRequest(coinGeckoUrl,"GET", nil, headers)
	//if err != nil {
	//	return nil, err
	//}

	response, err:= http.Get(coinGeckoUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var respMap interface{}
	err = json.Unmarshal(body, &respMap)
	if err != nil {
		return nil, err
	}

	log.Print("respBody: ", respMap, "resp",  string(body), "err: ", err)

	var respString string

	respString+= string(body)

	//for _, crypto := range allCrytpoCurrencies {
	//	if cryptoInCurrenciesMap,ok := respMap[crypto]; ok{
	//		respString = respString + "For " + crypto + ": \n"
	//		for currency, val := range cryptoInCurrenciesMap {
	//			respString = respString + currency + ": " + fmt.Sprintf("%v", val) + "\n"
	//		}
	//	}
	//}

	log.Print("resp string: ", respString)

	token := Token

	bearer := "Bearer " + token

	headers = map[string]string{
		"Authorization": bearer,
		"Content-Type":  "application/json",
	}

	return ExecuteRequest(slackUrl, "POST", map[string]interface{}{
		"text":      "<@" + user.(string) + ">, here's the value for today. \n" + respString,
		"channel":   channel.(string),
		"thread_ts": ts.(string),
		"username":  "AngryBots",
		"icon_url":  "https://image.flaticon.com/icons/png/512/528/528076.png",
	}, headers)

}