# GoとGraphQLのディレクトリ例
## ディレクトリ構成
```
.
├── Dockerfile               
├── README.md                 
└── api                       
    ├── Makefile              
    ├── cmd
    │   └── main.go           # アプリケーションのエントリーポイント
    ├── go.mod                # Goの依存関係管理ファイル
    ├── go.sum                # Goの依存関係のチェックサムファイル
    ├── gqlgen.yml            # GraphQL Code Generator (gqlgen) の設定ファイル
    ├── pkg                   # 共有パッケージディレクトリ
    │   ├── di
    │   │   └── wire.go       # 依存性注入の設定ファイル
    │   ├── graphql
    │   │   ├── generated
    │   │   │   └── generated.go  # gqlgenによって生成されたGraphQLコード
    │   │   └── model
    │   │       └── models_gen.go # gqlgenによって生成されたモデル定義
    │   └── orm
    │       └── model.go      # ORMのモデル定義ファイル
    ├── src                   # アプリケーションのソースコードディレクトリ
    │   ├── application       # アプリケーション層
    │   │   ├── dto
    │   │   │   ├── input.go  # 入力データ転送オブジェクト定義
    │   │   │   └── output.go # 出力データ転送オブジェクト定義
    │   │   └── usecase
    │   │       └── sample.go # ユースケースの実装例
    │   ├── domain            # ドメイン層
    │   │   ├── model
    │   │   │   └── sample.go # ドメインモデルの定義例
    │   │   └── repository
    │   │       └── sample.go # リポジトリインターフェースの定義例
    │   ├── infrastructure    # インフラストラクチャ層
    │   │   ├── connect
    │   │   │   └── db_connect.go  # データベース接続の設定
    │   │   └── repository
    │   │       └── sample.go # リポジトリの実装例
    │   └── presentation      # プレゼンテーション層
    │       ├── converter
    │       │   └── output.go # 出力データの変換ロジック
    │       ├── dataloader
    │       │   └── sample.go # DataLoaderの実装例
    │       ├── directive
    │       │   └── sample.go # GraphQLディレクティブの実装例
    │       ├── middleware
    │       │   └── sample.go # ミドルウェアの実装例
    │       ├── resolver
    │       │   ├── mervel.go 
    │       │   ├── resolver.go 
    │       │   └── schema.go 
    │       └── server.go     # HTTPサーバーの設定と起動ファイル
    └── schema                # GraphQLスキーマ定義ディレクトリ
        ├── directive.graphql 
        ├── mervel.graphql    
        └── schema.graphql    

```
## 注意
- sample.goという名前はそれぞれの適切なものに書き換えてくださ
- あくまでアーキテクチャ例なので,サンプルコードはありません