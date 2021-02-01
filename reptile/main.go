package main

import (
	"fmt"	"net/url"
)


func main1() {
	// html := `<htm
	// l>
	//         <body>
	//             <h1 id="title">春晓</h1>
	//             <p class="content1">
	//             春眠不觉晓，
	//             处处闻啼鸟。
	//             夜来风雨声，
	//             花落知多少。
	//             </p>
	//         </body>
	//         </html>
	//         `
	// dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// dom.Find("p").Each(func(i int, selection *goquery.Selection) {
	// 	fmt.Println(selection.Text())
	// })
}

func main() {
	fmt.Println("Start......")

	values :=url.Values{}

	req,err :=http.NewRequest("GET","https://royaleapi.com/cards/popular",strings.NewReader(values.Encode()))
	if err!= nil{
		panic(err)
	}
	//TODO 
	req.Header.Add("User-Agent")


}
