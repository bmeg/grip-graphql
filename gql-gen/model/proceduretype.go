
package model

import (
	"encoding/json"
	"fmt"
)

type SafeProcedureType struct {
	Recorder TypedObject `json:"recorder"`
	Report TypedObject `json:"report"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Note TypedObject `json:"note"`
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	OccurrenceAge *Age `json:"occurrenceAge,omitempty"`
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	ReportedBoolean *string `json:"reportedBoolean,omitempty"`
	Subject TypedObject `json:"subject"`
	OccurrenceRange *Range `json:"occurrenceRange,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Status *string `json:"status,omitempty"`
	ReportedReference TypedObject `json:"reportedReference"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	PartOf TypedObject `json:"partOf"`
	Used []*CodeableReference `json:"used,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	SupportingInfo TypedObject `json:"supportingInfo"`
	Language *string `json:"language,omitempty"`
	FocalDevice []*ProcedureFocalDevice `json:"focalDevice,omitempty"`
	ID *string `json:"id,omitempty"`
	OccurrenceString *string `json:"occurrenceString,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	Performer TypedObject `json:"performer"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	FollowUp []*CodeableConcept `json:"followUp,omitempty"`
	Complication []*CodeableReference `json:"complication,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Focus TypedObject `json:"focus"`
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
}

func (o *ProcedureType) UnmarshalJSON(b []byte) error {
	var safe SafeProcedureType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ProcedureType{
		ImplicitRules: safe.ImplicitRules,
		Meta: safe.Meta,
		BodySite: safe.BodySite,
		InstantiatesURI: safe.InstantiatesURI,
		OccurrenceTiming: safe.OccurrenceTiming,
		OccurrenceAge: safe.OccurrenceAge,
		OccurrenceDateTime: safe.OccurrenceDateTime,
		ReportedBoolean: safe.ReportedBoolean,
		OccurrenceRange: safe.OccurrenceRange,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Text: safe.Text,
		Status: safe.Status,
		Reason: safe.Reason,
		Outcome: safe.Outcome,
		Used: safe.Used,
		Code: safe.Code,
		Identifier: safe.Identifier,
		Category: safe.Category,
		Language: safe.Language,
		FocalDevice: safe.FocalDevice,
		ID: safe.ID,
		OccurrenceString: safe.OccurrenceString,
		Recorded: safe.Recorded,
		StatusReason: safe.StatusReason,
		FollowUp: safe.FollowUp,
		Complication: safe.Complication,
		Extension: safe.Extension,
		OccurrencePeriod: safe.OccurrencePeriod,
		InstantiatesCanonical: safe.InstantiatesCanonical,
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "report", safe.Report.Typename, &o.Report); err != nil {
		return fmt.Errorf("failed to unmarshal Report: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "reportedReference", safe.ReportedReference.Typename, &o.ReportedReference); err != nil {
		return fmt.Errorf("failed to unmarshal ReportedReference: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInfo", safe.SupportingInfo.Typename, &o.SupportingInfo); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInfo: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}

	return nil
}
