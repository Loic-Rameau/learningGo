package options

import "testing"

func TestOptions_Parse(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  Options
	}{
		{
			name:  "empty",
			input: []string{},
			want: Options{
				ActionType: ActionHelp,
			},
		},
		{
			name:  "unknown should print help",
			input: []string{"unknown"},
			want: Options{
				ActionType: ActionHelp,
			},
		},
		{
			name:  "help",
			input: []string{"help"},
			want: Options{
				ActionType: ActionHelp,
			},
		},
		{
			name:  "help should not parse extra",
			input: []string{"help", "extra"},
			want: Options{
				ActionType: ActionHelp,
			},
		},
		{
			name:  "add with no parameters",
			input: []string{"add"},
			want: Options{
				ActionType: ActionAdd,
			},
		},
		{
			name:  "add with parameters",
			input: []string{"add", "test"},
			want: Options{
				ActionType: ActionAdd,
				extra:      "test",
			},
		},
		{
			name:  "list",
			input: []string{"list"},
			want: Options{
				ActionType: ActionList,
			},
		},
		{
			name:  "list should not parse extra",
			input: []string{"list", "extra"},
			want: Options{
				ActionType: ActionList,
			},
		},
		{
			name:  "delete with no parameters",
			input: []string{"delete"},
			want: Options{
				ActionType: ActionDelete,
			},
		},
		{
			name:  "delete with parameters",
			input: []string{"delete", "test"},
			want: Options{
				ActionType: ActionDelete,
				extra:      "test",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			get := Parse(test.input)
			if get != test.want {
				t.Errorf("got %v; want %v", get, test.want)
			}
		})
	}
}
