package main

import (
	"context"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	ENDPOINT   = "localhost:9000"
	ACCESS_KEY = "HE0BGWT4CTYQ90AQ9A8D"
	SECRET     = "vk7J+Bat6I6NoQi17LNiEe4zOUyv7N+6d7xYJSgD"
)

func main() {
	client, err := minio.New(ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(ACCESS_KEY, SECRET, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalf("Error initial connection: %v \n", err)
	}

	//CreateBucket(client, "test-bucket")
	ListBucket(client)
}

func ListBucket(client *minio.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	buckets, _ := client.ListBuckets(ctx)

	for _, bucket := range buckets {
		log.Println("Bucket Name:", bucket.Name)
	}
}

func CreateBucket(client *minio.Client, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := client.MakeBucket(ctx, name, minio.MakeBucketOptions{
		Region: "us-east-1",
	})

	if err != nil {
		log.Fatalf("Failed to make bucket because error: %v \n", err)
	}

	log.Println("Bucket has been created!")
}
