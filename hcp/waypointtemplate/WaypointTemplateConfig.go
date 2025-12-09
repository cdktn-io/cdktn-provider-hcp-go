// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaypointTemplateConfig struct {
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
	// The name of the Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#name WaypointTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A brief description of the template, up to 110 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#summary WaypointTemplate#summary}
	Summary *string `field:"required" json:"summary" yaml:"summary"`
	// The ID of the Terraform no-code module to use for running Terraform operations.
	//
	// This is in the format of 'nocode-<ID>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#terraform_no_code_module_id WaypointTemplate#terraform_no_code_module_id}
	TerraformNoCodeModuleId *string `field:"required" json:"terraformNoCodeModuleId" yaml:"terraformNoCodeModuleId"`
	// Terraform Cloud No-Code Module details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#terraform_no_code_module_source WaypointTemplate#terraform_no_code_module_source}
	TerraformNoCodeModuleSource *string `field:"required" json:"terraformNoCodeModuleSource" yaml:"terraformNoCodeModuleSource"`
	// The ID of the Terraform Cloud Project to create workspaces in.
	//
	// The ID is found on the Terraform Cloud Project settings page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#terraform_project_id WaypointTemplate#terraform_project_id}
	TerraformProjectId *string `field:"required" json:"terraformProjectId" yaml:"terraformProjectId"`
	// List of actions by 'ID' to assign to this Template.
	//
	// Applications created from this Template will have these actions assigned to them. Only 'ID' is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#actions WaypointTemplate#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// A description of the template, along with when and why it should be used, up to 500 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#description WaypointTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of labels attached to this Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#labels WaypointTemplate#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the HCP project where the Waypoint Template is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#project_id WaypointTemplate#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Instructions for using the template (markdown format supported).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#readme_markdown_template WaypointTemplate#readme_markdown_template}
	ReadmeMarkdownTemplate *string `field:"optional" json:"readmeMarkdownTemplate" yaml:"readmeMarkdownTemplate"`
	// The ID of the agent pool to use for Terraform operations, for workspaces created for applications using this template.
	//
	// Required if terraform_execution_mode is set to 'agent'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#terraform_agent_pool_id WaypointTemplate#terraform_agent_pool_id}
	TerraformAgentPoolId *string `field:"optional" json:"terraformAgentPoolId" yaml:"terraformAgentPoolId"`
	// Terraform Cloud Workspace details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#terraform_cloud_workspace_details WaypointTemplate#terraform_cloud_workspace_details}
	TerraformCloudWorkspaceDetails *WaypointTemplateTerraformCloudWorkspaceDetails `field:"optional" json:"terraformCloudWorkspaceDetails" yaml:"terraformCloudWorkspaceDetails"`
	// The execution mode of the HCP Terraform workspaces created for applications using this template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#terraform_execution_mode WaypointTemplate#terraform_execution_mode}
	TerraformExecutionMode *string `field:"optional" json:"terraformExecutionMode" yaml:"terraformExecutionMode"`
	// If true, will auto-import the readme form the Terraform module used.
	//
	// If this is set to true, users should not also set `readme_markdown_template`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#use_module_readme WaypointTemplate#use_module_readme}
	UseModuleReadme interface{} `field:"optional" json:"useModuleReadme" yaml:"useModuleReadme"`
	// List of variable options for the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template#variable_options WaypointTemplate#variable_options}
	VariableOptions interface{} `field:"optional" json:"variableOptions" yaml:"variableOptions"`
}

