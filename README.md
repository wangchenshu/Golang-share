# Golang-share 
一個分享 Golang 基礎的文件

以下範例參考: [https://openhome.cc/Gossip/Go](https://openhome.cc/Gossip/Go)

## Go 是什麼? 能吃嗎?
---
[Go](https://zh.wikipedia.org/wiki/Go)

Go（又稱Golang）是Google開發的一种静态强类型、編譯型、并发型，并具有垃圾回收功能的编程语言。

羅伯特·格瑞史莫，羅勃·派克（Rob Pike）及肯·汤普逊於2007年9月开始设计Go，稍後Ian Lance Taylor、Russ Cox加入專案。
Go是基於Inferno作業系統所開發的。Go於2009年11月正式宣布推出，成為開放原始碼專案，支援Linux、macOS、Windows等作業系統。
在2016年，Go被軟體評價公司TIOBE 選為「TIOBE 2016 年最佳語言」。

## 描述
---
Go的語法接近C語言，但對於變數的聲明有所不同。Go支援垃圾回收功能。
Go的平行計算模型是以東尼·霍爾的交談循序程式（CSP）為基礎，採取類似模型的其他語言套件括Occam和Limbo，
Go也具有這個模型的特徵，比如通道傳輸。通過 goroutine 和通道等並列構造可以建造執行緒池和管道等。
在1.8版本中開放外掛程式（Plugin）的支援，這意味著現在能從Go中動態載入部分函式。

與C++相比，Go 並不包括如列舉、例外處理、繼承、泛型、斷言、虛擬函式等功能，
但增加了 切片(Slice) 型、並行、管道、垃圾回收功能、介面等特性的語言級支援。
Go 2.0版本將支援泛型，對於斷言的存在，則持負面態度，同時也為自己不提供型別繼承來辯護。

不同於Java，Go原生提供了關聯陣列（也稱為雜湊表（Hashes）或字典（Dictionaries）），就像字串類型一樣。

## 來個 Hello, World 如何
```
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}
```

### 如何運行? 就是這麼簡單
```
$ go run main.go
```

### package 與 GOPATH
記得，package 中定義的函式，名稱必須是以大寫開頭，其他套件外的程式，
才能進行呼叫，若函式名稱是小寫，那麼會是套件中才可以使用的函式。

為了方便，通常會設定 GOPATH，例如，指向目前的工作目錄：
> set GOPATH=/home/walter/go

如果沒有設定 GOPATH 的話，Go 預設會是使用者目錄的 go 目錄。

### go build
如果想編譯原始碼為可執行檔，那麼可以使用 go build 
> go build main.go

### 本地套件
一個相應於 package 的目錄底下，可以有許多個原始碼，而每個原始碼開頭，
只要 package 。
例如，你可以有個原始碼是 hello.go，位於 hello 底下：
```
package hello

func SayHello() string {
        return "Hello, World!"
}
```

主程式就可以這麼引入:
```
package main

import (
        "fmt"
        "sample/hello"
)

func main() {
        // fmt.Println("Hello, World!")
        helloStr := hello.SayHello()
        fmt.Println(helloStr)
}
```

### 遠端套件
由於 Go 的 workspace 設置，都必須是如此規範，因此，
若你想將原始碼發佈給他人使用時就很方便，例如，
你可以建立 src/github.com/JustinSDK 目錄，
然後將方才的 goexample 目錄移到 src/github.com/JustinSDK 當中，
這麼一來，顯然地，你的 main.go import 就要改成：
```
import "github.com/walter/example"
```


## gofmt 格式化原始碼

如果你是個有點責任感的開發者，在新接觸一門語言的時候，
應該會問一個問題：「我該用什麼格式寫程式？」所以了，
在 Go 裏要用什麼格式寫程式？這個問題可以直接請 gofmt 來幫你解答。

... 因為時間關係，以下請自行查找


## go doc 註解即文件
### 註解即文件
實際上，go doc 的文件說明來自於原始碼中的註解，這樣的概念有點類似 Java 的 JavaDoc，
或者是 Python 的 DocString，不過 Go 的理念是讓它更簡單，不使用特殊標記，不使用特別的格式，
希望可以在沒有 go doc 的場合中，也可以藉由閱讀原始碼中的註解，輕易地得到文件說明。

當然，基本上還是要有一些約定，例如，在函式之前，緊接著函式的註解，中間沒有空白行，就是函式的文件說明來源。
> $ go doc hello/hello.go
```

package hello // import "sample/hello"

func SayHello() string

```


## 認識預定義型態
Go 包括了一些預先定義型態（Pre-declared Type），這包括了布林、數字與字串型態。

## 變數宣告、常數宣告
變數（Variable）是儲存值的位置，變數宣告可以給予識別名稱（Identifier）、型態與初始值，在 Go 中寫下的 10、3.14、true、"Justin" 等稱之為常數（Constant），常數宣告可給予這些常數識別名稱。

### 基本變數宣告
要在 Go 中進行變數宣告有多種形式，使用 var 是最基本的方式。例如，宣告一個 x 變數，型態為 int，初始值為 10：

` var x int = 10 `


這麼一來，從 x 這個位置開始，儲存了 int 長度的值 10，在宣告變數時，型態是寫在名稱之後。你也可以同時建立多個變數並指定初值：

`var x, y, z int = 10, 20, 30`


這樣的話，x、y、z 的型態都是 int，值分別是 10、20、30。如果宣告多個變數時，想要指定不同的型態，可以使用批量宣告：

```
var (
    x int = 10
    y string = "Walter"
    z bool = true
)
```

### 自動推斷型態
在 Go 中宣告變數並指定值時，可以不用提供型態，由編譯器自動推斷型態，例如：

`var x = 10`

上頭的 x 型態會是 int，而底下的宣告：

```
var x, y, z = 10, 3.14, "Walter"
x、y、z 的型態分別會是 int、float64 與 string，批量宣告時也可以自動推斷型態，例如：

var (
    x = 10        // int 型態
    y = 3.14      // float32 型態
    z = "Walter"  // string 型態
)
```

短變數宣告
在函式中，想要定義變數值的場合，可以使用短變數宣告，例如：

```
x := 10
y := 3.14
z := "Walter"
```

如果 x 是首次定義，就等於是宣告變數並指定值。上例也可以寫成一行：

`x, y, z := 10, 3.14, "Walter"`

由於 Go 的函式外，每個語法都必須以關鍵字開始，因此短變數宣告不能在函式外使用。

### 調換變數值
在 Go 中，要調換兩變數的值很簡單，例如底下的程式執行過後，x 會是 20，而 y 會是 10：

```
var x = 10
var y = 20
x, y = y, x
```

### 基本常數宣告
如一開始談到的，在 Go 中寫下的 10、3.14、true、"Justin" 等稱之為常數（Constant），常數宣告可給予這些常數識別名稱，要給予名稱時使用的是 const 關鍵字，例如：

`const x = 10`

正如〈認識預定義型態〉中談過的，10 會是一個整數常數，不過型態未定，如果要定義一個常數的型態，可以使用 int32()、int64() 之類的函式進行型態轉換，或者是在使用 const 宣告常數名稱時指定型態，例如：

`const x int32 = 10`

這邊的 10 就是 int32 型態了，注意，這邊的 x 並不是一個變數，而是一個識別名稱罷了，因此，會說 x 常數的型態是 int32，而不能說 x 變數的型態是 int32。

如果有多個常數要宣告，也可以批量宣告，例如：

```
const (
    x = 10
    y = 3.14
    z = "Walter"
)
```

再次地，x、y、z 分別是未定型態的整數、浮點數與字串常數（而不是 int、float64、string 這三個 Go 中定義的型態），如果你想要讓他們為已定義型態的整數、浮點數與字串常數，可以在宣告時指定型態：

```
const (
    x int = 10
    y float32 = 3.14
    z string = "Walter"
)
```

## 位元組構成的字串
> 在〈認識預定義型態〉中略略談過字串，表面看來，用雙引號（"）或反引號（`）
>
> 括起來的文字就是字串，預設型態為 string，實際在 Go 中，字串是由唯讀的 UTF-8 編碼位元組所組成。
>


### 字串入門
先從簡單的開始，在 Go 原始碼中，如果你撰寫 "Go語言" 這麼一段文字，那麼會產生一個字串，預設型態為 string，字串是唯讀的，一旦建立，就無法改變字串內容。

> 使用 string 宣告變數若無指定初值，預設是空字串 ""，可以使用 + 對兩個字串進行串接，
>
> 由於字串是唯讀的，因此實際上串接的動作，會產生新的字串，如果想比較兩個字串的相等性，
>
> 可以使用 ==、!=、<、<=、>、>= 依字典順序比較。


```
package main

import "fmt"

func main() {
    text1 := "Go語言"
    text2 := "Cool"
    var text3 string
    fmt.Println(text1 + text2) // Go語言Cool
    fmt.Printf("%q\n", text3)  // ""
    fmt.Println(text1 > text2) // true
}
```

... 詳細請自行爬文 ^_^


## 身為複合值的陣列

在 Go 中，陣列的長度固定，是個複合值，元素的型態及個數決定了陣列的型態，在記憶體中使用連續空間配置。

### 建立與存取陣列
建立陣列的方式是 [n]type，其中 n 為陣列的元素數量，type 是元素的型態。例如：

```
package main

import "fmt"

func main() {
    var scores [10]int
    scores[0] = 90
    scores[1] = 89
    fmt.Println(scores)      // [90 89 0 0 0 0 0 0 0 0]
    fmt.Println(len(scores)) // 10
}
```

### 走訪陣列
想要逐一走訪陣列的話，基本上可以使用 for 迴圈，例如：

```
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for i := 0; i < len(arr); i++ {
        fmt.Printf("%d\n", arr[i])
    }
}
```

另一個方式是使用 for range：

```
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3}
    for index, element := range arr {
        fmt.Printf("%d: %d\n", index, element)
    }
}
```

## 底層為陣列的 slice

在〈身為複合值的陣列〉中看過陣列，有的場合需要陣列，然而，若只想處理陣列中某片區域，或者以更高階的觀點看待一片資料（而不是從固定長度的陣列觀點），那麼可以使用 slice。

### 建立一個 slice
> 如果需要一個 slice，可以使用 make 函式，舉個例子來說，可以如下建立一個長度與容量皆為 5 的 slice，並傳回 slice 的參考，型態為 []int：

```
package main

import "fmt"

func main() {
    s1 := make([]int, 5)
    s2 := s1
    fmt.Println(s1) // [0 0 0 0 0]
    fmt.Println(s2) // [0 0 0 0 0]
    s1[0] = 1
    fmt.Println(s1) // [1 0 0 0 0]
    fmt.Println(s2) // [1 0 0 0 0]
    s2[1] = 2
    fmt.Println(s1) // [1 2 0 0 0]
    fmt.Println(s2) // [1 2 0 0 0]
}
```

... 詳細請自行爬文 ^_^

## 成對鍵值的 map
> 許多語言中都會有的成對鍵值資料結構，在 Go 中是以內建型態 map 來實作，格式為 map[keyType]valueType。

### 建立與初始 map
想要建立例一個 map 實例，但尚無任何鍵值對，可以使用 make 函式，例如：

```
package main

import "fmt"

func main() {
    passwords := make(map[string]int)
    fmt.Println(passwords)      // map[]
    fmt.Println(len(passwords)) // 0

    passwords["caterpillar"] = 123456
    passwords["monica"] = 54321
    fmt.Println(passwords)                // map[caterpillar:123456 monica:54321]
    fmt.Println(len(passwords))           // 2
    fmt.Println(passwords["caterpillar"]) // 123456
    fmt.Println(passwords["monica"])      // 54321
}
```

在上例中，passwords 是個參考，指向 make(map[string]int) 建立的 map 實例。

... 詳細請自行爬文 ^_^

## 運算子

... 詳細請自行爬文 ^_^

## if..else、switch 條件式
在分支判斷的控制上，Go 提供了 if...else、switch 語法，相較於其他提供類似語法的語言，在 Go 中 if...else、switch 語法的相似性更高。

### if..else 語法
直接來看個 if..else 的實例：


```
package main

import "fmt"

func main() {
    input := 10
    remain := input % 2
    if remain == 1 {
        fmt.Printf("%d 為奇數\n", input)
    } else {
        fmt.Printf("%d 為偶數\n", input)
    }
}
```

在 Go 中，if 之後直接寫判斷式可以不用使用 () 括號，而 {} 是必要的，
這樣應該是比較能避免 Apple 曾經發生某個函式中有兩個連續縮排而引發的問題.

### switch 語法
實際上，對於上頭的範例，可以改用 switch 來撰寫，程式會更為簡潔：


```
package main

import "fmt"

func main() {
    var level rune
    score := 88

    switch score / 10 {
    case 10, 9:
        level = 'A'
    case 8:

        level = 'B'
    case 7:
        level = 'C'
    case 6:
        level = 'D'
    default:
        level = 'E'
    }
    fmt.Printf("得分等級：%c\n", level)
}


```

... 詳細請自行爬文 ^_^


## for 迴圈
在 Go 中唯一的迴圈語法是 for，然而，它也擔任了一些語言中 while 的功能，並可搭配 range 來使用。

### 有分號的 for
for 最基本的使用形式，與 C/C++、Java 等語言類似，具有初始式、條件式、後置式三個部份，中間使用分號加以區隔，不必使用 () 括號包住這三個式子，同樣地，for 迴圈本體一定要使用 {}。

初始式只執行一次，通常用來宣告或初始變數，若是宣告變數，可見範圍僅在 for 中。第一個分號後是每次執行迴圈本體前會執行一次，且必須是 true 或 false 的結果，true 就會執行迴圈本體，false 就會結束迴圈，第二個分號後，則是每次執行完迴圈本體後會執行一次。

實際來看個 for 迴圈範例，在文字模式下從 1 顯示到 10：


```
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```


... 詳細請自行爬文 ^_^


## 函式入門

在 Go 中要定義函式，是使用 func 來定義，其基本格式如下：

```
func funcName(param1 type1, param2 type2) (return1 type1, return2 type2) {
    // 一些程式碼...
    return value1, value2
}
```

### 定義函式
可以看到，Go 定義函式時，參數的型態宣告同樣地是放在名稱之後，如果多個參數有同樣的型態，
那麼只要最右邊同型態的名稱右方加上型態就可以了，比較特別的地方在於，可以有兩個以上的傳回值，且傳回值可以設定名稱。

來看個簡單的函式定義，

```
package main

import fmt

func SayHello() {
    fmt.Println("Hello")
}

func main() {
  SayHello()
}

```


