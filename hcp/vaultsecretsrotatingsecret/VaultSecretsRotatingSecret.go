// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsrotatingsecret

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/vaultsecretsrotatingsecret/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_rotating_secret hcp_vault_secrets_rotating_secret}.
type VaultSecretsRotatingSecret interface {
	cdktf.TerraformResource
	AppName() *string
	SetAppName(val *string)
	AppNameInput() *string
	AwsAccessKeys() VaultSecretsRotatingSecretAwsAccessKeysOutputReference
	AwsAccessKeysInput() interface{}
	AzureApplicationPassword() VaultSecretsRotatingSecretAzureApplicationPasswordOutputReference
	AzureApplicationPasswordInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfluentServiceAccount() VaultSecretsRotatingSecretConfluentServiceAccountOutputReference
	ConfluentServiceAccountInput() interface{}
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
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpServiceAccountKey() VaultSecretsRotatingSecretGcpServiceAccountKeyOutputReference
	GcpServiceAccountKeyInput() interface{}
	IntegrationName() *string
	SetIntegrationName(val *string)
	IntegrationNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongodbAtlasUser() VaultSecretsRotatingSecretMongodbAtlasUserOutputReference
	MongodbAtlasUserInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OrganizationId() *string
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
	// Experimental.
	RawOverrides() interface{}
	RotationPolicyName() *string
	SetRotationPolicyName(val *string)
	RotationPolicyNameInput() *string
	SecretProvider() *string
	SetSecretProvider(val *string)
	SecretProviderInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TwilioApiKey() VaultSecretsRotatingSecretTwilioApiKeyOutputReference
	TwilioApiKeyInput() interface{}
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
	PutAwsAccessKeys(value *VaultSecretsRotatingSecretAwsAccessKeys)
	PutAzureApplicationPassword(value *VaultSecretsRotatingSecretAzureApplicationPassword)
	PutConfluentServiceAccount(value *VaultSecretsRotatingSecretConfluentServiceAccount)
	PutGcpServiceAccountKey(value *VaultSecretsRotatingSecretGcpServiceAccountKey)
	PutMongodbAtlasUser(value *VaultSecretsRotatingSecretMongodbAtlasUser)
	PutTwilioApiKey(value *VaultSecretsRotatingSecretTwilioApiKey)
	ResetAwsAccessKeys()
	ResetAzureApplicationPassword()
	ResetConfluentServiceAccount()
	ResetGcpServiceAccountKey()
	ResetMongodbAtlasUser()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetTwilioApiKey()
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

// The jsii proxy struct for VaultSecretsRotatingSecret
type jsiiProxy_VaultSecretsRotatingSecret struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) AppNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) AwsAccessKeys() VaultSecretsRotatingSecretAwsAccessKeysOutputReference {
	var returns VaultSecretsRotatingSecretAwsAccessKeysOutputReference
	_jsii_.Get(
		j,
		"awsAccessKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) AwsAccessKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccessKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) AzureApplicationPassword() VaultSecretsRotatingSecretAzureApplicationPasswordOutputReference {
	var returns VaultSecretsRotatingSecretAzureApplicationPasswordOutputReference
	_jsii_.Get(
		j,
		"azureApplicationPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) AzureApplicationPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureApplicationPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) ConfluentServiceAccount() VaultSecretsRotatingSecretConfluentServiceAccountOutputReference {
	var returns VaultSecretsRotatingSecretConfluentServiceAccountOutputReference
	_jsii_.Get(
		j,
		"confluentServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) ConfluentServiceAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluentServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) GcpServiceAccountKey() VaultSecretsRotatingSecretGcpServiceAccountKeyOutputReference {
	var returns VaultSecretsRotatingSecretGcpServiceAccountKeyOutputReference
	_jsii_.Get(
		j,
		"gcpServiceAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) GcpServiceAccountKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpServiceAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) IntegrationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) IntegrationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) MongodbAtlasUser() VaultSecretsRotatingSecretMongodbAtlasUserOutputReference {
	var returns VaultSecretsRotatingSecretMongodbAtlasUserOutputReference
	_jsii_.Get(
		j,
		"mongodbAtlasUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) MongodbAtlasUserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbAtlasUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) RotationPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) RotationPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) SecretProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) SecretProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) TwilioApiKey() VaultSecretsRotatingSecretTwilioApiKeyOutputReference {
	var returns VaultSecretsRotatingSecretTwilioApiKeyOutputReference
	_jsii_.Get(
		j,
		"twilioApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsRotatingSecret) TwilioApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twilioApiKeyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_rotating_secret hcp_vault_secrets_rotating_secret} Resource.
func NewVaultSecretsRotatingSecret(scope constructs.Construct, id *string, config *VaultSecretsRotatingSecretConfig) VaultSecretsRotatingSecret {
	_init_.Initialize()

	if err := validateNewVaultSecretsRotatingSecretParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultSecretsRotatingSecret{}

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultSecretsRotatingSecret.VaultSecretsRotatingSecret",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/vault_secrets_rotating_secret hcp_vault_secrets_rotating_secret} Resource.
func NewVaultSecretsRotatingSecret_Override(v VaultSecretsRotatingSecret, scope constructs.Construct, id *string, config *VaultSecretsRotatingSecretConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultSecretsRotatingSecret.VaultSecretsRotatingSecret",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetAppName(val *string) {
	if err := j.validateSetAppNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetIntegrationName(val *string) {
	if err := j.validateSetIntegrationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationName",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetRotationPolicyName(val *string) {
	if err := j.validateSetRotationPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPolicyName",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsRotatingSecret)SetSecretProvider(val *string) {
	if err := j.validateSetSecretProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretProvider",
		val,
	)
}

// Generates CDKTF code for importing a VaultSecretsRotatingSecret resource upon running "cdktf plan <stack-name>".
func VaultSecretsRotatingSecret_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVaultSecretsRotatingSecret_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsRotatingSecret.VaultSecretsRotatingSecret",
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
func VaultSecretsRotatingSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultSecretsRotatingSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsRotatingSecret.VaultSecretsRotatingSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultSecretsRotatingSecret_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultSecretsRotatingSecret_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsRotatingSecret.VaultSecretsRotatingSecret",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultSecretsRotatingSecret_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultSecretsRotatingSecret_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsRotatingSecret.VaultSecretsRotatingSecret",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VaultSecretsRotatingSecret_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcp.vaultSecretsRotatingSecret.VaultSecretsRotatingSecret",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultSecretsRotatingSecret) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) PutAwsAccessKeys(value *VaultSecretsRotatingSecretAwsAccessKeys) {
	if err := v.validatePutAwsAccessKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAwsAccessKeys",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) PutAzureApplicationPassword(value *VaultSecretsRotatingSecretAzureApplicationPassword) {
	if err := v.validatePutAzureApplicationPasswordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAzureApplicationPassword",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) PutConfluentServiceAccount(value *VaultSecretsRotatingSecretConfluentServiceAccount) {
	if err := v.validatePutConfluentServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putConfluentServiceAccount",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) PutGcpServiceAccountKey(value *VaultSecretsRotatingSecretGcpServiceAccountKey) {
	if err := v.validatePutGcpServiceAccountKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGcpServiceAccountKey",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) PutMongodbAtlasUser(value *VaultSecretsRotatingSecretMongodbAtlasUser) {
	if err := v.validatePutMongodbAtlasUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMongodbAtlasUser",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) PutTwilioApiKey(value *VaultSecretsRotatingSecretTwilioApiKey) {
	if err := v.validatePutTwilioApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTwilioApiKey",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetAwsAccessKeys() {
	_jsii_.InvokeVoid(
		v,
		"resetAwsAccessKeys",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetAzureApplicationPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetAzureApplicationPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetConfluentServiceAccount() {
	_jsii_.InvokeVoid(
		v,
		"resetConfluentServiceAccount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetGcpServiceAccountKey() {
	_jsii_.InvokeVoid(
		v,
		"resetGcpServiceAccountKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetMongodbAtlasUser() {
	_jsii_.InvokeVoid(
		v,
		"resetMongodbAtlasUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetProjectId() {
	_jsii_.InvokeVoid(
		v,
		"resetProjectId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ResetTwilioApiKey() {
	_jsii_.InvokeVoid(
		v,
		"resetTwilioApiKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsRotatingSecret) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

