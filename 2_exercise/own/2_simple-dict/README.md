# simple-dict
https://fanyi.caiyunapp.com/#/

# 1. step 1 ： 浏览器查找post
- 右键点击 “检查”，然后 点击网页 的跳转
- 关注 “检查” 中的 “Network”
  - “Name” 里面查找 “Dict” 中的 “Headers” 和 “Payload”
    - “Headers” 是 `头`
    - “Payload” 个人理解是 `内容`
- 右键点击“copy as PowerShell”

# 2. step 2 ： 代码生成的网址 -> 1_get_request
https://curlconverter.com/go/

# 3. step 3 : 生成request body -> 2_request_body
```go
json.Marshal // 将数据结构转换为json字符串

// 抛出异常 error
if err != nil {
		log.Fatal(err)
}
```

# 4. 解析response body
https://oktools.net/json2go
- 从chorme“检查” ->  “Network” -> “preview” 中 copy 所有object。
- 粘贴到JSON框里，并点“粘贴-展开”

# 5. 压缩到一个main函数中
- 将原本的main函数编程main函数，然后新建一个main函数，将query函数封装起来！