// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationswebhook


type NotificationsWebhookSubscriptions struct {
	// The information about the events of a webhook subscription.
	//
	// The service that owns the resource is responsible for maintaining events. Refer to the service's webhook documentation for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/notifications_webhook#events NotificationsWebhook#events}
	Events interface{} `field:"required" json:"events" yaml:"events"`
	// Refers to the resource the webhook is subscribed to.
	//
	// If not set, the webhook subscribes to the emitted events listed in events for any resource in the webhook's project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/notifications_webhook#resource_id NotificationsWebhook#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

