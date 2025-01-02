package model

import (
	"encoding/json"
	"fmt"
)


func UnmarshalOrganizationTypeSlice(input []any) ([]OrganizationType, error) {
	out := []OrganizationType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := OrganizationType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalGroupTypeSlice(input []any) ([]GroupType, error) {
	out := []GroupType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := GroupType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalPractitionerTypeSlice(input []any) ([]PractitionerType, error) {
	out := []PractitionerType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := PractitionerType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalPractitionerRoleTypeSlice(input []any) ([]PractitionerRoleType, error) {
	out := []PractitionerRoleType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := PractitionerRoleType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalResearchStudyTypeSlice(input []any) ([]ResearchStudyType, error) {
	out := []ResearchStudyType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := ResearchStudyType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalPatientTypeSlice(input []any) ([]PatientType, error) {
	out := []PatientType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := PatientType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalResearchSubjectTypeSlice(input []any) ([]ResearchSubjectType, error) {
	out := []ResearchSubjectType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := ResearchSubjectType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalSubstanceTypeSlice(input []any) ([]SubstanceType, error) {
	out := []SubstanceType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := SubstanceType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalSubstanceDefinitionTypeSlice(input []any) ([]SubstanceDefinitionType, error) {
	out := []SubstanceDefinitionType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := SubstanceDefinitionType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalSpecimenTypeSlice(input []any) ([]SpecimenType, error) {
	out := []SpecimenType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := SpecimenType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalObservationTypeSlice(input []any) ([]ObservationType, error) {
	out := []ObservationType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := ObservationType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalDiagnosticReportTypeSlice(input []any) ([]DiagnosticReportType, error) {
	out := []DiagnosticReportType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := DiagnosticReportType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalConditionTypeSlice(input []any) ([]ConditionType, error) {
	out := []ConditionType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := ConditionType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalMedicationTypeSlice(input []any) ([]MedicationType, error) {
	out := []MedicationType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := MedicationType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalMedicationAdministrationTypeSlice(input []any) ([]MedicationAdministrationType, error) {
	out := []MedicationAdministrationType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := MedicationAdministrationType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalMedicationStatementTypeSlice(input []any) ([]MedicationStatementType, error) {
	out := []MedicationStatementType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := MedicationStatementType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalMedicationRequestTypeSlice(input []any) ([]MedicationRequestType, error) {
	out := []MedicationRequestType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := MedicationRequestType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalProcedureTypeSlice(input []any) ([]ProcedureType, error) {
	out := []ProcedureType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := ProcedureType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalDocumentReferenceTypeSlice(input []any) ([]DocumentReferenceType, error) {
	out := []DocumentReferenceType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := DocumentReferenceType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalTaskTypeSlice(input []any) ([]TaskType, error) {
	out := []TaskType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := TaskType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalImagingStudyTypeSlice(input []any) ([]ImagingStudyType, error) {
	out := []ImagingStudyType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := ImagingStudyType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalFamilyMemberHistoryTypeSlice(input []any) ([]FamilyMemberHistoryType, error) {
	out := []FamilyMemberHistoryType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := FamilyMemberHistoryType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}

func UnmarshalBodyStructureTypeSlice(input []any) ([]BodyStructureType, error) {
	out := []BodyStructureType{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := BodyStructureType{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}
