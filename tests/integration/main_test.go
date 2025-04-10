package integration_test

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"testing"
)

// TestMain is the entry point for all tests in the package
func TestMain(m *testing.M) {
	if !serverIsRunning() {
		fmt.Println("Server not running at http://localhost:8201")
		os.Exit(1)
	}
	code := runTests(m)
	os.Exit(code)
}

// runTests orchestrates the test execution order
func runTests(m *testing.M) int {
	writerTests := []func(t *testing.T){
		Test_Writer_Load_Ok,
		Test_Writer_Json_Schema_Load_Ok,
		Test_Writer_Json_Load_Ok,
		Test_Writer_Json_Schema_Not_Found,
		Test_Writer_Json_Load_Validation_Errors,
		Test_Writer_Json_Load_Invalid_Json,
		Test_Writer_Load_Malformed_Token,
		Test_Writer_Load_Expired_Token,
		Test_Writer_Load_No_Writer_Perms,
		Test_Writer_Get_Vertex_No_Reader_Perms,
		Test_Writer_Get_Vertex_Ok,
		Test_Writer_Get_Project_Vertices_Ok,
		Test_Writer_Delete_Edge_Ok,
	}

	gqlGenTests := []func(t *testing.T){
		Test_GqlGen_Basic_Graphql,
		Test_GqlGen_Bad_Filter_Logical_Operator,
		Test_GqlGen_Filter_Ok,
		Test_GqlGen_Edge_Traversal_Ok,
		Test_GqlGen_Nested_Edge_Traversal_Ok,
		Test_GqlGen_Nested_Edge_Traversal_Filter_Ok,
	}

	gripGraphqlTests := []func(t *testing.T){
		Test_Graphql_Query_Forbidden_Perms,
		Test_Graphql_Query_Proj,
		Test_Graphql_Mutation_AddSpecimen,
	}

	fmt.Println("Starting writer tests")
	for _, test := range writerTests {
		t := &testing.T{}
		fmt.Printf("Running %s\n", getFunctionName(test))
		test(t)
		if t.Failed() {
			fmt.Printf("Writer test %s failed\n", getFunctionName(test))
			return 1
		}
	}

	fmt.Println("Starting gql-gen tests")
	for _, test := range gqlGenTests {
		t := &testing.T{}
		fmt.Printf("Running %s\n", getFunctionName(test))
		test(t)
		if t.Failed() {
			fmt.Printf("Gql-gen test %s failed\n", getFunctionName(test))
			return 1
		}
	}

	fmt.Println("Starting grip-graphql tests")
	for _, test := range gripGraphqlTests {
		t := &testing.T{}
		fmt.Printf("Running %s\n", getFunctionName(test))
		test(t)
		if t.Failed() {
			fmt.Printf("Grip-graphql test %s failed\n", getFunctionName(test))
			return 1
		}
	}

	fmt.Println("Starting cleanup")
	cleanupT := &testing.T{}
	Test_Writer_Delete_Proj(cleanupT)
	if cleanupT.Failed() {
		fmt.Println("Cleanup test failed")
		return 1
	}

	return 0
}

// getFunctionName extracts the name of a function for use in t.Run
func getFunctionName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func serverIsRunning() bool {
	resp, err := http.Get("http://localhost:8201/writer/_status")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}
