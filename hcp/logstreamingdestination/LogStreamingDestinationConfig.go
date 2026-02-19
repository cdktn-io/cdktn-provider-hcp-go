// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logstreamingdestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogStreamingDestinationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The HCP Log Streaming Destination’s name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/log_streaming_destination#name LogStreamingDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/log_streaming_destination#cloudwatch LogStreamingDestination#cloudwatch}.
	Cloudwatch *LogStreamingDestinationCloudwatch `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/log_streaming_destination#datadog LogStreamingDestination#datadog}.
	Datadog *LogStreamingDestinationDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/log_streaming_destination#splunk_cloud LogStreamingDestination#splunk_cloud}.
	SplunkCloud *LogStreamingDestinationSplunkCloud `field:"optional" json:"splunkCloud" yaml:"splunkCloud"`
}

