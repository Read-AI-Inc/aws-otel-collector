// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsApiHandleAttributes Handle attributes.
type MicrosoftTeamsApiHandleAttributes struct {
	// Channel id.
	ChannelId *string `json:"channel_id,omitempty"`
	// Handle name.
	Name *string `json:"name,omitempty"`
	// Team id.
	TeamId *string `json:"team_id,omitempty"`
	// Tenant id.
	TenantId *string `json:"tenant_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftTeamsApiHandleAttributes instantiates a new MicrosoftTeamsApiHandleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsApiHandleAttributes() *MicrosoftTeamsApiHandleAttributes {
	this := MicrosoftTeamsApiHandleAttributes{}
	return &this
}

// NewMicrosoftTeamsApiHandleAttributesWithDefaults instantiates a new MicrosoftTeamsApiHandleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsApiHandleAttributesWithDefaults() *MicrosoftTeamsApiHandleAttributes {
	this := MicrosoftTeamsApiHandleAttributes{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleAttributes) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleAttributes) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleAttributes) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *MicrosoftTeamsApiHandleAttributes) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftTeamsApiHandleAttributes) SetName(v string) {
	o.Name = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleAttributes) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleAttributes) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleAttributes) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *MicrosoftTeamsApiHandleAttributes) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *MicrosoftTeamsApiHandleAttributes) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsApiHandleAttributes) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *MicrosoftTeamsApiHandleAttributes) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *MicrosoftTeamsApiHandleAttributes) SetTenantId(v string) {
	o.TenantId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsApiHandleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChannelId != nil {
		toSerialize["channel_id"] = o.ChannelId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsApiHandleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelId *string `json:"channel_id,omitempty"`
		Name      *string `json:"name,omitempty"`
		TeamId    *string `json:"team_id,omitempty"`
		TenantId  *string `json:"tenant_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel_id", "name", "team_id", "tenant_id"})
	} else {
		return err
	}
	o.ChannelId = all.ChannelId
	o.Name = all.Name
	o.TeamId = all.TeamId
	o.TenantId = all.TenantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
