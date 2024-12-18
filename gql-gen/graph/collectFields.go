package graph

import (
	"context"
	"sort"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func refinePaths(fields []string) []string {
	/*Remove fields that aren't neccessary in traversing the graph*/
	var refined []string
	for _, field := range fields {
		segments := strings.Split(field, ".")
		if len(segments) > 1 && !strings.HasSuffix(segments[0], "Type") {
			for i := 1; i < len(segments); i++ {
				if strings.HasSuffix(segments[i], "Type") {
					segments = RemoveIndex(segments, i-1)
				}
			}
		}
		refinedField := strings.Join(segments, ".")
		refined = append(refined, refinedField)

	}

	return refined
}

func SortFieldPaths(fields []string, rootType string) []string {
	//refined := refinePaths(fields)
	sort.Slice(fields, func(i, j int) bool {
		countTypeI, nestingLevelI := countTypeOccurrencesAndLevels(fields[i])
		countTypeJ, nestingLevelJ := countTypeOccurrencesAndLevels(fields[j])

		// Prioritize fewer "Type" segments
		if countTypeI != countTypeJ {
			return countTypeI < countTypeJ
		}

		// Within same Type count, prioritize fewer nesting levels
		if nestingLevelI != nestingLevelJ {
			return nestingLevelI < nestingLevelJ
		}
		return fields[i] < fields[j]
	})
	return fields
}

// Helper function to count "Type" segments and nesting levels
func countTypeOccurrencesAndLevels(field string) (typeCount int, nestingLevel int) {
	typeCount = 0
	segments := strings.Split(field, ".")
	for _, segment := range segments {
		if strings.HasSuffix(segment, "Type") {
			typeCount++
		}
	}
	nestingLevel = len(segments) - 1
	return typeCount, nestingLevel
}

func GetQueryFields(ctx context.Context, rootType string) []string {
	fields := GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, []string{}),
		"",
		rootType,
	)
	return SortFieldPaths(fields, rootType)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string, rootType string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column, rootType)
		nestedFields := graphql.CollectFields(ctx, column.Selections, []string{})
		if len(nestedFields) == 0 {
			preloads = append(preloads, prefixColumn)
		} else {
			preloads = append(preloads, GetNestedPreloads(ctx, nestedFields, prefixColumn, rootType)...)
		}
	}
	return preloads
}

func GetPreloadString(prefix string, name graphql.CollectedField, rootType string) string {
	// If edge out to another type, traverse to that type
	if strings.HasSuffix(name.ObjectDefinition.Name, "Type") && name.ObjectDefinition.Name != rootType {
		return prefix + "." + name.ObjectDefinition.Name + "." + name.Name
	}
	if len(prefix) > 0 {
		return prefix + "." + name.Name
	}
	return name.Name
}
