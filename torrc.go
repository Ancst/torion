package torion

type Torrc struct {
	Options       Options
	Client        Client
	Server        Server
	DataDir       DataDir
	HiddenService HiddenService
}

type Options struct{}

type Client struct {
	ExcludeNodes []string
	EntryNodes   []string
	ExitNodes    []string
}

type Server struct{}

type DataDir struct{}

type HiddenService struct{}
