
package model

import (
	"encoding/json"
	"fmt"
)

type SafeProcedureType struct {
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	Subject TypedObject `json:"subject"`
	Code *CodeableConcept `json:"code,omitempty"`
	ReportedBoolean *string `json:"reportedBoolean,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Status *string `json:"status,omitempty"`
	OccurrenceRange *Range `json:"occurrenceRange,omitempty"`
	Used []*CodeableReference `json:"used,omitempty"`
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	ID *string `json:"id,omitempty"`
	BodySite []*CodeableConcept `json:"bodySite,omitempty"`
	OccurrenceString *string `json:"occurrenceString,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	Recorder TypedObject `json:"recorder"`
	Focus TypedObject `json:"focus"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ReportedReference TypedObject `json:"reportedReference"`
	PartOf TypedObject `json:"partOf"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	Note TypedObject `json:"note"`
	Extension []*Extension `json:"extension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	OccurrenceAge *Age `json:"occurrenceAge,omitempty"`
	Report TypedObject `json:"report"`
	FocalDevice []*ProcedureFocalDevice `json:"focalDevice,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Language *string `json:"language,omitempty"`
	FollowUp []*CodeableConcept `json:"followUp,omitempty"`
	SupportingInfo TypedObject `json:"supportingInfo"`
	Complication []*CodeableReference `json:"complication,omitempty"`
	Performer TypedObject `json:"performer"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ProcedureType) UnmarshalJSON(b []byte) error {
	var safe SafeProcedureType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ProcedureType{
		OccurrenceTiming: safe.OccurrenceTiming,
		Code: safe.Code,
		ReportedBoolean: safe.ReportedBoolean,
		Meta: safe.Meta,
		Status: safe.Status,
		OccurrenceRange: safe.OccurrenceRange,
		Used: safe.Used,
		OccurrenceDateTime: safe.OccurrenceDateTime,
		ID: safe.ID,
		BodySite: safe.BodySite,
		OccurrenceString: safe.OccurrenceString,
		Recorded: safe.Recorded,
		ResourceType: safe.ResourceType,
		InstantiatesURI: safe.InstantiatesURI,
		Outcome: safe.Outcome,
		ImplicitRules: safe.ImplicitRules,
		ModifierExtension: safe.ModifierExtension,
		OccurrencePeriod: safe.OccurrencePeriod,
		Identifier: safe.Identifier,
		StatusReason: safe.StatusReason,
		Extension: safe.Extension,
		Text: safe.Text,
		OccurrenceAge: safe.OccurrenceAge,
		FocalDevice: safe.FocalDevice,
		Category: safe.Category,
		Reason: safe.Reason,
		Language: safe.Language,
		FollowUp: safe.FollowUp,
		Complication: safe.Complication,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}
	if err := unmarshalUnion(b, "reportedReference", safe.ReportedReference.Typename, &o.ReportedReference); err != nil {
		return fmt.Errorf("failed to unmarshal ReportedReference: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "report", safe.Report.Typename, &o.Report); err != nil {
		return fmt.Errorf("failed to unmarshal Report: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInfo", safe.SupportingInfo.Typename, &o.SupportingInfo); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInfo: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}

	return nil
}
