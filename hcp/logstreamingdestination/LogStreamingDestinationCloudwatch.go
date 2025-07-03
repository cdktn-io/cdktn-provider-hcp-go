// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logstreamingdestination


type LogStreamingDestinationCloudwatch struct {
	// The external_id to provide when assuming the aws IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/log_streaming_destination#external_id LogStreamingDestination#external_id}
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// The region the CloudWatch destination is set up to stream to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/log_streaming_destination#region LogStreamingDestination#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The role_arn that will be assumed to stream logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/log_streaming_destination#role_arn LogStreamingDestination#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The log_group_name of the CloudWatch destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/log_streaming_destination#log_group_name LogStreamingDestination#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

