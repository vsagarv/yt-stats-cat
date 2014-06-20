package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	youtube "code.google.com/p/google-api-go-client/youtube/v3"
)

func init() {
	registerDemo("youtube", youtube.YoutubeReadonlyScope, youtubeMetricsMain)
}

func youtubeMetricsMain(client *http.Client, argv []string) {

	ytService, _ := youtube.New(client)

	resp, err := ytService.Videos.List("statistics").Id("RU_GBcdxT84").Do()
	if err != nil {
		log.Fatalf("List failed: %v", err)
	}

	js, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("Video stats for RU_GBcdxT84:\n%s\n", js);

	// fmt.Printf("Video stats for RU_GBcdxT84: %u\n", resp.Items[0].Statistics)

	return
}
