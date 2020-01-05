package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/anaskhan96/soup"

	"go-crawler/mycrawler/db"
	"go-crawler/mycrawler/model"
)

func main() {
	baseURL := "http://www.atomy.com"
	productURL := baseURL + "/tw/Home/Product"

	for i := 1; i < 7; i++ {
		time.Sleep(time.Duration(1) * time.Second)

		postParm := fmt.Sprintf("LargeClass=00&MiddleClass=00&SmallClass=00&MatLevel=0&Order=basic&CurrentPageNo=%d&ListType=&CountPerPage=24", i)
		resp, err := http.Post(productURL+"/MallMain",
			"application/x-www-form-urlencoded",
			strings.NewReader(postParm))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		doc := soup.HTMLParse(string(body))
		titles := doc.FindAll("li", "class", "ptitle")
		prices := doc.FindAll("li", "class", "pprice")
		points := doc.FindAll("li", "class", "ppoint")

		for index, title := range titles {
			titleA := title.Find("a")
			priceSpan := prices[index].Find("span", "class", "numberic")
			pointSpan := points[index].Find("span", "class", "numberic")
			LinkURL := fmt.Sprintf("%s/%s", productURL, strings.Split(titleA.Attrs()["href"], "./")[1])

			product := model.Products{
				Name:  titleA.Text(),
				Price: priceSpan.Text(),
				Point: pointSpan.Text(),
				Link:  LinkURL,
			}
			product1 := model.Products{}

			db.Db.Model(&product).Where("name = ?", titleA.Text()).Update("link", LinkURL)
			db.Db.Where("name = ?", titleA.Text()).Find(&product1)
			if product1.ID == 0 {
				fmt.Println("find new product: ", titleA.Text())
				product1.Name = titleA.Text()
				product1.Price = priceSpan.Text()
				product1.Point = pointSpan.Text()
				product1.Link = LinkURL
				db.Db.Create(&product1)
			}
		}
	}

}
