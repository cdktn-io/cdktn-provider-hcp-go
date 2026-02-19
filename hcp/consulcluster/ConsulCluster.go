// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package consulcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-hcp-go/hcp/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-hcp-go/hcp/v11/consulcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster hcp_consul_cluster}.
type ConsulCluster interface {
	cdktn.TerraformResource
	AutoHvnToHvnPeering() interface{}
	SetAutoHvnToHvnPeering(val interface{})
	AutoHvnToHvnPeeringInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CloudProvider() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ConnectEnabled() interface{}
	SetConnectEnabled(val interface{})
	ConnectEnabledInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsulAutomaticUpgrades() cdktn.IResolvable
	ConsulCaFile() *string
	ConsulConfigFile() *string
	ConsulPrivateEndpointUrl() *string
	ConsulPublicEndpointUrl() *string
	ConsulRootTokenAccessorId() *string
	ConsulRootTokenSecretId() *string
	ConsulSnapshotInterval() *string
	ConsulSnapshotRetention() *string
	ConsulVersion() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Datacenter() *string
	SetDatacenter(val *string)
	DatacenterInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
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
	IpAllowlist() ConsulClusterIpAllowlistStructList
	IpAllowlistInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MinConsulVersion() *string
	SetMinConsulVersion(val *string)
	MinConsulVersionInput() *string
	// The tree node.
	Node() constructs.Node
	OrganizationId() *string
	PrimaryLink() *string
	SetPrimaryLink(val *string)
	PrimaryLinkInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicEndpoint() interface{}
	SetPublicEndpoint(val interface{})
	PublicEndpointInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	Scale() *float64
	SelfLink() *string
	Size() *string
	SetSize(val *string)
	SizeInput() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() ConsulClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
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
	PutIpAllowlist(value interface{})
	PutTimeouts(value *ConsulClusterTimeouts)
	ResetAutoHvnToHvnPeering()
	ResetConnectEnabled()
	ResetDatacenter()
	ResetId()
	ResetIpAllowlist()
	ResetMinConsulVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryLink()
	ResetProjectId()
	ResetPublicEndpoint()
	ResetSize()
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

// The jsii proxy struct for ConsulCluster
type jsiiProxy_ConsulCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ConsulCluster) AutoHvnToHvnPeering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHvnToHvnPeering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) AutoHvnToHvnPeeringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoHvnToHvnPeeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) CloudProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConnectEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConnectEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulAutomaticUpgrades() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"consulAutomaticUpgrades",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulCaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulCaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulConfigFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulConfigFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulPrivateEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulPrivateEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulPublicEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulPublicEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulRootTokenAccessorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulRootTokenAccessorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulRootTokenSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulRootTokenSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulSnapshotInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulSnapshotInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulSnapshotRetention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulSnapshotRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ConsulVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consulVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Datacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) DatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) HvnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvnId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) HvnIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvnIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) IpAllowlist() ConsulClusterIpAllowlistStructList {
	var returns ConsulClusterIpAllowlistStructList
	_jsii_.Get(
		j,
		"ipAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) IpAllowlistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) MinConsulVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minConsulVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) MinConsulVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minConsulVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) PrimaryLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) PrimaryLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) PublicEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) PublicEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Scale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) SizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) Timeouts() ConsulClusterTimeoutsOutputReference {
	var returns ConsulClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster hcp_consul_cluster} Resource.
func NewConsulCluster(scope constructs.Construct, id *string, config *ConsulClusterConfig) ConsulCluster {
	_init_.Initialize()

	if err := validateNewConsulClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConsulCluster{}

	_jsii_.Create(
		"@cdktn/provider-hcp.consulCluster.ConsulCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/consul_cluster hcp_consul_cluster} Resource.
func NewConsulCluster_Override(c ConsulCluster, scope constructs.Construct, id *string, config *ConsulClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-hcp.consulCluster.ConsulCluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConsulCluster)SetAutoHvnToHvnPeering(val interface{}) {
	if err := j.validateSetAutoHvnToHvnPeeringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoHvnToHvnPeering",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetConnectEnabled(val interface{}) {
	if err := j.validateSetConnectEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectEnabled",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetDatacenter(val *string) {
	if err := j.validateSetDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datacenter",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetHvnId(val *string) {
	if err := j.validateSetHvnIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hvnId",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetMinConsulVersion(val *string) {
	if err := j.validateSetMinConsulVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minConsulVersion",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetPrimaryLink(val *string) {
	if err := j.validateSetPrimaryLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryLink",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetPublicEndpoint(val interface{}) {
	if err := j.validateSetPublicEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicEndpoint",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetSize(val *string) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_ConsulCluster)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

// Generates CDKTN code for importing a ConsulCluster resource upon running "cdktn plan <stack-name>".
func ConsulCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateConsulCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.consulCluster.ConsulCluster",
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
func ConsulCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.consulCluster.ConsulCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsulCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.consulCluster.ConsulCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsulCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.consulCluster.ConsulCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConsulCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-hcp.consulCluster.ConsulCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConsulCluster) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ConsulCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConsulCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ConsulCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConsulCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ConsulCluster) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConsulCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConsulCluster) PutIpAllowlist(value interface{}) {
	if err := c.validatePutIpAllowlistParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIpAllowlist",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConsulCluster) PutTimeouts(value *ConsulClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConsulCluster) ResetAutoHvnToHvnPeering() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoHvnToHvnPeering",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetConnectEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetDatacenter() {
	_jsii_.InvokeVoid(
		c,
		"resetDatacenter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetIpAllowlist() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAllowlist",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetMinConsulVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMinConsulVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetPrimaryLink() {
	_jsii_.InvokeVoid(
		c,
		"resetPrimaryLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetProjectId() {
	_jsii_.InvokeVoid(
		c,
		"resetProjectId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetPublicEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetPublicEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetSize() {
	_jsii_.InvokeVoid(
		c,
		"resetSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

