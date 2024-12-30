
package model

import (
	"encoding/json"
	"fmt"
)

type SafeConditionType struct {
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	VerificationStatus *CodeableConcept `json:"verificationStatus,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	OnsetRange *Range `json:"onsetRange,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Language *string `json:"language,omitempty"`
	AbatementAge *Age `json:"abatementAge,omitempty"`
	ID *string `json:"id,omitempty"`
	OnsetAge *Age `json:"onsetAge,omitempty"`
	ClinicalStatus *CodeableConcept `json:"clinicalStatus,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AbatementPeriod *Period `json:"abatementPeriod,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	RecordedDate *string `json:"recordedDate,omitempty"`
	Stage TypedObject `json:"stage"`
	OnsetDateTime *string `json:"onsetDateTime,omitempty"`
	OnsetString *string `json:"onsetString,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Subject TypedObject `json:"subject"`
	Code *CodeableConcept `json:"code,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Note TypedObject `json:"note"`
	Evidence []*CodeableReference `json:"evidence,omitempty"`
	OnsetPeriod *Period `json:"onsetPeriod,omitempty"`
	AbatementString *string `json:"abatementString,omitempty"`
	Severity *CodeableConcept `json:"severity,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	AbatementDateTime *string `json:"abatementDateTime,omitempty"`
	AbatementRange *Range `json:"abatementRange,omitempty"`
	Participant TypedObject `json:"participant"`
	Extension []*Extension `json:"extension,omitempty"`
}

func (o *ConditionType) UnmarshalJSON(b []byte) error {
	var safe SafeConditionType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ConditionType{
		BodySite: safe.BodySite,
		VerificationStatus: safe.VerificationStatus,
		Identifier: safe.Identifier,
		OnsetRange: safe.OnsetRange,
		Category: safe.Category,
		Language: safe.Language,
		AbatementAge: safe.AbatementAge,
		ID: safe.ID,
		OnsetAge: safe.OnsetAge,
		ClinicalStatus: safe.ClinicalStatus,
		ModifierExtension: safe.ModifierExtension,
		AbatementPeriod: safe.AbatementPeriod,
		ImplicitRules: safe.ImplicitRules,
		RecordedDate: safe.RecordedDate,
		OnsetDateTime: safe.OnsetDateTime,
		OnsetString: safe.OnsetString,
		Meta: safe.Meta,
		Code: safe.Code,
		ResourceType: safe.ResourceType,
		Evidence: safe.Evidence,
		OnsetPeriod: safe.OnsetPeriod,
		AbatementString: safe.AbatementString,
		Severity: safe.Severity,
		Text: safe.Text,
		AbatementDateTime: safe.AbatementDateTime,
		AbatementRange: safe.AbatementRange,
		Extension: safe.Extension,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "stage", safe.Stage.Typename, &o.Stage); err != nil {
		return fmt.Errorf("failed to unmarshal Stage: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "participant", safe.Participant.Typename, &o.Participant); err != nil {
		return fmt.Errorf("failed to unmarshal Participant: %w", err)
	}

	return nil
}
