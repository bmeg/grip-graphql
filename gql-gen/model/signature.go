
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSignature struct {
	Data *string `json:"data,omitempty"`
	OnBehalfOf TypedObject `json:"onBehalfOf"`
	When *string `json:"when,omitempty"`
	SigFormat *string `json:"sigFormat,omitempty"`
	TargetFormat *string `json:"targetFormat,omitempty"`
	Who TypedObject `json:"who"`
	ID *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Type []*Coding `json:"type,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *Signature) UnmarshalJSON(b []byte) error {
	var safe SafeSignature
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = Signature{
		Data: safe.Data,
		When: safe.When,
		SigFormat: safe.SigFormat,
		TargetFormat: safe.TargetFormat,
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		Type: safe.Type,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "onBehalfOf", safe.OnBehalfOf.Typename, &o.OnBehalfOf); err != nil {
		return fmt.Errorf("failed to unmarshal OnBehalfOf: %w", err)
	}
	if err := unmarshalUnion(b, "who", safe.Who.Typename, &o.Who); err != nil {
		return fmt.Errorf("failed to unmarshal Who: %w", err)
	}

	return nil
}
