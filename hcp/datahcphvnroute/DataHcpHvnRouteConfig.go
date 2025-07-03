// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcphvnroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpHvnRouteConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/data-sources/hvn_route#hvn_link DataHcpHvnRoute#hvn_link}
	HvnLink *string `field:"required" json:"hvnLink" yaml:"hvnLink"`
	// The ID of the HVN route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/data-sources/hvn_route#hvn_route_id DataHcpHvnRoute#hvn_route_id}
	HvnRouteId *string `field:"required" json:"hvnRouteId" yaml:"hvnRouteId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/data-sources/hvn_route#id DataHcpHvnRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where the HVN route is located.
	//
	// Always matches the project ID in `hvn_link`. Setting this attribute is deprecated, but it will remain usable in read-only form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/data-sources/hvn_route#project_id DataHcpHvnRoute#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/data-sources/hvn_route#timeouts DataHcpHvnRoute#timeouts}
	Timeouts *DataHcpHvnRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

