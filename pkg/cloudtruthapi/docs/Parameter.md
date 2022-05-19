# Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** | A unique identifier for the parameter. | [readonly] 
**Name** | **string** | The parameter name. | 
**Description** | Pointer to **string** | A description of the parameter.  You may find it helpful to document how this parameter is used to assist others when they need to maintain software that uses this content. | [optional] 
**Secret** | Pointer to **bool** | Indicates if this content is secret or not.  When a parameter is considered to be a secret, any internal values are stored in a dedicated vault for your organization for maximum security.  External values are inspected on-demand to ensure they align with the parameter&#39;s secret setting and if they do not, those external values are not allowed to be used. | [optional] 
**Type** | Pointer to **string** | The type of this Parameter. | [optional] [default to "string"]
**Rules** | [**[]ParameterRule**](ParameterRule.md) | Rules applied to this parameter. | [readonly] 
**Project** | **string** | The project that the parameter is within. | [readonly] 
**ProjectName** | **string** | The project name that the parameter is within. | [readonly] 
**ReferencingTemplates** | **[]string** | Templates that reference this Parameter. | [readonly] 
**ReferencingValues** | **[]string** | Dynamic values that reference this Parameter. | [readonly] 
**Values** | [**map[string]ParameterValuesValue**](ParameterValuesValue.md) |              This dictionary has keys that correspond to environment urls, and values             that correspond to the effective value for this parameter in that environment.             Each parameter has an effective value in every environment based on             project dependencies and environment inheritance.              The effective value is found by looking (within the keyed environment) up             the project dependencies by parameter name.  If a value is not found, the             parent environment is consulted with the same logic to locate a value.  It             is possible for there to be a &#x60;null&#x60; value record for an environment, which             means there is no value set; it is also possible for there to be a value record             with a &#x60;value&#x60; of &#x60;null&#x60;, which means the value was explicitly set to &#x60;null&#x60;.              If the value&#39;s parameter does not match the enclosing parameter (holding the             values array) then that value is flowing in through project dependencies.             Clients must recognize this in case the user asks to modify the value; in this             case the client must POST a new Value to the current parameter to override the             value coming in from the project dependency.              If the Value.environment matches the key, then it is an explicit value set for             that environment.  If they differ, the value was obtained from a parent             environment (directly or indirectly).  If the value is None then no value has             ever been set in any environment for this parameter within all the project             dependencies.          | [readonly] 
**Overrides** | **NullableString** | If this parameter&#39;s project depends on another project which provides a parameter of the same name, this parameter overrides the one provided by the dependee.  You can use this field to determine if there will be side-effects the user should know about when deleting a parameter.  Deleting a parameter that overrides another one due to an identical name will uncover the one from the dependee project. | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewParameter

`func NewParameter(url string, id string, name string, rules []ParameterRule, project string, projectName string, referencingTemplates []string, referencingValues []string, values map[string]ParameterValuesValue, overrides NullableString, createdAt time.Time, modifiedAt time.Time, ) *Parameter`

NewParameter instantiates a new Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterWithDefaults

`func NewParameterWithDefaults() *Parameter`

NewParameterWithDefaults instantiates a new Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Parameter) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Parameter) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Parameter) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *Parameter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Parameter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Parameter) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Parameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Parameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Parameter) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Parameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Parameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Parameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Parameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *Parameter) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Parameter) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Parameter) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Parameter) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *Parameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Parameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Parameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Parameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRules

`func (o *Parameter) GetRules() []ParameterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Parameter) GetRulesOk() (*[]ParameterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Parameter) SetRules(v []ParameterRule)`

SetRules sets Rules field to given value.


### GetProject

`func (o *Parameter) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Parameter) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Parameter) SetProject(v string)`

SetProject sets Project field to given value.


### GetProjectName

`func (o *Parameter) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Parameter) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Parameter) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReferencingTemplates

`func (o *Parameter) GetReferencingTemplates() []string`

GetReferencingTemplates returns the ReferencingTemplates field if non-nil, zero value otherwise.

### GetReferencingTemplatesOk

`func (o *Parameter) GetReferencingTemplatesOk() (*[]string, bool)`

GetReferencingTemplatesOk returns a tuple with the ReferencingTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingTemplates

`func (o *Parameter) SetReferencingTemplates(v []string)`

SetReferencingTemplates sets ReferencingTemplates field to given value.


### GetReferencingValues

`func (o *Parameter) GetReferencingValues() []string`

GetReferencingValues returns the ReferencingValues field if non-nil, zero value otherwise.

### GetReferencingValuesOk

`func (o *Parameter) GetReferencingValuesOk() (*[]string, bool)`

GetReferencingValuesOk returns a tuple with the ReferencingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingValues

`func (o *Parameter) SetReferencingValues(v []string)`

SetReferencingValues sets ReferencingValues field to given value.


### GetValues

`func (o *Parameter) GetValues() map[string]ParameterValuesValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Parameter) GetValuesOk() (*map[string]ParameterValuesValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Parameter) SetValues(v map[string]ParameterValuesValue)`

SetValues sets Values field to given value.


### GetOverrides

`func (o *Parameter) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *Parameter) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *Parameter) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.


### SetOverridesNil

`func (o *Parameter) SetOverridesNil(b bool)`

 SetOverridesNil sets the value for Overrides to be an explicit nil

### UnsetOverrides
`func (o *Parameter) UnsetOverrides()`

UnsetOverrides ensures that no value is present for Overrides, not even an explicit nil
### GetCreatedAt

`func (o *Parameter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Parameter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Parameter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Parameter) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Parameter) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Parameter) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


