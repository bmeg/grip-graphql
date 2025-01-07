
package model

import (
	"encoding/json"
	"fmt"
)

type SafeConditionType struct {
	Contained TypedObject `json:"contained,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Participant []*ConditionParticipant `json:"participant,omitempty"`
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AbatementDateTime *string `json:"abatementDateTime,omitempty"`
	Severity *CodeableConcept `json:"severity,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	RecordedDate *string `json:"recordedDate,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Stage []*ConditionStage `json:"stage,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	OnsetPeriod *Period `json:"onsetPeriod,omitempty"`
	Subject TypedObject `json:"subject"`
	OnsetDateTime *string `json:"onsetDateTime,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Evidence []*CodeableReference `json:"evidence,omitempty"`
	ClinicalStatus *CodeableConcept `json:"clinicalStatus,omitempty"`
	OnsetAge *Age `json:"onsetAge,omitempty"`
	ID *string `json:"id,omitempty"`
	OnsetString *string `json:"onsetString,omitempty"`
	AbatementRange *Range `json:"abatementRange,omitempty"`
	AbatementPeriod *Period `json:"abatementPeriod,omitempty"`
	AbatementAge *Age `json:"abatementAge,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	VerificationStatus *CodeableConcept `json:"verificationStatus,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Language *string `json:"language,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	AbatementString *string `json:"abatementString,omitempty"`
	OnsetRange *Range `json:"onsetRange,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ConditionType) UnmarshalJSON(b []byte) error {
	var safe SafeConditionType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ConditionType{
		Meta: safe.Meta,
		Participant: safe.Participant,
		BodySite: safe.BodySite,
		ModifierExtension: safe.ModifierExtension,
		AbatementDateTime: safe.AbatementDateTime,
		Severity: safe.Severity,
		Code: safe.Code,
		RecordedDate: safe.RecordedDate,
		Text: safe.Text,
		ResourceType: safe.ResourceType,
		Stage: safe.Stage,
		Extension: safe.Extension,
		OnsetPeriod: safe.OnsetPeriod,
		OnsetDateTime: safe.OnsetDateTime,
		ImplicitRules: safe.ImplicitRules,
		Evidence: safe.Evidence,
		ClinicalStatus: safe.ClinicalStatus,
		OnsetAge: safe.OnsetAge,
		ID: safe.ID,
		OnsetString: safe.OnsetString,
		AbatementRange: safe.AbatementRange,
		AbatementPeriod: safe.AbatementPeriod,
		AbatementAge: safe.AbatementAge,
		Identifier: safe.Identifier,
		VerificationStatus: safe.VerificationStatus,
		Category: safe.Category,
		Language: safe.Language,
		Note: safe.Note,
		AbatementString: safe.AbatementString,
		OnsetRange: safe.OnsetRange,
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
