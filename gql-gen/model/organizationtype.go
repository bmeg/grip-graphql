
package model

import (
	"encoding/json"
	"fmt"
)

type SafeOrganizationType struct {
	Identifier []*Identifier `json:"identifier,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Active *string `json:"active,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	PartOf *OrganizationType `json:"partOf"`
	ResourceType *string `json:"resourceType,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Description *string `json:"description,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Name *string `json:"name,omitempty"`
	Type []*CodeableConcept `json:"type,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ID *string `json:"id,omitempty"`
	Qualification []*OrganizationQualification `json:"qualification,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Contact []*ExtendedContactDetail `json:"contact,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Language *string `json:"language,omitempty"`
}

func (o *OrganizationType) UnmarshalJSON(b []byte) error {
	var safe SafeOrganizationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = OrganizationType{
		Identifier: safe.Identifier,
		Active: safe.Active,
		ModifierExtension: safe.ModifierExtension,
		PartOf: safe.PartOf,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		Description: safe.Description,
		Meta: safe.Meta,
		Name: safe.Name,
		Type: safe.Type,
		ImplicitRules: safe.ImplicitRules,
		ID: safe.ID,
		Qualification: safe.Qualification,
		Alias: safe.Alias,
		Contact: safe.Contact,
		Text: safe.Text,
		Language: safe.Language,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
