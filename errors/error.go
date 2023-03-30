package errors

import "net/http"

type ProxyError struct {
	ErrCode  int    `json:"err_code"`
	HttpCode int    `json:"-"`
	ErrMsg   string `json:"err_msg"`
}

// 服务器内部错误
var (
	UnknownErr         = NewProxyError(1, http.StatusInternalServerError, "")
	CacheGetErr        = NewProxyError(11, http.StatusInternalServerError, "Cache获取错误/未命中")
	CacheSetErr        = NewProxyError(12, http.StatusInternalServerError, "Cache设置错误")
	CacheNotFound      = NewProxyError(13, http.StatusInternalServerError, "Cache丢失")
	JsonMarshalErr     = NewProxyError(21, http.StatusInternalServerError, "生成Json失败")
	JsonUnMarshalError = NewProxyError(22, http.StatusInternalServerError, "解析Json失败")
)

// HTTP 远程资源获取器获取器
var (
	RemoteReplyErr        = NewProxyError(101, http.StatusBadGateway, "Bestdori返回异常")
	RemoteReplyTimeout    = NewProxyError(102, http.StatusGatewayTimeout, "Bestdori返回超时")
	RemoteReplyReject     = NewProxyError(103, http.StatusNotFound, "Bestdori未找到资源")
	RemoteReplyParseErr   = NewProxyError(104, http.StatusInternalServerError, "Bestdori返回解析失败")
	RemoteReplyReadErr    = NewProxyError(105, http.StatusInternalServerError, "Bestdori返回读取失败")
	RemoteRequestParseErr = NewProxyError(106, http.StatusInternalServerError, "Bestdori请求体解析错误")
)

// Bestdori 官方与自制谱面接口
var (
	ChartIDParseErr      = NewProxyError(201, http.StatusBadRequest, "谱面ID解析错误")
	DiffParseErr         = NewProxyError(202, http.StatusBadRequest, "谱面难度解析错误")
	MethodParseErr       = NewProxyError(203, http.StatusBadRequest, "谱面请求方法解析错误")
	ChartNotFound        = NewProxyError(211, http.StatusNotFound, "谱面未找到")
	BandNotFound         = NewProxyError(212, http.StatusNotFound, "乐团未找到")
	AssetTypeErr         = NewProxyError(213, http.StatusBadGateway, "乐曲资源类型错误")
	DirectionNoteTypeErr = NewProxyError(221, http.StatusInternalServerError, "无法识别侧划音符的标识符")
	BeatLessThanZero     = NewProxyError(222, http.StatusInternalServerError, "某个音符的节拍数小于0")
	BPMNotAtBeatZero     = NewProxyError(223, http.StatusInternalServerError, "BPM不在Beat 0上")
)

// Bestdori 谱面列表
var (
	OffsetParseErr = NewProxyError(301, http.StatusBadRequest, "谱面列表offset解析错误")
	LimitParseErr  = NewProxyError(302, http.StatusBadRequest, "谱面列表单页限制解析错误")
)

// Bestdori 反向代理
var (
	URLParseErr = NewProxyError(401, http.StatusBadGateway, "URL解析出错")
)

func (e *ProxyError) Error() string {
	return e.ErrMsg
}

func NewProxyError(errorCode int, httpCode int, errMsg string) *ProxyError {
	return &ProxyError{
		ErrCode:  errorCode,
		HttpCode: httpCode,
		ErrMsg:   errMsg,
	}
}
