package main

type Args struct {
	Input         string `arg:"-i,required"  placeholder:"输入文件"`
	Output        string `arg:"-o" default:"output" placeholder:"输出目录"`
	Package       string `arg:"-p" default:"main" placeholder:"生成文件的包名"`
	AutoAddPerfix bool   `arg:"-a" default:"false" placeholder:"自动添加前缀，解决重名问题。默认为false，如无法确保是否有重名变量，可设置为true"`
	GoRoot        string `arg:"-g,env:GOROOT" placeholder:"go安装目录，取环境变量GOROOT"`
}

func (Args) Version() string {
	return "cs 0.0.1"
}
