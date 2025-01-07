
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDiagnosticReportSupportingInfo struct {
	Type *CodeableConcept `json:"type,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Reference TypedObject `json:"reference"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DiagnosticReportSupportingInfo) UnmarshalJSON(b []byte) error {
	var safe SafeDiagnosticReportSupportingInfo
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DiagnosticReportSupportingInfo{
		Type: safe.Type,
		Extension: safe.Extension,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "reference", safe.Reference.Typename, &o.Reference); err != nil {
		return fmt.Errorf("failed to unmarshal Reference: %w", err)
	}

	return nil
}
