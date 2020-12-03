package main

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"os"

	"github.com/ajbosco/segment-config-go/segment"
	// "reflect"
)

func main() {
	client := segment.NewClient("access-token", "workspace")

	jsonFile, err := os.Open("segment/trackingplan.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("successfully opened json file")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	tp, err := client.CreateTrackingPlan(result)

	fmt.Println(tp)

}
