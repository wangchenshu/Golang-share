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

## Go 的優缺點 - 引用了[這篇文章](https://zhuanlan.zhihu.com/p/68483743)
### 優點
+ 性能(機器代碼)
> Go 極其地快。其性能與 Java 或 C++相似。在我們的使用中，Go 一般比 Python 要快 30 倍。
+ 並發支持
> 這可能是GoLang最受好評的特色。它支持並發。它可以充分利用多核功能。 
>
> Go 使用 goroutine 來實現並發性，它提供了一個非常優雅的 goroutine 調度程序系統，
>
> 可以很容易地生成數百萬個 goroutine。堆棧使用也可以動態擴展/收縮，這使內存使用更加智能。這與 Java 線程不同，後者通常只允許創建數千個線程。
+ 支持 GC
> Go 的作者都有 C 背景，Go 有 C 基因。有25個關鍵詞但具有豐富的表達能力。它可以支持幾乎所有在其他語言中看到的功能，如繼承，重載，對像等。
+ 工具鏈
> 有許多易於使用的內置工具可以幫助開發人員編寫可維護和可讀的代碼。效率大大提高。
>
> 這些包括 gofmt，goimport等。它們可以使我們的代碼看起來標準化，並且可以簡化審查工作。
+ 本機C支持
> 您可以在Go程序中嵌入C代碼，以便可以使用許多功能強大的C庫。
### 缺點
+ 缺乏框架
> GoLang開發人員沒有重要的框架。但是有其他語言。 Ruby有Ruby on Rails，Python有Django，PHP有Laravel。
+ 錯誤處理
> 如果可能出現錯誤，Go 程序需要函數來返回錯誤。這可能導致錯誤跟踪丟失導致缺少有用的錯誤處理邏輯的問題。
>
> 有些工具可以幫助檢測這種錯誤，例如 errcheck 和 megacheck。但它們更像是解決方法。
>
> 開發人員還需要編寫大量的if塊來檢查錯誤並處理它，這使得代碼不那麼乾淨。

其它參考: 
> [Go語言的9大優勢和3大缺點](https://www.itread01.com/content/1523029088.html)
>
> [PHP轉Golang一些感想](https://www.jishuwen.com/d/2j9e/zh-tw)

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
由於 Go 的 workspace設置，都必須是如此規範，因此，
若你想將原始碼發佈給他人使用時就很方便，例如，
你可以建立 src/github.com/walter 目錄，
然後將方才的 goexample 目錄移到 src/github.com/walter 當中，
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
變數（Variable）是儲存值的位置，變數宣告可以給予識別名稱（Identifier）、型態與初始值，在 Go 中寫下的 10、3.14、true、"Walter" 等稱之為常數（Constant），常數宣告可給予這些常數識別名稱。

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
如一開始談到的，在 Go 中寫下的 10、3.14、true、"Walter" 等稱之為常數（Constant），常數宣告可給予這些常數識別名稱，要給予名稱時使用的是 const 關鍵字，例如：

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
    text1 := "Go 語言 "
    text2 := "Nice"
    var text3 string

    fmt.Println(text1 + text2) // Go 語言 Nice
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

    passwords["walter1"] = 123456
    passwords["walter2"] = 654321

    fmt.Println(passwords)                // map[walter1:123456 walter2: 654321]
    fmt.Println(len(passwords))           // 2

    fmt.Println(passwords["walter1"]) // 123456
    fmt.Println(passwords["walter2"]) // 654321
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
    input := 12
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
    score := 89

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

### 多個傳回值
Go 中允許多個傳回值，例如，定義一個函式，可傳回2個元素：


```
package main

import fmt

func GetNums() (int, int) {
    return 1, 2
}

func main() {
    num1, num2 := GetNums()
    fmt.Println(num1, num2)
}

```

### 可變參數
在呼叫方法時，若方法的引數個數事先無法決定該如何處理？
在 Go 中支援不定長度引數（Variable-length Argument），
可以輕鬆的解決這個問題。直接來看示範：


```
package main

import "fmt"

func Sum(numbers ...int) int {
    var sum int

    for _, number := range numbers {
        sum += number
    }

    return sum
}

func main() {
    fmt.Println(Sum(1, 2))          // 3
    fmt.Println(Sum(1, 2, 3))       // 6
    fmt.Println(Sum(1, 2, 3, 4))    // 10
    fmt.Println(Sum(1, 2, 3, 4, 5)) // 15
}

```

可以看到，要使用不定長度引數，宣告參數時要於型態關鍵字前加上 ...，
此參數本質上是個 slice，因此可以使用 for range 來走訪元素，
可接受可變長度的參數只能有一個，而必須是最後一個參數。

### 函式與指標
Go 語言有指標，因此，在變數傳遞就多了一種選擇，直接來看個例子，以下的執行結果會顯示 1：


```
package main

import "fmt"

func add1To(n int) {
    n = n + 1
}

func main() {
    number := 1
    add1To(number)
    fmt.Println(number) // 1
}
```


這應該沒有問題，因為傳遞的是變數值給 n，函式中 n 的值加上 1 之後，
再指定回給 n，這對 main 中的 number 變數毫無影響，因此函式結束後，
顯示 number 的值，仍舊是 1。

那麼來看下面這個例子：


```
package main

import "fmt"

func add1To(n *int) {
    *n = *n + 1
}

func main() {
    number := 1
    add1To(&number)
    fmt.Println(number) // 2
}
```


這次使用了 &number 取得 number 的位址值再傳遞給 n，也就是傳遞了變數位址值給 n，
函式中使用 *n 取得位址處的值，加上 1 後再將值存回原位址處，
因此，透過 main 函式中的 number 取得的值，也會是加 1 後的值。


## 一級函式

作為一門現代語言，Go 的特色之一是函式為一級函式（First-class function），可以作為值來進行傳遞。

### 函式作為值
例如你定義一個取最大值的函式 max，你可以將此函式作為值傳遞給 maximum：


```
package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func main() {
    maximum := max

    fmt.Println(max(10, 5))     // 10
    fmt.Println(maximum(10, 5)) // 10
}
```


可以看到，被 max 參考的函式，也被 maximum 參考著，因而，
現在透過 max 或者 maximum，都可以呼叫函式。

因為 Go 型態推斷能力的關係，上頭的 maximum 並不用宣告型態，
而可以直接參考 max 函式的型態，那麼，max 或者是 maximum 的型態是什麼呢？

```
package main

import "fmt"
import "reflect"

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func main() {
    maximum := max

    fmt.Println(reflect.TypeOf(max))     // func(int, int) int
    fmt.Println(reflect.TypeOf(maximum)) // func(int, int) int
}
```

可以看到，函式的型態包括了 func、參數型態與傳回值型態，但不用宣告函式、參數與傳回值的名稱。

### 宣告函式變數
你可以僅宣告一個變數可用來參考特定型態的函式，例如：

```
package main

import "fmt"

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func main() {
    var maximum func(int, int) int
    fmt.Println(maximum) // nil

    maximum = max
    fmt.Println(maximum(10, 5)) // 10
}
```

若想先宣告一個 maximum 變數，可以在之後參考 max 函式，
可以使用型態 func(int, int) int 來宣告，通常，宣告函式變數時，
若想免於冗長的函式型態宣告，可以使用 type 來定義一個新的型態名稱：

```
package main

import "fmt"

type BiFunc func(int, int) int // 定義了新型態

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func main() {
    var maximum BiFunc
    fmt.Println(maximum) // nil

    maximum = max
    fmt.Println(maximum(10, 5)) // 10
}
```

在上例中，BiFunc 是個新的定義型態（defined type），
底層型態（underlying type）為 func(int, int) int，
Go 會認定兩者屬於不同型態，因為新的型態會擁有新的名稱，
在 Go 1.9 前，這是避免冗長函式型態宣告的唯一方式。

不過，就這邊來說，實際上只是想要 func(int, int) int 能有個簡短一點的名稱，
從 Go 1.9 開始，可以為型態取別名，別名就只是同一型態的另一個名稱，：

```
package main

import "fmt"

type BiFunc = func(int, int) int // 型態別名宣告

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func main() {
    var maximum BiFunc
    fmt.Println(maximum) // nil

    maximum = max
    fmt.Println(maximum(10, 5)) // 10
}
```

在這邊，BiFunc 只是 func(int, int) int 的另一個名稱，而不是新的型態。

函式變數既然是個變數，也就可以對它取指標，例如：

```
package main

import "fmt"

type BiFunc = func(int, int) int

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func main() {
    var maximum BiFunc
    fmt.Println(&maximum) // 0x1040a130
    // fmt.Println(&max)
}
```

### 回呼應用
因為函式可以當作值傳遞，因此，對於函式中流程幾乎相同，
只有少數操作不同的情況，就可以將操作不同的部份以回呼（Callback）函式取代。
例如，可以設計一個 filter 函式，用來過濾出符合特定條件的值：

```
package main

import "fmt"

type Predicate = func(int) bool

func filter(origin []int, predicate Predicate) []int {
    filtered := []int{}

    for _, elem := range origin {
        if predicate(elem) {
            filtered = append(filtered, elem)
        }
    }

    return filtered
}

func greaterThan7(n int) bool {
    return n > 7
}

func lessThan5(n int) bool {
    return n < 5
}

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    fmt.Println(filter(data, greaterThan7))
    fmt.Println(filter(data, lessThan5))
}
```

在這個例子中，filter 函式重用了 for range 與 if 等流程，只要傳入過濾用的函式，就可以讓 filter 具有各種的過濾用途。

除了作為值傳遞之外，Go 的函式還可以是匿名函式，且具有閉包（Closure）的特性，這將在下一篇文件加以說明。

## 匿名函式與閉包
除了作為值傳遞之外，Go 的函式還可以是匿名函式，且具有閉包（Closure）的特性，由於 Go 具有指標，在理解閉包時反而容易一些了。

### 匿名函式
在〈一級函式〉中，我們看過函式可作為值傳遞的一個應用是，可將函式傳入另一函式作為回呼（Callback），除了傳遞具名的函式之外，有時會想要臨時建立一個函式進行傳遞，例如：


```
package main

import "fmt"

type Predicate = func(int) bool

func filter(origin []int, predicate Predicate) []int {
    filtered := []int{}
    for _, elem := range origin {
        if predicate(elem) {
            filtered = append(filtered, elem)
        }
    }
    return filtered
}

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(filter(data, func(elem int) bool {
        return elem > 5
    }))
    fmt.Println(filter(data, func(elem int) bool {
        return elem <= 6
    }))
}
```

這個函式與〈一級函式〉中最後一個範例的作用相同，不過這次傳遞了匿名函式給 filter，可以看到，匿名函式可使用 func 建立，同樣必須指定參數與傳回值型態。

你也可以在函式中建立匿名函式，並將之傳回：

```
package main

import "fmt"

type Func1 = func(int) int

func funcA() Func1 {
    x := 10
    return func(n int) int {
        return x + n
    }
}

func main() {
    fmt.Println(funcA()(2)) // 12
}
```

### 閉包
可以在函式中建立匿名函式，引發了一個有趣的事實，先來看個例子：

```
package main

import "fmt"

type Consumer = func(int)

func forEach(elems []int, consumer Consumer) {
    for _, elem := range elems {
        consumer(elem)
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    forEach(numbers, func(elem int) {
        sum += elem
    })
    fmt.Println(sum) // 15
}
```

乍看之下，似乎有點像是：

```
package main

import "fmt"

type Consumer = func(int)

func forEach(elems []int, consumer Consumer) {
    for _, elem := range elems {
        consumer(elem)
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, elem := range numbers {
        sum += elem
    }
    fmt.Println(sum) // 15
}
```

然而意義完全不同。在使用 forEach 函式的範例中，sum 變數被匿名函式包覆並傳入 forEach 之中，
在 forEach 執行迴圈的過程中，每次呼叫傳入的函式（被 consumer 參考），就會改變 sum 的值，因此，最後得到的是加總後的值 15。

實際上，使用 forEach 函式的範例中，建立了一個閉包，閉包本質上就是一個匿名函式，
sum 變數被閉包包覆，讓 sum 變數可以存活於閉包的範疇中，其實，更之前從 funcA 傳回函式的範例中，
也建立了閉包，funcA 的 x 區域變數被閉包包覆，因此，你執行傳回的函式時，即使 funcA 已執行完畢，
x 變數依然是存活著在傳回的閉包範疇中，所以，你指定的引數總是會與 x 的值進行相加。
在上面的範例中，執行 funcA 會傳回一個函式，這個傳回的函式會將接受的引數指定給參數 n，並與 x 的值進行相加，因此最後顯示結果為 12。

重點在於，閉包將變數本身關閉在自己的範疇中，而不是變數的值，可以用以下這個範例來做個示範：

```
package main

import "fmt"

type Getter = func() int
type Setter = func(int)

func x_getter_setter(x int) (Getter, Setter) {
    getter := func() int {
        return x
    }
    setter := func(n int) {
        x = n
    }
    return getter, setter
}

func main() {
    getX, setX := x_getter_setter(10)

    fmt.Println(getX()) // 10
    setX(20)
    fmt.Println(getX()) // 20
}
```

對 x_getter_setter 來說，x 參數也是變數，x_getter_setter 傳回了兩個匿名函式，這兩個匿名函式都形成了閉包，
將 x 變數關閉在自己的範疇中，因此，你使用了 setX(20) 改變了 x 的值，使用 getX() 時取得的值，就會是修改後的值。

### 閉包與指標
如果你寫過 JavaScript，對於方才的範例，應該不會陌生，也因為 JavaScript 的普及，現在開發者多半對閉包不會覺得神秘難解了，而對於「閉包將變數本身關閉在自己的範疇中，而不是變數的值」，也比較瞭解其應用所在。

由於 Go 語言有指標，我們可以將指標的值顯示出來，這代表著變數的位址值，來看看被閉包關閉的變數，到底是怎麼一回事好了：

```
package main

import "fmt"

type Getter = func() int
type Setter = func(int)

func x_getter_setter(x int) (Getter, Setter) {
    fmt.Printf("the parameter :\tx (%p) = %d\n", &x, x)

    getter := func() int {
        fmt.Printf("getter invoked:\tx (%p) = %d\n", &x, x)
        return x
    }
    setter := func(n int) {
        x = n
        fmt.Printf("setter invoked:\tx (%p) = %d\n", &x, x)
    }
    return getter, setter
}

func main() {
    getX, setX := x_getter_setter(10)

    fmt.Println(getX())
    setX(20)
    fmt.Println(getX())
}
```

這個範例與前一個範例類似，只不過呼叫函式時，都會顯示 x 變數的位址值與儲存值，一個執行結果是：
```
the parameter : x (0x104382e0) = 10
getter invoked: x (0x104382e0) = 10
10
setter invoked: x (0x104382e0) = 20
getter invoked: x (0x104382e0) = 20
20
```
看到了嗎？顯示的變數的位址值都是相同的，閉包將變數本身關閉在自己的範疇中，而不是變數的值，就是這麼一回事。

## defer、panic、recover

就許多現代語言而言，例外處理機制是基本特性之一，然而，例外處理是好是壞，一直以來存在著各種不同的意見，
在 Go 語言中，沒有例外處理機制，取而代之的，是運用 defer、panic、recover 來滿足類似的處理需求。

### defer 延遲執行
在 Go 語言中，可以使用 defer 指定某個函式延遲執行，那麼延遲到哪個時機？簡單來說，在函式 return 之前，例如：

```
package main

import "fmt"

func deferredFunc() {
    fmt.Println("deferredFunc")
}

func main() {
    defer deferredFunc()
    fmt.Println("Hello, 世界")
}
```

這個範例執行時，deferredFunc() 前加上了 defer，因此，會在 main() 函式 return 前執行，
結果就是先顯示了 "Hello, 世界"，才顯示 "deferredFunc"。

如果有多個函式被 defer，那麼在函式 return 前，會依 defer 的相反順序執行，也就是 LIFO。

### 使用 defer 清除資源
... 詳細請自行爬文 ^_^

### panic
... 詳細請自行爬文 ^_^

### recover 恢復流程
如果發生了 panic，而你必須做一些處理，可以使用 recover，這個函式必須在被 defer 的函式中執行才有效果，
若在被 defer 的函式外執行，recover 一定是傳回 nil。

... 詳細請自行爬文 ^_^

什麼時候該用 error？什麼時候該用 panic？在 Go 的慣例中，鼓勵你使用 error，明確地進行錯誤檢查，
然而，就如方才所言，巢狀且深層的呼叫時，使用 panic 會比較便於傳播錯誤，
就 Go 的慣例來說，是以套件為界限，於套件之中，必要時可以使用 panic，而套件公開的函式，
建議以 error 來回報錯誤，若套件公開的函式可能會收到 panic，建議使用 recover 捕捉，並轉換為 error。

## 結構入門

有些資料會有相關性，例如，一個 XY 平面上的點可以使用 (x, y) 座標來表示；
名稱、郵件位址、電話可能代表著一張名片上的資訊。將相關聯的資料組織在一起，
對於資料本身的可用性或者是程式碼的可讀性，都會有所幫助。

### struct 組織資料
Go 語言中有 struct，可以用來將相關的資料組織在一起，如果你學過 C 語言，這對你應該不陌生。舉個例子來說，相對於個別地存取 x、y 變數：

```
package main

import "fmt"

func main() {
    x := 10
    y := 20
    fmt.Printf("{%d %d}\n", x, y) // {10 20}

    x := 20
    y := 30
    fmt.Printf("{%d %d}\n", x, y) // {20 30}
}
```

若 x 與 y 變數，相當於 XY 平面上的 (x, y) 座標，那麼將之組織在一起同時存取會比較好：

```
package main

import "fmt"

func main() {
    point := struct{ x, y int }{10, 20}
    fmt.Printf("{%d %d}\n", point.x, point.y) // {10 20}

    point.x = 20
    point.y = 30

    fmt.Printf("{%d %d}\n", point.x, point.y) // {20 30}
}
```

實際上，fmt.Println 可以直接處理 struct，因此，上面的例子，可以直接使用 fmt.Println(point) 來得到相同的顯示結果。

在上面的例子中，struct 定義了一個結構，當中包括了 x 與 y 兩個值域（field），接著馬上用它來建立了一個實例，
依順序指定了 x 與 y 的值是 10 與 20，可以看到，想要存取結構的值域，可以運過點運算子（.）。

### 基於結構定義新型態
上面的例子中，建立了一個匿名型態的結構，你可以使用 type 基於 struct 來定義新型態，例如：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    point1 := Point{10, 20}
    fmt.Println(point1) // {10 20}

    point2 := Point{Y: 20, X: 30}
    fmt.Println(point2) // {30 20}
}
```

在上面基於結構定義了新型態 Point，留意到名稱開頭的大小寫，若是大寫的話，就可以在其他套件中存取，
這點對於結構的值域也是成立，大寫名稱的值域，才可以在其他套件中存取。在範例中也可以看到，
建立並指定結構的值域時，可以直接指定值域名稱，而不一定要按照定義時的順序。

如果一開始不知道結構的值域數值為何，可以使用 var 宣告即可，那麼值域會依型態而有適當的預設值。例如：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    var point Point
    fmt.Println(point)      // {0 0}
}
```

point 並不是參考，point 的位置開始，有一片可以儲存結構的空間，
可以使用 & 來取得 point 的位址值，point 的位址值無法改變。

### 結構與指標
如果你建立了一個結構的實例，並將之指定給另一個結構變數，那麼會進行值域的複製。例如：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    point1 := Point{X: 10, Y: 20}
    point2 := point1

    point1.X = 20

    fmt.Println(point1)  // {20, 20}
    fmt.Println(point2)  // {10 20}
}
```

這對於函式的參數傳遞也是一樣的：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

func changeX(point Point) {
    point.X = 20
    fmt.Println(point)
}

func main() {
    point := Point{X: 10, Y: 20}

    changeX(point)     // {20 20}
    fmt.Println(point) // {10 20}
}
```

point 的位置開始儲存了結構，可以對 point 使用 & 取值，將位址值指定給指標，
因此若指定或傳遞結構時，不是想要複製值域，可以使用指標。例如：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

func main() {
    point1 := Point{X: 10, Y: 20}
    point2 := &point1

    point1.X = 20

    fmt.Println(point1) // {20, 20}
    fmt.Println(point2) // &{20 20}
}
```

注意到 point2 := &point1 多了個 &，這取得了 point1 實例的指標值，並傳遞給 point2，point2 的型態是 *Point，
也就是相當於 var point2 *Point = &point1，因此，當你透過 point1.X 改變了值，透過 point2 就能取得對應的改變。

類似地，也可以在傳遞參數給函式時使用指標：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

func changeX(point *Point) {
    point.X = 20
    fmt.Printf("&{%d %d}\n", point.X, point.Y)
}

func main() {
    point := Point{X: 10, Y: 20}

    changeX(&point)    // &{20 20}
    fmt.Println(point) // {20 20}
}
```

可以看到在 Go 語言中，即使是指標，也可以直接透過點運算子來存取值域，
這是 Go 提供的語法糖，point.X 在編譯過後，會被轉換為 (*point).X。

你也可以透過 new 來建立結構實例，這會傳回結構實例的位址：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

func default_point() *Point {
    point := new(Point)
    point.X = 10
    point.Y = 10
    return point
}

func main() {
    point := default_point()
    fmt.Println(point) // &{10 10}
}
```

在這邊，point 是個指標，也就是 *Point 型態，儲存了結構實例的位址。

結構的值域也可以是指標型態，也可以是結構自身型態之指標，因此可實現鏈狀參考，例如：

```
package main

import "fmt"

type Point struct {
    X, Y int
}

type Node struct {
    point *Point
    next  *Node
}

func main() {
    node := new(Node)

    node.point = &Point{10, 20}
    node.next = new(Node)

    node.next.point = &Point{10, 30}

    fmt.Println(node.point)      // &{10 20}
    fmt.Println(node.next.point) // &{10 30}
}
```

$T{} 的寫法與 new(T) 是等效的，使用 &Point{10, 20} 這類的寫法，可以同時指定結構的值域。

## 結構與方法

在〈結構入門〉中看過，有些資料會有相關性，相關聯的資料組織在一起，對於資料本身的可用性或者是程式碼的可讀性，都會有所幫助，實際上，有些資料與可處理它的函式也會有相關性，將相關聯的資料與函式組織在一起，對資料與函式本身的可用性或者是程式碼的可讀性，也有著極大的幫助。

### 建立方法
假設可能原本有如下的程式內容，負責銀行帳戶的建立、存款與提款：

```
package main

import (
    "errors"
    "fmt"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func Deposit(account *Account, amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    account.balance += amount
}

func Withdraw(account *Account, amount float64) error {
    if amount > account.balance {
        return errors.New("餘額不足")
    }
    account.balance -= amount
    return nil
}

func String(account *Account) string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        account.id, account.name, account.balance)
}

func main() {
    account := &Account{"8765-4321", "Walter Wang", 1000}
    Deposit(account, 500)
    Withdraw(account, 200)
    fmt.Println(String(account)) // Account{8765-4321 Walter Wang 1300.00}
}
```

實際上，Desposit、Withdraw、String 的函式操作，都是與傳入的 Account 實例有關，何不將它們組織在一起呢？
這樣比較容易使用些，在 Go 語言中，你可以重新修改函式如下：

```
package main

import (
    "errors"
    "fmt"
)

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}

func (ac *Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

func (ac *Account) String() string {
    return fmt.Sprintf("Account{%s %s %.2f}",
        ac.id, ac.name, ac.balance)
}

func main() {
    account := &Account{"8765-4321", "Walter Wang", 1000}
    account.Deposit(500)
    account.Withdraw(200)
    fmt.Println(account.String()) // Account{8765-4321 Walter Wang 1300.00}
}
```

簡單來說，只是將函式的第一個參數，移至方法名稱之前成為函式呼叫的接收者（Receiver），這麼一來，
就可以使用 account.Deposit(500)、account.Withdraw(200)、account.String() 這樣的方式來呼叫函式，
就像是物件導向程式語言中的方法（Method）。

### 值都能有方法
實際上，不只是結構的實例可以定義方法，在 Go 語言中，只要是值，就可以定義方法，
條件是必須是定義的型態（defined type），具體而言，就是使用 type 定義的新型態。

例如，以下的範例為 []int 定義了一個新的型態名稱，並定義了一個 ForEach 方法：

```
package main

import "fmt"

type IntList []int
type Funcint func(int)

func (lt IntList) ForEach(f Funcint) {
    for _, ele := range lt {
        f(ele)
    }
}

func main() {
    var lt IntList = []int{10, 20, 30, 40, 50}
    lt.ForEach(func(ele int) {
        fmt.Println(ele)
    })
}
```

這個範例會顯示 10 到 50 作為結果，必須留意的是，type 定義了新型態 Funcint，
因為 ForEach 是針對 Funcint 定義，而不是針對 []int，因此底下是行不通的：

```
lt2 := []int {10, 20, 30, 40, 50}

// lt2.ForEach undefined (type []int has no field or method ForEach)
lt2.ForEach(func(ele int) {
    fmt.Println(ele)
})
```

編譯器認為 []int 並沒有定義 ForEach，因此發生錯誤，想要通過編譯的話，可以進行型態轉換：

```
lt2 := IntList([]int {10, 20, 30, 40, 50})
lt2.ForEach(func(ele int) {
    fmt.Println(ele)
})
```

你甚至可以基於 int 等基本型態定義方法，同樣地，必須定義一個新的型態名稱：

```
package main

import (
    "fmt"
)

type Int int
type FuncInt func(Int)

func (n Int) Times(f FuncInt) {
    if n < 0 {
        panic("必須是正數")
    }

    var i Int
    for i = 0; i < n; i++ {
        f(i)
    }
}

func main() {
    var x Int = 10
    x.Times(func(n Int) {
        fmt.Println(n)
    })
}
```

像這樣基於某個基本型態定義新型態，並為其定義更多高階特性，在 Go 的領域是常見的做法。
這個範例會顯示 0 到 9，看起來就像是指定函式，要求執行 x 次吧！

## 介面入門
在〈結構與繼承〉的最後討論到了多型，倘若現在需要有個函式，可以接受 Account 與 CheckingAccount 實例，
或者是有個陣列或 slice，可以收集 Account 與 CheckingAccount實例，那該怎麼辦呢？

### 介面定義行為
在 Go 語言中，可以使用 interface 定義行為，舉例來說，若現在想要定義儲蓄的行為，可以如下：

```
type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}
```

注意，不必使用 func 關鍵字，也不用宣告接受者型態，只需要定義行為的名稱、參數與傳回值。
接著該怎麼實現這個介面呢？實際上，就〈結構與繼承〉，已經實現了這個介面，也就是說，
結構上不用任何關鍵字，只要有函式實現這兩個行為就可以了。

因此，現在可以寫個函式，同時接受 Account 與 CheckingAccount 實例，在提款後顯示餘額：

```
package main

import (
    "errors"
    "fmt"
)

type Savings interface {
    Deposit(amount float64)
    Withdraw(amount float64) error
}

type Account struct {
    id      string
    name    string
    balance float64
}

func (ac *Account) Deposit(amount float64) {
    if amount <= 0 {
        panic("必須存入正數")
    }
    ac.balance += amount
}

func (ac *Account) Withdraw(amount float64) error {
    if amount > ac.balance {
        return errors.New("餘額不足")
    }
    ac.balance -= amount
    return nil
}

type CheckingAccount struct {
    Account
    overdraftlimit float64
}

func (ac *CheckingAccount) Withdraw(amount float64) error {
    if amount > ac.balance+ac.overdraftlimit {
        return errors.New("超出信用額度")
    }
    ac.balance -= amount
    return nil
}

func Withdraw(savings Savings) {
    if err := savings.Withdraw(500); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(savings)
    }
}

func main() {
    account1 := Account{"8765-4321", "Walter Wang", 1000}
    account2 := CheckingAccount{Account{"8765-4321", "Walter Wang", 1000}, 30000}
    Withdraw(&account1) // 顯示 &{8765-4321 Walter Wang 500}
    Withdraw(&account2) // 顯示 &{{8765-4321 Walter Wang 500} 30000}
}
```

雖然沒有定義接收者為 *CheckingAccount 的 Deposit 方法，然而，作為內部型態的 Account 有定義 
Deposit（並且沒有使用到 CheckingAccount 定義的值域），這個實現被提昇至外部型態，也就滿足了 Savings 要求的行為規範。

注意！由於在實作 Withdraw 與 Deposit 方法時，都是用指標 (ac *Account) 或 (ac *CheckingAccount) 
宣告了接受者型態，因此傳遞實例給 func Withdraw(savings Savings) 時，也就必須傳遞指標。

### 介面實例的型態與值
如果你定義了一個變數：

```
var savings Savings
```

那麼 savings 變數儲存了什麼？技術上來說，savings 變數儲存兩個資訊：型態與值。
就方才的savings 被指定為 nil 來說，代表著 savings 在底層儲存的型態為 nil，而值沒有指定，
這樣的介面實例稱為 nil interface，因為沒有型態資訊，也就不能透過 nil interface 呼叫方法。

### 空介面
那麼，如果想要建立一個實例容器，可以收集各種類型的實例，要怎麼做呢？
答案就是透過空介面，也就是沒有定義任何行為的 interface {}。

```
package main

import "fmt"

type Cat struct{}

func main() {
    instances := [](interface{}){
        &Cat{},
        [...]int{1, 2, 3, 4, 5, 6},
        map[string]int{"walter1": 123456, "walter2": 654321},
    }

    for _, instance := range instances {
        fmt.Println(instance)
    }
}
```

如果你查看 fmt.Println 的文件說明，可以發現，它的參數類型就是 interface {}：

```
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

## 型態斷言
宣告介面時使用的名稱，只是一個方便取用及閱讀的標示，最重要的是介面中定義的行為，
以及實際的接收者型態。因此，若你打算從一個介面轉換至另一個介面，只要行為符合就可以了。例如以下是可行的：

```
package main

import "fmt"

type ATester interface {
    test()
}

type BTester interface {
    test()
}

type Subject struct {
    name string
}

func (s *Subject) test() {
    fmt.Println(s)
}

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA
    testerA.test()
    testerB.test()
}
```

在第二個指定時，編譯器會檢查 testerA 的型態定義，也就是介面中，是否定義了 test() 行為，
若是則可通過編譯，若否就編譯錯誤。例如以下的情況：

```
package main

import "fmt"

type ATester interface {
    testA()
}

type BTester interface {
    testB()
}

type Subject struct {
    name string
}

func (s *Subject) testA() {
    fmt.Println(s)
}

func (s *Subject) testB() {
    fmt.Println(s)
}

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA // 錯誤：ATester does not implement BTester
    testerA.testA()
    testerB.testB()
}
```

就算 testerA 儲存的結構實例，確實有實作testB() 這個方法，然而從編譯器的角度來看，
testerA 的行為只有 testA()，而看不到它有 testB() 的行為，因此上面這個範例會編譯錯誤。

### Comma-ok 型態斷言

如果真的要通過編譯，可以使用型態斷言[（Type assertion）](https://golang.org/ref/spec#Type_assertions)：

```
...同前…略

func main() {
    var testerA ATester = &Subject{"Test"}
    var testerB BTester = testerA.(BTester)
    testerA.testA()
    testerB.testB()
}
```

x.(T) 這個語法，x 的型態是某介面，而 T 是預期的型態，或者是值實作的另一個介面名稱，
在〈介面入門〉中談過，介面底層儲存了型態與值的資訊，x.(T) 是在告知編譯器，
在執行時期再來斷言型態，也就是執行時期再來判斷 x 底層儲存的值，型態是否為 T，若是就傳回底層儲存的值。

型態斷言與型態轉換不同，型態轉換是將值的型態轉換為另一型態，編譯器會檢查兩個型態的資料結構是否相同，
若否會發生編譯錯誤。

斷言是執行時期進行的，在底下的範例中，執行時期會斷言 value 底層儲存的值，其型態為 Cat：

```
package main

import "fmt"

type Cat struct {
    name string
}

func main() {
    values := [...](interface{}){
        Cat{"Walter1"},
        Cat{"Walter2"},
    }

    for _, value := range values {
        cat := value.(Cat)
        fmt.Println(cat.name)
    }
}
```

如果 value 底層儲存的值，其型態為實際上不是 Cat 型態，那麼操作 cat 時會發生執行時期錯誤，
為了避免這類錯誤發生，可以進行 Comma-ok 型態斷言，例如：

```
package main

import "fmt"

type Cat struct {
    name string
}

func main() {
    values := [...](interface{}){
        Cat{"Walter1"},
        Cat{"Walter2"},
        [...]int{1, 2, 3, 4, 5, 6},
        map[string]int{"walter1": 123456, "walter2": 654321},
    }

    for _, value := range values {
        if cat, ok := value.(Cat); ok {
            fmt.Println(cat.name)
        }
    }
}
```

第一個 cat 變數是 Cat 型態，若 value 底層儲存的值確實是 Cat 型態，ok 變數會是 true，
否則 ok 會是 false，因此，在上面的例子中，只會針對 Cat 顯示其 name 的值。


### 型態 switch 測試
依照上面的說明，如果想測試多個型態，可以用多個 if...else if，例如：

```
package main

import "fmt"

type Cat struct {
    name string
}

func main() {
    values := [...](interface{}){
        Cat{"Walter1"},
        Cat{"Walter2"},
        [...]int{1, 2, 3, 4, 5},
        map[string]int{"walter1": 123456, "walter2": 654321},
        10,
    }

    for _, value := range values {
        if cat, ok := value.(Cat); ok {
            fmt.Println(cat.name)
        } else if arr, ok := value.([5]int); ok {
            fmt.Println(arr)
        } else if passwds, ok := value.(map[string]int); ok {
            fmt.Println(passwds)
        } else if i, ok := value.(int); ok {
            fmt.Println(i)
        } else {
            fmt.Println("非預期之型態")
        }
    }
}
```

不過，針對這個情況，使用型態 switch 測試會更為適合：

```
package main

import "fmt"

type Cat struct {
    name string
}

func main() {
    values := [...](interface{}){
        Cat{"Walter1"},
        Cat{"Walter2"},
        [...]int{1, 2, 3, 4, 5, 6},
        map[string]int{"walter1": 123456, "walter2": 654321},
        10,
    }

    for _, value := range values {
        switch v := value.(type) {
        case Cat:
            fmt.Println(v.name)
        case [5]int:
            fmt.Println(v[0])
        case map[string]int:
            fmt.Println(v["walter1"])
        case int:
            fmt.Println(v)
        default:
            fmt.Println("非預期之型態")
        }
    }
}
```

value.(type) 這樣的語法，只能用在 switch 之中。

## 介面組合
詊請請參考: [介面組合](https://openhome.cc/Gossip/Go/InterfaceComposition.html)

## Go 熱門的 Web Framework
Go 是一個快速增長的開源編程語言，用於構建簡單、快速和可靠的軟件。點這裡看有哪些大公司在使用Go語言來構建他們的服務。

本文提供了所有必要的信息，以幫助開發人員了解使用Go語言開發Web應用程序的最佳選項。 。

本文包含了最詳細的框架比較，通過盡可能多的角度(人氣，社區支持，內置功能等)來比較最知名的幾個Web 框架。

### Beego
一個Go語言下開源的，高性能Web框架

Beego 是一個快速開發 Go 應用的 HTTP 框架，他可以用來快速開發 API、Web 及後端服務等各種應用，

是一個 RESTful 的框架，主要設計靈感來源於 tornado、sinatra 和 flask 這三個框架，

但是結合了 Go 本身的一些特性（interface、struct 嵌入等）而設計的一個框架。

- https://github.com/astaxie/beego
- https://beego.me

### Buffalo
一個Go語言下快速Web開發框架

- https://github.com/gobuffalo/buffalo
- https://gobuffalo.io

### Echo

一個高性能，極簡的Web框架

- https://github.com/labstack/echo
- https://echo.labstack.com

### Gin

一個Go語言寫的HTTP Web框架。它提供了Martini風格的API並有更好的性能。

Gin 是用 Go 編寫的一個 Web 應用框架，對比其它主流的同類框架，他有更好的性能和更快的路由。

由於其本身只是在官方 net/http 包的基礎上做的完善，所以理解和上手很平滑。如果你現在開始做一套新的Api，我十分推薦你使用它。

- https://github.com/gin-gonic/gin
- https://gin-gonic.github.io/gin

### Iris

目前發展最快的Go Web框架。提供完整的MVC功能並且面向未來。

- https://github.com/kataras/iris
- https://iris-go.com

### Revel

一個高生產率，全棧Go語言的Web框架。

- https://github.com/revel/revel
- https://revel.github.io

## ORM - Object Relational Mapping (物件關係對映)
### gorm
- 全功能 ORM (無限接近)
- 關聯（包含一個，包含多個，屬於，多對多，多態）
- 鉤子 (在創建/保存/更新/刪除/查找之前或之後)
- 預加載
- 事務
- 複合主鍵
- SQL 生成器
- 數據庫自動遷移
- 自定義日誌
- 可擴展性, 可基於 GORM 回調編寫插件
- 所有功能都被測試覆蓋
- 開發者友好

#### 相關連結
- [gorm.io](https://gorm.io/)
- [中文指南](http://gorm.io/zh_CN/)
- [github](https://github.com/go-gorm/gorm)

### xorm
#### 相關連結
- [github](https://github.com/xormplus/xorm)

## 完... XD

