#  `ParseStrSize()`：字符串操作

## `fmt.Sscanf()`函数

`fmt.Sscanf()`是Go语言`fmt`包提供的一个非常有用的函数，用于根据指定的格式从字符串中解析并提取数据。它使得从格式化的字符串中读取数据变得简单且灵活。

### 功能

- 从一个字符串中按照指定的格式提取数据。
- 支持多种数据类型的提取，包括整数、浮点数、字符串等。
- 根据格式化字符串中的占位符，将数据读入到提供的变量中。

### 常用占位符

- `%d`：十进制整数。
- `%x`, `%o`, `%b`：分别表示十六进制、八进制、二进制整数。
- `%f`, `%g`, `%e`：表示浮点数。
- `%s`：表示字符串，直到遇到空白字符。
- `%c`：读取一个字符。
- `%v`：根据值的类型选择格式。
- `%%`：表示百分号本身。

### 使用示例

#### 基础用法

从字符串中提取整数和字符串：

```go
var i int
var s string
_, err := fmt.Sscanf("1234 is a number", "%d %s", &i, &s)
if err != nil {
    log.Fatal("Error:", err)
}
fmt.Println(i, s) // 输出: 1234 is
```

#  `sync`：了解锁包

在 Go 语言的并发编程中，`sync` 包提供了一系列的同步原语（synchronization primitives），用于在goroutines之间进行同步和数据保护。这些原语包括互斥锁（`Mutex`）、读写互斥锁（`RWMutex`）、等待组（`WaitGroup`）、一次性执行（`Once`）等，都是构建并发安全程序的基础工具。

## 读写互斥锁 `sync.RWMutex`（此项目中使用）

读写互斥锁（`RWMutex`）是一种特殊的锁，它允许多个读操作并发执行，但写操作是互斥的。这种锁特别适用于读多写少的场景，因为它可以增加程序的并发性能。

### 功能特点

- **读锁定（`RLock`）**：多个goroutine可以同时获得读锁，允许多个读操作并行执行。
- **读解锁（`RUnlock`）**：释放读锁，当所有的读锁都被释放，写锁才可以被获得。
- **写锁定（`Lock`）**：当写锁被获得时，任何其他尝试读锁定或写锁定的goroutine都会阻塞，直到写锁被释放。
- **写解锁（`Unlock`）**：释放写锁，使得其他的读锁或写锁可以被获得。

#### 使用示例

```go
var rw sync.RWMutex
var data int

func readData() int {
    rw.RLock() // 请求读锁
    defer rw.RUnlock() // 延迟释放读锁
    return data
}

func writeData(newData int) {
    rw.Lock() // 请求写锁
    defer rw.Unlock() // 延迟释放写锁
    data = newData
}
```
### 思考：既然读锁允许并发，其实我是没必要上读锁的啊
### 回答：Openai ChatGPT4回答(2024/03/25 中国)

#### 为什么需要读锁？

- **数据一致性**：在并发程序中，如果有其他goroutine可能会修改数据，那么在读取数据时使用读锁是必要的，以确保读操作期间数据不被修改，从而保证了操作的一致性。
- **防止写入时读取**：读锁还可以防止在进行写操作时进行读取。当一个goroutine持有写锁进行数据修改时，其他goroutine如果尝试获取读锁（即尝试读取数据），将会等待，直到写锁释放。这确保了在写入期间，读取操作不会得到部分写入或不一致的数据。

#### 什么时候可以不用读锁？

- **纯读场景**：如果您的应用场景中只有读操作，没有写操作，那么读锁不是必需的，因为没有写操作冲突的风险。
- **不变数据**：如果数据在初始化后不会再改变，或者对数据的任何修改都是在启动时（在多个goroutine运行之前）完成的，那么在并发读取时不需要读锁。

#### 使用读锁的考虑

使用读锁的决策应基于您的具体场景和需求。虽然读锁允许并发读取，但它们仍然会引入一定的开销，因为需要管理锁的状态。如果可以确保数据在读取期间不会被修改（例如，数据只在程序初始化时被设置），那么可能不需要上读锁。然而，在大多数动态的、多写操作的并发环境中，合理使用读锁是保证数据一致性和程序正确性的重要手段。

# `reflect`：反射

在Go语言中，`reflect`包提供了强大的反射功能，允许程序在运行时检查变量的`类型`和`值`，甚至修改它们。这为编写通用的代码和框架、处理未知类型的数据提供了可能。

## 功能概述

- **动态获取类型和值**：`reflect`允许在运行时动态地获取任何变量的类型(`Type`)和值(`Value`)信息。
- **修改变量值**：通过反射，你可以在运行时修改变量的值，包括私有字段。
- **调用方法**：反射还可以用于动态地调用对象的方法。

## 核心组件

### `Type`：理解变量类型

通过`reflect.TypeOf()`获取任何值的类型信息，返回一个`reflect.Type`对象，它包含了类型的所有信息。

### `Value`：操作变量值

`reflect.ValueOf()`函数接受任意参数，并返回一个`reflect.Value`类型的值，代表了运行时的数据。
可以从`reflect.Value`中读取具体值，甚至修改值（如果可设置）。

## 使用示例

### 获取类型信息

```go
var x float64 = 3.4
t := reflect.TypeOf(x)
fmt.Println("Type:", t.Name()) // 输出: float64
```

### 读取和设置值

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("Value:", v.Float()) // 读取值

// 修改值
p := reflect.ValueOf(&x) // 注意：需要传入变量的指针
e := p.Elem()            // 获取指针指向的元素
if e.CanSet() {
    e.SetFloat(7.1)
}
fmt.Println(x) // x的值被修改为7.1
```

### 调用方法

```go
type MyStruct struct {
    Value int
}

func (m *MyStruct) SetValue(v int) {
    m.Value = v
}

var ms MyStruct
v := reflect.ValueOf(&ms) // 获取指针的reflect.Value，因为SetValue是指针接收者
method := v.MethodByName("SetValue")
args := []reflect.Value{reflect.ValueOf(42)}
method.Call(args)

fmt.Println(ms.Value) // 输出: 42

```
