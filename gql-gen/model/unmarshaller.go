package model

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TypedObject struct {
	Typename string `json:"__typename"`
}

/*
can unmarshal to any type on every union because graphql will error out before
unmarshaller attempts to resolve a union edge type that is not allowed on a specific type
*/
func unmarshalUnion(b []byte, fieldName, typename string, target interface{}) error {
	// if typename is empty, then field was not returned in query, so nothing to unmarshal.
	if typename == "" {
		return nil
	}
	var partial map[string]json.RawMessage
	if err := json.Unmarshal(b, &partial); err != nil {
		return err
	}
	rawField, ok := partial[fieldName]
	if !ok {
		return fmt.Errorf("field %s not found in JSON", fieldName)
	}
	var obj interface{}
	switch typename {
	case "PatientType":
		obj = &PatientType{}
	case "SpecimenType":
		obj = &SpecimenType{}
	case "MedicationType":
		obj = &MedicationType{}
	case "OrganizationType":
		obj = &OrganizationType{}
	case "GroupType":
		obj = &GroupType{}
	case "PractitionerType":
		obj = &PractitionerType{}
	case "PractitionerRoleType":
		obj = &PractitionerRoleType{}
	case "ResearchStudyType":
		obj = &ResearchStudyType{}
	case "ResearchSubjectType":
		obj = &ResearchSubjectType{}
	case "SubstanceType":
		obj = &SubstanceType{}
	case "SubstanceDefinitionType":
		obj = &SubstanceDefinitionType{}
	case "ObservationType":
		obj = &ObservationType{}
	case "DiagnosticReportType":
		obj = &DiagnosticReportType{}
	case "ConditionType":
		obj = &ConditionType{}
	case "MedicationAdministrationType":
		obj = &MedicationAdministrationType{}
	case "MedicationStatementType":
		obj = &MedicationStatementType{}
	case "MedicationRequestType":
		obj = &MedicationRequestType{}
	case "ProcedureType":
		obj = &ProcedureType{}
	case "DocumentReferenceType":
		obj = &DocumentReferenceType{}
	case "TaskType":
		obj = &TaskType{}
	case "ImagingStudyType":
		obj = &ImagingStudyType{}
	case "FamilyMemberHistoryType":
		obj = &FamilyMemberHistoryType{}
	case "BodyStructureType":
		obj = &BodyStructureType{}
	default:
		return fmt.Errorf("unknown typename: %s", typename)
	}

	if err := json.Unmarshal(rawField, obj); err != nil {
		return fmt.Errorf("failed to unmarshal field %s as %s: %w", fieldName, typename, err)
	}
	reflect.ValueOf(target).Elem().Set(reflect.ValueOf(obj).Elem())
	return nil
}
