package cvr_test

import (
	"os"
	"testing"

	"github.com/cvr-dev/cvr.dev-go"
	"github.com/stretchr/testify/require"
)

func getAPIKey(t *testing.T) string {
	apiKey, ok := os.LookupEnv("CVR_DEV_TEST_API_KEY")
	require.True(t, ok, "failed to find valid API key in ENV")
	return apiKey
}

// TestTestAPIKey verifies that the client sends requests to the backend and
// correctly interprets the responses sent.
func TestTestAPIKey(t *testing.T) {
	if testing.Short() {
		t.Skip("Running short tests")
	}

	tests := map[string]struct {
		apiKey   string
		expected error
	}{
		"valid api key": {
			apiKey:   getAPIKey(t),
			expected: nil,
		},
		"invalid api key": {
			apiKey:   "invalid api key",
			expected: cvr.ErrUnauthorized,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := cvr.NewClient(test.apiKey)
			err := c.TestAPIKey()
			require.Equal(t, test.expected, err)
		})
	}
}

// TestCVRVirksomhedByNavn verifies that the client correctly sends requests to
// the backend and correctly interprets the responses sent.
func TestCVRVirksomhedByNavn(t *testing.T) {
	if testing.Short() {
		t.Skip("Running short tests")
	}

	tests := map[string]struct {
		navn     string
		assert   func([]cvr.Virksomhed) bool
		expected error
	}{
		"multiple virksomheder": {
			navn: "statsministeriet",
			assert: func(vs []cvr.Virksomhed) bool {
				return len(vs) > 0
			},
			expected: nil,
		},
		"not found": {
			navn: "kfgnkjdfgkjdgkdfhgkdjhgkdjgkjdfhgkjdfkdfngkdfjn",
			assert: func(vs []cvr.Virksomhed) bool {
				return len(vs) == 0
			},
			expected: cvr.ErrNotFound,
		},
	}

	apiKey := getAPIKey(t)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := cvr.NewClient(apiKey)
			vs, err := c.CVRVirksomhederByNavn(test.navn)
			require.True(t, test.assert(vs))
			require.IsType(t, test.expected, err)
		})
	}
}

// TestCVRVirksomhedByCVRNumre verifies that the client correctly sends requests
// to the backend and correctly interprets the responses sent.
func TestCVRVirksomhedByCVRNumre(t *testing.T) {
	if testing.Short() {
		t.Skip("Running short tests")
	}

	tests := map[string]struct {
		cvrNumre []int
		expected error
	}{
		"one virksomhed": {
			cvrNumre: []int{10103940},
			expected: nil,
		},
		"multiple virksomheder": {
			cvrNumre: []int{10103940, 10150817, 10213231},
			expected: nil,
		},
		"not found": {
			cvrNumre: []int{1337},
			expected: cvr.ErrNotFound,
		},
	}

	apiKey := getAPIKey(t)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := cvr.NewClient(apiKey)
			vs, err := c.CVRVirksomhederByCVRNumre(test.cvrNumre...)
			require.Equal(t, test.expected, err)

			if test.expected == nil {
				require.Equal(t, len(test.cvrNumre), len(vs))

				for i, v := range vs {
					require.Equal(t, test.cvrNumre[i], v.CVRNummer)
				}
			}
		})
	}
}

// TestProduktionsenhederByPNumre verifies that the client correctly sends
// requests to the backend and correctly interprets the responses sent.
func TestProduktionsenhederByPNumre(t *testing.T) {
	if testing.Short() {
		t.Skip("Running short tests")
	}

	tests := map[string]struct {
		pNumre   []int
		expected error
	}{
		"one produktionsenhed": {
			pNumre:   []int{1004862579},
			expected: nil,
		},
		"multiple produktionsenheder": {
			pNumre:   []int{1004862579, 1003388394, 1020852379},
			expected: nil,
		},
		"not found": {
			pNumre:   []int{1337},
			expected: cvr.ErrNotFound,
		},
	}

	apiKey := getAPIKey(t)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := cvr.NewClient(apiKey)
			vs, err := c.CVRProduktionsenhederByPNumre(test.pNumre...)
			require.Equal(t, test.expected, err)

			if test.expected == nil {
				require.Equal(t, len(test.pNumre), len(vs))

				for i, v := range vs {
					require.Equal(t, test.pNumre[i], v.PNummer)
				}
			}
		})
	}
}

// TestCVRProduktionsenhedByAdresse verifies that the client correctly sends
// requests to the backend and correctly interprets the responses sent.
func TestCVRProduktionsenhedByAdresse(t *testing.T) {
	if testing.Short() {
		t.Skip("Running short tests")
	}

	tests := map[string]struct {
		adresse  string
		assert   func([]cvr.Produktionsenhed) bool
		expected error
	}{
		"multiple produktionsenheder": {
			adresse: "Prins Jørgens Gård 11",
			assert: func(vs []cvr.Produktionsenhed) bool {
				return len(vs) > 0
			},
			expected: nil,
		},
		"not found": {
			adresse: "kfgnkjdfgkjdgkdfhgkdjhgkdjgkjdfhgkjdfkdfngkdfjn",
			assert: func(vs []cvr.Produktionsenhed) bool {
				return len(vs) == 0
			},
			expected: cvr.ErrNotFound,
		},
	}

	apiKey := getAPIKey(t)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := cvr.NewClient(apiKey)
			ps, err := c.CVRProduktionsenhederByAdresse(test.adresse)
			require.IsType(t, test.expected, err)
			require.True(t, test.assert(ps))
		})
	}
}
