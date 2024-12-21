# os(文件操作)

[os](https://www.liwenzhou.com/posts/Go/file/)

> os库提供了Go语言中文件读写的相关操作


## 打开和关闭文件

os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件。


## 读取文件

**file.Read()**

- 基本使用

  > Read方法定义如下：

  `func (f *File) Read(b []byte) (n int, err error)`

  > 接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF

- 循环读取
  
  > 使用for循环读取文件中的所有数据

**bufio读取文件**

> bufio 是在 file 的基础上封装了一层API，支持更多的功能。

`func bufio.NewReader(rd io.Reader) *bufio.Reader`

`func (b *bufio.Reader) ReadString(delim byte) (string, error)`

**读取整个文件**

> os（Go1.16之前io/ioutil）包的ReadFile函数能够读取完整的文件，只需要将文件名作为参数传入。

`func ioutil.ReadFile(filename string) ([]byte, error)`

## 文件写入操作

> os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}
```

- name：要打开的文件名 
- flag：打开文件的模式。 模式有以下几种：

| 模式        | 含义     |
| ----------- | -------- |
| os.O_WRONLY | 只写     |
| os.O_CREATE | 创建文件 |
| os.O_RDONLY | 只读     |
| os.O_RDWR   | 读写     |
| os.O_TRUNC  | 清空     |
| os.O_APPEND | 追加     |

- perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

