
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDiagnosticReportType struct {
	ConclusionCode []*CodeableConcept `json:"conclusionCode,omitempty"`
	PresentedForm []*Attachment `json:"presentedForm,omitempty"`
	ResultsInterpreter TypedObject `json:"resultsInterpreter"`
	Meta *Meta `json:"meta,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ID *string `json:"id,omitempty"`
	Specimen *SpecimenType `json:"specimen"`
	Issued *string `json:"issued,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	Performer TypedObject `json:"performer"`
	Result *ObservationType `json:"result"`
	Subject TypedObject `json:"subject"`
	Contained TypedObject `json:"contained,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Status *string `json:"status,omitempty"`
	Language *string `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Conclusion *string `json:"conclusion,omitempty"`
	Study TypedObject `json:"study"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	SupportingInfo []*DiagnosticReportSupportingInfo `json:"supportingInfo,omitempty"`
	Media []*DiagnosticReportMedia `json:"media,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DiagnosticReportType) UnmarshalJSON(b []byte) error {
	var safe SafeDiagnosticReportType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DiagnosticReportType{
		ConclusionCode: safe.ConclusionCode,
		PresentedForm: safe.PresentedForm,
		Meta: safe.Meta,
		Extension: safe.Extension,
		ResourceType: safe.ResourceType,
		ID: safe.ID,
		Specimen: safe.Specimen,
		Issued: safe.Issued,
		EffectiveDateTime: safe.EffectiveDateTime,
		Result: safe.Result,
		ImplicitRules: safe.ImplicitRules,
		Status: safe.Status,
		Language: safe.Language,
		Text: safe.Text,
		Conclusion: safe.Conclusion,
		EffectivePeriod: safe.EffectivePeriod,
		Identifier: safe.Identifier,
		Code: safe.Code,
		SupportingInfo: safe.SupportingInfo,
		Media: safe.Media,
		Note: safe.Note,
		Category: safe.Category,
		ModifierExtension: safe.ModifierExtension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "resultsInterpreter", safe.ResultsInterpreter.Typename, &o.ResultsInterpreter); err != nil {
		return fmt.Errorf("failed to unmarshal ResultsInterpreter: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "study", safe.Study.Typename, &o.Study); err != nil {
		return fmt.Errorf("failed to unmarshal Study: %w", err)
	}

	return nil
}
