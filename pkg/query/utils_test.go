package query

import "testing"

func Test_renderTemplate(t *testing.T) {
	tests := []struct {
		name         string
		aliasPattern string
		aliasData    map[string]string
		expected     string
	}{
		{
			name:         "Simple replacement",
			aliasPattern: "Hello, {{name}}!",
			aliasData:    map[string]string{"name": "Alice"},
			expected:     "Hello, Alice!",
		},
		{
			name:         "Multiple replacements",
			aliasPattern: "Hello, {{name}}! Today is {{day}}.",
			aliasData:    map[string]string{"name": "Bob", "day": "Tuesday"},
			expected:     "Hello, Bob! Today is Tuesday.",
		},
		{
			name:         "Missing key",
			aliasPattern: "Hello, {{name}}! Your age is {{age}}.",
			aliasData:    map[string]string{"name": "Charlie"},
			expected:     "Hello, Charlie! Your age is .",
		},
		{
			name:         "Extra keys in aliasData",
			aliasPattern: "Welcome, {{user}}!",
			aliasData:    map[string]string{"user": "Diana", "extra": "ignored"},
			expected:     "Welcome, Diana!",
		},
		{
			name:         "Empty aliasPattern",
			aliasPattern: "",
			aliasData:    map[string]string{"key": "value"},
			expected:     "",
		},
		{
			name:         "No placeholders in pattern",
			aliasPattern: "Just a static string.",
			aliasData:    map[string]string{"key": "value"},
			expected:     "Just a static string.",
		},
		{
			name:         "Whitespace around key",
			aliasPattern: "Hello, {{  name  }}!",
			aliasData:    map[string]string{"name": "Eve"},
			expected:     "Hello, Eve!",
		},
		{
			name:         "Unmatched placeholders",
			aliasPattern: "Hello, {{name}}! {{unmatched}} remains.",
			aliasData:    map[string]string{"name": "Frank"},
			expected:     "Hello, Frank!  remains.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := renderTemplate(tt.aliasPattern, tt.aliasData)
			if result != tt.expected {
				t.Errorf("expected: %q, got: %q", tt.expected, result)
			}
		})
	}
}
