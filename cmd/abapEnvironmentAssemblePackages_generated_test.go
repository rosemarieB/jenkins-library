package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbapEnvironmentAssemblePackagesCommand(t *testing.T) {

	testCmd := AbapEnvironmentAssemblePackagesCommand()

	// only high level testing performed - details are tested in step generation procedure
	assert.Equal(t, "abapEnvironmentAssemblePackages", testCmd.Use, "command name incorrect")

}
