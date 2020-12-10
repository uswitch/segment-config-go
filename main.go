package main

import (
	"encoding/json"

	"fmt"

	"github.com/ajbosco/segment-config-go/segment"
	// "reflect"
)

func main() {
	workspace := "workspace"
	accessToken := "access-token"
	client, err := segment.NewClient(&accessToken, &workspace)
	if err != nil {
		fmt.Println(err)
	}

	tp, err := client.GetTrackingPlan("rs_123abc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tp)

	// jsonFile, err := os.Open("segment/trackingplan.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	tp, err := client.CreateTrackingPlan(result)

	fmt.Println(tp)

	// fmt.Println(tp)
}
