# cs
config file to golang struct,support json ,yaml,toml

一、安装：

1.源码编译

```
  git clone git@github.com:jiy1012/cs.git && cd cs && go build ./ && chmod -x cs
```

2.brew安装

```
  brew tap jiy1012/brew
```

```
  brew install cs
```

二、执行

```
Usage: cs --input 输入文件 [--output 输出目录] [--package 生成文件的包名] [--goroot go安装目录，取环境变量GOROOT]

--input -i 输入文件 支持json,yaml,toml 详情见fileloader文件夹下init.go

--output -o 输出目录 生成golang struct的保存文件夹 默认当前路径下output文件夹

--package -p 生成go struct文件的包名，默认为main

--goroot -g go安装目录，默认取环境变量GOROOT。没有安装go则不会使用gofmt格式化代码，需自行格式化

-x 自动添加前缀，解决重名问题默认为false，如无法确保是否有重名变量，可设置为true

-e 自动添加omitempty，默认为false

-m 自动添加mapstructure，默认为false

```
