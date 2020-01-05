package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/anaskhan96/soup"

	"go-crawler/mycrawler/db"
	"go-crawler/mycrawler/model"
)

func main() {
	// saveProducts()
	saveEduCenter()
}

func saveEduCenter() {
	baseURL := "http://www.atomy.com/tw/Home/About/EduCenter"
	file, err := os.Create("center.csv")

	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	csvTitle := []string{
		"名稱", "地址", "電話號碼", "設立日期",
	}
	writer.Write(csvTitle)

	resp, err := soup.Get(baseURL)
	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	tbody := doc.Find("table", "class", "tableStyle2").Find("tbody")
	trs := tbody.FindAll("tr")
	for _, tr := range trs {
		th := tr.Find("th")
		centerName := th.Text()
		var centerAddr string
		var centerPhone string
		var centerDate string

		fmt.Println(centerName)

		tds := tr.FindAll("td")

		for k, v := range tds {
			switch k {
			case 0:
				centerAddr = v.Text()
				fmt.Println("地址: ", centerAddr)
			case 1:
				centerPhone = v.Text()
				fmt.Println("電話號碼: ", centerPhone)
			case 2:
				centerDate = v.Text()
				fmt.Println("設立日期: ", centerDate)
			}
		}

		value := []string{centerName, centerAddr, centerPhone, centerDate}
		err := writer.Write(value)
		checkError("Cannot write to file", err)

		fmt.Println("------")
		saveCenter(centerName, centerAddr, centerPhone, centerDate)
	}
}

func saveProducts() {
	baseURL := "http://www.atomy.com"
	productURL := baseURL + "/tw/Home/Product"
	file, err := os.Create("result.csv")

	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	csvTitle := []string{
		"名稱", "價格", "PV", "連結",
	}
	writer.Write(csvTitle)

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
			productName := titleA.Text()
			productPrice := priceSpan.Text()
			productPoint := pointSpan.Text()
			LinkURL := fmt.Sprintf("%s/%s", productURL, strings.Split(titleA.Attrs()["href"], "./")[1])

			value := []string{productName, productPrice, productPoint, LinkURL}
			err := writer.Write(value)
			checkError("Cannot write to file", err)

			saveProduct(productName, productPrice, productPoint, LinkURL)
		}
	}
}

func saveProduct(productName string, productPrice string, productPoint string, LinkURL string) {
	product := model.Products{
		Name:  productName,
		Price: productPrice,
		Point: productPoint,
		Link:  LinkURL,
	}
	db.Db.Model(&product).Where("name = ?", productName).Update("link", LinkURL)

	product1 := model.Products{}
	db.Db.Where("name = ?", productName).Find(&product1)
	if product1.ID == 0 {
		fmt.Println("find new product: ", productName)
		product1.Name = productName
		product1.Price = productPrice
		product1.Point = productPoint
		product1.Link = LinkURL
		db.Db.Create(&product1)
	}
}

func saveCenter(centerName string, centerAddr string, centerPhone string, centerDate string) {
	center := model.Centers{}
	db.Db.Where("name = ?", centerName).Find(&center)
	if center.ID == 0 {
		center.Name = centerName
		center.Address = centerAddr
		center.Phone = centerPhone
		center.Date = centerDate
		db.Db.Create(&center)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
