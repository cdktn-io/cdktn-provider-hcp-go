// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultplugin


type VaultPluginTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs/resources/vault_plugin#default VaultPlugin#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
}

