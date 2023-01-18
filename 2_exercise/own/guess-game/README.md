# guess-game
# 1. step 1 : 1-generate-random-num
- 需要确定随机种子
# 2. step 2 : 2-gene-num-and-input
- bufio.NewReader(os.Stdin)
  - 用来得到输入的字符串
- strings.TrimSuffix(input, "\r\n") 
  - 去掉回车 和 换行符
- strconv 
  - 解析字符串得到数字