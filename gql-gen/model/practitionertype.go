
package model

import (
	"encoding/json"
	"fmt"
)

type SafePractitionerType struct {
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Address []*Address `json:"address,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	Qualification []*PractitionerQualification `json:"qualification,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Active *string `json:"active,omitempty"`
	ID *string `json:"id,omitempty"`
	Communication []*PractitionerCommunication `json:"communication,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Language *string `json:"language,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *PractitionerType) UnmarshalJSON(b []byte) error {
	var safe SafePractitionerType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PractitionerType{
		DeceasedDateTime: safe.DeceasedDateTime,
		BirthDate: safe.BirthDate,
		DeceasedBoolean: safe.DeceasedBoolean,
		ModifierExtension: safe.ModifierExtension,
		Meta: safe.Meta,
		Address: safe.Address,
		Text: safe.Text,
		Identifier: safe.Identifier,
		ResourceType: safe.ResourceType,
		Photo: safe.Photo,
		Qualification: safe.Qualification,
		ImplicitRules: safe.ImplicitRules,
		Active: safe.Active,
		ID: safe.ID,
		Communication: safe.Communication,
		Gender: safe.Gender,
		Telecom: safe.Telecom,
		Extension: safe.Extension,
		Language: safe.Language,
		Name: safe.Name,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
