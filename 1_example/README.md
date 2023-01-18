# 1. hello world

# 2. var
- 变量的声明
  - 法一 ：var a = "initial"
  - 法二 ：var b, c int = 1, 2
  - 法三 ：f := float(e)
- 常量
  - const 
    - const没有确定的类型，根据上下文确定
    - const h = 10000

# 3. for 
- 只有for循环
  - 死循环 ：for {}
  - 经典C++循环
    - for i := 7; j <9; j++ {}
- continue
- break

# 4. if 
- if
  - if 后面没有括号
  - if 必须加 {}

# 5. switch
- switch
  - 与C++不一样的点：在case中不需要加break
  - 格式
```go
a := 2
switch a {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("other")
}
```

# 6. array（用的很少）
```go
var a [5] int
``` 
- 用的很少

# 7. slice（用的更多）
```go
s := make([]string, 3)
s[0] = "a"

s1 := []string{"a", "b", "c"}
// s1 = [abc]

// append
s = append(s, "d", "e")
// 注意需要赋值回s

// copy
c := make([]string, len(s))
copy(c, s)

// 切片
fmt.Println(s[2:5]) // 左闭右开 类似于python
```

# 8. map
```go
m := make(map[string]int)
// string 是key的类型； int是value的类型
m["one"] = 1
fmt.Prinlnin(m) // m[one:1]
fmt.Prinlnin(m["one"]) // 1
fmt.Prinlnin(m["unknown"]) // 0

// 确认是否存在
r, ok := m["unknown"]
fmt.Println(r, ok) // 0 false

// delete
delete(m, "one")

```

# 9. range
- range 返回的是两个值
  - for i, value := range(nums)
  - for k, value := range(map_var)
```go
// 遍历slice
nums := []int {2, 3, 4}
sum := 0
for i, value := range(nums) {
    sum += value
    if value == 2 {
        fmt.Println("index: ", i, "value: ", value)
    }
}
fmt.Println(sum)

// 遍历map
m := map[string]string
for k, v := range m {
    fmt.Println(k, v)
}
for k := range m {
    fmt.Println("key", k)
}
```

# 10. func
- 格式：
  - 即需要将返回什么变量类型 在func的位置说清楚
    - 所以说 void (C++) 的写法：func void_fn(a int) {}
      - 换句话说 不用写返回值的数据类型就是 void 函数
```go
func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}
```

# 11. point（用的不多）
- 定义
  - a *int
- 传参的时候需要& 
  - 即 &a
- 函数中使用的时候 必须加*
  - 即 *a

# 12. struct
- 定义
```go
type user struct {
  name string
  password string
}
```

- 初始化
  - 四种方法
- 函数调用的时候
  - struct在函数声明的时候 必须加*
  - 但是在 函数使用的时候 不加*
```go
// 法一
var a user
a.user = "wang"
a.password = "123"
// 法二：用key初始化
b := user{"wang" : "123"}
// 法三：用key:value初始化
c := user{name : "wang", password : "123"}
// 法四：混合初始化
d := user{name : "wang"}
d.password = "1024"
```

# 13，struct method
- 定义
  - 两种写法
    - 1. 不带指针
    - 2. 带指针
```go
func (u user) checkPassword(password string) bool {}
func (u *user) checkPassword(password string) bool {}
```

# 14. error
- import ("errors")
- error 的数据类型：error
- 新的error：error.New("not found!")

# 15. string
- import "strings"
- 各种method
  - strings.Contains
  - strings.Count()
  - ....

# 16. fmt
- Printf 类似于 C++的printf。但是go只需要 `%v` 就可以表示所有的数字
- `+%v` 更详细，可以表示 map的 key 和 value
- `%.2f` 表示两位小数
```go
s := "hello"
fmt.Printf("s = %v\n", s) // s = hello

map_var := {"name" : "alex"}
fmt.Printf("map_var = %+v", map_var) // map_var = {"name" : "alex"}
```

# 17. json
- 用 struct 且 每个key的首字母需要大写
  - 后面加`json:"age"`可以得到小写的key
- 用 json.Marshal()方法 序列化
- 用 json.Unmarshal()方法 反序列化

# 18. time

# 19，数字解析 strconv
- ParseInt 解析string 成 int
- ParseFloat 解析string 成 float

20. env 进程信息
