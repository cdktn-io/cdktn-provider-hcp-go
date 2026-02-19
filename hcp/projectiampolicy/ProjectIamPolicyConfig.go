// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package projectiampolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ProjectIamPolicyConfig struct {
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
	// The policy to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/project_iam_policy#policy_data ProjectIamPolicy#policy_data}
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// The ID of the HCP project to apply the IAM Policy to.
	//
	// If unspecified, the project configured on the provider is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/project_iam_policy#project_id ProjectIamPolicy#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

