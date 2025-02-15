package project

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		want  []string
	}{
		{
			name:  "Empty",
			slice: []string{},
			want:  []string{},
		},
		{
			name:  "SingleValue",
			slice: []string{"foo"},
			want:  []string{"foo"},
		},
		{
			name:  "SingleValue2",
			slice: []string{"foo", "foo"},
			want:  []string{},
		},
		{
			name:  "TwoValue",
			slice: []string{"foo", "bar"},
			want:  []string{"foo", "bar"},
		},
		{
			name:  "TwoValues2",
			slice: []string{"foo", "bar", "foo", "bar"},
			want:  []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unique(tt.slice)
			assert.Truef(t, reflect.DeepEqual(got, tt.want), "unique() = %v, want %v", got, tt.want)
		})
	}
}
