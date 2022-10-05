package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chmenegatti/get-nsxt-profile-id/configs"
)

type SubnetsField struct {
	GatwayAddress string `json:"gateway_address"`
	Network       string `json:"network"`
}

type Segment struct {
	Id           string         `json:"id,omitempty"`
	DisplayName  string         `json:"display_name"`
	Path         string         `json:"path"`
	ResourceType string         `json:"resource_type"`
	Subnets      []SubnetsField `json:"subnets"`
}

type Segments struct {
	Results     []Segment `json:"results,omitempty"`
	ResultCount int64     `json:"result_count,omitempty"`
}

func GetSegmentIds(edge string) (Segments, error) {

	env := fmt.Sprintf("%s_URL", edge)

	url := configs.GetEnvKeys(env) + "/policy/api/v1/infra/segments"

	res, err := RequestNSXTApi(url, edge)

	if err != nil {
		log.Fatal(err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	var responseObject Segments
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject, err
}
