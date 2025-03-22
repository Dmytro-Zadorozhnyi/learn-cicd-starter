package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"simple":           {input: http.Header{"Authorization": []string{"ApiKey 123"}}, want: "123", err: nil},
		"missing header":   {input: http.Header{}, want: "", err: ErrNoAuthHeaderIncluded},
		"no ApiKey":        {input: http.Header{"Authorization": []string{"123"}}, want: "", err: errors.New("malformed authorization header")},
		"multiple headers": {input: http.Header{"Authorization": []string{"ApiKey 456", "789"}}, want: "456", err: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)

			// Compare result using cmp for detailed diff
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Mismatch (-want +got):\n%s", diff)
			}

			// Compare error messages if both expected and actual are not nil
			if (err != nil || tc.err != nil) && (err == nil || tc.err == nil || err.Error() != tc.err.Error()) {
				t.Errorf("Error mismatch (-want +got):\n%s", cmp.Diff(tc.err, err))
			}
		})
	}
}
