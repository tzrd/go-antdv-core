package utils

import "testing"

func TestGetUrlParameters(t *testing.T) {
	Url := "https://root:123456@www.baidu.com:0000/login?name=xiaoming&name=xiaoqing&age=24&age1=23#fffffff"
	p := GetUrlParameters(Url)

	if p == nil {
		t.Fail()
	}

	t.Log(p)
}

func TestIsUrlFormator(t *testing.T) {
	Url1 := "https://www.baidu.com"
	Ret1 := IsUrlFormator(Url1)

	if Ret1 == false {
		t.Fail()
	} else {
		t.Logf("%s is url: %v\n", Url1, Ret1)
	}

	Url2 := "http://www.baidu.com"
	Ret2 := IsUrlFormator(Url2)

	if Ret2 == false {
		t.Fail()
	} else {
		t.Logf("%s is url: %v\n", Url2, Ret2)
	}

	Url3 := "http:/www.baidu.com"
	Ret3 := IsUrlFormator(Url3)

	if Ret3 == true {
		t.Fail()
	} else {
		t.Logf("%s is url: %v\n", Url3, Ret3)
	}
}

func TestIsHttpsUrl(t *testing.T) {
	Url1 := "https://www.baidu.com"
	Ret1 := IsHttpsUrl(Url1)

	if Ret1 == false {
		t.Fail()
	} else {
		t.Logf("%s is  https: %v\n", Url1, Ret1)
	}

	Url2 := "http://www.baidu.com"
	Ret2 := IsHttpsUrl(Url2)

	if Ret2 == true {
		t.Fail()
	} else {
		t.Logf("%s is https: %v\n", Url2, Ret2)
	}

	Url3 := "www.baidu.com"
	Ret3 := IsHttpsUrl(Url3)

	if Ret3 == true {
		t.Fail()
	} else {
		t.Logf("%s is https: %v\n", Url3, Ret3)
	}
}
