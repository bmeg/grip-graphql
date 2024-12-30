
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDiagnosticReportType struct {
	SupportingInfo TypedObject `json:"supportingInfo"`
	Specimen *SpecimenType `json:"specimen"`
	Contained TypedObject `json:"contained,omitempty"`
	Language *string `json:"language,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ResultsInterpreter TypedObject `json:"resultsInterpreter"`
	ID *string `json:"id,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Note TypedObject `json:"note"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	PresentedForm []*Attachment `json:"presentedForm,omitempty"`
	Subject TypedObject `json:"subject"`
	Study TypedObject `json:"study"`
	Media []*DiagnosticReportMedia `json:"media,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ConclusionCode []*CodeableConcept `json:"conclusionCode,omitempty"`
	Result *ObservationType `json:"result"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Issued *string `json:"issued,omitempty"`
	Performer TypedObject `json:"performer"`
	Status *string `json:"status,omitempty"`
	Conclusion *string `json:"conclusion,omitempty"`
}

func (o *DiagnosticReportType) UnmarshalJSON(b []byte) error {
	var safe SafeDiagnosticReportType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DiagnosticReportType{
		Specimen: safe.Specimen,
		Language: safe.Language,
		Meta: safe.Meta,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		EffectiveDateTime: safe.EffectiveDateTime,
		Category: safe.Category,
		Identifier: safe.Identifier,
		EffectivePeriod: safe.EffectivePeriod,
		ResourceType: safe.ResourceType,
		Code: safe.Code,
		Extension: safe.Extension,
		PresentedForm: safe.PresentedForm,
		Media: safe.Media,
		Text: safe.Text,
		ConclusionCode: safe.ConclusionCode,
		Result: safe.Result,
		ImplicitRules: safe.ImplicitRules,
		Issued: safe.Issued,
		Status: safe.Status,
		Conclusion: safe.Conclusion,
	}
	if err := unmarshalUnion(b, "supportingInfo", safe.SupportingInfo.Typename, &o.SupportingInfo); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInfo: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "resultsInterpreter", safe.ResultsInterpreter.Typename, &o.ResultsInterpreter); err != nil {
		return fmt.Errorf("failed to unmarshal ResultsInterpreter: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "study", safe.Study.Typename, &o.Study); err != nil {
		return fmt.Errorf("failed to unmarshal Study: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}

	return nil
}
