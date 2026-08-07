// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datahcppackerartifact

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataHcpPackerArtifactConfig struct {
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
	// The name of the HCP Packer Bucket where the Artifact is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#bucket_name DataHcpPackerArtifact#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Name of the platform where the HCP Packer Artifact is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#platform DataHcpPackerArtifact#platform}
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The Region where the HCP Packer Artifact is stored, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#region DataHcpPackerArtifact#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The name of the HCP Packer Channel the Version containing this Artifact is assigned to.
	//
	// The Version currently assigned to the Channel will be fetched.
	// Exactly one of `channel_name` or `version_fingerprint` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#channel_name DataHcpPackerArtifact#channel_name}
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Name of the Packer builder that built this Artifact. Ex: `amazon-ebs.example`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#component_type DataHcpPackerArtifact#component_type}
	ComponentType *string `field:"optional" json:"componentType" yaml:"componentType"`
	// When set (non-empty) with `channel_name`, the data source uses **GetImageByBuildLabels**: the channel's current version is scanned for builds whose labels contain every key/value you supply.
	//
	// If more than one build matches, the HCP Packer API returns the **first** candidate in **`updated_at` descending** order that also matches the **`platform`** and **`region`** you configure (sent as cloud provider and region). Always set `platform` and `region` so the result is predictable when multiple builds or clouds exist.
	//
	// If several builds still match the same platform and region (for example different Packer sources), use **`component_type`** to disambiguate, consistent with the non-label lookup path.
	//
	// When `labels` is unset or empty, this attribute is computed from the resolved build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#labels DataHcpPackerArtifact#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The ID of the HCP Organization where the Artifact is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#project_id DataHcpPackerArtifact#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The fingerprint of the HCP Packer Version where the Artifact is located.
	//
	// If provided in the config, it is used to fetch the Version.
	// Exactly one of `channel_name` or `version_fingerprint` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.113.0/docs/data-sources/packer_artifact#version_fingerprint DataHcpPackerArtifact#version_fingerprint}
	VersionFingerprint *string `field:"optional" json:"versionFingerprint" yaml:"versionFingerprint"`
}

