// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/vaultcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_cluster hcp_vault_cluster}.
type VaultCluster interface {
	cdktf.TerraformResource
	AuditLogConfig() VaultClusterAuditLogConfigOutputReference
	AuditLogConfigInput() *VaultClusterAuditLogConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudProvider() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HvnId() *string
	SetHvnId(val *string)
	HvnIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAllowlist() VaultClusterIpAllowlistStructList
	IpAllowlistInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MajorVersionUpgradeConfig() VaultClusterMajorVersionUpgradeConfigOutputReference
	MajorVersionUpgradeConfigInput() *VaultClusterMajorVersionUpgradeConfig
	MetricsConfig() VaultClusterMetricsConfigOutputReference
	MetricsConfigInput() *VaultClusterMetricsConfig
	MinVaultVersion() *string
	SetMinVaultVersion(val *string)
	MinVaultVersionInput() *string
	Namespace() *string
	// The tree node.
	Node() constructs.Node
	OrganizationId() *string
	PathsFilter() *[]*string
	SetPathsFilter(val *[]*string)
	PathsFilterInput() *[]*string
	PrimaryLink() *string
	SetPrimaryLink(val *string)
	PrimaryLinkInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProxyEndpoint() *string
	SetProxyEndpoint(val *string)
	ProxyEndpointInput() *string
	PublicEndpoint() interface{}
	SetPublicEndpoint(val interface{})
	PublicEndpointInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SelfLink() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() VaultClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VaultPrivateEndpointUrl() *string
	VaultProxyEndpointUrl() *string
	VaultPublicEndpointUrl() *string
	VaultVersion() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAuditLogConfig(value *VaultClusterAuditLogConfig)
	PutIpAllowlist(value interface{})
	PutMajorVersionUpgradeConfig(value *VaultClusterMajorVersionUpgradeConfig)
	PutMetricsConfig(value *VaultClusterMetricsConfig)
	PutTimeouts(value *VaultClusterTimeouts)
	ResetAuditLogConfig()
	ResetId()
	ResetIpAllowlist()
	ResetMajorVersionUpgradeConfig()
	ResetMetricsConfig()
	ResetMinVaultVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathsFilter()
	ResetPrimaryLink()
	ResetProjectId()
	ResetProxyEndpoint()
	ResetPublicEndpoint()
	ResetTier()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VaultCluster
type jsiiProxy_VaultCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VaultCluster) AuditLogConfig() VaultClusterAuditLogConfigOutputReference {
	var returns VaultClusterAuditLogConfigOutputReference
	_jsii_.Get(
		j,
		"auditLogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) AuditLogConfigInput() *VaultClusterAuditLogConfig {
	var returns *VaultClusterAuditLogConfig
	_jsii_.Get(
		j,
		"auditLogConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) CloudProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) HvnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvnId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) HvnIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvnIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) IpAllowlist() VaultClusterIpAllowlistStructList {
	var returns VaultClusterIpAllowlistStructList
	_jsii_.Get(
		j,
		"ipAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) IpAllowlistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) MajorVersionUpgradeConfig() VaultClusterMajorVersionUpgradeConfigOutputReference {
	var returns VaultClusterMajorVersionUpgradeConfigOutputReference
	_jsii_.Get(
		j,
		"majorVersionUpgradeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) MajorVersionUpgradeConfigInput() *VaultClusterMajorVersionUpgradeConfig {
	var returns *VaultClusterMajorVersionUpgradeConfig
	_jsii_.Get(
		j,
		"majorVersionUpgradeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) MetricsConfig() VaultClusterMetricsConfigOutputReference {
	var returns VaultClusterMetricsConfigOutputReference
	_jsii_.Get(
		j,
		"metricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) MetricsConfigInput() *VaultClusterMetricsConfig {
	var returns *VaultClusterMetricsConfig
	_jsii_.Get(
		j,
		"metricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) MinVaultVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minVaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) MinVaultVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minVaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) PathsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) PathsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) PrimaryLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) PrimaryLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ProxyEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) ProxyEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) PublicEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) PublicEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) Timeouts() VaultClusterTimeoutsOutputReference {
	var returns VaultClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) VaultPrivateEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultPrivateEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) VaultProxyEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultProxyEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) VaultPublicEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultPublicEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultCluster) VaultVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultVersion",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_cluster hcp_vault_cluster} Resource.
func NewVaultCluster(scope constructs.Construct, id *string, config *VaultClusterConfig) VaultCluster {
	_init_.Initialize()

	if err := validateNewVaultClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultCluster{}

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_cluster hcp_vault_cluster} Resource.
func NewVaultCluster_Override(v VaultCluster, scope constructs.Construct, id *string, config *VaultClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VaultCluster)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetHvnId(val *string) {
	if err := j.validateSetHvnIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hvnId",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetMinVaultVersion(val *string) {
	if err := j.validateSetMinVaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minVaultVersion",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetPathsFilter(val *[]*string) {
	if err := j.validateSetPathsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathsFilter",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetPrimaryLink(val *string) {
	if err := j.validateSetPrimaryLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryLink",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetProxyEndpoint(val *string) {
	if err := j.validateSetProxyEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyEndpoint",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetPublicEndpoint(val interface{}) {
	if err := j.validateSetPublicEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicEndpoint",
		val,
	)
}

func (j *jsiiProxy_VaultCluster)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

// Generates CDKTF code for importing a VaultCluster resource upon running "cdktf plan <stack-name>".
func VaultCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVaultCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func VaultCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VaultCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcp.vaultCluster.VaultCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VaultCluster) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VaultCluster) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VaultCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VaultCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VaultCluster) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultCluster) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VaultCluster) PutAuditLogConfig(value *VaultClusterAuditLogConfig) {
	if err := v.validatePutAuditLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAuditLogConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultCluster) PutIpAllowlist(value interface{}) {
	if err := v.validatePutIpAllowlistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpAllowlist",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultCluster) PutMajorVersionUpgradeConfig(value *VaultClusterMajorVersionUpgradeConfig) {
	if err := v.validatePutMajorVersionUpgradeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMajorVersionUpgradeConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultCluster) PutMetricsConfig(value *VaultClusterMetricsConfig) {
	if err := v.validatePutMetricsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMetricsConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultCluster) PutTimeouts(value *VaultClusterTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultCluster) ResetAuditLogConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetAuditLogConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetIpAllowlist() {
	_jsii_.InvokeVoid(
		v,
		"resetIpAllowlist",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetMajorVersionUpgradeConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetMajorVersionUpgradeConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetMetricsConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetMetricsConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetMinVaultVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetMinVaultVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetPathsFilter() {
	_jsii_.InvokeVoid(
		v,
		"resetPathsFilter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetPrimaryLink() {
	_jsii_.InvokeVoid(
		v,
		"resetPrimaryLink",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetProjectId() {
	_jsii_.InvokeVoid(
		v,
		"resetProjectId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetProxyEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetProxyEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetPublicEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetPublicEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetTier() {
	_jsii_.InvokeVoid(
		v,
		"resetTier",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

