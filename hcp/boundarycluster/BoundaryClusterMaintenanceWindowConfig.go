// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package boundarycluster


type BoundaryClusterMaintenanceWindowConfig struct {
	// The maintenance day of the week for scheduled upgrades.
	//
	// Valid options for maintenance window day - `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`, `SUNDAY`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/boundary_cluster#day BoundaryCluster#day}
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The end time which upgrades can be performed.
	//
	// Uses 24H clock and must be in UTC time zone. Valid options include - 1 to 24 (inclusive)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/boundary_cluster#end BoundaryCluster#end}
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// The start time which upgrades can be performed.
	//
	// Uses 24H clock and must be in UTC time zone. Valid options include - 0 to 23 (inclusive)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/boundary_cluster#start BoundaryCluster#start}
	Start *float64 `field:"optional" json:"start" yaml:"start"`
	// The upgrade type for the cluster. Valid options for upgrade type - `AUTOMATIC`, `SCHEDULED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/boundary_cluster#upgrade_type BoundaryCluster#upgrade_type}
	UpgradeType *string `field:"optional" json:"upgradeType" yaml:"upgradeType"`
}

