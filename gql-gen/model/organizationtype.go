
package model

import (
	"encoding/json"
	"fmt"
)

type SafeOrganizationType struct {
	Contact []*ExtendedContactDetail `json:"contact,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Name *string `json:"name,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Qualification []*OrganizationQualification `json:"qualification,omitempty"`
	PartOf *OrganizationType `json:"partOf"`
	Type []*CodeableConcept `json:"type,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Description *string `json:"description,omitempty"`
	Active *string `json:"active,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *OrganizationType) UnmarshalJSON(b []byte) error {
	var safe SafeOrganizationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = OrganizationType{
		Contact: safe.Contact,
		Extension: safe.Extension,
		Identifier: safe.Identifier,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		ID: safe.ID,
		Language: safe.Language,
		ImplicitRules: safe.ImplicitRules,
		Meta: safe.Meta,
		Name: safe.Name,
		Text: safe.Text,
		Qualification: safe.Qualification,
		PartOf: safe.PartOf,
		Type: safe.Type,
		Alias: safe.Alias,
		Description: safe.Description,
		Active: safe.Active,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
