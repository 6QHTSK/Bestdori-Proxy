package controller

import (
	"Bestdori-Proxy/errors"
	"Bestdori-Proxy/models"
	"Bestdori-Proxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func getAssets(c *gin.Context) (assets models.AssetsURL, err error) {
	strID := c.Param("chartID")
	chartID, err := strconv.Atoi(strID)
	if err != nil {
		return assets, errors.ChartIDParseErr
	}
	return service.FetchAssetsUrl(chartID)
}

func ReverseProxy(c *gin.Context, remote *url.URL, contentType string) {
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

func CoverProxy(c *gin.Context) {
	assets, err := getAssets(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	remote, err := url.Parse(assets.Cover)
	if err != nil {
		_ = c.Error(errors.URLParseErr)
		return
	}
	ReverseProxy(c, remote, "image/jpeg")
}

func MusicProxy(c *gin.Context) {
	assets, err := getAssets(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	remote, err := url.Parse(assets.Audio)
	if err != nil {
		_ = c.Error(errors.URLParseErr)
		return
	}
	ReverseProxy(c, remote, "audio/mp3")
}

func SonolusProxy(c *gin.Context) {
	var target = "https://servers.sonolus.com"
	path := c.Param("path")
	proxyUrl, _ := url.Parse(target + path)
	ReverseProxy(c, proxyUrl, "")
}
