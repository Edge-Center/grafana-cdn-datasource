package query

import (
	"math"
	"regexp"
)

func renderTemplate(aliasPattern string, aliasData map[string]string) string {
	aliasRegex := regexp.MustCompile(`\{\{\s*(.+?)\s*\}\}`)
	return aliasRegex.ReplaceAllStringFunc(aliasPattern, func(match string) string {
		key := aliasRegex.FindStringSubmatch(match)[1]
		return aliasData[key]
	})
}

func clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(value, max))
}
