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
- `set`: 设置一个环境变量

### 使用示例

#### 查看环境变量的值

```
envtool get PATH
```

#### 在环境变量中添加一个值

```
envtool add PATH C:\\go\\bin
```

#### 从环境变量中删除一个值

```
envtool remove PATH C:\\go\\bin
```

#### 删除特定的环境变量

```
envtool delete GOROOT
```

#### 列出环境变量的所有值

```
envtool list PATH
```

#### 设置一个环境变量

```
envtool set GOROOT C:\\go\\bin
```

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
- `set`: Set an environment variable

### Examples

#### View the value of an environment variable

```
envtool get PATH
```

#### Add a value to an environment variable

```
envtool add PATH C:\\go\\bin
```

#### Remove a value from an environment variable

```
envtool remove PATH C:\\go\\bin
```

#### Delete specific environment variable

```
envtool delete GOROOT
```

#### List all values of an environment variable

```
envtool list PATH
```

#### Set an environment variable

```
envtool set GOROOT C:\\go\\bin
```
