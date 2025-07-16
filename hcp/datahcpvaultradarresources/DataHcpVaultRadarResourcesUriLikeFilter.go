// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpvaultradarresources


type DataHcpVaultRadarResourcesUriLikeFilter struct {
	// URI like filters to apply radar resources. Each entry in the list will act like an or condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/vault_radar_resources#values DataHcpVaultRadarResources#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// If true, the uri like filter will be case insensitive. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/vault_radar_resources#case_insensitive DataHcpVaultRadarResources#case_insensitive}
	CaseInsensitive interface{} `field:"optional" json:"caseInsensitive" yaml:"caseInsensitive"`
}

