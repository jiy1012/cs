# cs
config file to golang struct,support json ,yaml,toml

使用方法：
1. clone 代码编译
```
git clone git@github.com:jiy1012/cs.git && cd cs && go build ./
```

2.执行
```
Usage: cs --input 输入文件 [--output 输出目录] [--package 生成文件的包名] [--goroot go安装目录，取环境变量GOROOT]

--input -i 输入文件 支持json,yaml,toml 详情见fileloader文件夹下init.go

--output -o 输出目录 生成golang struct的保存文件夹 默认当前路径下output文件夹

--package -p 生成go struct文件的包名，默认为main

--goroot -g go安装目录，默认取环境变量GOROOT。没有安装go则不会使用gofmt格式化代码，需自行格式化

```
