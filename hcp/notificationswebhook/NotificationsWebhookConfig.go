// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationswebhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationsWebhookConfig struct {
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
	// The webhook configuration used to deliver event payloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/notifications_webhook#config NotificationsWebhook#config}
	Config *NotificationsWebhookConfigA `field:"required" json:"config" yaml:"config"`
	// The webhook's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/notifications_webhook#name NotificationsWebhook#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The webhook's description. Descriptions are useful for helping others understand the purpose of the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/notifications_webhook#description NotificationsWebhook#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates if the webhook should receive payloads for the subscribed events. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/notifications_webhook#enabled NotificationsWebhook#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID of the project to create the webhook under.
	//
	// If unspecified, the webhook will be created in the project the provider is configured with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/notifications_webhook#project_id NotificationsWebhook#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Set of events to subscribe the webhook to all resources or a specific resource in the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/notifications_webhook#subscriptions NotificationsWebhook#subscriptions}
	Subscriptions interface{} `field:"optional" json:"subscriptions" yaml:"subscriptions"`
}

