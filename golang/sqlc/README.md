# sqlc

## Installation

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Commands

### `generate` - Generating Code

- SQLをパースし、結果を分析し、コードを生成する
  - ファイルはアルファベット順とか数字順にソートされて読み取られるっぽい
- デフォルトではビルトインのクエリ解析エンジンを使用する (早いけど複雑なクエリや型推論を扱うことはできない)
  - 実際のDB接続を用いてリッチな解析を行うことは可能
  - 実際のDBを用いた解析は以下のように `uri` を設定する
    - マイグレーションはやってくれないので、自身でup-to-dateにしておく必要がある

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "query.sql"
    schema: "schema.sql"
    database:
      uri: "postgres://postgres:${PG_PASSWORD}@localhost:5432/postgres"
    gen:
      go:
        out: "db"
        sql_package: "pgx/v5"
```
