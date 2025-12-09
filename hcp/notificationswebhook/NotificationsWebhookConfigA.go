// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationswebhook


type NotificationsWebhookConfigA struct {
	// The HTTP or HTTPS destination URL that HCP delivers the event payloads to.
	//
	// The destination must be able to use the HCP webhook
	// [payload](https://developer.hashicorp.com/hcp/docs/hcp/admin/projects/webhooks#webhook-payload).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/notifications_webhook#url NotificationsWebhook#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The arbitrary secret that HCP uses to sign all its webhook requests.
	//
	// This is a write-only field, it is written once and not visible thereafter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/notifications_webhook#hmac_key NotificationsWebhook#hmac_key}
	HmacKey *string `field:"optional" json:"hmacKey" yaml:"hmacKey"`
}

