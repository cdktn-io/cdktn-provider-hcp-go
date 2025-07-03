// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logstreamingdestination


type LogStreamingDestinationDatadog struct {
	// The value for the DD-API-KEY to send when making requests to DataDog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/log_streaming_destination#api_key LogStreamingDestination#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The Datadog endpoint to send logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/log_streaming_destination#endpoint LogStreamingDestination#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The value for the DD-APPLICATION-KEY to send when making requests to DataDog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/log_streaming_destination#application_key LogStreamingDestination#application_key}
	ApplicationKey *string `field:"optional" json:"applicationKey" yaml:"applicationKey"`
}

