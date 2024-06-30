# bun

- 生のSQLもORM的な書き方もできる

## テスト

- いい感じでできそう

## マイグレーション

```sql
ALTER TABLE users ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP();
```

- entityは修正しないでOK
- entity修正しても、fixtureはいじらなくても良さそう
