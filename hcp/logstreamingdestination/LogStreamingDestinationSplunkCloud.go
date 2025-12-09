// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logstreamingdestination


type LogStreamingDestinationSplunkCloud struct {
	// The Splunk Cloud endpoint to send logs to. Streaming to free trial instances is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/log_streaming_destination#endpoint LogStreamingDestination#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The authentication token that will be used by the platform to access Splunk Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/log_streaming_destination#token LogStreamingDestination#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

