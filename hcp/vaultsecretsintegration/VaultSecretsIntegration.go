// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecretsintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/vaultsecretsintegration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration hcp_vault_secrets_integration}.
type VaultSecretsIntegration interface {
	cdktf.TerraformResource
	AwsAccessKeys() VaultSecretsIntegrationAwsAccessKeysOutputReference
	AwsAccessKeysInput() interface{}
	AwsFederatedWorkloadIdentity() VaultSecretsIntegrationAwsFederatedWorkloadIdentityOutputReference
	AwsFederatedWorkloadIdentityInput() interface{}
	AzureClientSecret() VaultSecretsIntegrationAzureClientSecretOutputReference
	AzureClientSecretInput() interface{}
	AzureFederatedWorkloadIdentity() VaultSecretsIntegrationAzureFederatedWorkloadIdentityOutputReference
	AzureFederatedWorkloadIdentityInput() interface{}
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	CapabilitiesInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfluentStaticCredentials() VaultSecretsIntegrationConfluentStaticCredentialsOutputReference
	ConfluentStaticCredentialsInput() interface{}
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
	GcpFederatedWorkloadIdentity() VaultSecretsIntegrationGcpFederatedWorkloadIdentityOutputReference
	GcpFederatedWorkloadIdentityInput() interface{}
	GcpServiceAccountKey() VaultSecretsIntegrationGcpServiceAccountKeyOutputReference
	GcpServiceAccountKeyInput() interface{}
	GitlabAccess() VaultSecretsIntegrationGitlabAccessOutputReference
	GitlabAccessInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongodbAtlasStaticCredentials() VaultSecretsIntegrationMongodbAtlasStaticCredentialsOutputReference
	MongodbAtlasStaticCredentialsInput() interface{}
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
	ProviderType() *string
	SetProviderType(val *string)
	ProviderTypeInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceId() *string
	ResourceName() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TwilioStaticCredentials() VaultSecretsIntegrationTwilioStaticCredentialsOutputReference
	TwilioStaticCredentialsInput() interface{}
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
	PutAwsAccessKeys(value *VaultSecretsIntegrationAwsAccessKeys)
	PutAwsFederatedWorkloadIdentity(value *VaultSecretsIntegrationAwsFederatedWorkloadIdentity)
	PutAzureClientSecret(value *VaultSecretsIntegrationAzureClientSecret)
	PutAzureFederatedWorkloadIdentity(value *VaultSecretsIntegrationAzureFederatedWorkloadIdentity)
	PutConfluentStaticCredentials(value *VaultSecretsIntegrationConfluentStaticCredentials)
	PutGcpFederatedWorkloadIdentity(value *VaultSecretsIntegrationGcpFederatedWorkloadIdentity)
	PutGcpServiceAccountKey(value *VaultSecretsIntegrationGcpServiceAccountKey)
	PutGitlabAccess(value *VaultSecretsIntegrationGitlabAccess)
	PutMongodbAtlasStaticCredentials(value *VaultSecretsIntegrationMongodbAtlasStaticCredentials)
	PutTwilioStaticCredentials(value *VaultSecretsIntegrationTwilioStaticCredentials)
	ResetAwsAccessKeys()
	ResetAwsFederatedWorkloadIdentity()
	ResetAzureClientSecret()
	ResetAzureFederatedWorkloadIdentity()
	ResetConfluentStaticCredentials()
	ResetGcpFederatedWorkloadIdentity()
	ResetGcpServiceAccountKey()
	ResetGitlabAccess()
	ResetMongodbAtlasStaticCredentials()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetTwilioStaticCredentials()
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

// The jsii proxy struct for VaultSecretsIntegration
type jsiiProxy_VaultSecretsIntegration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VaultSecretsIntegration) AwsAccessKeys() VaultSecretsIntegrationAwsAccessKeysOutputReference {
	var returns VaultSecretsIntegrationAwsAccessKeysOutputReference
	_jsii_.Get(
		j,
		"awsAccessKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) AwsAccessKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccessKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) AwsFederatedWorkloadIdentity() VaultSecretsIntegrationAwsFederatedWorkloadIdentityOutputReference {
	var returns VaultSecretsIntegrationAwsFederatedWorkloadIdentityOutputReference
	_jsii_.Get(
		j,
		"awsFederatedWorkloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) AwsFederatedWorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsFederatedWorkloadIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) AzureClientSecret() VaultSecretsIntegrationAzureClientSecretOutputReference {
	var returns VaultSecretsIntegrationAzureClientSecretOutputReference
	_jsii_.Get(
		j,
		"azureClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) AzureClientSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) AzureFederatedWorkloadIdentity() VaultSecretsIntegrationAzureFederatedWorkloadIdentityOutputReference {
	var returns VaultSecretsIntegrationAzureFederatedWorkloadIdentityOutputReference
	_jsii_.Get(
		j,
		"azureFederatedWorkloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) AzureFederatedWorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureFederatedWorkloadIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) CapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ConfluentStaticCredentials() VaultSecretsIntegrationConfluentStaticCredentialsOutputReference {
	var returns VaultSecretsIntegrationConfluentStaticCredentialsOutputReference
	_jsii_.Get(
		j,
		"confluentStaticCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ConfluentStaticCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluentStaticCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) GcpFederatedWorkloadIdentity() VaultSecretsIntegrationGcpFederatedWorkloadIdentityOutputReference {
	var returns VaultSecretsIntegrationGcpFederatedWorkloadIdentityOutputReference
	_jsii_.Get(
		j,
		"gcpFederatedWorkloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) GcpFederatedWorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpFederatedWorkloadIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) GcpServiceAccountKey() VaultSecretsIntegrationGcpServiceAccountKeyOutputReference {
	var returns VaultSecretsIntegrationGcpServiceAccountKeyOutputReference
	_jsii_.Get(
		j,
		"gcpServiceAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) GcpServiceAccountKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpServiceAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) GitlabAccess() VaultSecretsIntegrationGitlabAccessOutputReference {
	var returns VaultSecretsIntegrationGitlabAccessOutputReference
	_jsii_.Get(
		j,
		"gitlabAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) GitlabAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitlabAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) MongodbAtlasStaticCredentials() VaultSecretsIntegrationMongodbAtlasStaticCredentialsOutputReference {
	var returns VaultSecretsIntegrationMongodbAtlasStaticCredentialsOutputReference
	_jsii_.Get(
		j,
		"mongodbAtlasStaticCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) MongodbAtlasStaticCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbAtlasStaticCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) TwilioStaticCredentials() VaultSecretsIntegrationTwilioStaticCredentialsOutputReference {
	var returns VaultSecretsIntegrationTwilioStaticCredentialsOutputReference
	_jsii_.Get(
		j,
		"twilioStaticCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultSecretsIntegration) TwilioStaticCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twilioStaticCredentialsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration hcp_vault_secrets_integration} Resource.
func NewVaultSecretsIntegration(scope constructs.Construct, id *string, config *VaultSecretsIntegrationConfig) VaultSecretsIntegration {
	_init_.Initialize()

	if err := validateNewVaultSecretsIntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultSecretsIntegration{}

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultSecretsIntegration.VaultSecretsIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.110.0/docs/resources/vault_secrets_integration hcp_vault_secrets_integration} Resource.
func NewVaultSecretsIntegration_Override(v VaultSecretsIntegration, scope constructs.Construct, id *string, config *VaultSecretsIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.vaultSecretsIntegration.VaultSecretsIntegration",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetCapabilities(val *[]*string) {
	if err := j.validateSetCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetProviderType(val *string) {
	if err := j.validateSetProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_VaultSecretsIntegration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a VaultSecretsIntegration resource upon running "cdktf plan <stack-name>".
func VaultSecretsIntegration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVaultSecretsIntegration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsIntegration.VaultSecretsIntegration",
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
func VaultSecretsIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultSecretsIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsIntegration.VaultSecretsIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultSecretsIntegration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultSecretsIntegration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsIntegration.VaultSecretsIntegration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultSecretsIntegration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultSecretsIntegration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.vaultSecretsIntegration.VaultSecretsIntegration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VaultSecretsIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcp.vaultSecretsIntegration.VaultSecretsIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VaultSecretsIntegration) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VaultSecretsIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VaultSecretsIntegration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VaultSecretsIntegration) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutAwsAccessKeys(value *VaultSecretsIntegrationAwsAccessKeys) {
	if err := v.validatePutAwsAccessKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAwsAccessKeys",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutAwsFederatedWorkloadIdentity(value *VaultSecretsIntegrationAwsFederatedWorkloadIdentity) {
	if err := v.validatePutAwsFederatedWorkloadIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAwsFederatedWorkloadIdentity",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutAzureClientSecret(value *VaultSecretsIntegrationAzureClientSecret) {
	if err := v.validatePutAzureClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAzureClientSecret",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutAzureFederatedWorkloadIdentity(value *VaultSecretsIntegrationAzureFederatedWorkloadIdentity) {
	if err := v.validatePutAzureFederatedWorkloadIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAzureFederatedWorkloadIdentity",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutConfluentStaticCredentials(value *VaultSecretsIntegrationConfluentStaticCredentials) {
	if err := v.validatePutConfluentStaticCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putConfluentStaticCredentials",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutGcpFederatedWorkloadIdentity(value *VaultSecretsIntegrationGcpFederatedWorkloadIdentity) {
	if err := v.validatePutGcpFederatedWorkloadIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGcpFederatedWorkloadIdentity",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutGcpServiceAccountKey(value *VaultSecretsIntegrationGcpServiceAccountKey) {
	if err := v.validatePutGcpServiceAccountKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGcpServiceAccountKey",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutGitlabAccess(value *VaultSecretsIntegrationGitlabAccess) {
	if err := v.validatePutGitlabAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putGitlabAccess",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutMongodbAtlasStaticCredentials(value *VaultSecretsIntegrationMongodbAtlasStaticCredentials) {
	if err := v.validatePutMongodbAtlasStaticCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMongodbAtlasStaticCredentials",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) PutTwilioStaticCredentials(value *VaultSecretsIntegrationTwilioStaticCredentials) {
	if err := v.validatePutTwilioStaticCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTwilioStaticCredentials",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetAwsAccessKeys() {
	_jsii_.InvokeVoid(
		v,
		"resetAwsAccessKeys",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetAwsFederatedWorkloadIdentity() {
	_jsii_.InvokeVoid(
		v,
		"resetAwsFederatedWorkloadIdentity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetAzureClientSecret() {
	_jsii_.InvokeVoid(
		v,
		"resetAzureClientSecret",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetAzureFederatedWorkloadIdentity() {
	_jsii_.InvokeVoid(
		v,
		"resetAzureFederatedWorkloadIdentity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetConfluentStaticCredentials() {
	_jsii_.InvokeVoid(
		v,
		"resetConfluentStaticCredentials",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetGcpFederatedWorkloadIdentity() {
	_jsii_.InvokeVoid(
		v,
		"resetGcpFederatedWorkloadIdentity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetGcpServiceAccountKey() {
	_jsii_.InvokeVoid(
		v,
		"resetGcpServiceAccountKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetGitlabAccess() {
	_jsii_.InvokeVoid(
		v,
		"resetGitlabAccess",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetMongodbAtlasStaticCredentials() {
	_jsii_.InvokeVoid(
		v,
		"resetMongodbAtlasStaticCredentials",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetProjectId() {
	_jsii_.InvokeVoid(
		v,
		"resetProjectId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) ResetTwilioStaticCredentials() {
	_jsii_.InvokeVoid(
		v,
		"resetTwilioStaticCredentials",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultSecretsIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultSecretsIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

