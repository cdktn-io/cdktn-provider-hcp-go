// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logstreamingdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogStreamingDestinationConfig struct {
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
	// The HCP Log Streaming Destination’s name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/log_streaming_destination#name LogStreamingDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/log_streaming_destination#cloudwatch LogStreamingDestination#cloudwatch}.
	Cloudwatch *LogStreamingDestinationCloudwatch `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/log_streaming_destination#datadog LogStreamingDestination#datadog}.
	Datadog *LogStreamingDestinationDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/log_streaming_destination#splunk_cloud LogStreamingDestination#splunk_cloud}.
	SplunkCloud *LogStreamingDestinationSplunkCloud `field:"optional" json:"splunkCloud" yaml:"splunkCloud"`
}

