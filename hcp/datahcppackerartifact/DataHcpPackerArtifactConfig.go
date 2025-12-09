// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datahcppackerartifact

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataHcpPackerArtifactConfig struct {
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
	// The name of the HCP Packer Bucket where the Artifact is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/packer_artifact#bucket_name DataHcpPackerArtifact#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Name of the platform where the HCP Packer Artifact is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/packer_artifact#platform DataHcpPackerArtifact#platform}
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The Region where the HCP Packer Artifact is stored, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/packer_artifact#region DataHcpPackerArtifact#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The name of the HCP Packer Channel the Version containing this Artifact is assigned to.
	//
	// The Version currently assigned to the Channel will be fetched.
	// Exactly one of `channel_name` or `version_fingerprint` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/packer_artifact#channel_name DataHcpPackerArtifact#channel_name}
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Name of the Packer builder that built this Artifact. Ex: `amazon-ebs.example`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/packer_artifact#component_type DataHcpPackerArtifact#component_type}
	ComponentType *string `field:"optional" json:"componentType" yaml:"componentType"`
	// The ID of the HCP Organization where the Artifact is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/packer_artifact#project_id DataHcpPackerArtifact#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The fingerprint of the HCP Packer Version where the Artifact is located.
	//
	// If provided in the config, it is used to fetch the Version.
	// Exactly one of `channel_name` or `version_fingerprint` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/data-sources/packer_artifact#version_fingerprint DataHcpPackerArtifact#version_fingerprint}
	VersionFingerprint *string `field:"optional" json:"versionFingerprint" yaml:"versionFingerprint"`
}

