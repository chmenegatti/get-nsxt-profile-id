package models

import "time"

type Organizations struct {
	Id                      string    `json:"id,omitempty"`
	Name                    string    `json:"name,omitempty"`
	Description             string    `json:"description,omitempty"`
	TierProvider            string    `json:"tier_provider,omitempty"`
	EdgeCluster             string    `json:"edge_cluster,omitempty"`
	TierProviderID          string    `json:"tier_provider_id"`
	Tier1GatewayID          string    `json:"tier1_gateway_id"`
	PolicyID                string    `json:"policy_id"`
	LocaleServiceID         string    `json:"locale_service_id"`
	BackupCluster           string    `json:"backup_cluster,omitempty"`
	LoadBalanceID           string    `json:"load_balance_id"`
	PhysicalFirewall        string    `json:"physical_firewall"`
	VirtualFirewall         string    `json:"virtual_firewall"`
	FirewallExternalAddress string    `json:"firewall_external_address"`
	LoadBalanceSize         string    `json:"load_balance_size"`
	Status                  string    `json:"status"`
	Created                 time.Time `json:"created"`
	UpdatedAt               time.Time `json:"updated_at"`
	DeletedAt               time.Time `json:"deleted_at"`
}
