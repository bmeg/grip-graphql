
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDiagnosticReportType struct {
	Performer TypedObject `json:"performer"`
	BasedOn TypedObject `json:"basedOn"`
	ResourceType *string `json:"resourceType,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	ID *string `json:"id,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Conclusion *string `json:"conclusion,omitempty"`
	Note TypedObject `json:"note"`
	ConclusionCode []*CodeableConcept `json:"conclusionCode,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Subject TypedObject `json:"subject"`
	Study TypedObject `json:"study"`
	Status *string `json:"status,omitempty"`
	Language *string `json:"language,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	Issued *string `json:"issued,omitempty"`
	SupportingInfo TypedObject `json:"supportingInfo"`
	Contained TypedObject `json:"contained,omitempty"`
	Specimen *SpecimenType `json:"specimen"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	Media []*DiagnosticReportMedia `json:"media,omitempty"`
	ResultsInterpreter TypedObject `json:"resultsInterpreter"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Result *ObservationType `json:"result"`
	Code *CodeableConcept `json:"code,omitempty"`
	PresentedForm []*Attachment `json:"presentedForm,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DiagnosticReportType) UnmarshalJSON(b []byte) error {
	var safe SafeDiagnosticReportType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DiagnosticReportType{
		ResourceType: safe.ResourceType,
		Category: safe.Category,
		ID: safe.ID,
		Meta: safe.Meta,
		Text: safe.Text,
		Conclusion: safe.Conclusion,
		ConclusionCode: safe.ConclusionCode,
		ModifierExtension: safe.ModifierExtension,
		Status: safe.Status,
		Language: safe.Language,
		ImplicitRules: safe.ImplicitRules,
		Extension: safe.Extension,
		EffectiveDateTime: safe.EffectiveDateTime,
		Issued: safe.Issued,
		Specimen: safe.Specimen,
		EffectivePeriod: safe.EffectivePeriod,
		Media: safe.Media,
		Identifier: safe.Identifier,
		Result: safe.Result,
		Code: safe.Code,
		PresentedForm: safe.PresentedForm,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
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
	if err := unmarshalUnion(b, "supportingInfo", safe.SupportingInfo.Typename, &o.SupportingInfo); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInfo: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "resultsInterpreter", safe.ResultsInterpreter.Typename, &o.ResultsInterpreter); err != nil {
		return fmt.Errorf("failed to unmarshal ResultsInterpreter: %w", err)
	}

	return nil
}
