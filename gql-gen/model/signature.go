
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSignature struct {
	TargetFormat *string `json:"targetFormat,omitempty"`
	Type []*Coding `json:"type,omitempty"`
	When *string `json:"when,omitempty"`
	Who TypedObject `json:"who"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Data *string `json:"data,omitempty"`
	OnBehalfOf TypedObject `json:"onBehalfOf"`
	SigFormat *string `json:"sigFormat,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *Signature) UnmarshalJSON(b []byte) error {
	var safe SafeSignature
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = Signature{
		TargetFormat: safe.TargetFormat,
		Type: safe.Type,
		When: safe.When,
		Extension: safe.Extension,
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		Data: safe.Data,
		SigFormat: safe.SigFormat,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "who", safe.Who.Typename, &o.Who); err != nil {
		return fmt.Errorf("failed to unmarshal Who: %w", err)
	}
	if err := unmarshalUnion(b, "onBehalfOf", safe.OnBehalfOf.Typename, &o.OnBehalfOf); err != nil {
		return fmt.Errorf("failed to unmarshal OnBehalfOf: %w", err)
	}

	return nil
}
