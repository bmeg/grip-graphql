
package model

import (
	"encoding/json"
	"fmt"
)

type SafeFamilyMemberHistoryType struct {
	AgeRange *Range `json:"ageRange,omitempty"`
	Patient *PatientType `json:"patient"`
	Name *string `json:"name,omitempty"`
	DeceasedBoolean *string `json:"deceasedBoolean,omitempty"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ID *string `json:"id,omitempty"`
	BornString *string `json:"bornString,omitempty"`
	DeceasedAge *Age `json:"deceasedAge,omitempty"`
	AgeAge *Age `json:"ageAge,omitempty"`
	Date *string `json:"date,omitempty"`
	Sex *CodeableConcept `json:"sex,omitempty"`
	Language *string `json:"language,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Note TypedObject `json:"note"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Relationship *CodeableConcept `json:"relationship,omitempty"`
	Condition []*FamilyMemberHistoryCondition `json:"condition,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	DeceasedString *string `json:"deceasedString,omitempty"`
	Status *string `json:"status,omitempty"`
	AgeString *string `json:"ageString,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Participant TypedObject `json:"participant"`
	DeceasedDate *string `json:"deceasedDate,omitempty"`
	EstimatedAge *string `json:"estimatedAge,omitempty"`
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	Procedure []*FamilyMemberHistoryProcedure `json:"procedure,omitempty"`
	BornPeriod *Period `json:"bornPeriod,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	DeceasedRange *Range `json:"deceasedRange,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	BornDate *string `json:"bornDate,omitempty"`
}

func (o *FamilyMemberHistoryType) UnmarshalJSON(b []byte) error {
	var safe SafeFamilyMemberHistoryType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = FamilyMemberHistoryType{
		AgeRange: safe.AgeRange,
		Patient: safe.Patient,
		Name: safe.Name,
		DeceasedBoolean: safe.DeceasedBoolean,
		InstantiatesURI: safe.InstantiatesURI,
		ModifierExtension: safe.ModifierExtension,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		Text: safe.Text,
		ID: safe.ID,
		BornString: safe.BornString,
		DeceasedAge: safe.DeceasedAge,
		AgeAge: safe.AgeAge,
		Date: safe.Date,
		Sex: safe.Sex,
		Language: safe.Language,
		Reason: safe.Reason,
		Relationship: safe.Relationship,
		Condition: safe.Condition,
		Extension: safe.Extension,
		Meta: safe.Meta,
		DeceasedString: safe.DeceasedString,
		Status: safe.Status,
		AgeString: safe.AgeString,
		Identifier: safe.Identifier,
		DeceasedDate: safe.DeceasedDate,
		EstimatedAge: safe.EstimatedAge,
		DataAbsentReason: safe.DataAbsentReason,
		Procedure: safe.Procedure,
		BornPeriod: safe.BornPeriod,
		ImplicitRules: safe.ImplicitRules,
		DeceasedRange: safe.DeceasedRange,
		ResourceType: safe.ResourceType,
		BornDate: safe.BornDate,
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
