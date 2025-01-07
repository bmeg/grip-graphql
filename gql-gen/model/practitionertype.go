
package model

import (
	"encoding/json"
	"fmt"
)

type SafePractitionerType struct {
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Qualification []*PractitionerQualification `json:"qualification,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ID *string `json:"id,omitempty"`
	Address []*Address `json:"address,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Communication []*PractitionerCommunication `json:"communication,omitempty"`
	Language *string `json:"language,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Active *string `json:"active,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *PractitionerType) UnmarshalJSON(b []byte) error {
	var safe SafePractitionerType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PractitionerType{
		ImplicitRules: safe.ImplicitRules,
		Meta: safe.Meta,
		Gender: safe.Gender,
		Qualification: safe.Qualification,
		ModifierExtension: safe.ModifierExtension,
		ID: safe.ID,
		Address: safe.Address,
		DeceasedBoolean: safe.DeceasedBoolean,
		Identifier: safe.Identifier,
		DeceasedDateTime: safe.DeceasedDateTime,
		ResourceType: safe.ResourceType,
		Communication: safe.Communication,
		Language: safe.Language,
		Name: safe.Name,
		Text: safe.Text,
		BirthDate: safe.BirthDate,
		Telecom: safe.Telecom,
		Extension: safe.Extension,
		Active: safe.Active,
		Photo: safe.Photo,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
