package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var client *http.Client
var cookies []*http.Cookie

const (
	xiaoyaoUrl string = "https://xiaoyaojp.com/"
)

func main() {

	login()

}

func login() {

	//获取登陆界面的cookie
	postURL := "https://xiaoyaojp.com/member.php?mod=logging&action=login&loginsubmit=yes&frommessage&loginhash=LjSeu&inajax=1"
	var username string = "dong1hang"
	var password string = "331801363"
	req, _ := http.NewRequest("GET", "https://xiaoyaojp.com/forum-130-1.html", nil)
	client = &http.Client{}
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
		log.Fatal(err)
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
		log.Fatal(err)
	}
	cookies = res.Cookies()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	// 获取一览下的body
	doc := getResultHtml("https://xiaoyaojp.com/forum-130-1.html")
	//在桌面创建文件夹
	ioserr := os.Mkdir("C:/Users/Administrator/Desktop/xiaoyaoImgs", os.ModePerm)
	if ioserr != nil {
		fmt.Println(err)
	}
	//find the review items
	//循环店铺
	doc.Find("div.fl_icn_g").EachWithBreak(func(i int, s *goquery.Selection) bool {
		a := s.Next().Find("a").First()
		a_href, _ := a.Attr("href")
		a_html, _ := a.Html()

		fmt.Println("------------------------------------")
		fmt.Printf("[%s]\n", a_href)
		fmt.Println(a_html)

		shopsUrl := getResultHtml(xiaoyaoUrl + a_href)

		//根据店面创建文件夹
		err := os.Mkdir("C:/Users/Administrator/Desktop/xiaoyaoImgs/"+a_html, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		// 循环每个帖子并找到下个页面的链接
		shopsUrl.Find("div.deanflists").EachWithBreak(func(i int, gs *goquery.Selection) bool {
			theSencodA := gs.Find("a").Eq(2)
			title, _ := theSencodA.Attr("title")
			fmt.Println("----------------theSencodA-------------------")
			fmt.Printf("[%s]\n", title)
			if "隐藏置顶帖" != title {
				a := gs.Find("a").First()
				a_html1, _ := a.Html()
				if "樱花岛美女" == a_html1 {
					a = gs.Find("a").Eq(1)
					a_html1, _ = a.Html()
				}
				a_href, _ := a.Attr("href")
				if strings.ContainsAny(a_html1, "/") {
					s := strings.Split(a_html1, "/")

					for _, str := range s {
						a_html1 = str
					}
				}
				fmt.Printf("[%s]\n", a_html1)
				shopsUrl1 := getResultHtml(xiaoyaoUrl + a_href)
				err := os.Mkdir("C:/Users/Administrator/Desktop/xiaoyaoImgs/"+a_html+"/"+a_html1, os.ModePerm)
				if err != nil {
					fmt.Println(err)

				}

				shopsUrl1.Find("ignore_js_op").EachWithBreak(func(i int, gks *goquery.Selection) bool {
					img := gks.Find("img").First()
					text, _ := img.Attr("zoomfile")
					//imgUrl, _ := img.Attr("src")

					fmt.Println("----------------imgUrl-------------------")
					fmt.Println(img)
					fmt.Printf("[%s]\n", text)
					s := strings.Split(text, "/")
					str := s[len(s)-1]
					if len(text) != 0 {

						DownloadFile(xiaoyaoUrl+text, "C:/Users/Administrator/Desktop/xiaoyaoImgs/"+a_html+"/"+a_html1+"/"+str)
					}
					return true
				})

			} else {
				fmt.Printf("[%s]\n", title)
			}

			return true

		})

		return true
	})
}

func gbk2utf8(str []byte) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder()))
}

// getResultHtml 获取url下的body
//return *goquery.Document
func getResultHtml(get_url string) *goquery.Document {

	values := url.Values{}
	// TODO  ?_dsign 这个参数随机的值不知道怎么获取
	//req1, err := http.NewRequest("GET", get_url+"?_dsign=9aaf12bf", strings.NewReader(values.Encode()))
	req1, err := http.NewRequest("GET", get_url, strings.NewReader(values.Encode()))
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
		log.Fatalf("status code error: %d %s", res1.StatusCode, res1.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res1.Body)

	if err != nil {
		log.Fatal(err)
	}

	printResultHtml(get_url)

	return doc

}

// debug用 打印body
func printResultHtml(get_url string) {

	values := url.Values{}
	// TODO  ?_dsign 这个参数随机的值不知道怎么获取
	//req1, err := http.NewRequest("GET", get_url+"?_dsign=9aaf12bf", strings.NewReader(values.Encode()))
	req1, err := http.NewRequest("GET", get_url, strings.NewReader(values.Encode()))
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
		log.Fatalf("status code error: %d %s", res1.StatusCode, res1.Status)
	}

	body := res1.Body
	bodytest, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodytest))

}

//DownloadFile 下载图片
// return error
func DownloadFile(url string, filepath string) error {
	fmt.Println(url, "->", filepath)
	//get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// create the file
	out, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
		fmt.Println(filepath)
	}
	defer out.Close()
	//write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
