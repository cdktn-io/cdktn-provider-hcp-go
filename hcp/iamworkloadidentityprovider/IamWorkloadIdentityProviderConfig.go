// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentityprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IamWorkloadIdentityProviderConfig struct {
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
	// conditional_access is a hashicorp/go-bexpr string that is evaluated when exchanging tokens.
	//
	// It restricts which upstream identities are allowed to access the service principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/iam_workload_identity_provider#conditional_access IamWorkloadIdentityProvider#conditional_access}
	ConditionalAccess *string `field:"required" json:"conditionalAccess" yaml:"conditionalAccess"`
	// The workload identity provider's name. Ideally, this should be descriptive of the workload being federated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/iam_workload_identity_provider#name IamWorkloadIdentityProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The service principal's resource name for which the workload identity provider will be created for.
	//
	// Only service principals created within a project are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/iam_workload_identity_provider#service_principal IamWorkloadIdentityProvider#service_principal}
	ServicePrincipal *string `field:"required" json:"servicePrincipal" yaml:"servicePrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/iam_workload_identity_provider#aws IamWorkloadIdentityProvider#aws}.
	Aws *IamWorkloadIdentityProviderAws `field:"optional" json:"aws" yaml:"aws"`
	// A description for the workload identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/iam_workload_identity_provider#description IamWorkloadIdentityProvider#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/iam_workload_identity_provider#oidc IamWorkloadIdentityProvider#oidc}.
	Oidc *IamWorkloadIdentityProviderOidc `field:"optional" json:"oidc" yaml:"oidc"`
}

