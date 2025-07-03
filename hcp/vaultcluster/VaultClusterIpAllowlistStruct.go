// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster


type VaultClusterIpAllowlistStruct struct {
	// IP address range in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_cluster#address VaultCluster#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Description to help identify source (maximum 255 chars).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/vault_cluster#description VaultCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

