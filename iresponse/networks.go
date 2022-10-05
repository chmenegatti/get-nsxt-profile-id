package iresponse

import "github.com/chmenegatti/get-nsxt-profile-id/models"

type Networks struct {
	Id                      string `json:"id"`
	Organization            string `json:"organization"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Address                 string `json:"address"`
	SegmentID               string `json:"segment_id"`
	SwitchID                string `json:"switch_id"`
	GroupID                 string `json:"group_id"`
	ProfileID               string `json:"profile_id"`
	DisplayName             string `json:"display_name"`
	EnableSideCommunication bool   `json:"enable_side_communication"`
	Status                  string `json:"status"`
}

func CreateNetworkResponse(network models.Networks) Networks {
	return Networks{Id: network.Id, Organization: network.Organization, Name: network.Name, Description: network.Description,
		Address: network.Address, SegmentID: network.SegmentID, SwitchID: network.SwitchID, GroupID: network.GroupID,
		ProfileID: network.ProfileID, DisplayName: network.DisplayName, EnableSideCommunication: network.EnableSideCommunication,
		Status: network.Status,
	}
}
