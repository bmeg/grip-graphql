
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDiagnosticReportType struct {
	Extension []*Extension `json:"extension,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Performer TypedObject `json:"performer"`
	ResultsInterpreter TypedObject `json:"resultsInterpreter"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Issued *string `json:"issued,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	Media []*DiagnosticReportMedia `json:"media,omitempty"`
	Result *ObservationType `json:"result"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ID *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	SupportingInfo []*DiagnosticReportSupportingInfo `json:"supportingInfo,omitempty"`
	PresentedForm []*Attachment `json:"presentedForm,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Specimen *SpecimenType `json:"specimen"`
	Code *CodeableConcept `json:"code,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Subject TypedObject `json:"subject"`
	Note []*Annotation `json:"note,omitempty"`
	Study TypedObject `json:"study"`
	ResourceType *string `json:"resourceType,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	ConclusionCode []*CodeableConcept `json:"conclusionCode,omitempty"`
	Conclusion *string `json:"conclusion,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DiagnosticReportType) UnmarshalJSON(b []byte) error {
	var safe SafeDiagnosticReportType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DiagnosticReportType{
		Extension: safe.Extension,
		ImplicitRules: safe.ImplicitRules,
		Issued: safe.Issued,
		Identifier: safe.Identifier,
		Language: safe.Language,
		Media: safe.Media,
		Result: safe.Result,
		ModifierExtension: safe.ModifierExtension,
		ID: safe.ID,
		Status: safe.Status,
		SupportingInfo: safe.SupportingInfo,
		PresentedForm: safe.PresentedForm,
		Text: safe.Text,
		Specimen: safe.Specimen,
		Code: safe.Code,
		Note: safe.Note,
		ResourceType: safe.ResourceType,
		Meta: safe.Meta,
		EffectivePeriod: safe.EffectivePeriod,
		Category: safe.Category,
		EffectiveDateTime: safe.EffectiveDateTime,
		ConclusionCode: safe.ConclusionCode,
		Conclusion: safe.Conclusion,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "resultsInterpreter", safe.ResultsInterpreter.Typename, &o.ResultsInterpreter); err != nil {
		return fmt.Errorf("failed to unmarshal ResultsInterpreter: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "study", safe.Study.Typename, &o.Study); err != nil {
		return fmt.Errorf("failed to unmarshal Study: %w", err)
	}

	return nil
}
