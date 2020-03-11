package gopackage

// go get -u github.com/easierway/concurrent_map

import(
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemotePackage_ConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(8)
	m.Set(cm.StrKey("key"), 20)
	t.Log(m.Get(cm.StrKey("key")))
}

// 查找依赖包路径 1.6及后
// 1、当前包下的vendor目录
// 2、向上级目录查找，直到找到src下的vendor目录
// 3、在GOPATH下面查找依赖包
// 4、在GOROOT下面查找依赖包

// 常用依赖管理工具
// godep https://github.com/tools/godep
// glide https://github.com/Masterminds/glide
// dep   https://github.com/golang/dep

// cd E:\src\go_learning\src\package

// brew install glide
// glide init
// vi glide.yaml

// glide install
//
