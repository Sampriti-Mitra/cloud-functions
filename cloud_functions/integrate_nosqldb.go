package cloud_functions

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"net/http"
	service "weekend.side/GCP/cloud_functions/services"
)

func CloudFunctionWithFirestore(w http.ResponseWriter, r *http.Request) {
	ProjectId := "" // project id

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, ProjectId)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	request := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	service.AddDataToCollectionInFirestore(ctx, client, request)

	service.WriteDataFromCollectionInFirestore(ctx, client, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
