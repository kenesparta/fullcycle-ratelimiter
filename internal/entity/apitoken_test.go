package entity

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateValue(t *testing.T) {
	at := &APIToken{}
	value, err := at.GenerateValue()

	require.NoError(t, err, "GenerateValue should not return an error")
	require.NotEmpty(t, value, "Generated value should not be empty")
	require.Len(t, value, 64, "Generated value should be 32 characters long")
}
