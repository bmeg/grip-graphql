
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDocumentReferenceAttester struct {
	ID *string `json:"id,omitempty"`
	Mode *CodeableConcept `json:"mode,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Party TypedObject `json:"party"`
	ResourceType *string `json:"resourceType,omitempty"`
	Time *string `json:"time,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DocumentReferenceAttester) UnmarshalJSON(b []byte) error {
	var safe SafeDocumentReferenceAttester
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DocumentReferenceAttester{
		ID: safe.ID,
		Mode: safe.Mode,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Time: safe.Time,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "party", safe.Party.Typename, &o.Party); err != nil {
		return fmt.Errorf("failed to unmarshal Party: %w", err)
	}

	return nil
}
