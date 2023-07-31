# ドメイン駆動設計について (DDD)

- ドメイン駆動設計とは？

## 用語の整理

| 用語 | 意味 | 具体例 | 補足 |
| --- | --- | --- | --- |
| ドメイン | ソフトウェアが動くビジネスの領域 | ラーメン通販 | **何が含まれるかが重要** |
| ドメインモデル | ドメインに含まれる概念を抽象化したもの | ラーメン、お店、購入者、トッピング | ドメインに含まれる概念の**抽象化** |
| ドメインオブジェクト | ドメインモデルをコードに落とし込んだもの | | 値オブジェクトなど |

## ドメインオブジェクトの種類

### 値オブジェクト

- システム固有の値を表したオブジェクト

```go
// ラーメンの名前をRamenNameという値オブジェクトを使って表現する
type RamenName string

type Ramen struct {
	Name RamenName
}
```

- 値オブジェクトの性質
  - 不変である
  - 交換が可能である
  - 等価性によって比較される

#### 値オブジェクトの性質1: 不変である

```go
name := "こんちには"
name.ChangeTo("Hello") // こんなメソッドはない
fmt.Println(greet)
// Output: Hello
```

↑のようなメソッドがあれば、どこかで `greet` が"Hello"になったり、0が1になったりする。
値は不変であることが求められる。

↓のようなコードは書けるが、DDDの設計としては良くない `Ramen` 構造体が `ChangeName` のメソッドを持つ方が自然

```go
ramenName := "とんこつラーメン"
ramenName.ChangeName("しょうゆラーメン") // OK??
```

#### 値オブジェクトの性質2: 交換が可能である

- 代入して値を更新するのはOK

```go
ramenName := "とんこつラーメン"
ramenName = "しょうゆラーメン" // OK
```

#### 値オブジェクトの性質3: 等価性によって比較される

#### いつ値オブジェクトを使うのか？

- ルールが存在しているか？
  - ex. ラーメンの名前に文字数制限や使える単語を限定したいか？
- それ単体で取り扱いたいか？

#### 他の例

- 振る舞いを持った値オブジェクトも定義できる
  - **つまり、値オブジェクトはただのデータ構造体のことではなく、オブジェクトに対する操作の振る舞いとして一緒にまとめることで、自身に関するルールを持つドメインオブジェクトになる**

```go
func NewMoney(amount int, currency string) Money {
	return Money{amount: amount, currency: currency}
}

type Money struct {
	amount int
	currency string // 通貨、これもCurrencyみたいな文字列型の値オブジェクトにした方が良いかもしれない
}

func (m Money) Amount() int {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Add(money Money) (Money, error) {
	// 通貨が一致しているかチェックする
	// 一致している場合は新しいMoney構造体を、一致していない場合はエラーを返す
}
```

### エンティティ

- 値オブジェクトのこと。違いは同一性によって区別されるかどうか。
  - 属性の変更によって、値オブジェクトそのものが別物になるかどうか？
  - ex. ユーザーというドメインモデルに誕生日 (年齢) や身長といった属性を付け加えて扱う時、誕生日を迎えて年齢が変化したり、身長が変化したときに全く別のユーザーとして扱うべきか？ということ

- エンティティの性質
  - 可変である
  - 同じ属性であっても区別される
  - 同一性により区別される

#### エンティティの性質1: 可変である

- `RamenShop` 構造体の `Name` は可変

```go
type RamenShop struct {
	Name string
}

func (m Ramen) ChangeName(newName string) {
	// も字数制限のバリデーションがあるかもしれない
	m.Name = newName 
}
```

- ↓みたいな感じには書かない

```go
func NewRamenShop(name string) RamenShop {
	return RamenShop{Name: name}
}

func main() {
	ramenShop := NewRamenShop("天下一品")
	ramenShop = NewRamenShop("我馬")
}
```

- 交換 (代入) によって変更を表現するのではなく、振る舞いを通じて属性を変更させる

#### エンティティの性質2: 同じ属性であっても区別される

- 値オブジェクトは同じ属性であれば同じものとして扱える (ex. `Money` 構造体)
- エンティティは同じ属性であっても区別される
  - ex. 同じラーメン店の名前でも違うお店として区別する


#### 値オブジェクト vs エンティティ

- ライフサイクルが存在し、そこに連続性が存在するか
  - ex. ユーザーは作成されて、ユーザー名が変更されて、削除されるのでエンティティが良さそう

## DDDにおけるサービス

- ドメインサービス: ドメインのためのサービス
- アプリケーションサービス: アプリケーションのためのサービス

### ドメインサービス

- 値オブジェクトやエンティティに記述するのが不自然な振る舞いをドメインサービスに逃す
  - ex. ラーメンショップの名前は重複が許されないというドメインルールがあるとすると、それはドメインサービスにおくと良さそう

```go
type RamenShop struct {
	Name string
}

func (rs RamenShop) IsDuplicatedWith(rs RamenShop) bool {
	// ...
}

func main() {
	ramenShopGaba := NewRamenShop("我馬")
	ramenShopGaba.IsDuplicatedWith(ramenShopGaba) // 生成したオブジェクト自身に問い合わせする？
}
```

- ドメインサービスにロジックを逃す
  - ドメインサービスは値オブジェクトやエンティティと異なり、地震の振る舞いを変更するようなインスタンス特有の状態を持たない
  - **ドメインサービスの濫用はロジックの点在を促すので使う時は要注意**

```go
type UserService struct {}

func (us UserService) IsDuplicatedWith(rs ramenShop) bool {
	// ...
}

func main() {
	ramenShopGaba := NewRamenShop("我馬")
	UserService{}.IsDuplicatedWith(ramenShopGaba)
}
```

### アプリケーションサービス

- ラーメンショップの登録のコード

```go
type UserService struct {}

func (us UserService) IsDuplicatedWith(rs ramenShop) bool {
	// DBに問い合わせる
}

func main() {
	ramenShopGaba := NewRamenShop("我馬")
	if UserService{}.IsDuplicatedWith(ramenShopGaba) {
		return 
	}
	// DBに永続化
}
```

## リポジトリ

- オブジェクトを繰り返し利用するには、何らかのデータストアにオブジェクトを永続化する必要がある
- リポジトリはデータを永続化し再構築するという処理を扱うためのオブジェクト