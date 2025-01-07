
package model

import (
	"encoding/json"
	"fmt"
)

type SafeConditionStage struct {
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Summary *CodeableConcept `json:"summary,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Assessment TypedObject `json:"assessment"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ConditionStage) UnmarshalJSON(b []byte) error {
	var safe SafeConditionStage
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ConditionStage{
		Extension: safe.Extension,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Summary: safe.Summary,
		Type: safe.Type,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "assessment", safe.Assessment.Typename, &o.Assessment); err != nil {
		return fmt.Errorf("failed to unmarshal Assessment: %w", err)
	}

	return nil
}
