// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsforwarding


type DnsForwardingForwardingRule struct {
	// The domain name for DNS forwarding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/dns_forwarding#domain_name DnsForwarding#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The list of inbound endpoint IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/dns_forwarding#inbound_endpoint_ips DnsForwarding#inbound_endpoint_ips}
	InboundEndpointIps *[]*string `field:"required" json:"inboundEndpointIps" yaml:"inboundEndpointIps"`
	// The ID of the forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/dns_forwarding#rule_id DnsForwarding#rule_id}
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
}

