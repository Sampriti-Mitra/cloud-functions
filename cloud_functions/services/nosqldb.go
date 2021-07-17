package services

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
)

func AddDataToCollectionInFirestore(ctx context.Context, client *firestore.Client, request map[string]interface{}) {
	// create table, and add data
	var err error
	for key, val := range request {
		_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
			key: val,
		})
	}
	if &err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func WriteDataFromCollectionInFirestore(ctx context.Context, client *firestore.Client, w http.ResponseWriter) {

	var err error
	// retrieve the table
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Fprintln(w, doc.Data())
	}
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

