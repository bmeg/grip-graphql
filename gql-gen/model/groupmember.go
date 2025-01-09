
package model

import (
	"encoding/json"
	"fmt"
)

type SafeGroupMember struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Entity TypedObject `json:"entity"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	Inactive *bool `json:"inactive,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Period *Period `json:"period,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *GroupMember) UnmarshalJSON(b []byte) error {
	var safe SafeGroupMember
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = GroupMember{
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		ID: safe.ID,
		Inactive: safe.Inactive,
		ModifierExtension: safe.ModifierExtension,
		Period: safe.Period,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "entity", safe.Entity.Typename, &o.Entity); err != nil {
		return fmt.Errorf("failed to unmarshal Entity: %w", err)
	}

	return nil
}
