package models

import "time"

type Networks struct {
	Id                      string    `json:"id"`
	Organization            string    `json:"organization"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	Address                 string    `json:"address"`
	SegmentID               string    `json:"segment_id"`
	SwitchID                string    `json:"switch_id"`
	GroupID                 string    `json:"group_id"`
	ProfileID               string    `json:"profile_id"`
	DisplayName             string    `json:"display_name"`
	EnableSideCommunication bool      `json:"enable_side_communication"`
	Status                  string    `json:"status"`
	Created                 time.Time `json:"created"`
	UpdatedAt               time.Time `json:"updated_at"`
	DeletedAt               time.Time `json:"deleted_at"`
}
