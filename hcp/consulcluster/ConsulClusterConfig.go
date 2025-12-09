// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consulcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConsulClusterConfig struct {
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
	// The ID of the HCP Consul cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#cluster_id ConsulCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The ID of the HVN this HCP Consul cluster is associated to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#hvn_id ConsulCluster#hvn_id}
	HvnId *string `field:"required" json:"hvnId" yaml:"hvnId"`
	// The tier that the HCP Consul cluster will be provisioned as.
	//
	// Only `development`, `standard`, `plus`, and `premium` are available at this time. See [pricing information](https://www.hashicorp.com/products/consul/pricing).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#tier ConsulCluster#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Enables automatic HVN to HVN peering when creating a secondary cluster in a federation.
	//
	// The alternative to using the auto-accept feature is to create an [`hcp_hvn_peering_connection`](hvn_peering_connection.md) resource that explicitly defines the HVN resources that are allowed to communicate with each other.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#auto_hvn_to_hvn_peering ConsulCluster#auto_hvn_to_hvn_peering}
	AutoHvnToHvnPeering interface{} `field:"optional" json:"autoHvnToHvnPeering" yaml:"autoHvnToHvnPeering"`
	// Denotes the Consul connect feature should be enabled for this cluster.  Default to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#connect_enabled ConsulCluster#connect_enabled}
	ConnectEnabled interface{} `field:"optional" json:"connectEnabled" yaml:"connectEnabled"`
	// The Consul data center name of the cluster. If not specified, it is defaulted to the value of `cluster_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#datacenter ConsulCluster#datacenter}
	Datacenter *string `field:"optional" json:"datacenter" yaml:"datacenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#id ConsulCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_allowlist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#ip_allowlist ConsulCluster#ip_allowlist}
	IpAllowlist interface{} `field:"optional" json:"ipAllowlist" yaml:"ipAllowlist"`
	// The minimum Consul patch version of the cluster.
	//
	// Allows only the rightmost version component to increment (E.g: `1.13.0` will allow installation of `1.13.2` and `1.13.3` etc., but not `1.14.0`). If not specified, it is defaulted to the version that is currently recommended by HCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#min_consul_version ConsulCluster#min_consul_version}
	MinConsulVersion *string `field:"optional" json:"minConsulVersion" yaml:"minConsulVersion"`
	// The `self_link` of the HCP Consul cluster which is the primary in the federation setup with this HCP Consul cluster.
	//
	// If not specified, it is a standalone cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#primary_link ConsulCluster#primary_link}
	PrimaryLink *string `field:"optional" json:"primaryLink" yaml:"primaryLink"`
	// The ID of the HCP project where the HCP Consul cluster is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#project_id ConsulCluster#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Denotes that the cluster has a public endpoint for the Consul UI. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#public_endpoint ConsulCluster#public_endpoint}
	PublicEndpoint interface{} `field:"optional" json:"publicEndpoint" yaml:"publicEndpoint"`
	// The t-shirt size representation of each server VM that this Consul cluster is provisioned with.
	//
	// Valid option for development tier - `x_small`. Valid options for other tiers - `small`, `medium`, `large`. For more details - https://cloud.hashicorp.com/pricing/consul. Upgrading the size of a cluster after creation is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#size ConsulCluster#size}
	Size *string `field:"optional" json:"size" yaml:"size"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster#timeouts ConsulCluster#timeouts}
	Timeouts *ConsulClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

