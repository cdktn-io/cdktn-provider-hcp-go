// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster


type VaultClusterMajorVersionUpgradeConfig struct {
	// The major upgrade type for the cluster. Valid options for upgrade type - `AUTOMATIC`, `SCHEDULED`, `MANUAL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_cluster#upgrade_type VaultCluster#upgrade_type}
	UpgradeType *string `field:"required" json:"upgradeType" yaml:"upgradeType"`
	// The maintenance day of the week for scheduled upgrades.
	//
	// Valid options for maintenance window day - `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_cluster#maintenance_window_day VaultCluster#maintenance_window_day}
	MaintenanceWindowDay *string `field:"optional" json:"maintenanceWindowDay" yaml:"maintenanceWindowDay"`
	// The maintenance time frame for scheduled upgrades. Valid options for maintenance window time - `WINDOW_12AM_4AM`, `WINDOW_6AM_10AM`, `WINDOW_12PM_4PM`, `WINDOW_6PM_10PM`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_cluster#maintenance_window_time VaultCluster#maintenance_window_time}
	MaintenanceWindowTime *string `field:"optional" json:"maintenanceWindowTime" yaml:"maintenanceWindowTime"`
}

