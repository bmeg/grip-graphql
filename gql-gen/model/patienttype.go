
package model

import (
	"encoding/json"
	"fmt"
)

type SafePatientType struct {
	ManagingOrganization *OrganizationType `json:"managingOrganization"`
	Contact []*PatientContact `json:"contact,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Link TypedObject `json:"link"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	MultipleBirthInteger *string `json:"multipleBirthInteger,omitempty"`
	ID *string `json:"id,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Communication []*PatientCommunication `json:"communication,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	GeneralPractitioner TypedObject `json:"generalPractitioner"`
	Language *string `json:"language,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	MultipleBirthBoolean *string `json:"multipleBirthBoolean,omitempty"`
	Address []*Address `json:"address,omitempty"`
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	Active *string `json:"active,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
}

func (o *PatientType) UnmarshalJSON(b []byte) error {
	var safe SafePatientType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PatientType{
		ManagingOrganization: safe.ManagingOrganization,
		Contact: safe.Contact,
		Gender: safe.Gender,
		ImplicitRules: safe.ImplicitRules,
		MultipleBirthInteger: safe.MultipleBirthInteger,
		ID: safe.ID,
		Meta: safe.Meta,
		Communication: safe.Communication,
		BirthDate: safe.BirthDate,
		Identifier: safe.Identifier,
		ResourceType: safe.ResourceType,
		Language: safe.Language,
		DeceasedDateTime: safe.DeceasedDateTime,
		Telecom: safe.Telecom,
		Photo: safe.Photo,
		Text: safe.Text,
		MultipleBirthBoolean: safe.MultipleBirthBoolean,
		Address: safe.Address,
		MaritalStatus: safe.MaritalStatus,
		ModifierExtension: safe.ModifierExtension,
		Extension: safe.Extension,
		Name: safe.Name,
		Active: safe.Active,
		DeceasedBoolean: safe.DeceasedBoolean,
	}
	if err := unmarshalUnion(b, "link", safe.Link.Typename, &o.Link); err != nil {
		return fmt.Errorf("failed to unmarshal Link: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "generalPractitioner", safe.GeneralPractitioner.Typename, &o.GeneralPractitioner); err != nil {
		return fmt.Errorf("failed to unmarshal GeneralPractitioner: %w", err)
	}

	return nil
}
