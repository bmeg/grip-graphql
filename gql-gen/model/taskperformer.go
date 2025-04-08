
package model

import (
	"encoding/json"
	"fmt"
)

type SafeTaskPerformer struct {
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Actor TypedObject `json:"actor"`
	Extension []*Extension `json:"extension,omitempty"`
	Function *CodeableConcept `json:"function,omitempty"`
	ID *string `json:"id,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *TaskPerformer) UnmarshalJSON(b []byte) error {
	var safe SafeTaskPerformer
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = TaskPerformer{
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		Function: safe.Function,
		ID: safe.ID,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "actor", safe.Actor.Typename, &o.Actor); err != nil {
		return fmt.Errorf("failed to unmarshal Actor: %w", err)
	}

	return nil
}
