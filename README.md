# MCP Server Sample

このプロジェクトは、Model Context Protocol (MCP) に準拠したサーバーの基本実装です。MCPは、モデルとローカルで実行されているMCPサーバー間の通信を可能にするプロトコルです。

## 機能

- 標準入出力を介したJSONメッセージの処理
- MCPプロトコルに準拠した基本的な型定義
- エラーハンドリング機能

## プロジェクト構造

```
.
├── main.go              # メインサーバーの実装
├── go.mod              # Goモジュール定義
└── server/
    └── types.go        # MCPプロトコルの型定義
```

## 主要コンポーネント

### ServerInfo (main.go)

サーバーの基本情報を管理する構造体です。

```go
type ServerInfo struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}
```

### Server (main.go)

MCPサーバーの基本機能を提供する構造体です。

```go
type Server struct {
    info         ServerInfo
    capabilities map[string]interface{}
}
```

### MCP メッセージ型 (server/types.go)

- `Request`: クライアントからのリクエストを表現
- `Response`: サーバーからのレスポンスを表現
- `Error`: エラー情報を表現

## セットアップ

1. リポジトリのクローン:
```bash
git clone <repository-url>
cd mcp-server-sample
```

2. 依存関係のインストール:
```bash
go mod tidy
```

3. サーバーの実行:
```bash
go run main.go
```

## 拡張

このサーバーは以下の機能を追加することで拡張できます：

1. ツールの実装
   - 特定の機能を提供するツールの追加
   - ツールのパラメータとレスポンスの定義

2. リソースの実装
   - 静的/動的リソースの追加
   - リソーステンプレートの定義

3. メッセージハンドラー
   - 特定のメソッドに対するハンドラーの実装
   - カスタムエラーハンドリング

## 生成AI

このドキュメントは生成AIの支援を受けて作成されています。

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。
