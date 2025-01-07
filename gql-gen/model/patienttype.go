
package model

import (
	"encoding/json"
	"fmt"
)

type SafePatientType struct {
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
	Communication []*PatientCommunication `json:"communication,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`
	Language *string `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	GeneralPractitioner TypedObject `json:"generalPractitioner"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ManagingOrganization *OrganizationType `json:"managingOrganization"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	Active *string `json:"active,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Address []*Address `json:"address,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	MultipleBirthBoolean *string `json:"multipleBirthBoolean,omitempty"`
	Contact []*PatientContact `json:"contact,omitempty"`
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	MultipleBirthInteger *string `json:"multipleBirthInteger,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *PatientType) UnmarshalJSON(b []byte) error {
	var safe SafePatientType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PatientType{
		Telecom: safe.Telecom,
		DeceasedBoolean: safe.DeceasedBoolean,
		Communication: safe.Communication,
		Extension: safe.Extension,
		Name: safe.Name,
		BirthDate: safe.BirthDate,
		DeceasedDateTime: safe.DeceasedDateTime,
		Language: safe.Language,
		Text: safe.Text,
		Identifier: safe.Identifier,
		ManagingOrganization: safe.ManagingOrganization,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		ImplicitRules: safe.ImplicitRules,
		Photo: safe.Photo,
		Active: safe.Active,
		Address: safe.Address,
		Gender: safe.Gender,
		Meta: safe.Meta,
		MultipleBirthBoolean: safe.MultipleBirthBoolean,
		Contact: safe.Contact,
		MaritalStatus: safe.MaritalStatus,
		ResourceType: safe.ResourceType,
		MultipleBirthInteger: safe.MultipleBirthInteger,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "generalPractitioner", safe.GeneralPractitioner.Typename, &o.GeneralPractitioner); err != nil {
		return fmt.Errorf("failed to unmarshal GeneralPractitioner: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
