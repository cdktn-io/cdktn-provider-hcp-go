// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dnsforwarding

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DnsForwardingConfig struct {
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
	// The connection type for DNS forwarding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#connection_type DnsForwarding#connection_type}
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// The ID of the DNS forwarding configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#dns_forwarding_id DnsForwarding#dns_forwarding_id}
	DnsForwardingId *string `field:"required" json:"dnsForwardingId" yaml:"dnsForwardingId"`
	// forwarding_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#forwarding_rule DnsForwarding#forwarding_rule}
	ForwardingRule *DnsForwardingForwardingRule `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// The ID of the HVN that this DNS forwarding belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#hvn_id DnsForwarding#hvn_id}
	HvnId *string `field:"required" json:"hvnId" yaml:"hvnId"`
	// The ID of the peering connection for DNS forwarding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#peering_id DnsForwarding#peering_id}
	PeeringId *string `field:"required" json:"peeringId" yaml:"peeringId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#id DnsForwarding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where the DNS forwarding is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#project_id DnsForwarding#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding#timeouts DnsForwarding#timeouts}
	Timeouts *DnsForwardingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

