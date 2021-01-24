package main_test

import (
	"testing"

	"github.com/Hybrid-Performance-Method/interval/interval"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	t.Parallel()
	assert.Equal(t, interval.ParseRequirements("test==1.1.0 test==2.0.0")[1], "test==2.0.0")
	assert.Equal(t, interval.ParseRequirements("test1, test2, test3")[2], "test3")
}
