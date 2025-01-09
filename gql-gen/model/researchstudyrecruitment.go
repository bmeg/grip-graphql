
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchStudyRecruitment struct {
	ResourceType *string `json:"resourceType,omitempty"`
	TargetNumber *int `json:"targetNumber,omitempty"`
	ActualGroup *GroupType `json:"actualGroup"`
	ActualNumber *int `json:"actualNumber,omitempty"`
	Eligibility TypedObject `json:"eligibility"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchStudyRecruitment) UnmarshalJSON(b []byte) error {
	var safe SafeResearchStudyRecruitment
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchStudyRecruitment{
		ResourceType: safe.ResourceType,
		TargetNumber: safe.TargetNumber,
		ActualGroup: safe.ActualGroup,
		ActualNumber: safe.ActualNumber,
		Extension: safe.Extension,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "eligibility", safe.Eligibility.Typename, &o.Eligibility); err != nil {
		return fmt.Errorf("failed to unmarshal Eligibility: %w", err)
	}

	return nil
}
