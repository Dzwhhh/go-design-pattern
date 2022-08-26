package design_pattern

type server interface {
	HandleRequest(string, string) (int, string)
}

type Application struct{}

func (a *Application) HandleRequest(url, method string) (int, string) {
	if url == "/create/user" && method == "POST" {
		return 200, "User Created"
	}
	return 404, "Not OK"
}

type Nginx struct {
	App               *Application
	MaxAllowedRequest int
	RateLimiter       map[string]int
}

func NewNginx(maxAllowed int) *Nginx {
	return &Nginx{
		App:               &Application{},
		MaxAllowedRequest: maxAllowed,
		RateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiter(url)
	if !allowed {
		return 403, "request over limit"
	}
	return n.App.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiter(url string) bool {
	if _, exist := n.RateLimiter[url]; !exist {
		n.RateLimiter[url] = 0
	}
	if n.RateLimiter[url] > n.MaxAllowedRequest {
		return false
	}
	n.RateLimiter[url] += 1
	return true
}
