package torion

type Torrc struct {
	Options       Options
	Client        Client
	Server        Server
	DataDir       DataDir
	HiddenService HiddenService
}

func NewTorrcConf() *Torrc {
	return &Torrc{}
}
