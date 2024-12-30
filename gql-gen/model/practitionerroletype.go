
package model

import (
	"encoding/json"
	"fmt"
)

type SafePractitionerRoleType struct {
	Communication []*CodeableConcept `json:"communication,omitempty"`
	Active *string `json:"active,omitempty"`
	ID *string `json:"id,omitempty"`
	Organization *OrganizationType `json:"organization"`
	Contained TypedObject `json:"contained,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Period *Period `json:"period,omitempty"`
	Code []*CodeableConcept `json:"code,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Specialty []*CodeableConcept `json:"specialty,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Language *string `json:"language,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Contact []*ExtendedContactDetail `json:"contact,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Practitioner *PractitionerType `json:"practitioner"`
	Availability []*Availability `json:"availability,omitempty"`
	Characteristic []*CodeableConcept `json:"characteristic,omitempty"`
}

func (o *PractitionerRoleType) UnmarshalJSON(b []byte) error {
	var safe SafePractitionerRoleType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PractitionerRoleType{
		Communication: safe.Communication,
		Active: safe.Active,
		ID: safe.ID,
		Organization: safe.Organization,
		Extension: safe.Extension,
		Period: safe.Period,
		Code: safe.Code,
		ModifierExtension: safe.ModifierExtension,
		Specialty: safe.Specialty,
		ResourceType: safe.ResourceType,
		Identifier: safe.Identifier,
		ImplicitRules: safe.ImplicitRules,
		Language: safe.Language,
		Meta: safe.Meta,
		Contact: safe.Contact,
		Text: safe.Text,
		Practitioner: safe.Practitioner,
		Availability: safe.Availability,
		Characteristic: safe.Characteristic,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
