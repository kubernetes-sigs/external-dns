package profile

type HttpProfile struct {
<<<<<<< HEAD
	ReqMethod  string
	ReqTimeout int
	Scheme     string
	RootDomain string
	Endpoint   string
	// Deprecated, use Scheme instead
	Protocol string
}

func NewHttpProfile() *HttpProfile {
	return &HttpProfile{
		ReqMethod:  "POST",
		ReqTimeout: 60,
		Scheme:     "HTTPS",
		RootDomain: "",
		Endpoint:   "",
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	ReqMethod     string
	ReqTimeout    int
	Scheme        string
	RootDomain    string
	Endpoint      string
	ApigwEndpoint string
	// Deprecated, use Scheme instead
	Protocol string
	Proxy    string
}

func NewHttpProfile() *HttpProfile {
	return &HttpProfile{
		ReqMethod:     "POST",
		ReqTimeout:    60,
		Scheme:        "HTTPS",
		RootDomain:    "",
		Endpoint:      "",
		ApigwEndpoint: "",
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
}
