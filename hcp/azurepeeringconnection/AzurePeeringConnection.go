// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azurepeeringconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/azurepeeringconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/azure_peering_connection hcp_azure_peering_connection}.
type AzurePeeringConnection interface {
	cdktf.TerraformResource
	AllowForwardedTraffic() interface{}
	SetAllowForwardedTraffic(val interface{})
	AllowForwardedTrafficInput() interface{}
	ApplicationId() *string
	AzurePeeringId() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	ExpiresAt() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HvnLink() *string
	SetHvnLink(val *string)
	HvnLinkInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OrganizationId() *string
	PeeringId() *string
	SetPeeringId(val *string)
	PeeringIdInput() *string
	PeerResourceGroupName() *string
	SetPeerResourceGroupName(val *string)
	PeerResourceGroupNameInput() *string
	PeerSubscriptionId() *string
	SetPeerSubscriptionId(val *string)
	PeerSubscriptionIdInput() *string
	PeerTenantId() *string
	SetPeerTenantId(val *string)
	PeerTenantIdInput() *string
	PeerVnetName() *string
	SetPeerVnetName(val *string)
	PeerVnetNameInput() *string
	PeerVnetRegion() *string
	SetPeerVnetRegion(val *string)
	PeerVnetRegionInput() *string
	ProjectId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SelfLink() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AzurePeeringConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	UseRemoteGateways() interface{}
	SetUseRemoteGateways(val interface{})
	UseRemoteGatewaysInput() interface{}
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
	PutTimeouts(value *AzurePeeringConnectionTimeouts)
	ResetAllowForwardedTraffic()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetUseRemoteGateways()
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

// The jsii proxy struct for AzurePeeringConnection
type jsiiProxy_AzurePeeringConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AzurePeeringConnection) AllowForwardedTraffic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForwardedTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) AllowForwardedTrafficInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowForwardedTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) AzurePeeringId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azurePeeringId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) ExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) HvnLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvnLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) HvnLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hvnLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeeringId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeeringIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerSubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSubscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerSubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSubscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerVnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerVnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerVnetRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVnetRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) PeerVnetRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVnetRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) Timeouts() AzurePeeringConnectionTimeoutsOutputReference {
	var returns AzurePeeringConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) UseRemoteGateways() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRemoteGateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurePeeringConnection) UseRemoteGatewaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useRemoteGatewaysInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/azure_peering_connection hcp_azure_peering_connection} Resource.
func NewAzurePeeringConnection(scope constructs.Construct, id *string, config *AzurePeeringConnectionConfig) AzurePeeringConnection {
	_init_.Initialize()

	if err := validateNewAzurePeeringConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzurePeeringConnection{}

	_jsii_.Create(
		"@cdktf/provider-hcp.azurePeeringConnection.AzurePeeringConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.108.0/docs/resources/azure_peering_connection hcp_azure_peering_connection} Resource.
func NewAzurePeeringConnection_Override(a AzurePeeringConnection, scope constructs.Construct, id *string, config *AzurePeeringConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.azurePeeringConnection.AzurePeeringConnection",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetAllowForwardedTraffic(val interface{}) {
	if err := j.validateSetAllowForwardedTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowForwardedTraffic",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetHvnLink(val *string) {
	if err := j.validateSetHvnLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hvnLink",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetPeeringId(val *string) {
	if err := j.validateSetPeeringIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peeringId",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetPeerResourceGroupName(val *string) {
	if err := j.validateSetPeerResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetPeerSubscriptionId(val *string) {
	if err := j.validateSetPeerSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerSubscriptionId",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetPeerTenantId(val *string) {
	if err := j.validateSetPeerTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerTenantId",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetPeerVnetName(val *string) {
	if err := j.validateSetPeerVnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVnetName",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetPeerVnetRegion(val *string) {
	if err := j.validateSetPeerVnetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVnetRegion",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AzurePeeringConnection)SetUseRemoteGateways(val interface{}) {
	if err := j.validateSetUseRemoteGatewaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useRemoteGateways",
		val,
	)
}

// Generates CDKTF code for importing a AzurePeeringConnection resource upon running "cdktf plan <stack-name>".
func AzurePeeringConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAzurePeeringConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.azurePeeringConnection.AzurePeeringConnection",
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
func AzurePeeringConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurePeeringConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.azurePeeringConnection.AzurePeeringConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzurePeeringConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurePeeringConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.azurePeeringConnection.AzurePeeringConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzurePeeringConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurePeeringConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.azurePeeringConnection.AzurePeeringConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzurePeeringConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcp.azurePeeringConnection.AzurePeeringConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) PutTimeouts(value *AzurePeeringConnectionTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AzurePeeringConnection) ResetAllowForwardedTraffic() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowForwardedTraffic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurePeeringConnection) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurePeeringConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurePeeringConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurePeeringConnection) ResetUseRemoteGateways() {
	_jsii_.InvokeVoid(
		a,
		"resetUseRemoteGateways",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurePeeringConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurePeeringConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

