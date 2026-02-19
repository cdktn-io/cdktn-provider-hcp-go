// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dnsforwardingrule


type DnsForwardingRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding_rule#create DnsForwardingRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding_rule#default DnsForwardingRule#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding_rule#delete DnsForwardingRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/dns_forwarding_rule#read DnsForwardingRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

