// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awstransitgatewayattachment


type AwsTransitGatewayAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/aws_transit_gateway_attachment#create AwsTransitGatewayAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/aws_transit_gateway_attachment#default AwsTransitGatewayAttachment#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/aws_transit_gateway_attachment#delete AwsTransitGatewayAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

