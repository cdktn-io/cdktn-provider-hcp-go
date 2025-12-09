// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcpiampolicy


type DataHcpIamPolicyBindings struct {
	// The set of principals to bind to the given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/iam_policy#principals DataHcpIamPolicy#principals}
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
	// The role name to bind to the given principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/iam_policy#role DataHcpIamPolicy#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

