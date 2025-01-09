
package model

import (
	"encoding/json"
	"fmt"
)

type SafeFamilyMemberHistoryParticipant struct {
	Function *CodeableConcept `json:"function,omitempty"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Actor TypedObject `json:"actor"`
	Extension []*Extension `json:"extension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *FamilyMemberHistoryParticipant) UnmarshalJSON(b []byte) error {
	var safe SafeFamilyMemberHistoryParticipant
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = FamilyMemberHistoryParticipant{
		Function: safe.Function,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "actor", safe.Actor.Typename, &o.Actor); err != nil {
		return fmt.Errorf("failed to unmarshal Actor: %w", err)
	}

	return nil
}
