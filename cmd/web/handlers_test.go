package main

import (
	"net/http"
	"testing"

	"snippetbox.rubacodes.com/internal/assert"
)

func TestPing(t *testing.T) {
	// Create a new instance of our application struct. For now, this just
	// contains a couple of mock loggers (which discard anything written to
	// them).
	app := newTestApplication(t)
	// We then use the httptest.NewTLSServer() function to create a new test
	// server, passing in the value returned by our app.routes() method as the
	// handler for the server. This starts up a HTTPS server which listens on a
	// randomly-chosen port of your local machine for the duration of the test.
	// Notice that we defer a call to ts.Close() so that the server is shutdown
	// when the test finishes.
	ts := newTestServer(t, app.routes())
	defer ts.Close()
	// The network address that the test server is listening on is contained in
	// the ts.URL field. We can use this along with the ts.Client().Get() method
	// to make a GET /ping request against the test server. This returns a
	// http.Response struct containing the response.
	code, _, body := ts.get(t, "/ping")

	// We can then check the value of the response status code and body using
	// the same pattern as before.
	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	// Create a new instance of our application struct which uses the mocked
	// dependencies.
	app := newTestApplication(t)
	// Establish a new test server for running end-to-end tests.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)
			t.Log(code, body)
			assert.Equal(t, code, tt.wantCode)
			if tt.wantBody != "" {
				assert.StringContains(t, body, tt.wantBody)
			}
		})
	}
}
