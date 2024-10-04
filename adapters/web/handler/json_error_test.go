package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	message := "Hello JSON"
	result := jsonError(message)
	require.Equal(t, `{"message":"Hello JSON"}`, string(result))
}
