// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package privatelink

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PrivateLinkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the HVN associated with the private link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#hvn_id PrivateLink#hvn_id}
	HvnId *string `field:"required" json:"hvnId" yaml:"hvnId"`
	// The ID of the private link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#private_link_id PrivateLink#private_link_id}
	PrivateLinkId *string `field:"required" json:"privateLinkId" yaml:"privateLinkId"`
	// The ID of the HCP Vault cluster associated with the private link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#vault_cluster_id PrivateLink#vault_cluster_id}
	VaultClusterId *string `field:"required" json:"vaultClusterId" yaml:"vaultClusterId"`
	// The list of consumer accounts allowed to connect to the private link.
	//
	// In AWS, these are IAM Principals. In Azure, these are Azure Subscription/Resource IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#consumer_accounts PrivateLink#consumer_accounts}
	ConsumerAccounts *[]*string `field:"optional" json:"consumerAccounts" yaml:"consumerAccounts"`
	// The list of consumer IP ranges or CIDRs allowed to connect to the HVD cluster associated with the private link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#consumer_ip_ranges PrivateLink#consumer_ip_ranges}
	ConsumerIpRanges *[]*string `field:"optional" json:"consumerIpRanges" yaml:"consumerIpRanges"`
	// The cloud provider regions from which consumers can connect to the private link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#consumer_regions PrivateLink#consumer_regions}
	ConsumerRegions *[]*string `field:"optional" json:"consumerRegions" yaml:"consumerRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#id PrivateLink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where the private link is located.
	//
	// If not specified, the project configured in the provider is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#project_id PrivateLink#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/private_link#timeouts PrivateLink#timeouts}
	Timeouts *PrivateLinkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

