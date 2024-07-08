package underscore

import "testing"

func TestCamel(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"CamelCase", "camel_case"},
		{"CamelCaseName", "camel_case_name"},
		{"CamelCaseNameID", "camel_case_name_id"},
		{"CamelCaseNameID", "camel_case_name_id"},
	}

	for _, tt := range tests {
		got := Camel(tt.name)
		if got != tt.want {
			t.Errorf("Camel() = %v, want %v", got, tt.want)
		}
	}
}
