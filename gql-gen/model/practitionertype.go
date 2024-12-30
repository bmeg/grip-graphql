
package model

import (
	"encoding/json"
	"fmt"
)

type SafePractitionerType struct {
	BirthDate *string `json:"birthDate,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	Qualification []*PractitionerQualification `json:"qualification,omitempty"`
	Address []*Address `json:"address,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Active *string `json:"active,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Communication []*PractitionerCommunication `json:"communication,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
}

func (o *PractitionerType) UnmarshalJSON(b []byte) error {
	var safe SafePractitionerType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PractitionerType{
		BirthDate: safe.BirthDate,
		ID: safe.ID,
		Language: safe.Language,
		Qualification: safe.Qualification,
		Address: safe.Address,
		ModifierExtension: safe.ModifierExtension,
		Active: safe.Active,
		Gender: safe.Gender,
		Name: safe.Name,
		Text: safe.Text,
		DeceasedDateTime: safe.DeceasedDateTime,
		Extension: safe.Extension,
		Meta: safe.Meta,
		Photo: safe.Photo,
		Telecom: safe.Telecom,
		ResourceType: safe.ResourceType,
		Communication: safe.Communication,
		Identifier: safe.Identifier,
		DeceasedBoolean: safe.DeceasedBoolean,
		ImplicitRules: safe.ImplicitRules,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
