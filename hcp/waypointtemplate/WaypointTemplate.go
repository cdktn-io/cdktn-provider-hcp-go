// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package waypointtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-hcp-go/hcp/v10/waypointtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template hcp_waypoint_template}.
type WaypointTemplate interface {
	cdktf.TerraformResource
	Actions() *[]*string
	SetActions(val *[]*string)
	ActionsInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Labels() *[]*string
	SetLabels(val *[]*string)
	LabelsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	ReadmeMarkdownTemplate() *string
	SetReadmeMarkdownTemplate(val *string)
	ReadmeMarkdownTemplateInput() *string
	Summary() *string
	SetSummary(val *string)
	SummaryInput() *string
	TerraformAgentPoolId() *string
	SetTerraformAgentPoolId(val *string)
	TerraformAgentPoolIdInput() *string
	TerraformCloudWorkspaceDetails() WaypointTemplateTerraformCloudWorkspaceDetailsOutputReference
	TerraformCloudWorkspaceDetailsInput() interface{}
	TerraformExecutionMode() *string
	SetTerraformExecutionMode(val *string)
	TerraformExecutionModeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	TerraformNoCodeModuleId() *string
	SetTerraformNoCodeModuleId(val *string)
	TerraformNoCodeModuleIdInput() *string
	TerraformNoCodeModuleSource() *string
	SetTerraformNoCodeModuleSource(val *string)
	TerraformNoCodeModuleSourceInput() *string
	TerraformProjectId() *string
	SetTerraformProjectId(val *string)
	TerraformProjectIdInput() *string
	// Experimental.
	TerraformResourceType() *string
	UseModuleReadme() interface{}
	SetUseModuleReadme(val interface{})
	UseModuleReadmeInput() interface{}
	VariableOptions() WaypointTemplateVariableOptionsList
	VariableOptionsInput() interface{}
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
	PutTerraformCloudWorkspaceDetails(value *WaypointTemplateTerraformCloudWorkspaceDetails)
	PutVariableOptions(value interface{})
	ResetActions()
	ResetDescription()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetReadmeMarkdownTemplate()
	ResetTerraformAgentPoolId()
	ResetTerraformCloudWorkspaceDetails()
	ResetTerraformExecutionMode()
	ResetUseModuleReadme()
	ResetVariableOptions()
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

// The jsii proxy struct for WaypointTemplate
type jsiiProxy_WaypointTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WaypointTemplate) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) ReadmeMarkdownTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readmeMarkdownTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) ReadmeMarkdownTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readmeMarkdownTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) Summary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) SummaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformAgentPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAgentPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformAgentPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAgentPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformCloudWorkspaceDetails() WaypointTemplateTerraformCloudWorkspaceDetailsOutputReference {
	var returns WaypointTemplateTerraformCloudWorkspaceDetailsOutputReference
	_jsii_.Get(
		j,
		"terraformCloudWorkspaceDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformCloudWorkspaceDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terraformCloudWorkspaceDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformExecutionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformExecutionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformExecutionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformExecutionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformNoCodeModuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformNoCodeModuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformNoCodeModuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformNoCodeModuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformNoCodeModuleSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformNoCodeModuleSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformNoCodeModuleSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformNoCodeModuleSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) UseModuleReadme() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useModuleReadme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) UseModuleReadmeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useModuleReadmeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) VariableOptions() WaypointTemplateVariableOptionsList {
	var returns WaypointTemplateVariableOptionsList
	_jsii_.Get(
		j,
		"variableOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaypointTemplate) VariableOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template hcp_waypoint_template} Resource.
func NewWaypointTemplate(scope constructs.Construct, id *string, config *WaypointTemplateConfig) WaypointTemplate {
	_init_.Initialize()

	if err := validateNewWaypointTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WaypointTemplate{}

	_jsii_.Create(
		"@cdktf/provider-hcp.waypointTemplate.WaypointTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/hcp/0.111.0/docs/resources/waypoint_template hcp_waypoint_template} Resource.
func NewWaypointTemplate_Override(w WaypointTemplate, scope constructs.Construct, id *string, config *WaypointTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-hcp.waypointTemplate.WaypointTemplate",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetLabels(val *[]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetReadmeMarkdownTemplate(val *string) {
	if err := j.validateSetReadmeMarkdownTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readmeMarkdownTemplate",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetSummary(val *string) {
	if err := j.validateSetSummaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summary",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetTerraformAgentPoolId(val *string) {
	if err := j.validateSetTerraformAgentPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAgentPoolId",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetTerraformExecutionMode(val *string) {
	if err := j.validateSetTerraformExecutionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformExecutionMode",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetTerraformNoCodeModuleId(val *string) {
	if err := j.validateSetTerraformNoCodeModuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformNoCodeModuleId",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetTerraformNoCodeModuleSource(val *string) {
	if err := j.validateSetTerraformNoCodeModuleSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformNoCodeModuleSource",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetTerraformProjectId(val *string) {
	if err := j.validateSetTerraformProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformProjectId",
		val,
	)
}

func (j *jsiiProxy_WaypointTemplate)SetUseModuleReadme(val interface{}) {
	if err := j.validateSetUseModuleReadmeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useModuleReadme",
		val,
	)
}

// Generates CDKTF code for importing a WaypointTemplate resource upon running "cdktf plan <stack-name>".
func WaypointTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWaypointTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.waypointTemplate.WaypointTemplate",
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
func WaypointTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaypointTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.waypointTemplate.WaypointTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WaypointTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaypointTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.waypointTemplate.WaypointTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WaypointTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaypointTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-hcp.waypointTemplate.WaypointTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WaypointTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-hcp.waypointTemplate.WaypointTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WaypointTemplate) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WaypointTemplate) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WaypointTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WaypointTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WaypointTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WaypointTemplate) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WaypointTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WaypointTemplate) PutTerraformCloudWorkspaceDetails(value *WaypointTemplateTerraformCloudWorkspaceDetails) {
	if err := w.validatePutTerraformCloudWorkspaceDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTerraformCloudWorkspaceDetails",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WaypointTemplate) PutVariableOptions(value interface{}) {
	if err := w.validatePutVariableOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVariableOptions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetActions() {
	_jsii_.InvokeVoid(
		w,
		"resetActions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetLabels() {
	_jsii_.InvokeVoid(
		w,
		"resetLabels",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetProjectId() {
	_jsii_.InvokeVoid(
		w,
		"resetProjectId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetReadmeMarkdownTemplate() {
	_jsii_.InvokeVoid(
		w,
		"resetReadmeMarkdownTemplate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetTerraformAgentPoolId() {
	_jsii_.InvokeVoid(
		w,
		"resetTerraformAgentPoolId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetTerraformCloudWorkspaceDetails() {
	_jsii_.InvokeVoid(
		w,
		"resetTerraformCloudWorkspaceDetails",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetTerraformExecutionMode() {
	_jsii_.InvokeVoid(
		w,
		"resetTerraformExecutionMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetUseModuleReadme() {
	_jsii_.InvokeVoid(
		w,
		"resetUseModuleReadme",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) ResetVariableOptions() {
	_jsii_.InvokeVoid(
		w,
		"resetVariableOptions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WaypointTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaypointTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

