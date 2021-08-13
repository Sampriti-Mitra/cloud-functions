package cloud_functions

import (
	//"cloud.google.com/go"
	dialogflowcx "cloud.google.com/go/dialogflow/cx/apiv3"
	"context"
	"google.golang.org/api/option/internaloption"

	//"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
	"log"
	"net/http"
)

func SimpleBotFunction(w http.ResponseWriter, r *http.Request) {
	ProjectId := PROJECT_ID // project id

	agent := AGENT_NAME

	ctx := context.Background()

	sa := option.WithCredentialsFile("./serverless_function_source_code/dialogflowcx.json")

	detectIntentReq := cx.DetectIntentRequest{
		Session: agent + "/sessions/" + ProjectId,
		QueryInput: &cx.QueryInput{
			Input: &cx.QueryInput_Text{
				&cx.TextInput{
					Text: "Hi",
				},
			},
			LanguageCode: "en",
		},
	}

	opts := []option.ClientOption{
		internaloption.WithDefaultEndpoint("us-central1-dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://us-central1-dialogflow.googleapis.com/"),
	}

	dialogFlowClient, err := dialogflowcx.NewSessionsClient(ctx, append(opts, sa)...)

	if err != nil {
		log.Print(err)
	}

	resp, err := dialogFlowClient.DetectIntent(ctx, &detectIntentReq)

	if err != nil {
		log.Print(err)
	}

	queryResult := resp.GetQueryResult()

	responseMessages := queryResult.GetResponseMessages()

	log.Print(responseMessages)

}
