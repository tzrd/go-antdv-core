package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// 获取 url参数
func GetUrlParameters(Url string) map[string][]string {
	u, err := url.Parse(Url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	urlParam := u.RawQuery
	m, err := url.ParseQuery(urlParam)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	return m
}

// 判断 是否为http或https url格式
func IsUrlFormator(url string) bool {
	if strings.Contains(url, "http://") && strings.Index(url, "http://") == 0 {
		return true
	}

	if strings.Contains(url, "https://") && strings.Index(url, "https://") == 0 {
		return true
	}

	return false
}

// 判断 是否为https url格式
func IsHttpsUrl(url string) bool {
	if strings.Contains(url, "https://") && strings.Index(url, "https://") == 0 {
		return true
	}

	return false
}
