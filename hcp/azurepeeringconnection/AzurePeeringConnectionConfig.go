// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azurepeeringconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AzurePeeringConnectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The `self_link` of the HashiCorp Virtual Network (HVN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#hvn_link AzurePeeringConnection#hvn_link}
	HvnLink *string `field:"required" json:"hvnLink" yaml:"hvnLink"`
	// The ID of the peering connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#peering_id AzurePeeringConnection#peering_id}
	PeeringId *string `field:"required" json:"peeringId" yaml:"peeringId"`
	// The resource group name of the peer VNet in Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#peer_resource_group_name AzurePeeringConnection#peer_resource_group_name}
	PeerResourceGroupName *string `field:"required" json:"peerResourceGroupName" yaml:"peerResourceGroupName"`
	// The subscription ID of the peer VNet in Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#peer_subscription_id AzurePeeringConnection#peer_subscription_id}
	PeerSubscriptionId *string `field:"required" json:"peerSubscriptionId" yaml:"peerSubscriptionId"`
	// The tenant ID of the peer VNet in Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#peer_tenant_id AzurePeeringConnection#peer_tenant_id}
	PeerTenantId *string `field:"required" json:"peerTenantId" yaml:"peerTenantId"`
	// The name of the peer VNet in Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#peer_vnet_name AzurePeeringConnection#peer_vnet_name}
	PeerVnetName *string `field:"required" json:"peerVnetName" yaml:"peerVnetName"`
	// The region of the peer VNet in Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#peer_vnet_region AzurePeeringConnection#peer_vnet_region}
	PeerVnetRegion *string `field:"required" json:"peerVnetRegion" yaml:"peerVnetRegion"`
	// Whether the forwarded traffic originating from the peered VNet is allowed in the HVN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#allow_forwarded_traffic AzurePeeringConnection#allow_forwarded_traffic}
	AllowForwardedTraffic interface{} `field:"optional" json:"allowForwardedTraffic" yaml:"allowForwardedTraffic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#id AzurePeeringConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#timeouts AzurePeeringConnection#timeouts}
	Timeouts *AzurePeeringConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// If the HVN should use the gateway of the peered VNet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/azure_peering_connection#use_remote_gateways AzurePeeringConnection#use_remote_gateways}
	UseRemoteGateways interface{} `field:"optional" json:"useRemoteGateways" yaml:"useRemoteGateways"`
}

