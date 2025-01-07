
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchSubjectType struct {
	AssignedComparisonGroup *string `json:"assignedComparisonGroup,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Status *string `json:"status,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ID *string `json:"id,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Subject TypedObject `json:"subject"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Period *Period `json:"period,omitempty"`
	Progress []*ResearchSubjectProgress `json:"progress,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Study *ResearchStudyType `json:"study"`
	ActualComparisonGroup *string `json:"actualComparisonGroup,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchSubjectType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchSubjectType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchSubjectType{
		AssignedComparisonGroup: safe.AssignedComparisonGroup,
		Status: safe.Status,
		Identifier: safe.Identifier,
		Language: safe.Language,
		Text: safe.Text,
		ID: safe.ID,
		Extension: safe.Extension,
		ImplicitRules: safe.ImplicitRules,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Period: safe.Period,
		Progress: safe.Progress,
		Meta: safe.Meta,
		Study: safe.Study,
		ActualComparisonGroup: safe.ActualComparisonGroup,
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
