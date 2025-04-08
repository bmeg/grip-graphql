
package model

import (
	"encoding/json"
	"fmt"
)

type SafePatientLink struct {
	Type *string `json:"type,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Other TypedObject `json:"other"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *PatientLink) UnmarshalJSON(b []byte) error {
	var safe SafePatientLink
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PatientLink{
		Type: safe.Type,
		Extension: safe.Extension,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "other", safe.Other.Typename, &o.Other); err != nil {
		return fmt.Errorf("failed to unmarshal Other: %w", err)
	}

	return nil
}
