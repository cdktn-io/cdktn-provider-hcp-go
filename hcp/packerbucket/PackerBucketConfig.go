// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package packerbucket

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PackerBucketConfig struct {
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
	// The bucket's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_bucket#name PackerBucket#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the project to create the bucket under.
	//
	// If unspecified, the bucket will be created in the project the provider is configured with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_bucket#project_id PackerBucket#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

