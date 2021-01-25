package main_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Hybrid-Performance-Method/interval/interval"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	t.Parallel()
	assert.Equal(t, interval.ParseRequirements("test==1.1.0 test==2.0.0")[1], "test==2.0.0")
	assert.Equal(t, interval.ParseRequirements("test1, test2, test3")[2], "test3")

	paramsFile := "parameters.yml"
	assert.Equal(t, strings.HasPrefix(paramsFile, "parameters"), true)

	nb := "notebook.ipynb"
	assert.Equal(t, strings.HasSuffix(nb, "ipynb"), true)
}

func TestAddDateString(t *testing.T) {
	t.Parallel()
	testFile := "output.ipynb"

	// sample date and layout
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"

	d, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, interval.AddDateString(testFile, d), "output-2014-11-12.ipynb")
}
