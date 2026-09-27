package fluxa

import (
	"errors"
	"net"
	"net/url"
	"strings"
)

func ValidateBaseURL(value string) error {
	_, err := parseBaseURL(value)
	return err
}

func parseBaseURL(value string) (*url.URL, error) {
	u, err := url.Parse(strings.TrimSpace(value))
	if err != nil || u == nil || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" || u.Path != "" && u.Path != "/" {
		return nil, errors.New("base URL must be a host URL without a path, query, or credentials")
	}
	if u.Scheme != "https" && !(u.Scheme == "http" && isLoopback(u.Hostname())) {
		return nil, errors.New("base URL must use HTTPS (HTTP is allowed for localhost)")
	}
	return u, nil
}

func isLoopback(host string) bool {
	if strings.EqualFold(host, "localhost") {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}
