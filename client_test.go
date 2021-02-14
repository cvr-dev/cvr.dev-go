package cvr_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cvr-dev/cvr.dev-go"
	"github.com/stretchr/testify/require"
)

// TestExpectedAPIKey verifies that the client correctly authorizes with the
// server for all endpoints.
func TestExpectedAPIKey(t *testing.T) {
	const expectedAPIKey = "sekkrit-api-key"

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, expectedAPIKey, r.Header.Get(cvr.HTTPAuthorizationHeader))

		w.Write([]byte("null"))
	}))
	defer s.Close()

	c := cvr.NewClientBaseAddress(expectedAPIKey, s.URL)

	tests := map[string]struct {
		apiCall  func() error
		endpoint string
	}{
		"TestAPIKey": {
			apiCall: func() error {
				return c.TestAPIKey()
			},
		},
		"CVRVirksomhederByCVRNumre": {
			apiCall: func() error {
				_, err := c.CVRVirksomhederByCVRNumre(1234)
				return err
			},
		},
		"CVRVirksomhederByNavn": {
			apiCall: func() error {
				_, err := c.CVRVirksomhederByNavn("navn")
				return err
			},
		},
		"ProduktionsenhederByPNumre": {
			apiCall: func() error {
				_, err := c.CVRProduktionsenhederByPNumre(1234)
				return err
			},
		},
		"ProduktionsenhederByAdresse": {
			apiCall: func() error {
				_, err := c.CVRProduktionsenhederByAdresse("adresse")
				return err
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.NoError(t, test.apiCall())
		})
	}
}

// TestTestAPIKeyResponses verifies that TestAPIKey returns true when the server
// responds with HTTP status ok.
func TestTestAPIKeyResponses(t *testing.T) {
	tests := map[string]struct {
		statusCode int
		expected   error
	}{
		"success": {
			statusCode: http.StatusOK,
			expected:   nil,
		},
		"failure": {
			statusCode: http.StatusUnauthorized,
			expected:   cvr.ErrUnauthorized,
		},
		"internal server error": {
			statusCode: http.StatusInternalServerError,
			expected:   cvr.ErrServerError{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
			}))
			defer s.Close()

			c := cvr.NewClientBaseAddress("some-api-key", s.URL)
			err := c.TestAPIKey()
			require.IsType(t, test.expected, err)
		})
	}
}

// TestExpectedEndpoints verifies that all client calls send requests to the
// expected endpoint.
func TestExpectedEndpoints(t *testing.T) {
	tests := map[string]struct {
		expectedEndpoint string
		handler          func(c *cvr.Client) error
	}{
		"TestAPIKeyExpectedURL": {
			expectedEndpoint: "/api/test/apikey",
			handler: func(c *cvr.Client) error {
				return c.TestAPIKey()
			},
		},
		"CVRVirksomhederByCVRNumre": {
			expectedEndpoint: "/api/cvr/virksomhed",
			handler: func(c *cvr.Client) error {
				_, err := c.CVRVirksomhederByCVRNumre(1234)
				return err
			},
		},
		"CVRVirksomhederByNavn": {
			expectedEndpoint: "/api/cvr/virksomhed",
			handler: func(c *cvr.Client) error {
				_, err := c.CVRVirksomhederByNavn("navn")
				return err
			},
		},
		"CVRProduktionsenhederByPNumre": {
			expectedEndpoint: "/api/cvr/produktionsenhed",
			handler: func(c *cvr.Client) error {
				_, err := c.CVRProduktionsenhederByPNumre(1234)
				return err
			},
		},
		"CVRProduktionsenhederByAdresse": {
			expectedEndpoint: "/api/cvr/produktionsenhed",
			handler: func(c *cvr.Client) error {
				_, err := c.CVRProduktionsenhederByAdresse("adresse")
				return err
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, test.expectedEndpoint, r.URL.Path)

				// Return valid JSON so that endpoints don't error because of
				// missing data
				w.Write([]byte("null"))
			}))
			defer s.Close()

			c := cvr.NewClientBaseAddress("some-api-key", s.URL)
			err := test.handler(c)
			require.NoError(t, err)
		})
	}
}

// TestCVRVirksomhedByCVRNumreResponses verifies that expected HTTP status codes
// are handled, and that returned data is deserialized as expected.
func TestCVRVirksomhedByCVRNumreResponses(t *testing.T) {
	tests := map[string]struct {
		virksomheder []cvr.Virksomhed
		statusCode   int
		expected     error
	}{
		"success": {
			virksomheder: []cvr.Virksomhed{
				{CVRNummer: 1234},
				{CVRNummer: 2345},
				{CVRNummer: 3456},
			},
			statusCode: http.StatusOK,
			expected:   nil,
		},
		"not found": {
			virksomheder: nil,
			statusCode:   http.StatusNotFound,
			expected:     cvr.ErrNotFound,
		},
		"unauthorized": {
			virksomheder: nil,
			statusCode:   http.StatusUnauthorized,
			expected:     cvr.ErrUnauthorized,
		},
		"server error": {
			virksomheder: nil,
			statusCode:   http.StatusInternalServerError,
			expected:     cvr.ErrServerError{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
				bs, err := json.Marshal(test.virksomheder)
				require.NoError(t, err)
				w.Write(bs)
			}))
			defer s.Close()

			c := cvr.NewClientBaseAddress("some-api-key", s.URL)
			vs, err := c.CVRVirksomhederByCVRNumre(0)
			require.IsType(t, test.expected, err)
			require.Equal(t, len(test.virksomheder), len(vs))

			for i, v := range test.virksomheder {
				require.Equal(t, test.virksomheder[i].CVRNummer, v.CVRNummer)
			}
		})
	}
}

// TestCVRVirksomhedByNavnRNumreResponses verifies that expected HTTP status
// codes are handled, and that returned data is deserialized as expected.
func TestCVRVirksomhedByNavnRNumreResponses(t *testing.T) {
	tests := map[string]struct {
		virksomheder []cvr.Virksomhed
		statusCode   int
		expected     error
	}{
		"success": {
			virksomheder: []cvr.Virksomhed{
				{CVRNummer: 1234},
				{CVRNummer: 2345},
				{CVRNummer: 3456},
			},
			statusCode: http.StatusOK,
			expected:   nil,
		},
		"not found": {
			virksomheder: nil,
			statusCode:   http.StatusNotFound,
			expected:     cvr.ErrNotFound,
		},
		"unauthorized": {
			virksomheder: nil,
			statusCode:   http.StatusUnauthorized,
			expected:     cvr.ErrUnauthorized,
		},
		"server error": {
			virksomheder: nil,
			statusCode:   http.StatusInternalServerError,
			expected:     cvr.ErrServerError{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
				bs, err := json.Marshal(test.virksomheder)
				require.NoError(t, err)
				w.Write(bs)
			}))
			defer s.Close()

			c := cvr.NewClientBaseAddress("some-api-key", s.URL)
			vs, err := c.CVRVirksomhederByNavn("")
			require.IsType(t, test.expected, err)
			require.Equal(t, len(test.virksomheder), len(vs))

			for i, v := range test.virksomheder {
				require.Equal(t, test.virksomheder[i].CVRNummer, v.CVRNummer)
			}
		})
	}
}

// TestCVRProduktionsenhedByPNumreResponses verifies that expected HTTP status
// codes are handled, and that returned data is deserialized as expected.
func TestCVRProduktionsenhedByPNumreResponses(t *testing.T) {
	tests := map[string]struct {
		produktionsenheder []cvr.Produktionsenhed
		statusCode         int
		expected           error
	}{
		"success": {
			produktionsenheder: []cvr.Produktionsenhed{
				{PNummer: 1234},
				{PNummer: 2345},
				{PNummer: 3456},
			},
			statusCode: http.StatusOK,
			expected:   nil,
		},
		"not found": {
			produktionsenheder: nil,
			statusCode:         http.StatusNotFound,
			expected:           cvr.ErrNotFound,
		},
		"unauthorized": {
			produktionsenheder: nil,
			statusCode:         http.StatusUnauthorized,
			expected:           cvr.ErrUnauthorized,
		},
		"server error": {
			produktionsenheder: nil,
			statusCode:         http.StatusInternalServerError,
			expected:           cvr.ErrServerError{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
				bs, err := json.Marshal(test.produktionsenheder)
				require.NoError(t, err)
				w.Write(bs)
			}))
			defer s.Close()

			c := cvr.NewClientBaseAddress("some-api-key", s.URL)
			ps, err := c.CVRProduktionsenhederByPNumre(0)
			require.IsType(t, test.expected, err)
			require.Equal(t, len(test.produktionsenheder), len(ps))

			for i, p := range test.produktionsenheder {
				require.Equal(t, test.produktionsenheder[i].PNummer, p.PNummer)
			}
		})
	}
}

func TestCVRProduktionsenhedByAdresseResponses(t *testing.T) {
	tests := map[string]struct {
		produktionsenheder []cvr.Produktionsenhed
		statusCode         int
		expected           error
	}{
		"success": {
			produktionsenheder: []cvr.Produktionsenhed{
				{PNummer: 1234},
				{PNummer: 2345},
				{PNummer: 3456},
			},
			statusCode: http.StatusOK,
			expected:   nil,
		},
		"not found": {
			produktionsenheder: nil,
			statusCode:         http.StatusNotFound,
			expected:           cvr.ErrNotFound,
		},
		"unauthorized": {
			produktionsenheder: nil,
			statusCode:         http.StatusUnauthorized,
			expected:           cvr.ErrUnauthorized,
		},
		"server error": {
			produktionsenheder: nil,
			statusCode:         http.StatusInternalServerError,
			expected:           cvr.ErrServerError{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.statusCode)
				bs, err := json.Marshal(test.produktionsenheder)
				require.NoError(t, err)
				w.Write(bs)
			}))
			defer s.Close()

			c := cvr.NewClientBaseAddress("some-api-key", s.URL)
			ps, err := c.CVRProduktionsenhederByAdresse("")
			require.IsType(t, test.expected, err)
			require.Equal(t, len(test.produktionsenheder), len(ps))

			for i, p := range test.produktionsenheder {
				require.Equal(t, test.produktionsenheder[i].PNummer, p.PNummer)
			}
		})
	}
}
