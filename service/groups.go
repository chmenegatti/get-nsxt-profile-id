package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chmenegatti/get-nsxt-profile-id/configs"
)

type Group struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"display_name"`
	Path        string `json:"path"`
}

type Groups struct {
	Results     []Group `json:"results,omitempty"`
	ResultCount int64   `json:"result_count,omitempty"`
}

func GetGroupsIds(edge string) (Groups, error) {
	env := fmt.Sprintf("%s_URL", edge)

	url := configs.GetEnvKeys(env) + "/policy/api/v1/infra/domains/default/groups"

	res, err := RequestNSXTApi(url, edge)

	if err != nil {
		log.Fatal(err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	var responseObject Groups
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject, err
}
