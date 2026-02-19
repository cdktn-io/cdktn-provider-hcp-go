// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datahcpdnsforwardingrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataHcpDnsForwardingRuleConfig struct {
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
	// The ID of the DNS forwarding configuration this rule belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/dns_forwarding_rule#dns_forwarding_id DataHcpDnsForwardingRule#dns_forwarding_id}
	DnsForwardingId *string `field:"required" json:"dnsForwardingId" yaml:"dnsForwardingId"`
	// The ID of the DNS forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/dns_forwarding_rule#dns_forwarding_rule_id DataHcpDnsForwardingRule#dns_forwarding_rule_id}
	DnsForwardingRuleId *string `field:"required" json:"dnsForwardingRuleId" yaml:"dnsForwardingRuleId"`
	// The ID of the HVN that this DNS forwarding rule belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/dns_forwarding_rule#hvn_id DataHcpDnsForwardingRule#hvn_id}
	HvnId *string `field:"required" json:"hvnId" yaml:"hvnId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/dns_forwarding_rule#id DataHcpDnsForwardingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where the DNS forwarding rule is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/dns_forwarding_rule#project_id DataHcpDnsForwardingRule#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

