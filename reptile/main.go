package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	values := url.Values{}

	req, err := http.NewRequest("GET", "https://royaleapi.com/cards/popular", strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}
	//模拟
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36)")
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

	//load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	fmt.Println(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//find the review items
	doc.Find("div#card_container > div.grid_item").EachWithBreak(func(i int, s *goquery.Selection) bool {
		a := s.Find("a").First()
		a_href, _ := a.Attr("href")
		img := s.Find("img").First()
		img_alt, _ := img.Attr("alt")
		img_src, _ := img.Attr("src")
		img_src = img_src[:strings.Index(img_src, "?")]
		img_url, err := url.Parse(img_src)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------------------------------------")
		fmt.Printf("[%s]\n", img_alt)
		dl_filename := img_url.Path[strings.LastIndex(img_url.Path, "/")+1:]
		DownloadFile(img_src, "C:\\Users\\Administrator\\Desktop\\imgs"+dl_filename)
		fmt.Println(img_alt, a_href, img_src)

		//load detail page

		//detail_url := "https://royaleapi.com" + a_href
		//loadDetailPage(detail_url)
		return true

	})

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
