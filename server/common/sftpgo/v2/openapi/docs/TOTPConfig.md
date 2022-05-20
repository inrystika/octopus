# TOTPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Algo** | Pointer to [**TOTPHMacAlgo**](TOTPHMacAlgo.md) |  | [optional] 

## Methods

### NewTOTPConfig

`func NewTOTPConfig() *TOTPConfig`

NewTOTPConfig instantiates a new TOTPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTOTPConfigWithDefaults

`func NewTOTPConfigWithDefaults() *TOTPConfig`

NewTOTPConfigWithDefaults instantiates a new TOTPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TOTPConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TOTPConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TOTPConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TOTPConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIssuer

`func (o *TOTPConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TOTPConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TOTPConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *TOTPConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAlgo

`func (o *TOTPConfig) GetAlgo() TOTPHMacAlgo`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *TOTPConfig) GetAlgoOk() (*TOTPHMacAlgo, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *TOTPConfig) SetAlgo(v TOTPHMacAlgo)`

SetAlgo sets Algo field to given value.

### HasAlgo

`func (o *TOTPConfig) HasAlgo() bool`

HasAlgo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


