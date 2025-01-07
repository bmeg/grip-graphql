
package model

import (
	"encoding/json"
	"fmt"
)

type SafeTaskPerformer struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Actor TypedObject `json:"actor"`
	Extension []*Extension `json:"extension,omitempty"`
	Function *CodeableConcept `json:"function,omitempty"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *TaskPerformer) UnmarshalJSON(b []byte) error {
	var safe SafeTaskPerformer
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = TaskPerformer{
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		Function: safe.Function,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "actor", safe.Actor.Typename, &o.Actor); err != nil {
		return fmt.Errorf("failed to unmarshal Actor: %w", err)
	}

	return nil
}
