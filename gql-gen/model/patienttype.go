
package model

import (
	"encoding/json"
	"fmt"
)

type SafePatientType struct {
	Language *string `json:"language,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Contact []*PatientContact `json:"contact,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ManagingOrganization *OrganizationType `json:"managingOrganization"`
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	Name []*HumanName `json:"name,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
	MultipleBirthBoolean *string `json:"multipleBirthBoolean,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	GeneralPractitioner TypedObject `json:"generalPractitioner"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Active *string `json:"active,omitempty"`
	Address []*Address `json:"address,omitempty"`
	MultipleBirthInteger *string `json:"multipleBirthInteger,omitempty"`
	Communication []*PatientCommunication `json:"communication,omitempty"`
	Link TypedObject `json:"link"`
	Telecom []*ContactPoint `json:"telecom,omitempty"`
	Gender *string `json:"gender,omitempty"`
	ID *string `json:"id,omitempty"`
	Photo []*Attachment `json:"photo,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *PatientType) UnmarshalJSON(b []byte) error {
	var safe SafePatientType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = PatientType{
		Language: safe.Language,
		Contact: safe.Contact,
		Extension: safe.Extension,
		Text: safe.Text,
		DeceasedBoolean: safe.DeceasedBoolean,
		Meta: safe.Meta,
		ManagingOrganization: safe.ManagingOrganization,
		MaritalStatus: safe.MaritalStatus,
		Name: safe.Name,
		DeceasedDateTime: safe.DeceasedDateTime,
		BirthDate: safe.BirthDate,
		MultipleBirthBoolean: safe.MultipleBirthBoolean,
		Identifier: safe.Identifier,
		ModifierExtension: safe.ModifierExtension,
		ImplicitRules: safe.ImplicitRules,
		ResourceType: safe.ResourceType,
		Active: safe.Active,
		Address: safe.Address,
		MultipleBirthInteger: safe.MultipleBirthInteger,
		Communication: safe.Communication,
		Telecom: safe.Telecom,
		Gender: safe.Gender,
		ID: safe.ID,
		Photo: safe.Photo,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "generalPractitioner", safe.GeneralPractitioner.Typename, &o.GeneralPractitioner); err != nil {
		return fmt.Errorf("failed to unmarshal GeneralPractitioner: %w", err)
	}
	if err := unmarshalUnion(b, "link", safe.Link.Typename, &o.Link); err != nil {
		return fmt.Errorf("failed to unmarshal Link: %w", err)
	}

	return nil
}
