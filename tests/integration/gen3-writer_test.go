package integration_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/bmeg/grip-graphql/tests/integration"
	"github.com/bmeg/grip/gripql"
)

// --- Helper Functions for Test Isolation (Setup/Teardown) ---

const (
	GraphTEST     = "TEST"
	GraphJSONTEST = "JSONTEST"
	ProjectOHSU   = "ohsu-test"
	TestVertexID  = "cec32723-9ede-5f24-ba63-63cb8c6a02cf"
	TestEdgeID    = "e2867c6d-db7e-5d6e-87d8-b1c293f5b47e"
)

// deleteTestProject attempts to delete a project and logs the result.
func deleteTestProject(t *testing.T, graph, project string) {
	t.Helper()
	req := &integration.Request{
		Url:    fmt.Sprintf("http://localhost:8201/writer/%s/proj-delete/%s", graph, project),
		Method: "DELETE",
		// Use a writer token for deletion
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	response, _ := integration.TemplateRequest(req, t)
	status := response["status"].(float64)

	// We only log if deletion failed or was not found (which can happen if a previous test deleted it)
	if status != 200 {
		t.Logf("Project cleanup for %s/%s returned status %f: %v", graph, project, status, response)
	}
}

// loadTestProjectData creates the graph and bulk-loads the GRIP data.
func loadTestProjectData(t *testing.T, graph, project string) {
	t.Helper()
	// 1. Create Graph
	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/add-graph/%s/%s", graph, project),
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	response_json, pass := integration.TemplateRequest(req, t)
	if !pass {
		t.Errorf("Failed to create graph %s/%s: %v", graph, project, response_json)
		return
	}
	t.Logf("Created Graph %s/%s: %v", graph, project, response_json)

	// 2. Bulk Load Data
	err, responses := integration.BulkLoad(
		fmt.Sprintf("http://localhost:8201/writer/%s/bulk-load/%s", graph, project),
		"../fixtures/combio-examples-grip",
		integration.CreateToken(false, true, true),
	)
	if err != nil {
		t.Error("BulkLoad failed:", err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 200 {
			t.Errorf("BulkLoad response error: %v", resp)
		}
	}
}

// loadJSONSchemaData creates the JSONTEST graph, downloads the schema, and loads it.
func loadJSONSchemaData(t *testing.T) {
	t.Helper()

	// 1. Create Graph JSONTEST/ohsu-test
	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/add-graph/%s/%s", GraphJSONTEST, ProjectOHSU),
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	response_json, pass := integration.TemplateRequest(req, t)
	if !pass {
		t.Errorf("Failed to create graph %s/%s: %v", GraphJSONTEST, ProjectOHSU, response_json)
		return
	}
	t.Logf("Created Graph %s/%s: %v", GraphJSONTEST, ProjectOHSU, response_json)

	// 2. Download Schema (Existing logic)
	tempDir, err := os.MkdirTemp("", "jsonschema")
	if err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(tempDir) })
	tempFilePath := fmt.Sprintf("%s/graph-fhir.json", tempDir)
	curlCmd := exec.Command("curl", "-o", tempFilePath, "https://raw.githubusercontent.com/bmeg/iceberg/refs/heads/main/schemas/graph/graph-fhir.json")
	if err := curlCmd.Run(); err != nil {
		t.Fatalf("Error running curl command: %v", err)
	}
	t.Logf("Downloaded schema to: %s", tempFilePath)

	// 3. Add JSON Schema
	err, response := integration.MultipartFormTest(
		fmt.Sprintf("http://localhost:8201/writer/%s/add-json-schema/%s", GraphJSONTEST, ProjectOHSU),
		tempFilePath,
		integration.CreateToken(false, true, true),
	)
	if err != nil {
		t.Error("Add JSON schema failed:", err)
	}
	t.Log("RESP: ", response)
}

// --- Test Functions ---

func Test_Writer_Load_Ok(t *testing.T) {
	// Set up cleanup first: delete project regardless of test result
	t.Cleanup(func() {
		deleteTestProject(t, GraphTEST, ProjectOHSU)
	})

	// Perform the loading steps (now simplified by the helper)
	loadTestProjectData(t, GraphTEST, ProjectOHSU)
}

// This test now relies on the loadJSONSchemaData helper
func Test_Writer_Json_Schema_Load_Ok(t *testing.T) {
	// Cleanup: delete project regardless of test result
	t.Cleanup(func() {
		deleteTestProject(t, GraphJSONTEST, ProjectOHSU)
	})

	loadJSONSchemaData(t) // This helper performs the create and load
}

func Test_Writer_Json_Load_Ok(t *testing.T) {
	// This test needs the schema loaded, so we run the setup.
	// Since Test_Writer_Json_Schema_Load_Ok handles cleanup, we just run the setup.
	// Note: In a production test suite, you'd call a dedicated setup helper here,
	// but we call the schema load one for simplicity.
	loadJSONSchemaData(t)
	// Add cleanup for the setup data in case the main schema test didn't run.
	t.Cleanup(func() { deleteTestProject(t, GraphJSONTEST, ProjectOHSU) })

	// Original assertions
	err, response := integration.MultipartFormTest(
		fmt.Sprintf("http://localhost:8201/writer/%s/bulk-load-raw/%s", GraphJSONTEST, ProjectOHSU),
		"../fixtures/compbio-examples-fhir/DocumentReference.ndjson",
		integration.CreateToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	if response["status"].(float64) != 200 {
		t.Errorf("Response %f != 200\n", response["status"].(float64))
	}
}

// ... (Test_Writer_Json_Schema_Not_Found remains largely the same, no data dependency) ...

func Test_Writer_Json_Schema_Not_Found(t *testing.T) {
	// Note: This test still relies on the JSONTEST/ohsu-test project existing,
	// but since it attempts to load an invalid schema file, it's mostly self-contained.
	// We'll ensure the project exists and is cleaned up.
	t.Cleanup(func() { deleteTestProject(t, GraphJSONTEST, ProjectOHSU) })
	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/add-graph/%s/%s", GraphJSONTEST, ProjectOHSU),
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	integration.TemplateRequest(req, t)

	err, response := integration.MultipartFormTest(
		fmt.Sprintf("http://localhost:8201/writer/%s/add-json-schema/%s", GraphJSONTEST, ProjectOHSU),
		"../fixtures/not_found_schema.json",
		integration.CreateToken(false, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	if response["status"].(float64) == 500 {
		t.Log("Expected outcome:", response["message"].(string))
	} else {
		t.Error("Status shouldn't equal 500 or message suggests error was not about file not found", response["status"])
	}
}

// ... (Tests for validation/invalid JSON errors should also run loadJSONSchemaData first and add cleanup) ...
func Test_Writer_Json_Load_Validation_Errors(t *testing.T) {
	loadJSONSchemaData(t)
	t.Cleanup(func() { deleteTestProject(t, GraphJSONTEST, ProjectOHSU) })

	// Original logic...
	file, err := os.Open("../fixtures/compbio-examples-fhir/ValidationErrors.ndjson")
	if err != nil {
		t.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("file", filepath.Base("../fixtures/compbio-examples-fhir/ValidationErrors.ndjson"))
	io.Copy(part, file)
	writer.Close()

	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/bulk-load-raw/%s", GraphJSONTEST, ProjectOHSU),
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}

	// ... (rest of the HTTP request execution logic remains the same) ...
	request, err := http.NewRequest(req.Method, req.Url, body)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	for key, val := range req.Headers {
		request.Header.Set(key, val.(string))
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Logf("Server responded with status: %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	resp.Body.Close()

	var data map[string]any
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		t.Errorf("Error unmarshalling response: %v\nResponse Body: %s", errors, buf.String())
	}

	// ... (rest of the assertions remain the same) ...
	message, ok := data["message"].([]any)
	if !ok {
		t.Errorf("Message expected []any, got %T", data["message"])
	} else if len(message) != 3 {
		t.Errorf("Expected validation message length to be 3 but got %d instead", len(message))
	}

	t.Log("MESSAGE: ", message)

	status, ok := data["status"].(float64)
	if !ok {
		var actualType string
		if rawStatus, found := data["status"]; found {
			actualType = fmt.Sprintf("%T", rawStatus)
		} else {
			actualType = "missing"
		}
		t.Errorf("Validation Error: Expected 'status' field to be a float64, got %s", actualType)
	} else if status != 206 {
		t.Errorf("Validation Error: Expected status 206, got %f", status)
	} else {
		t.Log("Status field value OK.")
	}
}

func Test_Writer_Json_Load_Invalid_Json(t *testing.T) {
	loadJSONSchemaData(t)
	t.Cleanup(func() { deleteTestProject(t, GraphJSONTEST, ProjectOHSU) })

	// Original logic...
	file, err := os.Open("../fixtures/compbio-examples-fhir/invalid_json.ndjson")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("file", filepath.Base("../fixtures/compbio-examples-fhir/ValidationErrors.ndjson"))
	io.Copy(part, file)
	writer.Close()

	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/bulk-load-raw/%s", GraphJSONTEST, ProjectOHSU),
		Method:  "POST",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}

	// ... (rest of the HTTP request execution logic remains the same) ...
	request, err := http.NewRequest(req.Method, req.Url, body)
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.Headers {
		request.Header.Set(key, val.(string))
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error("Error sending request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Logf("server responded with status: %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		t.Error("Error reading response:", err)
		return
	}

	var data map[string]any
	errors := json.Unmarshal([]byte(buf.String()), &data)
	if errors != nil {
		t.Error("Error: ", errors)
		return
	}
	t.Log(data)
	if data["status"].(float64) != 206 {
		t.Error()
	}
}

// --- Dependency-Free Tests (Access/Token Errors) ---
// These don't modify data, but still rely on the project existing for the request URL to be valid.

func Test_Writer_Load_Malformed_Token(t *testing.T) {
	// No data setup needed, but a quick cleanup in case this is the first test run.
	t.Cleanup(func() { deleteTestProject(t, GraphTEST, ProjectOHSU) })

	// Server returns a 400 given an unparsable token
	err, responses := integration.BulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"../fixtures/combio-examples-grip",
		integration.CreateToken(false, true, true)[2:50],
	)
	if err != nil {
		t.Error(err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 400 {
			t.Error(resp)
		} else {
			t.Log(resp)
		}
	}
}

func Test_Writer_Load_Expired_Token(t *testing.T) {
	t.Cleanup(func() { deleteTestProject(t, GraphTEST, ProjectOHSU) })

	// Server returns a 401 given an expired token
	err, responses := integration.BulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"../fixtures/combio-examples-grip",
		integration.CreateToken(true, true, true),
	)
	if err != nil {
		t.Error(err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 401 {
			t.Error(resp)
		} else {
			t.Log(resp)
		}
	}
}

func Test_Writer_Load_No_Writer_Perms(t *testing.T) {
	t.Cleanup(func() { deleteTestProject(t, GraphTEST, ProjectOHSU) })

	// Server returns a 403 given a token that represents a user with no writer perms on the specified project
	err, responses := integration.BulkLoad("http://localhost:8201/writer/TEST/bulk-load/ohsu-test",
		"../fixtures/combio-examples-grip",
		integration.CreateToken(false, false, true),
	)
	if err != nil {
		t.Error(err)
	}
	for _, resp := range responses {
		if resp["status"].(float64) != 403 {
			t.Error(resp)
		} else {
			t.Log(resp)
		}
	}
}

// --- Data-Dependent Read Tests (Need data loaded) ---

func Test_Writer_Get_Vertex_No_Reader_Perms(t *testing.T) {
	// Setup: Ensure data is loaded
	loadTestProjectData(t, GraphTEST, ProjectOHSU)
	t.Cleanup(func() { deleteTestProject(t, GraphTEST, ProjectOHSU) })

	// Server returns a 403 given a token that represents a user with no reader perms on the specified project
	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/get-vertex/%s/%s", GraphTEST, TestVertexID, ProjectOHSU),
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, false)},
	}

	response, status := integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 403) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Writer_Get_Vertex_Ok(t *testing.T) {
	// Setup: Ensure data is loaded
	loadTestProjectData(t, GraphTEST, ProjectOHSU)
	t.Cleanup(func() { deleteTestProject(t, GraphTEST, ProjectOHSU) })

	// Given appropriate perms, vertex should be retrieved
	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/get-vertex/%s/%s", GraphTEST, TestVertexID, ProjectOHSU),
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, false, true)},
	}
	response, status := integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

func Test_Writer_Get_Project_Vertices_Ok(t *testing.T) {
	// Setup: Ensure data is loaded
	loadTestProjectData(t, GraphTEST, ProjectOHSU)
	t.Cleanup(func() { deleteTestProject(t, GraphTEST, ProjectOHSU) })

	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/get-vertices/%s", GraphTEST, ProjectOHSU),
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, false, true)},
	}

	request, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer(req.Body))
	if err != nil {
		t.Error("Error creating request:", err)
		return
	}
	for key, val := range req.Headers {
		request.Header.Set(key, val.(string))
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Error("Error sending request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Logf("Server responded with status: %s", resp.Status)
	}

	reader := bufio.NewReader(resp.Body)
	jum := gripql.NewFlattenMarshaler()
	// ... (rest of the read loop logic remains the same) ...
	count := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Errorf("Error reading response: %s", err)
		}
		v := &gripql.Vertex{}
		err = jum.Unmarshal([]byte(line), v)
		if err != nil {
			t.Error(err)
		}
		if v.Id == "" {
			t.Error("Gid should be populated if unmarshal was successful")
		}
		mappedData := v.Data.AsMap()

		t.Logf("Received Vertex: %s\n", mappedData["resourceType"])

		if mappedData["auth_resource_path"] != "/programs/ohsu/projects/test" {
			t.Error("Returned data should have resource path: /programs/ohsu/projects/test")
		}
		count++
	}
	if count == 0 {
		t.Error("Expected to receive vertices, but none were returned.")
	}
}

func Test_Writer_Delete_Edge_Ok(t *testing.T) {
	// Setup: Ensure data is loaded (including edges)
	loadTestProjectData(t, GraphTEST, ProjectOHSU)
	t.Cleanup(func() { deleteTestProject(t, GraphTEST, ProjectOHSU) })

	// Regular delete should return 200
	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/del-edge/%s/%s", GraphTEST, TestEdgeID, ProjectOHSU),
		Method:  "DELETE",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	response, status := integration.TemplateRequest(req, t)
	if !(response["status"].(float64) == 200) || !status {
		t.Error(response)
	}
	t.Log(status, response)
}

// FIX: Rewritten to be self-contained and run its own setup.
func Test_Writer_Delete_Proj_Endpoint_Ok(t *testing.T) {
	// 1. SETUP: Guarantee the project exists immediately before attempting deletion.
	loadTestProjectData(t, GraphTEST, ProjectOHSU)

	// 2. Delete everything from test graph project ohsu-test. Should return 200
	req := &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/proj-delete/%s", GraphTEST, ProjectOHSU),
		Method:  "DELETE",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, true, true)},
	}
	response, status := integration.TemplateRequest(req, t)
	// IMPORTANT: Check for 200, which confirms the project was deleted.
	if !(response["status"].(float64) == 200) {
		t.Error("Project deletion failed with status:", response)
	}
	t.Log(status, response)

	// 3. Look for the vertex that was retrieved earlier. It should be gone (404)
	req = &integration.Request{
		Url:     fmt.Sprintf("http://localhost:8201/writer/%s/get-vertex/%s/%s", GraphTEST, TestVertexID, ProjectOHSU),
		Method:  "GET",
		Headers: map[string]any{"Authorization": integration.CreateToken(false, false, true)},
	}
	response, status = integration.TemplateRequest(req, t)
	t.Log("RESPONSE: ", response)
	if !(response["status"].(float64) == 404) {
		t.Errorf("Vertex was unexpectedly found after project deletion. Status: %f, Response: %v", response["status"].(float64), response)
	}

	// No t.Cleanup() is needed here, as the test successfully deleted the project.
}
