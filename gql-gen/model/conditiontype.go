
package model

import (
	"encoding/json"
	"fmt"
)

type SafeConditionType struct {
	ClinicalStatus *CodeableConcept `json:"clinicalStatus,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	OnsetRange *Range `json:"onsetRange,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Stage []*ConditionStage `json:"stage,omitempty"`
	RecordedDate *string `json:"recordedDate,omitempty"`
	AbatementDateTime *string `json:"abatementDateTime,omitempty"`
	OnsetPeriod *Period `json:"onsetPeriod,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Evidence []*CodeableReference `json:"evidence,omitempty"`
	AbatementRange *Range `json:"abatementRange,omitempty"`
	AbatementPeriod *Period `json:"abatementPeriod,omitempty"`
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	Language *string `json:"language,omitempty"`
	OnsetAge *Age `json:"onsetAge,omitempty"`
	Participant []*ConditionParticipant `json:"participant,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	OnsetDateTime *string `json:"onsetDateTime,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AbatementAge *Age `json:"abatementAge,omitempty"`
	AbatementString *string `json:"abatementString,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Subject TypedObject `json:"subject"`
	ID *string `json:"id,omitempty"`
	OnsetString *string `json:"onsetString,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	VerificationStatus *CodeableConcept `json:"verificationStatus,omitempty"`
	Severity *CodeableConcept `json:"severity,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ConditionType) UnmarshalJSON(b []byte) error {
	var safe SafeConditionType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ConditionType{
		ClinicalStatus: safe.ClinicalStatus,
		Meta: safe.Meta,
		OnsetRange: safe.OnsetRange,
		Identifier: safe.Identifier,
		Stage: safe.Stage,
		RecordedDate: safe.RecordedDate,
		AbatementDateTime: safe.AbatementDateTime,
		OnsetPeriod: safe.OnsetPeriod,
		Extension: safe.Extension,
		Evidence: safe.Evidence,
		AbatementRange: safe.AbatementRange,
		AbatementPeriod: safe.AbatementPeriod,
		BodySite: safe.BodySite,
		Language: safe.Language,
		OnsetAge: safe.OnsetAge,
		Participant: safe.Participant,
		ImplicitRules: safe.ImplicitRules,
		OnsetDateTime: safe.OnsetDateTime,
		ResourceType: safe.ResourceType,
		AbatementAge: safe.AbatementAge,
		AbatementString: safe.AbatementString,
		Code: safe.Code,
		ID: safe.ID,
		OnsetString: safe.OnsetString,
		Category: safe.Category,
		Text: safe.Text,
		ModifierExtension: safe.ModifierExtension,
		Note: safe.Note,
		VerificationStatus: safe.VerificationStatus,
		Severity: safe.Severity,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}

	return nil
}
