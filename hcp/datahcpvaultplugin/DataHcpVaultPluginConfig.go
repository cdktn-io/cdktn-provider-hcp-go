// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datahcpvaultplugin

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataHcpVaultPluginConfig struct {
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
	// The ID of the HCP Vault cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_plugin#cluster_id DataHcpVaultPlugin#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The name of the plugin - Valid options for plugin name - 'venafi-pki-backend'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_plugin#plugin_name DataHcpVaultPlugin#plugin_name}
	PluginName *string `field:"required" json:"pluginName" yaml:"pluginName"`
	// The type of the plugin - Valid options for plugin type - 'SECRET', 'AUTH', 'DATABASE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_plugin#plugin_type DataHcpVaultPlugin#plugin_type}
	PluginType *string `field:"required" json:"pluginType" yaml:"pluginType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_plugin#id DataHcpVaultPlugin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the HCP project where the HCP Vault cluster is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_plugin#project_id DataHcpVaultPlugin#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/vault_plugin#timeouts DataHcpVaultPlugin#timeouts}
	Timeouts *DataHcpVaultPluginTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

