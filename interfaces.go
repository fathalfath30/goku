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

package goku

type (
	IPayment interface{}

	IGoku interface {
		// BaseUrl will return base url for production or development depends
		// on config.Mode
		BaseUrl() string

		// Mode will return current mode
		Mode() Mode

		// IsDevelopmentMode will return true if config.Mode is "goku.ModeDevelopment"
		IsDevelopmentMode() bool

		// LogEnabled will return true if config.EnableLog is true
		LogEnabled() bool
	}
)
