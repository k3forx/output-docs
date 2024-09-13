# casbin

# How it works

- Request: 
- Matcher: リクエストとポリシーのマッチングルールを定義する。結果は `p.eft` に保存される。
- Effect: Matcherの結果に基づいて論理的な判断を下す。 `e = some(where(p.eft == allow))` は、`p.eft` が1つでもallowであれば、trueを返す。