// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package notificationswebhook


type NotificationsWebhookSubscriptionsEvents struct {
	// The list of event actions subscribed for the resource type set as the [source](#source).
	//
	// For example, `["create", "update"]`. When the action is '*', it means that the webhook is subscribed to all event actions for the event source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/notifications_webhook#actions NotificationsWebhook#actions}
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The resource type of the source of the event.
	//
	// For example, `hashicorp.packer.version`. Event source might not be the same type as the resource that the webhook is subscribed to ([resource_id](#resource_id)) if the event is from a descendant resource. For example, webhooks are subscribed to a `hashicorp.packer.registry` and receive events for descendent resources such as a `hashicorp.packer.version`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/notifications_webhook#source NotificationsWebhook#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

