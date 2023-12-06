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

const (
	BaseUrlProduction  BaseUrl = "https://api-sandbox.doku.com"
	BaseUrlDevelopment BaseUrl = "https://api.doku.com"

	ModeProduction  Mode = "PRODUCTION"
	ModeDevelopment Mode = "DEVELOPMENT"
)

type (
	Endpoint string
	BaseUrl  string
	Mode     string

	Config struct {
		Mode      Mode
		EnableLog bool
	}

	Goku struct {
		mode      Mode
		enableLog bool
		baseUrl   string
	}
)

func New(config *Config) (IGoku, error) {
	if config == nil {
		config = &Config{
			// set default mode to "DEVELOPMENT"
			Mode: ModeDevelopment,

			// as default log will be printed into console
			EnableLog: true,
		}
	}

	// setup base url based on config.Mode
	var baseUrl = BaseUrlDevelopment
	if config.Mode == ModeProduction {
		// for production mode
		baseUrl = BaseUrlProduction
	}

	return &Goku{
		mode:      config.Mode,
		baseUrl:   string(baseUrl),
		enableLog: config.EnableLog,
	}, nil
}

// BaseUrl will return base url for production or development depends
// on config.Mode
func (gk *Goku) BaseUrl() string {
	return gk.baseUrl
}

// Mode will return current mode
func (gk *Goku) Mode() Mode {
	return gk.mode
}

// IsDevelopmentMode will return true if config.Mode is "goku.ModeDevelopment"
func (gk *Goku) IsDevelopmentMode() bool {
	return gk.mode == ModeDevelopment
}

// LogEnabled will return true if config.EnableLog is true
func (gk *Goku) LogEnabled() bool {
	return gk.enableLog
}
