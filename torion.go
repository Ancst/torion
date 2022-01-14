package torion

import "github.com/Ancst/torion/models"

// NewTorion returns a new blank Configuration instance, which is not configured and, you have to do it so.
func NewTorion() *Configuration {
	return &Configuration{}
}

// Default return Configuration instance configured for performing normal http requests over TOR.
func Default() *Configuration {
	return &Configuration{
		BootstrapTimeout: 0,
		MaxRetry:         0,
		DisableDownload:  false,
		WorkingDirectory: "",
		TorrcFromFile:    "",
		TorrcFromConf: &models.Torrc{
			Options:       models.Options{},
			Client:        models.Client{},
			Server:        models.Server{},
			DataDir:       models.DataDir{},
			HiddenService: models.HiddenService{},
		},
	}
}

// Fire up the TORION
func (config *Configuration) Fire() {

}
