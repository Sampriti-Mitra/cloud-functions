package cloud_functions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"weekend.side/GCP/cloud_functions/externals"
)

func SimpleSlackFunction(w http.ResponseWriter, r *http.Request) {
	request := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Fprint(w, err)
		log.Print(err)
		return
	}
	defer r.Body.Close()

	w.Header().Set("X-Slack-No-Retry", "1")

	log.Print(request)

	if typeMessage, ok := request["type"]; ok {

		if typeMessage == "event_callback" {
			_, isBot := request["event"].(map[string]interface{})["bot_id"]

			if ok && !isBot {
				externals.SlackWebhook(request["event"])
				_, err := externals.CoinGecko(request["event"])
				log.Print("error:", err)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprint(w, request)
}
