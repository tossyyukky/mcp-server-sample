package server

// Request はMCPリクエストの基本構造を定義します
type Request struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
	ID     interface{}            `json:"id,omitempty"`
}

// Response はMCPレスポンスの基本構造を定義します
type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  *Error      `json:"error,omitempty"`
	ID     interface{} `json:"id,omitempty"`
}

// Error はMCPエラーの構造を定義します
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorCode は標準的なMCPエラーコードを定義します
const (
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
)
