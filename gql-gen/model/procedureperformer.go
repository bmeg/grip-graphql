
package model

import (
	"encoding/json"
	"fmt"
)

type SafeProcedurePerformer struct {
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	OnBehalfOf *OrganizationType `json:"onBehalfOf"`
	Period *Period `json:"period,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Actor TypedObject `json:"actor"`
	Extension []*Extension `json:"extension,omitempty"`
	Function *CodeableConcept `json:"function,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ProcedurePerformer) UnmarshalJSON(b []byte) error {
	var safe SafeProcedurePerformer
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ProcedurePerformer{
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		OnBehalfOf: safe.OnBehalfOf,
		Period: safe.Period,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		Function: safe.Function,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "actor", safe.Actor.Typename, &o.Actor); err != nil {
		return fmt.Errorf("failed to unmarshal Actor: %w", err)
	}

	return nil
}
