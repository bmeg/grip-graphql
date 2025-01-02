
package model

import (
	"encoding/json"
	"fmt"
)

type SafeConditionType struct {
	OnsetString *string `json:"onsetString,omitempty"`
	OnsetPeriod *Period `json:"onsetPeriod,omitempty"`
	OnsetDateTime *string `json:"onsetDateTime,omitempty"`
	Severity *CodeableConcept `json:"severity,omitempty"`
	Note TypedObject `json:"note"`
	Subject TypedObject `json:"subject"`
	AbatementAge *Age `json:"abatementAge,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	AbatementDateTime *string `json:"abatementDateTime,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Stage TypedObject `json:"stage"`
	AbatementString *string `json:"abatementString,omitempty"`
	OnsetRange *Range `json:"onsetRange,omitempty"`
	Participant TypedObject `json:"participant"`
	ID *string `json:"id,omitempty"`
	Evidence []*CodeableReference `json:"evidence,omitempty"`
	AbatementPeriod *Period `json:"abatementPeriod,omitempty"`
	ClinicalStatus *CodeableConcept `json:"clinicalStatus,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	AbatementRange *Range `json:"abatementRange,omitempty"`
	RecordedDate *string `json:"recordedDate,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	OnsetAge *Age `json:"onsetAge,omitempty"`
	VerificationStatus *CodeableConcept `json:"verificationStatus,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ConditionType) UnmarshalJSON(b []byte) error {
	var safe SafeConditionType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ConditionType{
		OnsetString: safe.OnsetString,
		OnsetPeriod: safe.OnsetPeriod,
		OnsetDateTime: safe.OnsetDateTime,
		Severity: safe.Severity,
		AbatementAge: safe.AbatementAge,
		Meta: safe.Meta,
		AbatementDateTime: safe.AbatementDateTime,
		Category: safe.Category,
		AbatementString: safe.AbatementString,
		OnsetRange: safe.OnsetRange,
		ID: safe.ID,
		Evidence: safe.Evidence,
		AbatementPeriod: safe.AbatementPeriod,
		ClinicalStatus: safe.ClinicalStatus,
		ModifierExtension: safe.ModifierExtension,
		Text: safe.Text,
		BodySite: safe.BodySite,
		AbatementRange: safe.AbatementRange,
		RecordedDate: safe.RecordedDate,
		ResourceType: safe.ResourceType,
		ImplicitRules: safe.ImplicitRules,
		OnsetAge: safe.OnsetAge,
		VerificationStatus: safe.VerificationStatus,
		Extension: safe.Extension,
		Identifier: safe.Identifier,
		Language: safe.Language,
		Code: safe.Code,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "stage", safe.Stage.Typename, &o.Stage); err != nil {
		return fmt.Errorf("failed to unmarshal Stage: %w", err)
	}
	if err := unmarshalUnion(b, "participant", safe.Participant.Typename, &o.Participant); err != nil {
		return fmt.Errorf("failed to unmarshal Participant: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
