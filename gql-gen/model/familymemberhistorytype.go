
package model

import (
	"encoding/json"
	"fmt"
)

type SafeFamilyMemberHistoryType struct {
	Sex *CodeableConcept `json:"sex,omitempty"`
	Status *string `json:"status,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ID *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AgeString *string `json:"ageString,omitempty"`
	DeceasedRange *Range `json:"deceasedRange,omitempty"`
	BornString *string `json:"bornString,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	AgeAge *Age `json:"ageAge,omitempty"`
	Name *string `json:"name,omitempty"`
	Relationship *CodeableConcept `json:"relationship,omitempty"`
	Condition []*FamilyMemberHistoryCondition `json:"condition,omitempty"`
	BornDate *string `json:"bornDate,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	Note TypedObject `json:"note"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Procedure []*FamilyMemberHistoryProcedure `json:"procedure,omitempty"`
	Participant TypedObject `json:"participant"`
	BornPeriod *Period `json:"bornPeriod,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	AgeRange *Range `json:"ageRange,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	EstimatedAge *string `json:"estimatedAge,omitempty"`
	Date *string `json:"date,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Patient *PatientType `json:"patient"`
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	DeceasedString *string `json:"deceasedString,omitempty"`
	DeceasedDate *string `json:"deceasedDate,omitempty"`
	DeceasedAge *Age `json:"deceasedAge,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *FamilyMemberHistoryType) UnmarshalJSON(b []byte) error {
	var safe SafeFamilyMemberHistoryType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = FamilyMemberHistoryType{
		Sex: safe.Sex,
		Status: safe.Status,
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		AgeString: safe.AgeString,
		DeceasedRange: safe.DeceasedRange,
		BornString: safe.BornString,
		Reason: safe.Reason,
		Identifier: safe.Identifier,
		Language: safe.Language,
		AgeAge: safe.AgeAge,
		Name: safe.Name,
		Relationship: safe.Relationship,
		Condition: safe.Condition,
		BornDate: safe.BornDate,
		DeceasedBoolean: safe.DeceasedBoolean,
		InstantiatesURI: safe.InstantiatesURI,
		ModifierExtension: safe.ModifierExtension,
		Extension: safe.Extension,
		Procedure: safe.Procedure,
		BornPeriod: safe.BornPeriod,
		Text: safe.Text,
		AgeRange: safe.AgeRange,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		Meta: safe.Meta,
		EstimatedAge: safe.EstimatedAge,
		Date: safe.Date,
		ImplicitRules: safe.ImplicitRules,
		Patient: safe.Patient,
		DataAbsentReason: safe.DataAbsentReason,
		DeceasedString: safe.DeceasedString,
		DeceasedDate: safe.DeceasedDate,
		DeceasedAge: safe.DeceasedAge,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "participant", safe.Participant.Typename, &o.Participant); err != nil {
		return fmt.Errorf("failed to unmarshal Participant: %w", err)
	}

	return nil
}
