package pihole

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

func (p *PiholeProvider) retrieveNewToken(ctx context.Context) error {
	if p.cfg.Password == "" {
		return nil
	}

	form := &url.Values{}
	form.Add("pw", p.cfg.Password)
	url := fmt.Sprintf("%s/admin/index.php?login", p.cfg.Server)
	log.Debugf("Fetching new token from %s", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	body, err := p.do(req)
	if err != nil {
		return err
	}
	defer body.Close()

	// If successful the request will redirect us to an HTML page with a hidden
	// div containing the token...The token gives us access to other PHP
	// endpoints via a form value.
	p.token, err = parseTokenFromLogin(body)
	return err
}

func parseTokenFromLogin(body io.ReadCloser) (string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	tokenNode := getElementById(doc, "token")
	if tokenNode == nil {
		return "", errors.New("could not parse token from login response")
	}

	return tokenNode.FirstChild.Data, nil
}

func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func hasID(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, "id")
		if ok && s == id {
			return true
		}
	}
	return false
}

func traverse(n *html.Node, id string) *html.Node {
	if hasID(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := traverse(c, id)
		if result != nil {
			return result
		}
	}

	return nil
}

func getElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id)
}
