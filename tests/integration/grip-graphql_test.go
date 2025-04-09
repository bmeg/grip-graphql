package integration_test

import (
	"testing"

	"github.com/bmeg/grip-graphql/tests/integration"
)

func Test_Graphql_Query_Forbidden_Perms(t *testing.T) {
	/* Testing do a query, but with a token that doesn't have read perms */
	payload := []byte(`{
	    "query": "query PatientIdsWithSpecimenEdge { PatientIdsWithSpecimenEdge {    id  }}",
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
	    "query": "query PatientIdsWithSpecimenEdge { PatientIdsWithSpecimenEdge {    id  }}",
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
	correct_response, ok := response["data"].(map[string]any)["PatientIdsWithSpecimenEdge"].([]any)
	if !ok {
		t.Error("Response not indexable for 'data' and/or 'PatientIdsWithSpecimenEdge' keys")
	}
	for _, resp := range correct_response {
		val, ok := resp.(map[string]any)["id"].(string)
		if !ok {
			t.Error("Response not indexable on 'id': ", resp)
		}
		t.Log("Return VAL: ", val)
	}
}
