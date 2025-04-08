
package model

import (
	"encoding/json"
	"fmt"
)

type SafeImagingStudySeriesPerformer struct {
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Actor TypedObject `json:"actor"`
	Extension []*Extension `json:"extension,omitempty"`
	Function *CodeableConcept `json:"function,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ImagingStudySeriesPerformer) UnmarshalJSON(b []byte) error {
	var safe SafeImagingStudySeriesPerformer
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ImagingStudySeriesPerformer{
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
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
