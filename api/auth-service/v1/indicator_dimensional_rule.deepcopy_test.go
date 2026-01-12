package v1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndicatorDimensionalRuleSpec_DeepCopyInto(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "indicator_dimensional_rule_spec.json"))
	require.NoError(t, err)

	var a, b IndicatorDimensionalRuleSpec
	require.NoError(t, json.Unmarshal(data, &a))

	a.DeepCopyInto(&b)
	assert.Equal(t, a, b)
}

func TestIndicatorDimensionalRuleSpec_DeepCopy(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "indicator_dimensional_rule_spec.json"))
	require.NoError(t, err)

	a := new(IndicatorDimensionalRuleSpec)
	require.NoError(t, json.Unmarshal(data, a))

	b := a.DeepCopy()
	assert.Equal(t, a, b)
}
