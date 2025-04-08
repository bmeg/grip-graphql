
package model

import (
	"encoding/json"
	"fmt"
)

type SafePatientType struct {
	Language *string `json:"language,omitempty"`
	DeceasedBoolean *bool `json:"deceasedBoolean,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Address []*Address `json:"address,omitempty"`
	Communication []*PatientCommunication `json:"communication,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Active *bool `json:"active,omitempty"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Contact []*PatientContact `json:"contact,omitempty"`
	ID *string `json:"id,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	MultipleBirthInteger *int `json:"multipleBirthInteger,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`
	GeneralPractitioner TypedObject `json:"generalPractitioner"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`
	ManagingOrganization *OrganizationType `json:"managingOrganization"`
	Meta *Meta `json:"meta,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *PatientType) UnmarshalJSON(b []byte) error {
	var safe SafePatientType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PatientType{
		Language: safe.Language,
		DeceasedBoolean: safe.DeceasedBoolean,
		Identifier: safe.Identifier,
		Address: safe.Address,
		Communication: safe.Communication,
		Gender: safe.Gender,
		Name: safe.Name,
		Extension: safe.Extension,
		ImplicitRules: safe.ImplicitRules,
		Text: safe.Text,
		MaritalStatus: safe.MaritalStatus,
		ResourceType: safe.ResourceType,
		Active: safe.Active,
		Telecom: safe.Telecom,
		Contact: safe.Contact,
		ID: safe.ID,
		Photo: safe.Photo,
		MultipleBirthInteger: safe.MultipleBirthInteger,
		BirthDate: safe.BirthDate,
		ModifierExtension: safe.ModifierExtension,
		MultipleBirthBoolean: safe.MultipleBirthBoolean,
		DeceasedDateTime: safe.DeceasedDateTime,
		ManagingOrganization: safe.ManagingOrganization,
		Meta: safe.Meta,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "generalPractitioner", safe.GeneralPractitioner.Typename, &o.GeneralPractitioner); err != nil {
		return fmt.Errorf("failed to unmarshal GeneralPractitioner: %w", err)
	}

	return nil
}
