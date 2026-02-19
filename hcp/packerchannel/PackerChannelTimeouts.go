// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package packerchannel


type PackerChannelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_channel#create PackerChannel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_channel#default PackerChannel#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_channel#delete PackerChannel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/packer_channel#update PackerChannel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

