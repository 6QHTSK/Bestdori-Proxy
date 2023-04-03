package controller

import (
	"github.com/6QHTSK/Bestdori-Proxy/errors"
	"github.com/6QHTSK/Bestdori-Proxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func reverseProxy(c *gin.Context, remote *url.URL, contentType string) {
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(request *http.Request) {
		request.Header = c.Request.Header
		request.Host = remote.Host
		request.URL.Scheme = remote.Scheme
		request.URL.Host = remote.Host
		request.URL.Path = remote.Path
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		if contentType != "" {
			response.Header.Set("Content-Type", contentType)
		}
		return nil
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func CoverProxy(ctx *gin.Context) {
	server := ctx.GetInt("server")
	postID := ctx.GetInt("postID")

	assets, err := service.FetchAssetsUrl(postID, server)
	if abortWhenErr(ctx, err, err) {
		return
	}

	remote, err := url.Parse(assets.Cover)
	if abortWhenErr(ctx, err, errors.RemoteReplyURLParseErr) {
		return
	}

	reverseProxy(ctx, remote, "image/jpeg")
}

func AudioProxy(ctx *gin.Context) {
	server := ctx.GetInt("server")
	postID := ctx.GetInt("postID")

	assets, err := service.FetchAssetsUrl(postID, server)
	if abortWhenErr(ctx, err, err) {
		return
	}

	remote, err := url.Parse(assets.Audio)
	if abortWhenErr(ctx, err, errors.RemoteReplyURLParseErr) {
		return
	}

	reverseProxy(ctx, remote, "audio/mp3")
}
