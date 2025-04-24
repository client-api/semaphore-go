# LoginMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OidcProviders** | Pointer to [**[]LoginMetadataOidcProvidersInner**](LoginMetadataOidcProvidersInner.md) | List of OIDC providers | [optional] 

## Methods

### NewLoginMetadata

`func NewLoginMetadata() *LoginMetadata`

NewLoginMetadata instantiates a new LoginMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginMetadataWithDefaults

`func NewLoginMetadataWithDefaults() *LoginMetadata`

NewLoginMetadataWithDefaults instantiates a new LoginMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidcProviders

`func (o *LoginMetadata) GetOidcProviders() []LoginMetadataOidcProvidersInner`

GetOidcProviders returns the OidcProviders field if non-nil, zero value otherwise.

### GetOidcProvidersOk

`func (o *LoginMetadata) GetOidcProvidersOk() (*[]LoginMetadataOidcProvidersInner, bool)`

GetOidcProvidersOk returns a tuple with the OidcProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProviders

`func (o *LoginMetadata) SetOidcProviders(v []LoginMetadataOidcProvidersInner)`

SetOidcProviders sets OidcProviders field to given value.

### HasOidcProviders

`func (o *LoginMetadata) HasOidcProviders() bool`

HasOidcProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


