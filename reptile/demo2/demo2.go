package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {

	//获取登陆界面的cookie
	postURL := "https://xiaoyaojp.com/member.php?mod=logging&action=login&loginsubmit=yes&frommessage&loginhash=LjSeu&inajax=1"
	var username string = "dong1hang"
	var password string = "331801363"
	req, _ := http.NewRequest("GET", "https://xiaoyaojp.com/forum-130-1.html", nil)
	client := &http.Client{}
	res, _ := client.Do(req)

	var tempCookies = res.Cookies()
	for _, v := range res.Cookies() {
		req.AddCookie(v)
	}

	Jar, _ := cookiejar.New(nil)
	postURLJSON, _ := url.Parse(postURL)
	Jar.SetCookies(postURLJSON, tempCookies)
	client.Jar = Jar

	var resp *http.Response
	req, err := http.NewRequest("POST", postURL, strings.NewReader("username="+username+"&password="+password+"&referer=https://xiaoyaojp.com/forum-130-1.html&questionid=0&formhash=4efa916d&answer="))

	if err != nil {
		return
	}

	req.Header.Set("Host", "xiaoyaojp.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "131")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Origin", "https://xiaoyaojp.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "xiaoyaojp.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Referer", "https://xiaoyaojp.com/forum-130-1.html")
	req.Header.Set("Accept-Language", "ja,zh-CN;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Cookie", "UM_distinctid=17775548d0b789-0014634a9a91f5-13e3563-1fa400-17775548d0c29; _ga=GA1.2.1380242526.1612581211; _gid=GA1.2.698374539.1612581211; haspost_11530=9; CNZZDATA1275248188=1848093610-1612577998-%7C1612611522; xdD4_2132_sid=Eeo9d5; xdD4_2132_lastvisit=1612608945; _gat=1; xdD4_2132_sendmail=1; xdD4_2132_noticeTitle=1; xdD4_2132_saltkey=onD76Cjk; xdD4_2132_st_t=0%7C1612612547%7C78a2ca1b9e83cf5439809cda99ce0fb9; xdD4_2132_lastact=1612612548%09member.php%09logging")

	resp, err = client.Do(req)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	} else {
		fmt.Println("ERROR  返回为空 ")
	}
	if resp == nil || resp.Body == nil || err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	values := url.Values{}
	req1, err := http.NewRequest("GET", "https://xiaoyaojp.com/forum-130-1.html", strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}

	req1.Header.Set("Cookie", "UM_distinctid=17775548d0b789-0014634a9a91f5-13e3563-1fa400-17775548d0c29; _ga=GA1.2.1380242526.1612581211; _gid=GA1.2.698374539.1612581211; haspost_11530=9; CNZZDATA1275248188=1848093610-1612577998-%7C1612611522; xdD4_2132_sid=Eeo9d5; xdD4_2132_lastvisit=1612608945; _gat=1; xdD4_2132_sendmail=1; xdD4_2132_noticeTitle=1; xdD4_2132_saltkey=onD76Cjk; xdD4_2132_st_t=0%7C1612612547%7C78a2ca1b9e83cf5439809cda99ce0fb9; xdD4_2132_lastact=1612612548%09member.php%09logging")

	res1, err := client.Do(req1)
	// Request the HTML page.
	if err != nil {
		log.Fatal(err)
	}

	defer res1.Body.Close()
	if res1.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body1))

	// // 打印所有请求头信息
	// for k, v := range req.Header {
	// 	fmt.Println("请求头信息 :", k, "=", v)
	// }
	// // 打印所有响应头信息
	// for k, v := range resp.Header {
	// 	fmt.Println("响应头信息 :", k, "=", v)
	// }

	// for _, v := range client.Jar.Cookies(req.URL) {
	// 	fmt.Println("响应头信息- :", v.Name, "=", v.Value)
	// }
	// 打印正文信息
	// var buf []byte
	// buf, err = ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	// buf, _ = gbk2utf8(buf)
	// content := string(buf)
	// fmt.Println("正文信息:", content)
}
func gbk2utf8(str []byte) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder()))
}
