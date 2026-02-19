// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hvnroute


type HvnRouteAzureConfig struct {
	// The type of Azure hop the packet should be sent to.
	//
	// Valid options for Next Hop Type - `VIRTUAL_APPLIANCE` or `VIRTUAL_NETWORK_GATEWAY`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/hvn_route#next_hop_type HvnRoute#next_hop_type}
	NextHopType *string `field:"required" json:"nextHopType" yaml:"nextHopType"`
	// Contains the IP address packets should be forwarded to.
	//
	// Next hop values are only allowed in routes where the next hop type is VIRTUAL_APPLIANCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/hvn_route#next_hop_ip_address HvnRoute#next_hop_ip_address}
	NextHopIpAddress *string `field:"optional" json:"nextHopIpAddress" yaml:"nextHopIpAddress"`
}

