package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

type ClassDefinition struct {
	Name   string
	Fields []FieldDefinition
}

type FieldDefinition struct {
	Name     string
	Type     string
	Tag      string
	IsUnion  bool
	UnionTag string
}

type SafeFieldDefinition struct {
	Name     string
	Type     string
	Tag      string
	UnionTag string
}

const safeStructTemplate = `
package model

import (
	"encoding/json"
	"fmt"
)

type Safe{{ .Name }} struct {
{{- range .SafeFields }}
	{{ .Name }} {{ .Type }} {{ .Tag }}
{{- end }}
}

func (o *{{ .Name }}) UnmarshalJSON(b []byte) error {
	var safe Safe{{ .Name }}
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = {{ .Name }}{
{{- range .NonUnionFields }}
		{{ .Name }}: safe.{{ .Name }},
{{- end }}
	}

{{- range .UnionFields }}
	if err := unmarshalUnion(b, "{{ .UnionTag }}", safe.{{ .Name }}.Typename, &o.{{ .Name }}); err != nil {
		return fmt.Errorf("failed to unmarshal {{ .Name }}: %w", err)
	}
{{- end }}

	return nil
}
`

const unmarshalFuncTemplate = `
func Unmarshal{{ .Name }}Slice(input []any) ([]{{ .Name }}, error) {
	out := []{{ .Name }}{}
	for _, item := range input {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal item to JSON: %w", err)
		}
		typedItem := {{ .Name }}{}
		if err := json.Unmarshal(jsonData, &typedItem); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON into target type: %w", err)
		}
		out = append(out, typedItem)
	}
	return out, nil
}
`

func main() {
	var schemaPath string
	var modelFile string
	var classNames []string

	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "",
		Short: "Generate Safe struct files from model definitions",
	}

	// Define the "generate" subcommand
	var generateCmd = &cobra.Command{
		Use:   "union",
		Short: "Generate Safe struct files from model definitions",
		Run: func(cmd *cobra.Command, args []string) {
			if modelFile == "" {
				log.Fatal("Model file path is required")
			}

			// Read the model file
			content, err := ioutil.ReadFile(modelFile)
			if err != nil {
				log.Fatalf("Failed to read model file: %v", err)
			}

			modelText := string(content)

			schema, err := loadSchema(schemaPath)
			if err != nil {
				fmt.Printf("Error loading schema: %s\n", err)
				os.Exit(1)
			}

			resolvers := findTypesWithResolvers(schema)

			classDefinitions := extractClassDefinitions(modelText, resolvers)

			for _, classDef := range classDefinitions {
				generateSafeStruct(classDef)
			}
		},
	}

	// Define the "unmarshal" subcommand
	var unmarshalCmd = &cobra.Command{
		Use:   "unmarshal",
		Short: "Generate Unmarshal functions for predefined types",
		Run: func(cmd *cobra.Command, args []string) {

			generateUnmarshalFuncs(classNames)
		},
	}

	generateCmd.Flags().StringVarP(&schemaPath, "schemaPath", "c", "", "Path to graphql schema")
	generateCmd.Flags().StringVarP(&modelFile, "model-file", "m", "", "Path to the model file")

	unmarshalCmd.Flags().StringSliceVarP(&classNames, "classes", "c", nil, "List of class names to process")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(unmarshalCmd)

	// Execute the root command, which will handle both subcommands
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// Extract class definitions for the given class names
func extractClassDefinitions(modelText string, classNames []string) []ClassDefinition {
	var classDefs []ClassDefinition

	for _, className := range classNames {
		re := regexp.MustCompile(fmt.Sprintf(`type %s struct \{([\s\S]*?)\}`, className))
		match := re.FindStringSubmatch(modelText)
		if len(match) < 2 {
			continue
		}

		fields := extractFields(match[1])
		classDefs = append(classDefs, ClassDefinition{
			Name:   className,
			Fields: fields,
		})
	}

	return classDefs
}

// Extract fields from a struct body
func extractFields(body string) []FieldDefinition {
	lines := strings.Split(body, "\n")
	var fields []FieldDefinition

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		fieldType := parts[1]
		if strings.HasSuffix(fieldType, "Union") {
			fieldType = "TypedObject"
		}

		field := FieldDefinition{
			Name:     parts[0],
			Type:     fieldType,
			IsUnion:  strings.HasSuffix(parts[1], "Union"),
			Tag:      extractTag(line),
			UnionTag: extractFieldNameFromTag(line),
		}

		fields = append(fields, field)
	}

	return fields
}

// Extract JSON tag from a line
func extractTag(line string) string {
	re := regexp.MustCompile("`json:\"[^\"]*\"`")
	tag := re.FindString(line)
	return tag
}

func extractFieldNameFromTag(line string) string {
	re := regexp.MustCompile("`json:\"([^\"]*)\"`")
	tag := re.FindStringSubmatch(line)
	fieldName := tag[1]
	fieldName = strings.ReplaceAll(fieldName, ",omitempty", "")
	return fieldName
}

// Generate a Safe struct and its Unmarshaller
func generateSafeStruct(classDef ClassDefinition) {
	var safeFields []SafeFieldDefinition
	var nonUnionFields, unionFields []FieldDefinition

	for _, field := range classDef.Fields {
		safeFields = append(safeFields, SafeFieldDefinition{
			Name:     field.Name,
			Type:     field.Type,
			Tag:      field.Tag,
			UnionTag: field.UnionTag,
		})

		if field.IsUnion {
			unionFields = append(unionFields, field)
		} else {
			nonUnionFields = append(nonUnionFields, field)
		}
	}

	data := struct {
		Name           string
		SafeFields     []SafeFieldDefinition
		NonUnionFields []FieldDefinition
		UnionFields    []FieldDefinition
	}{
		Name:           classDef.Name,
		SafeFields:     safeFields,
		NonUnionFields: nonUnionFields,
		UnionFields:    unionFields,
	}

	var buf bytes.Buffer
	tmpl := template.Must(template.New("safeStruct").Parse(safeStructTemplate))
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	outputFile := fmt.Sprintf("../%s.go", strings.ToLower(classDef.Name))
	if err := ioutil.WriteFile(outputFile, buf.Bytes(), 0644); err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Printf("Generated: %s\n", outputFile)
}

// Generate Unmarshal functions for the predefined types and write them to a single file
func generateUnmarshalFuncs(classNames []string) {

	var buf bytes.Buffer

	buf.WriteString("package model\n\n")
	buf.WriteString("import (\n")
	buf.WriteString("\t\"encoding/json\"\n")
	buf.WriteString("\t\"fmt\"\n")
	buf.WriteString(")\n\n")

	for _, t := range classNames {
		data := struct {
			Name string
		}{
			Name: t,
		}

		tmpl := template.Must(template.New("unmarshalFunc").Parse(unmarshalFuncTemplate))
		if err := tmpl.Execute(&buf, data); err != nil {
			log.Fatalf("Failed to execute template: %v", err)
		}
	}

	outputFile := "../unmarshalSlice.go"
	if err := ioutil.WriteFile(outputFile, buf.Bytes(), 0644); err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Printf("Generated: %s\n", outputFile)
}

// findTypesWithResolvers parses a GraphQL schema and identifies types with resolver fields
func findTypesWithResolvers(schema string) []string {
	typePattern := regexp.MustCompile(`type\s+(\w+)\s*\{([^}]*)\}`)
	fieldPattern := regexp.MustCompile(`(\w+)\s*\(.*\):\s*[\w\[\]!]+`)

	resolvers := []string{}

	matches := typePattern.FindAllStringSubmatch(schema, -1)
	for _, match := range matches {
		typeName := match[1]
		fields := match[2]
		fieldMatches := fieldPattern.FindAllStringSubmatch(fields, -1)
		for _, fieldMatch := range fieldMatches {
			if len(fieldMatch) > 0 {
				if strings.HasSuffix(fieldMatch[0], "Union!") || strings.HasSuffix(fieldMatch[0], "Union") {
					if !slices.Contains(resolvers, typeName) && typeName != "Query" {
						resolvers = append(resolvers, typeName)
					}
				}
			}
		}
	}

	return resolvers
}

// loadSchema reads the GraphQL schema from the specified file path
func loadSchema(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return sb.String(), nil
}
