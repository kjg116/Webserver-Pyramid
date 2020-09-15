package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AreTheyPyramids(t *testing.T) {
	// Bad Request
	payload := `{"words": ["this", "is", "sparta", banana"]}`
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("", "", bytes.NewBuffer([]byte(payload)))

	AreTheyPyramids(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Good Request
	payload = `{
		"words": [
			"banana",
			"bs"
		]
	}`
	w = httptest.NewRecorder()
	r, _ = http.NewRequest("", "", bytes.NewBuffer([]byte(payload)))
	q := url.Values{}
	q.Set("show_details", "true")
	r.URL.RawQuery = q.Encode()

	AreTheyPyramids(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_isThisAPyramid(t *testing.T) {
	testCases := []struct {
		word     string
		expected bool
	}{
		{
			word:     "this",
			expected: false,
		},
		{
			word:     "foo",
			expected: true,
		},
		{
			word:     "banana",
			expected: true,
		},
		{
			word:     "qwweeerrrrtttttyyyyyy",
			expected: true,
		},
		{
			word:     "do",
			expected: false,
		},
		{
			word:     "I",
			expected: false,
		},
	}
	for _, tc := range testCases {
		result, _ := isThisAPyramid(tc.word)
		assert.Equal(t, tc.expected, result)
	}
}

func Test_isThisPlusOne(t *testing.T) {
	assert.False(t, isThisPlusOne(2, 5))
	assert.True(t, isThisPlusOne(6, 5))
}

func Test_getBoolQueryParam(t *testing.T) {
	assert.False(t, getBoolQueryParam(""))
	assert.True(t, getBoolQueryParam("true"))
}
