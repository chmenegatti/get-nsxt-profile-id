package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chmenegatti/get-nsxt-profile-id/configs"
)

type ProfileId struct {
	MacDiscoveryProfilePath string `json:"mac_discovery_profile_path,omitempty"`
	IpDiscoveryProfilePath  string `json:"ip_discovery_profile_path,omitempty"`
	ResourceType            string `json:"resource_type,omitempty"`
	Id                      string `json:"id,omitempty"`
	DisplayName             string `json:"display_name,omitempty"`
	Path                    string `json:"path,omitempty"`
	RelativePath            string `json:"relative_path,omitempty"`
	ParentPath              string `json:"parent_path,omitempty"`
	MarkedForDelete         bool   `json:"marked_for_delete,omitempty"`
	CreateUser              string `json:"_create_user,omitempty"`
	CreateTime              int64  `json:"_create_time,omitempty"`
	LastModifiedTime        int64  `json:"_last_modified_time,omitempty"`
	LastModifiedUser        string `json:"_last_modified_user,omitempty"`
	SystemOwned             bool   `json:"_system_owned,omitempty"`
	Protection              string `json:"protection,omitempty"`
	Revision                int64  `json:"_revision,omitempty"`
}

type ProfileIdList struct {
	Results     []ProfileId `json:"results,omitempty"`
	ResultCount int64       `json:"result_count,omitempty"`
}

func GetProfileIds(nsxt, edge string) (ProfileIdList, error) {

	env := fmt.Sprintf("%s_URL", edge)

	url := configs.GetEnvKeys(env) +
		"/policy/api/v1/infra/segments/" +
		nsxt + "/segment-discovery-profile-binding-maps"

	res, err := RequestNSXTApi(url, edge)

	if err != nil {
		log.Fatal(err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer res.Body.Close()

	var responseObject ProfileIdList
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject, err
}
