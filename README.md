# envtool

envtool 是一个使用 Go 语言编写的 Windows 环境变量设置工具。

## 用法

使用方法如下：

```
envtool OPTION [KEY] [VALUE]
```

其中，`OPTION` 为选项，可以是以下之一：

- `get`, `view`: 查看环境变量的值
- `add`, `append`: 在环境变量中添加一个值
- `rm`, `remove`: 从环境变量中删除一个值
- `del`, `delete`: 删除特定的环境变量
- `list`: 列出环境变量的所有值

## Usage

envtool is a Windows environment variable setting tool written in Go.

### Usage

The usage is as follows:

```
envtool OPTION [KEY] [VALUE]
```

Where `OPTION` is an option, which can be one of the following:

- `get`, `view`: View the value of an environment variable
- `add`, `append`: Add a value to an environment variable
- `rm`, `remove`: Remove a value from an environment variable
- `del`, `delete`: Delete specific environment variable
- `list`: List all values of an environment variable
