package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ServerInfo は、サーバーの基本情報を保持する構造体です
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Server は、MCPサーバーの基本構造を定義します
type Server struct {
	info         ServerInfo
	capabilities map[string]interface{}
}

// NewServer は、新しいMCPサーバーインスタンスを作成します
func NewServer(info ServerInfo) *Server {
	return &Server{
		info: info,
		capabilities: map[string]interface{}{
			"resources": map[string]interface{}{},
			"tools":     map[string]interface{}{},
		},
	}
}

// handleStdin は標準入力からのメッセージを処理します
func (s *Server) handleStdin() error {
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		var message map[string]interface{}
		if err := decoder.Decode(&message); err != nil {
			return fmt.Errorf("failed to decode message: %v", err)
		}

		// TODO: メッセージの処理を実装

		response := map[string]interface{}{
			"status": "success",
			"data":   nil,
		}

		if err := encoder.Encode(response); err != nil {
			return fmt.Errorf("failed to encode response: %v", err)
		}
	}
}

func main() {
	server := NewServer(ServerInfo{
		Name:    "sample-mcp-server",
		Version: "0.1.0",
	})

	if err := server.handleStdin(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
