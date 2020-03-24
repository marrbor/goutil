package http

import (
	"net/url"
	"path"
)

// GetURLPathBase returns base(last part) of URL. ex '/api/v1/users/12345' => 12345
func GetURLPathBase(u *url.URL) string {
	return path.Base(u.Path)
}
