// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package boundarycluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BoundaryClusterConfig struct {
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
	// The ID of the Boundary cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#cluster_id BoundaryCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The password of the initial admin user.
	//
	// This must be at least 8 characters in length. Note that this may show up in logs, and it will be stored in the state file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#password BoundaryCluster#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The tier that the HCP Boundary cluster will be provisioned as, 'Standard' or 'Plus'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#tier BoundaryCluster#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// The username of the initial admin user.
	//
	// This must be at least 3 characters in length, alphanumeric, hyphen, or period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#username BoundaryCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The time to live for the auth token in golang's time.Duration string format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#auth_token_time_to_live BoundaryCluster#auth_token_time_to_live}
	AuthTokenTimeToLive *string `field:"optional" json:"authTokenTimeToLive" yaml:"authTokenTimeToLive"`
	// The time to stale for the auth token in golang's time.Duration string format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#auth_token_time_to_stale BoundaryCluster#auth_token_time_to_stale}
	AuthTokenTimeToStale *string `field:"optional" json:"authTokenTimeToStale" yaml:"authTokenTimeToStale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#id BoundaryCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_window_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#maintenance_window_config BoundaryCluster#maintenance_window_config}
	MaintenanceWindowConfig *BoundaryClusterMaintenanceWindowConfig `field:"optional" json:"maintenanceWindowConfig" yaml:"maintenanceWindowConfig"`
	// The ID of the HCP project where the Boundary cluster is located.
	//
	// If not specified, the project specified in the HCP Provider config block will be used, if configured.
	// If a project is not configured in the HCP Provider config block, the oldest project in the organization will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#project_id BoundaryCluster#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/boundary_cluster#timeouts BoundaryCluster#timeouts}
	Timeouts *BoundaryClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

