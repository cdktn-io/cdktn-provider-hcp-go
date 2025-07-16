// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpawstransitgatewayattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpAwsTransitGatewayAttachmentConfig struct {
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
	// The ID of the HashiCorp Virtual Network (HVN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/aws_transit_gateway_attachment#hvn_id DataHcpAwsTransitGatewayAttachment#hvn_id}
	HvnId *string `field:"required" json:"hvnId" yaml:"hvnId"`
	// The user-settable name of the transit gateway attachment in HCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/aws_transit_gateway_attachment#transit_gateway_attachment_id DataHcpAwsTransitGatewayAttachment#transit_gateway_attachment_id}
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/aws_transit_gateway_attachment#id DataHcpAwsTransitGatewayAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where the transit gateway attachment is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/aws_transit_gateway_attachment#project_id DataHcpAwsTransitGatewayAttachment#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/aws_transit_gateway_attachment#timeouts DataHcpAwsTransitGatewayAttachment#timeouts}
	Timeouts *DataHcpAwsTransitGatewayAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// If `true`, Terraform will wait for the transit gateway attachment to reach an `ACTIVE` state before continuing. Default `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/data-sources/aws_transit_gateway_attachment#wait_for_active_state DataHcpAwsTransitGatewayAttachment#wait_for_active_state}
	WaitForActiveState interface{} `field:"optional" json:"waitForActiveState" yaml:"waitForActiveState"`
}

