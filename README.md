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

    fmt.Println(passwords)                // map[walter1:123456 walter2:654321]
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

