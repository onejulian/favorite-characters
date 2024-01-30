package util_test

import (
	"favorite-characters/src/infraestructure/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmailValid(t *testing.T) {
	assert.True(t, util.IsEmailValid("ame@mail.co"))
	assert.False(t, util.IsEmailValid("ame@mail"))
	assert.False(t, util.IsEmailValid("ame@mail."))
}