/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package goku_test

import (
	"github.com/fathalfath30/goku"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoku(t *testing.T) {
	t.Run("it can create new Goku with default configuration",
		func(t *testing.T) {
			gk, err := goku.New(nil)

			assert.NotNil(t, gk)
			assert.Nil(t, err)

			assert.IsType(t, &goku.Goku{}, gk)

			// check default on each variable
			assert.Equal(t, string(goku.BaseUrlDevelopment), gk.BaseUrl())
			assert.Equal(t, goku.ModeDevelopment, gk.Mode())
			assert.True(t, gk.IsDevelopmentMode())
			assert.True(t, gk.LogEnabled())
		})
	t.Run("it can create new Goku with config",
		func(t *testing.T) {
			gk, err := goku.New(&goku.Config{
				Mode:      goku.ModeProduction,
				EnableLog: false,
			})

			assert.NotNil(t, gk)
			assert.Nil(t, err)

			assert.IsType(t, &goku.Goku{}, gk)

			// check default on each variable
			assert.Equal(t, string(goku.BaseUrlProduction), gk.BaseUrl())
			assert.Equal(t, goku.ModeProduction, gk.Mode())
			assert.False(t, gk.IsDevelopmentMode())
			assert.False(t, gk.LogEnabled())
		})
}
