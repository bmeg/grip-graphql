
package model

import (
	"encoding/json"
	"fmt"
)

type SafeTaskType struct {
	Meta *Meta `json:"meta,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Description *string `json:"description,omitempty"`
	Input []*TaskInput `json:"input,omitempty"`
	PartOf *TaskType `json:"partOf"`
	Owner TypedObject `json:"owner"`
	Intent *string `json:"intent,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Status *string `json:"status,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Priority *string `json:"priority,omitempty"`
	RequestedPeriod *Period `json:"requestedPeriod,omitempty"`
	DoNotPerform *string `json:"doNotPerform,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Focus TypedObject `json:"focus"`
	Restriction *TaskRestriction `json:"restriction,omitempty"`
	Output []*TaskOutput `json:"output,omitempty"`
	Requester TypedObject `json:"requester"`
	Contained TypedObject `json:"contained,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	BusinessStatus *CodeableConcept `json:"businessStatus,omitempty"`
	StatusReason *CodeableReference `json:"statusReason,omitempty"`
	ExecutionPeriod *Period `json:"executionPeriod,omitempty"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	Performer []*TaskPerformer `json:"performer,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	RequestedPerformer []*CodeableReference `json:"requestedPerformer,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *TaskType) UnmarshalJSON(b []byte) error {
	var safe SafeTaskType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = TaskType{
		Meta: safe.Meta,
		Description: safe.Description,
		Input: safe.Input,
		PartOf: safe.PartOf,
		Intent: safe.Intent,
		Text: safe.Text,
		ImplicitRules: safe.ImplicitRules,
		Status: safe.Status,
		Code: safe.Code,
		Priority: safe.Priority,
		RequestedPeriod: safe.RequestedPeriod,
		DoNotPerform: safe.DoNotPerform,
		Identifier: safe.Identifier,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		InstantiatesURI: safe.InstantiatesURI,
		Note: safe.Note,
		Restriction: safe.Restriction,
		Output: safe.Output,
		GroupIdentifier: safe.GroupIdentifier,
		BusinessStatus: safe.BusinessStatus,
		StatusReason: safe.StatusReason,
		ExecutionPeriod: safe.ExecutionPeriod,
		AuthoredOn: safe.AuthoredOn,
		Performer: safe.Performer,
		Reason: safe.Reason,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		LastModified: safe.LastModified,
		RequestedPerformer: safe.RequestedPerformer,
		ID: safe.ID,
		Language: safe.Language,
		ModifierExtension: safe.ModifierExtension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "owner", safe.Owner.Typename, &o.Owner); err != nil {
		return fmt.Errorf("failed to unmarshal Owner: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
