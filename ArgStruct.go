package main

type Args struct {
	Input   string `arg:"-i,required"  placeholder:"输入文件"`
	Output  string `arg:"-o" default:"output" placeholder:"输出目录"`
	Package string `arg:"-p" default:"main" placeholder:"生成文件的包名"`
	GoRoot  string `arg:"-g,env:GOROOT" placeholder:"go安装目录，取环境变量GOROOT"`
}
