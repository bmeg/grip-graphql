
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchSubjectType struct {
	Meta *Meta `json:"meta,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Subject TypedObject `json:"subject"`
	AssignedComparisonGroup *string `json:"assignedComparisonGroup,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Period *Period `json:"period,omitempty"`
	Study *ResearchStudyType `json:"study"`
	Text *Narrative `json:"text,omitempty"`
	Language *string `json:"language,omitempty"`
	ID *string `json:"id,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ActualComparisonGroup *string `json:"actualComparisonGroup,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Status *string `json:"status,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Progress []*ResearchSubjectProgress `json:"progress,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchSubjectType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchSubjectType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchSubjectType{
		Meta: safe.Meta,
		ModifierExtension: safe.ModifierExtension,
		AssignedComparisonGroup: safe.AssignedComparisonGroup,
		ResourceType: safe.ResourceType,
		Period: safe.Period,
		Study: safe.Study,
		Text: safe.Text,
		Language: safe.Language,
		ID: safe.ID,
		Identifier: safe.Identifier,
		ActualComparisonGroup: safe.ActualComparisonGroup,
		ImplicitRules: safe.ImplicitRules,
		Status: safe.Status,
		Extension: safe.Extension,
		Progress: safe.Progress,
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
