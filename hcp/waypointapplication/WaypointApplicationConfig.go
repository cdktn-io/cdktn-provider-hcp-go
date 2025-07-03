// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WaypointApplicationConfig struct {
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
	// The name of the Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_application#name WaypointApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the Template this Application is based on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_application#template_id WaypointApplication#template_id}
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
	// List of actions by 'ID' to assign to this Template.
	//
	// Applications created from this Template will have these actions assigned to them. Only 'ID' is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_application#actions WaypointApplication#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Input variables set for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_application#application_input_variables WaypointApplication#application_input_variables}
	ApplicationInputVariables interface{} `field:"optional" json:"applicationInputVariables" yaml:"applicationInputVariables"`
	// The ID of the HCP project where the Waypoint Application is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_application#project_id WaypointApplication#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Instructions for using the Application (markdown format supported).
	//
	// Note: this is a base64 encoded string, and can only be set in configuration after initial creation. The initial version of the README is generated from the README Template from source Template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/waypoint_application#readme_markdown WaypointApplication#readme_markdown}
	ReadmeMarkdown *string `field:"optional" json:"readmeMarkdown" yaml:"readmeMarkdown"`
}

