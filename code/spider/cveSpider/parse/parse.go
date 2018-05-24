package parse

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/tcsecchen/bs"
)

//Info 是cve详细信息结构体
type Info struct {
	ID             string
	Name           string
	Assortment     string
	BugtraqID      string
	RemoteOverflow string
	LocalOverflow  string
	ReleaseDate    string
	UpdateDate     string
	Author         string
	Version        string
	Discuss        string
	Exploit        string
	Solution       string
	Reference      string
}

//GetLinks 函数向给定URL发起HTTP GET请求
//解析HTML并返回HTML文档中存在的cve链接
func GetLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	var links []string

	soup := bs.Init(resp.Body)
	// 找出所有 a 标签的链接
	for _, j := range soup.SelByTag("a") {
		pattern, _ := regexp.MatchString("^http://cve.scap.org.cn/CVE", (*j.Attrs)["href"])
		if pattern {
			links = append(links, (*j.Attrs)["href"])
		}
	}
	resp.Body.Close()

	return links, nil
}

//GetContent 函数获取cve链接中的内容
func GetContent(url string) (Info, error) {
	resp, err := http.Get(url)
	var info Info
	if err != nil {
		return info, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return info, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	soup := bs.Init(resp.Body)
	for _, j := range soup.SelByClass("cve_id") {
		info.ID = j.Value
	}
	for _, j := range soup.Sel("table", &map[string]string{"id": "info_sf"}) {
		for _, v := range j.SelByTag("td") {
			s := strings.SplitN(v.Value, ":", 2)
			for i := range s {
				s[i] = strings.TrimSpace(s[i])
			}
			switch s[0] {
			case "漏洞名称":
				info.Name = s[1]
			case "漏洞分类":
				info.Assortment = s[1]
			case "BugtraqID":
				info.BugtraqID = s[1]
			case "远程溢出":
				info.RemoteOverflow = s[1]
			case "本地溢出":
				info.LocalOverflow = s[1]
			case "发布日期":
				info.ReleaseDate = s[1]
			case "更新日期":
				info.UpdateDate = s[1]
			case "漏洞作者":
				info.Author = s[1]
			}
		}
	}
	//获取影响版本
	for _, j := range soup.Sel("table", &map[string]string{"id": "info_ver1"}) {
		for _, v := range j.SelByTag("td") {
			s := strings.Replace(v.Value, "\n", "$", -1)
			s = strings.Replace(s, "\t", "", -1)
			s = strings.Replace(s, "$$$", ",", -1)
			s = strings.Replace(s, "$$+ $$", "+ ", -1)
			info.Version = s
		}
	}

	//获取讨论
	for _, j := range soup.Sel("table", &map[string]string{"id": "info_discuss"}) {
		for _, v := range j.SelByTag("td") {
			info.Discuss = v.Value
		}
	}
	//获取漏洞利用
	for _, j := range soup.Sel("table", &map[string]string{"id": "info_exploit"}) {
		for _, v := range j.SelByTag("td") {
			s := strings.TrimSpace(v.Value)
			info.Exploit = s
		}
	}
	//获取解决方案
	for _, j := range soup.Sel("table", &map[string]string{"id": "info_solution"}) {
		for _, v := range j.SelByTag("td") {
			info.Solution = v.Value
		}
	}
	//获取相关参考
	for _, j := range soup.Sel("table", &map[string]string{"id": "info_ref"}) {
		var ref []string
		for _, v := range j.SelByTag("a") {
			s := v.Value + "(" + (*v.Attrs)["href"] + ")"
			ref = append(ref, s)
		}
		s := strings.Join(ref, ", ")
		info.Reference = s
	}

	resp.Body.Close()
	if err != nil {
		return info, fmt.Errorf("parsing %s as HTML：%v", url, err)
	}
	return info, nil
}
