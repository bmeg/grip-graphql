
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchStudyAssociatedParty struct {
	Role *CodeableConcept `json:"role,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Period []*Period `json:"period,omitempty"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Party TypedObject `json:"party"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchStudyAssociatedParty) UnmarshalJSON(b []byte) error {
	var safe SafeResearchStudyAssociatedParty
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchStudyAssociatedParty{
		Role: safe.Role,
		Extension: safe.Extension,
		ID: safe.ID,
		Name: safe.Name,
		Period: safe.Period,
		Classifier: safe.Classifier,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "party", safe.Party.Typename, &o.Party); err != nil {
		return fmt.Errorf("failed to unmarshal Party: %w", err)
	}

	return nil
}
