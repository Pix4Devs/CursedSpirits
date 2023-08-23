package globals

var (
	REFS    = make([]string, 0)
	ACCEPTS = make([]string, 0)
	PROXIES = make([]string, 0)

	TARGET          string
	CONCURRENCY     int
	DURATION        int
	PROXY_TYPE      string
	VALID_PROTOCOLS = []string{"socks4", "socks5"}
)