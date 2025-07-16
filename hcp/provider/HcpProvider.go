// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs hcp}.
type HcpProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CredentialFile() *string
	SetCredentialFile(val *string)
	CredentialFileInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	SkipStatusCheck() interface{}
	SetSkipStatusCheck(val interface{})
	SkipStatusCheckInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	WorkloadIdentity() interface{}
	SetWorkloadIdentity(val interface{})
	WorkloadIdentityInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetClientId()
	ResetClientSecret()
	ResetCredentialFile()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetSkipStatusCheck()
	ResetWorkloadIdentity()
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

// The jsii proxy struct for HcpProvider
type jsiiProxy_HcpProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_HcpProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) CredentialFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) CredentialFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) SkipStatusCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStatusCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) SkipStatusCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStatusCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) WorkloadIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HcpProvider) WorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workloadIdentityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs hcp} Resource.
func NewHcpProvider(scope constructs.Construct, id *string, config *HcpProviderConfig) HcpProvider {
	_init_.Initialize()

	if err := validateNewHcpProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HcpProvider{}

	_jsii_.Create(
		"@cdktf/provider-hcp.provider.HcpProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.109.0/docs hcp} Resource.
func NewHcpProvider_Override(h HcpProvider, scope constructs.Construct, id *string, config *HcpProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.provider.HcpProvider",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HcpProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_HcpProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_HcpProvider)SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_HcpProvider)SetCredentialFile(val *string) {
	_jsii_.Set(
		j,
		"credentialFile",
		val,
	)
}

func (j *jsiiProxy_HcpProvider)SetProjectId(val *string) {
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_HcpProvider)SetSkipStatusCheck(val interface{}) {
	if err := j.validateSetSkipStatusCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipStatusCheck",
		val,
	)
}

func (j *jsiiProxy_HcpProvider)SetWorkloadIdentity(val interface{}) {
	if err := j.validateSetWorkloadIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadIdentity",
		val,
	)
}

// Generates CDKTF code for importing a HcpProvider resource upon running "cdktf plan <stack-name>".
func HcpProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHcpProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.provider.HcpProvider",
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
func HcpProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHcpProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.provider.HcpProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HcpProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHcpProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.provider.HcpProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HcpProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHcpProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.provider.HcpProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HcpProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcp.provider.HcpProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HcpProvider) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HcpProvider) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HcpProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		h,
		"resetAlias",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		h,
		"resetClientId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) ResetClientSecret() {
	_jsii_.InvokeVoid(
		h,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) ResetCredentialFile() {
	_jsii_.InvokeVoid(
		h,
		"resetCredentialFile",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) ResetProjectId() {
	_jsii_.InvokeVoid(
		h,
		"resetProjectId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) ResetSkipStatusCheck() {
	_jsii_.InvokeVoid(
		h,
		"resetSkipStatusCheck",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) ResetWorkloadIdentity() {
	_jsii_.InvokeVoid(
		h,
		"resetWorkloadIdentity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HcpProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcpProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcpProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcpProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcpProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HcpProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

