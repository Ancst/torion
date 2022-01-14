package torion

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
		TorrcFromConf: &Torrc{
			Port: 9050,
		},
	}
}

// Fire up the TORION
func (config *Configuration) Fire() {

}
