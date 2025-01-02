
package model

import (
	"encoding/json"
	"fmt"
)

type SafePractitionerRoleType struct {
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Code []*CodeableConcept `json:"code,omitempty"`
	Availability []*Availability `json:"availability,omitempty"`
	Practitioner *PractitionerType `json:"practitioner"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ID *string `json:"id,omitempty"`
	Active *string `json:"active,omitempty"`
	Communication []*CodeableConcept `json:"communication,omitempty"`
	Language *string `json:"language,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Characteristic []*CodeableConcept `json:"characteristic,omitempty"`
	Specialty []*CodeableConcept `json:"specialty,omitempty"`
	Contact []*ExtendedContactDetail `json:"contact,omitempty"`
	Period *Period `json:"period,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Organization *OrganizationType `json:"organization"`
	Meta *Meta `json:"meta,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *PractitionerRoleType) UnmarshalJSON(b []byte) error {
	var safe SafePractitionerRoleType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PractitionerRoleType{
		ImplicitRules: safe.ImplicitRules,
		Code: safe.Code,
		Availability: safe.Availability,
		Practitioner: safe.Practitioner,
		Identifier: safe.Identifier,
		ID: safe.ID,
		Active: safe.Active,
		Communication: safe.Communication,
		Language: safe.Language,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Characteristic: safe.Characteristic,
		Specialty: safe.Specialty,
		Contact: safe.Contact,
		Period: safe.Period,
		Text: safe.Text,
		Extension: safe.Extension,
		Organization: safe.Organization,
		Meta: safe.Meta,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
