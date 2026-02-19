// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package consulsnapshot


type ConsulSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_snapshot#create ConsulSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_snapshot#default ConsulSnapshot#default}.
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_snapshot#delete ConsulSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_snapshot#update ConsulSnapshot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

