package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//"github.com/PuerkitoBio/goquery"
// func main1() {
// 	html := `<htm
// 	l>
// 	        <body>
// 	            <h1 id="title">春晓</h1>
// 	            <p class="content1">
// 	            春眠不觉晓，
// 	            处处闻啼鸟。
// 	            夜来风雨声，
// 	            花落知多少。
// 	            </p>
// 	        </body>
// 	        </html>
// 	        `
// 	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	dom.Find("p").Each(func(i int, selection *goquery.Selection) {
// 		fmt.Println(selection.Text())
// 	})
// }

func main() {

	fmt.Println("Start......")

	//fmt.Println(getUrlRespHtml())
	fmt.Println("Start......")

	//var cookieJar, _ = browsercookie.Chrome("https://account.geekbang.org")
	//cookieJar, _ := cookiejar.New(nil)
	values := url.Values{}

	var cookie string = "UM_distinctid=17775548d0b789-0014634a9a91f5-13e3563-1fa400-17775548d0c29; CNZZDATA1275248188=1848093610-1612577998-|1612577998; _ga=GA1.2.1380242526.1612581211; _gid=GA1.2.698374539.1612581211; haspost_11530=9; xdD4_2132_lastvisit=1612579053; xdD4_2132_sendmail=1; xdD4_2132_saltkey=aTHBQ31b; xdD4_2132_noticeTitle=1; xdD4_2132_st_t=0|1612582655|4618301a37ec4ccfd76cfa1d0c5cfd47; xdD4_2132_sid=V5g7jk; xdD4_2132_lastact=1612582780	member.php	logging; xdD4_2132_ulastactivity=2747yDTA6jTNBsQtxRnYcShQj0DgLb8U5hpmfOvpP2wGc6CMjr+L; xdD4_2132_lastcheckfeed=11530|1612582780; xdD4_2132_checkfollow=1; xdD4_2132_lip=126.51.91.20,1612582676; xdD4_2132_auth=d8a6ExNZJLw9lhUJtQ+S9n+PilaULrZ/aP+MD4mthNmJ65hK6JSnRuoKqobscTnZNrfypoPxFuLuNh2dAc3jAIGvQA; xdD4_2132_tshuz_accountlogin=11530"

	req, err := http.NewRequest("GET", "https://xiaoyaojp.com/forum-130-1.html", strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", cookie)
	//req.Header.Add("Agent", GetRandomUserAgent())

	//模拟
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Cache-Control", "max-age=0")
	//req.Header.Set("Origin","http://center.qianlima.com")
	req.Header.Set("Accept-Language", "ja,zh-CN;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6")
	//	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36")
	//req.Header.Set("Host","center.qianlima.com")
	// req.Header.Set("Connection","keep-alive")
	//req.Header.Set("Upgrade-Insecure-Requests","1")
	//req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	req.Header.Set("remoteAddress", "192.254.222.138:443")

	client := &http.Client{}
	res, err := client.Do(req)
	// Request the HTML page.
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	// //load the HTML document
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// fmt.Println(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //find the review items
	// doc.Find("div#card_container > div.grid_item").EachWithBreak(func(i int, s *goquery.Selection) bool {
	// 	a := s.Find("a").First()
	// 	a_href, _ := a.Attr("href")
	// 	img := s.Find("img").First()
	// 	img_alt, _ := img.Attr("alt")
	// 	img_src, _ := img.Attr("src")
	// 	img_src = img_src[:strings.Index(img_src, "?")]
	// 	img_url, err := url.Parse(img_src)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("------------------------------------")
	// 	fmt.Printf("[%s]\n", img_alt)
	// 	dl_filename := img_url.Path[strings.LastIndex(img_url.Path, "/")+1:]
	// 	DownloadFile(img_src, "C:\\Users\\Administrator\\Desktop\\imgs"+dl_filename)
	// 	fmt.Println(img_alt, a_href, img_src)

	// 	//load detail page

	// 	//detail_url := "https://royaleapi.com" + a_href
	// 	//loadDetailPage(detail_url)
	// 	return true

	// })

}

// getUrlRespHtml 获取网站信息
// return string body
func getUrlRespHtml() string {
	url := "https://xiaoyaojp.com/forum-130-1.html"
	var cookie string = "UM_distinctid=17775548d0b789-0014634a9a91f5-13e3563-1fa400-17775548d0c29; CNZZDATA1275248188=1848093610-1612577998-|1612577998; _ga=GA1.2.1380242526.1612581211; _gid=GA1.2.698374539.1612581211; haspost_11530=9; xdD4_2132_lastvisit=1612579053; xdD4_2132_sendmail=1; xdD4_2132_saltkey=aTHBQ31b; xdD4_2132_noticeTitle=1; xdD4_2132_st_t=0|1612582655|4618301a37ec4ccfd76cfa1d0c5cfd47; xdD4_2132_sid=V5g7jk; xdD4_2132_lastact=1612582780	member.php	logging; xdD4_2132_ulastactivity=2747yDTA6jTNBsQtxRnYcShQj0DgLb8U5hpmfOvpP2wGc6CMjr+L; xdD4_2132_lastcheckfeed=11530|1612582780; xdD4_2132_checkfollow=1; xdD4_2132_lip=126.51.91.20,1612582676; xdD4_2132_auth=d8a6ExNZJLw9lhUJtQ+S9n+PilaULrZ/aP+MD4mthNmJ65hK6JSnRuoKqobscTnZNrfypoPxFuLuNh2dAc3jAIGvQA; xdD4_2132_tshuz_accountlogin=11530"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("获取地址错误")
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Add("Agent", GetRandomUserAgent())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("登录错误")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)

}

// GetRandomUserAgent 模拟用户
// return 随机
func GetRandomUserAgent() string {

	var userAgentList = []string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
		"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
		"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
		"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
		"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
		"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgentList[r.Intn(len(userAgentList))]
}

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
	}
	defer out.Close()
	//write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
