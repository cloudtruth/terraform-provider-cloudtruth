# PatchedParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the parameter. | [optional] [readonly] 
**Name** | Pointer to **string** | The parameter name. | [optional] 
**Description** | Pointer to **string** | A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content. | [optional] 
**Secret** | Pointer to **bool** | Indicates if this content is secret or not.  When a parameter is considered to be a secret, any internal values are stored in a dedicated vault for your organization for maximum security.  External values are inspected on-demand to ensure they align with the parameter&#39;s secret setting and if they do not, those external values are not allowed to be used. | [optional] 
**Type** | Pointer to **string** | The type of this Parameter. | [optional] [default to "string"]
**Rules** | Pointer to [**[]ParameterRule**](ParameterRule.md) | Rules applied to this parameter. | [optional] [readonly] 
**Project** | Pointer to **string** | The project that the parameter is within. | [optional] [readonly] 
**ProjectName** | Pointer to **string** | The project name that the parameter is within. | [optional] [readonly] 
**ReferencingTemplates** | Pointer to **[]string** | Templates that reference this Parameter. | [optional] [readonly] 
**ReferencingValues** | Pointer to **[]string** | Dynamic values that reference this Parameter. | [optional] [readonly] 
**Values** | Pointer to [**map[string]ParameterValuesValue**](ParameterValuesValue.md) |              This dictionary has keys that correspond to environment urls, and values             that correspond to the effective value for this parameter in that environment.             Each parameter has an effective value in every environment based on             project dependencies and environment inheritance.              The effective value is found by looking (within the keyed environment) up             the project dependencies by parameter name.  If a value is not found, the             parent environment is consulted with the same logic to locate a value.  It             is possible for there to be a &#x60;null&#x60; value record for an environment, which             means there is no value set; it is also possible for there to be a value record             with a &#x60;value&#x60; of &#x60;null&#x60;, which means the value was explicitly set to &#x60;null&#x60;.              If the value&#39;s parameter does not match the enclosing parameter (holding the             values array) then that value is flowing in through project dependencies.             Clients must recognize this in case the user asks to modify the value; in this             case the client must POST a new Value to the current parameter to override the             value coming in from the project dependency.              If the Value.environment matches the key, then it is an explicit value set for             that environment.  If they differ, the value was obtained from a parent             environment (directly or indirectly).  If the value is None then no value has             ever been set in any environment for this parameter within all the project             dependencies.          | [optional] [readonly] 
**Overrides** | Pointer to **NullableString** | If this parameter&#39;s project depends on another project which provides a parameter of the same name, this parameter overrides the one provided by the dependee.  You can use this field to determine if there will be side-effects the user should know about when deleting a parameter.  Deleting a parameter that overrides another one due to an identical name will uncover the one from the dependee project. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedParameter

`func NewPatchedParameter() *PatchedParameter`

NewPatchedParameter instantiates a new PatchedParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedParameterWithDefaults

`func NewPatchedParameterWithDefaults() *PatchedParameter`

NewPatchedParameterWithDefaults instantiates a new PatchedParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PatchedParameter) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedParameter) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedParameter) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedParameter) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetId

`func (o *PatchedParameter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedParameter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedParameter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedParameter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedParameter) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedParameter) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedParameter) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedParameter) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *PatchedParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedParameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedParameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRules

`func (o *PatchedParameter) GetRules() []ParameterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PatchedParameter) GetRulesOk() (*[]ParameterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PatchedParameter) SetRules(v []ParameterRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PatchedParameter) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetProject

`func (o *PatchedParameter) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedParameter) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedParameter) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedParameter) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectName

`func (o *PatchedParameter) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *PatchedParameter) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *PatchedParameter) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *PatchedParameter) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetReferencingTemplates

`func (o *PatchedParameter) GetReferencingTemplates() []string`

GetReferencingTemplates returns the ReferencingTemplates field if non-nil, zero value otherwise.

### GetReferencingTemplatesOk

`func (o *PatchedParameter) GetReferencingTemplatesOk() (*[]string, bool)`

GetReferencingTemplatesOk returns a tuple with the ReferencingTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingTemplates

`func (o *PatchedParameter) SetReferencingTemplates(v []string)`

SetReferencingTemplates sets ReferencingTemplates field to given value.

### HasReferencingTemplates

`func (o *PatchedParameter) HasReferencingTemplates() bool`

HasReferencingTemplates returns a boolean if a field has been set.

### GetReferencingValues

`func (o *PatchedParameter) GetReferencingValues() []string`

GetReferencingValues returns the ReferencingValues field if non-nil, zero value otherwise.

### GetReferencingValuesOk

`func (o *PatchedParameter) GetReferencingValuesOk() (*[]string, bool)`

GetReferencingValuesOk returns a tuple with the ReferencingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingValues

`func (o *PatchedParameter) SetReferencingValues(v []string)`

SetReferencingValues sets ReferencingValues field to given value.

### HasReferencingValues

`func (o *PatchedParameter) HasReferencingValues() bool`

HasReferencingValues returns a boolean if a field has been set.

### GetValues

`func (o *PatchedParameter) GetValues() map[string]ParameterValuesValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PatchedParameter) GetValuesOk() (*map[string]ParameterValuesValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PatchedParameter) SetValues(v map[string]ParameterValuesValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *PatchedParameter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOverrides

`func (o *PatchedParameter) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *PatchedParameter) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *PatchedParameter) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *PatchedParameter) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### SetOverridesNil

`func (o *PatchedParameter) SetOverridesNil(b bool)`

 SetOverridesNil sets the value for Overrides to be an explicit nil

### UnsetOverrides
`func (o *PatchedParameter) UnsetOverrides()`

UnsetOverrides ensures that no value is present for Overrides, not even an explicit nil
### GetCreatedAt

`func (o *PatchedParameter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedParameter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedParameter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedParameter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PatchedParameter) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PatchedParameter) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PatchedParameter) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *PatchedParameter) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


