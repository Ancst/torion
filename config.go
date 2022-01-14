package torion

import (
	"github.com/Ancst/torion/models"
	"time"
)

type Configuration struct {
	// BootstrapTimeout determine how long TOR bundle allowed to reach the network, if this timeout reach, TORION will panic(Bootstrapping timeout reached).
	//
	// if MaxRetry option is set then, this will not panic until all requested amount of re-try fails.
	//
	// NOTE: When using bridges for tor, the bundle may have need longer time to connect to the network. By default, TORION will wait for bundle to get (100% Bootstrapped).
	BootstrapTimeout time.Duration

	// MaxRetry is the number of max TORION retry in case of timeout reach or connection-failure.
	//
	// MaxRetry should only be used when BootstrapTimeout is set.
	MaxRetry int

	// DisableDownload will disable downloading tor compressed bundle from torproject.org and mirrors.
	//
	// DO NOT FORGET TO SET THE WorkingDirectory IF YOU WANT TO DISABLE DOWNLOADS.
	DisableDownload bool

	// WorkingDirectory will be used for downloading and extracting the compressed TOR bundle, by default it's set to (%TEMP%) on Windows and (CURRENT DIR OF PROGRAM EXECUTION) for linux/unix like operating systems.
	//
	// If DisableDownload options set then, the TORION will look for extracted ABS path of tor bundle executable, and it's required files in this directory.
	WorkingDirectory string

	// TorrcFromFile will load your bundle-configuration from given absolute torrc configuration file path.
	//
	// TorrcFromFile will always overwrite the TorrcFromConf option, this means that you have to only use one of these at the same time.
	TorrcFromFile string

	// TorrcFromConf allow you to pass the torrc file option in easy way and, without file and, you don't need to care about its syntax.
	//
	// TorrcFromConf will always overwrite by TorrcFromFile option, this means that you have to only use one of these at the same time.
	TorrcFromConf *models.Torrc
}
