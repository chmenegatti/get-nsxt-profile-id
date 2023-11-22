package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chmenegatti/get-nsxt-profile-id/configs"
)

type Switch struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type Switches struct {
	Results     []Switch `json:"results,omitempty"`
	ResultCount int64    `json:"result_count,omitempty"`
}

func GetSwitchIds(edge string) (Switches, error) {

	env := fmt.Sprintf("%s_URL", edge)

	url := configs.GetEnvKeys(env) + "/api/v1/logical-switches"

	res, err := RequestNSXTApi(url, edge)

	if err != nil {
		log.Println(err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	var responseObject Switches
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject, err
}
