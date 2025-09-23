// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultradarintegrationslacksubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VaultRadarIntegrationSlackSubscriptionConfig struct {
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
	// Name of the Slack channel that messages should be sent to.
	//
	// Note that HashiCorp Vault Radar will send a test message to verify the channel. Example: dev-ops-team
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_radar_integration_slack_subscription#channel VaultRadarIntegrationSlackSubscription#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// id of the integration slack connection to use for the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_radar_integration_slack_subscription#connection_id VaultRadarIntegrationSlackSubscription#connection_id}
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Name of subscription. Name must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_radar_integration_slack_subscription#name VaultRadarIntegrationSlackSubscription#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the HCP project where Vault Radar is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_radar_integration_slack_subscription#project_id VaultRadarIntegrationSlackSubscription#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

