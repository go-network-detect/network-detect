package detector

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestOnline(t *testing.T) {
  d := NewNetworkDetector()
  value, err := d.up()
  assert.Nil(t, err)
  assert.True(t, value)
}
