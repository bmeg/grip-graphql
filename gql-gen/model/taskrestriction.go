
package model

import (
	"encoding/json"
	"fmt"
)

type SafeTaskRestriction struct {
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Period *Period `json:"period,omitempty"`
	Recipient TypedObject `json:"recipient"`
	Repetitions *int `json:"repetitions,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *TaskRestriction) UnmarshalJSON(b []byte) error {
	var safe SafeTaskRestriction
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = TaskRestriction{
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		Period: safe.Period,
		Repetitions: safe.Repetitions,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "recipient", safe.Recipient.Typename, &o.Recipient); err != nil {
		return fmt.Errorf("failed to unmarshal Recipient: %w", err)
	}

	return nil
}
