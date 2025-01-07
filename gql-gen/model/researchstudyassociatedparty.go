
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchStudyAssociatedParty struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	Name *string `json:"name,omitempty"`
	Role *CodeableConcept `json:"role,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ID *string `json:"id,omitempty"`
	Party TypedObject `json:"party"`
	Period []*Period `json:"period,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchStudyAssociatedParty) UnmarshalJSON(b []byte) error {
	var safe SafeResearchStudyAssociatedParty
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchStudyAssociatedParty{
		ResourceType: safe.ResourceType,
		Classifier: safe.Classifier,
		Name: safe.Name,
		Role: safe.Role,
		ModifierExtension: safe.ModifierExtension,
		ID: safe.ID,
		Period: safe.Period,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "party", safe.Party.Typename, &o.Party); err != nil {
		return fmt.Errorf("failed to unmarshal Party: %w", err)
	}

	return nil
}
