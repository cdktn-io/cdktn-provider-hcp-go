// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vaultradarsecretmanagervaultdedicated

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-hcp-go/hcp/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-hcp-go/hcp/v11/vaultradarsecretmanagervaultdedicated/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated hcp_vault_radar_secret_manager_vault_dedicated}.
type VaultRadarSecretManagerVaultDedicated interface {
	cdktn.TerraformResource
	AccessReadWrite() interface{}
	SetAccessReadWrite(val interface{})
	AccessReadWriteInput() interface{}
	ApprolePush() VaultRadarSecretManagerVaultDedicatedApprolePushOutputReference
	ApprolePushInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	Id() *string
	Kubernetes() VaultRadarSecretManagerVaultDedicatedKubernetesOutputReference
	KubernetesInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Token() VaultRadarSecretManagerVaultDedicatedTokenOutputReference
	TokenInput() interface{}
	VaultUrl() *string
	SetVaultUrl(val *string)
	VaultUrlInput() *string
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
	PutApprolePush(value *VaultRadarSecretManagerVaultDedicatedApprolePush)
	PutKubernetes(value *VaultRadarSecretManagerVaultDedicatedKubernetes)
	PutToken(value *VaultRadarSecretManagerVaultDedicatedToken)
	ResetAccessReadWrite()
	ResetApprolePush()
	ResetKubernetes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetToken()
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

// The jsii proxy struct for VaultRadarSecretManagerVaultDedicated
type jsiiProxy_VaultRadarSecretManagerVaultDedicated struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) AccessReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) AccessReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ApprolePush() VaultRadarSecretManagerVaultDedicatedApprolePushOutputReference {
	var returns VaultRadarSecretManagerVaultDedicatedApprolePushOutputReference
	_jsii_.Get(
		j,
		"approlePush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ApprolePushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approlePushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Kubernetes() VaultRadarSecretManagerVaultDedicatedKubernetesOutputReference {
	var returns VaultRadarSecretManagerVaultDedicatedKubernetesOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) KubernetesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) Token() VaultRadarSecretManagerVaultDedicatedTokenOutputReference {
	var returns VaultRadarSecretManagerVaultDedicatedTokenOutputReference
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) TokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) VaultUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated) VaultUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultUrlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated hcp_vault_radar_secret_manager_vault_dedicated} Resource.
func NewVaultRadarSecretManagerVaultDedicated(scope constructs.Construct, id *string, config *VaultRadarSecretManagerVaultDedicatedConfig) VaultRadarSecretManagerVaultDedicated {
	_init_.Initialize()

	if err := validateNewVaultRadarSecretManagerVaultDedicatedParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultRadarSecretManagerVaultDedicated{}

	_jsii_.Create(
		"@cdktn/provider-hcp.vaultRadarSecretManagerVaultDedicated.VaultRadarSecretManagerVaultDedicated",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_radar_secret_manager_vault_dedicated hcp_vault_radar_secret_manager_vault_dedicated} Resource.
func NewVaultRadarSecretManagerVaultDedicated_Override(v VaultRadarSecretManagerVaultDedicated, scope constructs.Construct, id *string, config *VaultRadarSecretManagerVaultDedicatedConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-hcp.vaultRadarSecretManagerVaultDedicated.VaultRadarSecretManagerVaultDedicated",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetAccessReadWrite(val interface{}) {
	if err := j.validateSetAccessReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessReadWrite",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VaultRadarSecretManagerVaultDedicated)SetVaultUrl(val *string) {
	if err := j.validateSetVaultUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vaultUrl",
		val,
	)
}

// Generates CDKTN code for importing a VaultRadarSecretManagerVaultDedicated resource upon running "cdktn plan <stack-name>".
func VaultRadarSecretManagerVaultDedicated_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVaultRadarSecretManagerVaultDedicated_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.vaultRadarSecretManagerVaultDedicated.VaultRadarSecretManagerVaultDedicated",
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
func VaultRadarSecretManagerVaultDedicated_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultRadarSecretManagerVaultDedicated_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.vaultRadarSecretManagerVaultDedicated.VaultRadarSecretManagerVaultDedicated",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultRadarSecretManagerVaultDedicated_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultRadarSecretManagerVaultDedicated_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.vaultRadarSecretManagerVaultDedicated.VaultRadarSecretManagerVaultDedicated",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultRadarSecretManagerVaultDedicated_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultRadarSecretManagerVaultDedicated_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-hcp.vaultRadarSecretManagerVaultDedicated.VaultRadarSecretManagerVaultDedicated",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VaultRadarSecretManagerVaultDedicated_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-hcp.vaultRadarSecretManagerVaultDedicated.VaultRadarSecretManagerVaultDedicated",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) PutApprolePush(value *VaultRadarSecretManagerVaultDedicatedApprolePush) {
	if err := v.validatePutApprolePushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putApprolePush",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) PutKubernetes(value *VaultRadarSecretManagerVaultDedicatedKubernetes) {
	if err := v.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) PutToken(value *VaultRadarSecretManagerVaultDedicatedToken) {
	if err := v.validatePutTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putToken",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ResetAccessReadWrite() {
	_jsii_.InvokeVoid(
		v,
		"resetAccessReadWrite",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ResetApprolePush() {
	_jsii_.InvokeVoid(
		v,
		"resetApprolePush",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ResetKubernetes() {
	_jsii_.InvokeVoid(
		v,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ResetProjectId() {
	_jsii_.InvokeVoid(
		v,
		"resetProjectId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ResetToken() {
	_jsii_.InvokeVoid(
		v,
		"resetToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultRadarSecretManagerVaultDedicated) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

