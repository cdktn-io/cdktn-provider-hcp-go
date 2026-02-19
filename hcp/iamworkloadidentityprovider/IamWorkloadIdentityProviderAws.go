// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentityprovider


type IamWorkloadIdentityProviderAws struct {
	// The AWS Account ID that is allowed to exchange workload identities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/iam_workload_identity_provider#account_id IamWorkloadIdentityProvider#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

