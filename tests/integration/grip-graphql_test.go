package integration_test

import (
	"testing"

	"github.com/bmeg/grip-graphql/tests/integration"
)

func Test_Graphql_Query_Forbidden_Perms(t *testing.T) {
	/* Testing do a query, but with a token that doesn't have read perms */
	payload := []byte(`{
	    "query": "query IdsWithSpecimenEdge { IdsWithSpecimenEdge {    id  }}",
	    "variables": {
	        "limit": 1000
	    }
	}`)

	req := &integration.Request{
		Url:    "http://localhost:8201/reader/api",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, false),
			"Content-Type":  "application/json"},
		Body: payload,
	}
	response, status := integration.TemplateRequest(req, t)
	if !(response["StatusCode"].(float64) == 403) {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Graphql_Query_Proj(t *testing.T) {
	/* A basic test for Graphql style query with mock auth */
	payload := []byte(`{
	    "query": "query IdsWithSpecimenEdge { IdsWithSpecimenEdge {    id  }}",
	    "variables": {
	        "limit": 1000
	    }
	}`)

	req := &integration.Request{
		Url:    "http://localhost:8201/reader/api",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, false, true),
			"Content-Type":  "application/json"},
		Body: payload,
	}

	response, status := integration.TemplateRequest(req, t)
	if !status {
		t.Error("Status returned false: ", response)
	}
	correct_response, ok := response["data"].(map[string]any)["IdsWithSpecimenEdge"].([]any)
	if !ok {
		t.Error("Response not indexable for 'data' and/or 'IdsWithSpecimenEdge' keys")
	}
	for _, resp := range correct_response {
		val, ok := resp.(map[string]any)["id"].(string)
		if !ok {
			t.Error("Response not indexable on 'id': ", resp)
		}
		t.Log("Return VAL: ", val)
	}
}

func Test_Graphql_Mutation_AddSpecimen(t *testing.T) {
	/* A basic test for Graphql style mutation with mock auth */
	payload := []byte(`{
	    "query": "mutation AddSpecimen($id: String, $auth_resource_path: String!, $collection: CollectionInput, $identifier: [IdentifierInput], $processing: [ProcessingInput], $resourceType: String, $subject: SubjectInput) { AddSpecimen(id: $id, auth_resource_path: $auth_resource_path, collection: $collection, identifier: $identifier, processing: $processing, resourceType: $resourceType, subject: $subject) { id } }",
	    "variables": {
	        "id": "53c67a06-ea2d-4d24-9249-418dc77a16a9",
	        "auth_resource_path": "/programs/ohsu/projects/test-invalid",
	        "collection": {
	            "bodySite": {
	                "concept": {
	                    "coding": [{"code": "76752008", "display": "Breast", "system": "http://snomed.info/sct"}],
	                    "text": "Breast"
	                }
	            },
	            "collector": {"reference": "Organization/89c8dc4c-2d9c-48c7-8862-241a49a78f14"}
	        },
	        "identifier": [
	            {"system": "https://my_demo.org/labA", "use": "official", "value": "specimen_1234_labA"}
	        ],
	        "processing": [
	            {
	                "method": {
	                    "coding": [
	                        {"code": "117032008", "display": "Spun specimen (procedure)", "system": "http://snomed.info/sct"},
	                        {"code": "Double-Spun", "display": "Double-Spun", "system": "https://my_demo.org/labA"}
	                    ],
	                    "text": "Spun specimen (procedure)"
	                }
	            }
	        ],
	        "resourceType": "Specimen",
	        "subject": {"reference": "Patient/ac4e1aa6-cb52-40e9-8f20-594d9c84f920"}
	    }
	}`)

	req := &integration.Request{
		Url:    "http://localhost:8201/reader/api",
		Method: "POST",
		Headers: map[string]any{
			"Authorization": integration.CreateToken(false, true, true),
			"Content-Type":  "application/json",
		},
		Body: payload,
	}

	response, status := integration.TemplateRequest(req, t)
	if !status {
		t.Errorf("Status returned false: %v", response)
		return
	}

	// Parse the response into a map
	respData, ok := response["data"].(map[string]any)
	if !ok {
		t.Errorf("Response 'data' not found or not a map: %v", response)
		return
	}

	// Check AddSpecimen field
	addSpecimen, ok := respData["AddSpecimen"].(map[string]any)
	if !ok {
		t.Errorf("Response 'AddSpecimen' not found or not a map: %v", respData)
		return
	}

	// Check id field
	id, ok := addSpecimen["id"].(string)
	if !ok {
		t.Errorf("Response 'id' not found or not a string: %v", addSpecimen)
		return
	}

	if id != "53c67a06-ea2d-4d24-9249-418dc77a16a9" {
		t.Errorf("Response '%s' != 53c67a06-ea2d-4d24-9249-418dc77a16a9", id)
		return
	}
}
