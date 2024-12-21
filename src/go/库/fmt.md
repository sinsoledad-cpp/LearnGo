# fmt

[fmt](https://www.liwenzhou.com/posts/Go/fmt/)
> fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。

## 向外输出

## 获取输入

**bufio.NewReader**

> 输入的内容可能包含空格
> 可以使用bufio包来实现

```go
reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
fmt.Print("请输入内容：")
text, _ := reader.ReadString('\n') // 读到换行
text = strings.TrimSpace(text)
```
