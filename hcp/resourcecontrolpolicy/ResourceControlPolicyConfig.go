// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourcecontrolpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResourceControlPolicyConfig struct {
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
	// The list of constraint IDs to enable for the organization.
	//
	// Each constraint ID must be a recognized constraint returned by ListConstraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.115.0/docs/resources/resource_control_policy#enabled_constraints ResourceControlPolicy#enabled_constraints}
	EnabledConstraints *[]*string `field:"required" json:"enabledConstraints" yaml:"enabledConstraints"`
	// The ID of the organization to manage the resource control policy for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.115.0/docs/resources/resource_control_policy#organization_id ResourceControlPolicy#organization_id}
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
}

