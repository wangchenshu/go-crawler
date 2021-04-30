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

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// WriteToSQL - WriteToSQL
const WriteToSQL = false

func saveEduCenter() []model.CenterCommon {
	baseURL := "https://www.atomy.com/tw/Home/About/EduCenter" // base url
	file, err := os.Create("center.csv")
	var centers []model.CenterCommon

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
		name := th.Text()
		var addr string
		var phone string
		var date string

		tds := tr.FindAll("td")

		for k, v := range tds {
			switch k {
			case 0:
				addr = v.Text()
				// Debug
				// fmt.Println("地址: ", centerAddr)
			case 1:
				phone = v.Text()
				// Debug
				// fmt.Println("電話號碼: ", centerPhone)
			case 2:
				date = v.Text()
				// Debug
				// fmt.Println("設立日期: ", centerDate)
			}
		}

		// 定義結構
		c := model.CenterCommon{
			Name:    name,
			Address: addr,
			Phone:   phone,
			Date:    date,
		}
		// 增加
		centers = append(centers, c)

		value := []string{name, addr, phone, date}
		err := writer.Write(value)
		checkError("Cannot write to file", err)

		// Debug
		// fmt.Println("------")
		if WriteToSQL {
			saveEduCenterToSQL(name, addr, phone, date)
		}
	}

	return centers
}

func saveProducts() []model.ProductCommon {
	baseURL := "https://www.atomy.com"         // base url
	productURL := baseURL + "/tw/Home/Product" // product url
	var products []model.ProductCommon
	var findProducts []string

	findProductFunc := func(name string) bool {
		for _, n := range findProducts {
			if name == n {
				return true
			}
		}
		return false
	}

	// 建立檔案
	file, err := os.Create("result.csv")

	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 定義 title
	csvTitle := []string{
		"名稱", "價格", "PV", "連結",
	}
	// 寫 title
	writer.Write(csvTitle)

	for i := 1; i < 7; i++ {
		// 休一下
		time.Sleep(time.Duration(1) * time.Second)
		// post 參數
		postParm := fmt.Sprintf("LargeClass=00&MiddleClass=00&SmallClass=00&MatLevel=0&Order=basic&CurrentPageNo=%d&ListType=&CountPerPage=24", i)
		// 訪問
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
		// 找出所有 title, prices, points
		titles := doc.FindAll("li", "class", "ptitle")
		prices := doc.FindAll("li", "class", "pprice")
		points := doc.FindAll("li", "class", "ppoint")

		for index, title := range titles {
			titleA := title.Find("a")
			priceSpan := prices[index].Find("span", "class", "numberic")
			pointSpan := points[index].Find("span", "class", "numberic")

			name := titleA.Text()     // 名稱
			price := priceSpan.Text() // 價格
			point := pointSpan.Text() // PV
			// 連結
			link := fmt.Sprintf("%s/%s", productURL, strings.Split(titleA.Attrs()["href"], "./")[1])

			// 定義結構
			p := model.ProductCommon{
				Name:  name,
				Price: price,
				Point: point,
				Link:  link,
			}

			// 如果 還未存過
			if !findProductFunc(name) {
				// 增加
				products = append(products, p)
				// 增加找到的產品名稱
				findProducts = append(findProducts, name)
				// 定義值
				value := []string{name, price, point, link}
				// 寫入
				err := writer.Write(value)
				checkError("Cannot write to file", err)

				// 寫到 sql
				if WriteToSQL {
					saveProductToSQL(name, price, point, link)
				}
			}
		}
	}

	return products
}

// saveProductToSQL - save product to sql
func saveProductToSQL(productName string, productPrice string, productPoint string, LinkURL string) {
	product := model.Product{
		Name:  productName,
		Price: productPrice,
		Point: productPoint,
		Link:  LinkURL,
	}
	db.Db.Model(&product).Where("name = ?", productName).Update("link", LinkURL)

	product1 := model.Product{}
	db.Db.Where("name = ?", productName).Find(&product1)
	if product1.ID == 0 {
		// fmt.Println("find new product: ", productName)
		product1.Name = productName
		product1.Price = productPrice
		product1.Point = productPoint
		product1.Link = LinkURL
		db.Db.Create(&product1)
	}
}

// saveEduCenterToSQL - save edu center to sql
func saveEduCenterToSQL(centerName string, centerAddr string, centerPhone string, centerDate string) {
	center := model.Center{}
	db.Db.Where("name = ?", centerName).Find(&center)
	if center.ID == 0 {
		center.Name = centerName
		center.Address = centerAddr
		center.Phone = centerPhone
		center.Date = centerDate
		db.Db.Create(&center)
	}
}

// checkError - check error
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// writeProductToExcel - write product to excel
func writeProductToExcel(products []model.ProductCommon) error {
	// Debug
	// fmt.Println(products)

	sheetName := "Sheet1"
	// 開新檔
	f := excelize.NewFile()
	// 定義 sheet name
	index := f.NewSheet(sheetName)
	// 設定 title
	f.SetCellValue(sheetName, "A1", "名稱")
	f.SetCellValue(sheetName, "B1", "價格")
	f.SetCellValue(sheetName, "C1", "PV")
	f.SetCellValue(sheetName, "D1", "連結")

	// 處理每筆資料
	for i, p := range products {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), p.Name)  // 名稱
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), p.Price) // 價格
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), p.Point) // PV
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), p.Link)  // 連結
	}

	// 設定 sheet
	f.SetActiveSheet(index)

	// 儲存檔案
	if err := f.SaveAs("product.xlsx"); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// writeCenterToExcel - write center to excel
func writeCenterToExcel(centers []model.CenterCommon) error {
	// Debug
	// fmt.Println(products)
	sheetName := "Sheet1"
	// 開新檔
	f := excelize.NewFile()
	// 定義 sheet name
	index := f.NewSheet(sheetName)
	// 設定 title
	f.SetCellValue(sheetName, "A1", "名稱")
	f.SetCellValue(sheetName, "B1", "地址")
	f.SetCellValue(sheetName, "C1", "電話號碼")
	f.SetCellValue(sheetName, "D1", "設立日期")

	// 處理每筆資料
	for i, c := range centers {
		// 設定值
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), c.Name)    // 名稱
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), c.Address) // 地址
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), c.Phone)   // 電話號碼
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), c.Date)    // 設立日期
	}

	// 設定 sheet
	f.SetActiveSheet(index)

	// 儲存檔案
	if err := f.SaveAs("center.xlsx"); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func main() {
	// 儲存產品
	products := saveProducts()
	// 儲存中心
	centers := saveEduCenter()
	// 存成 Excel
	if len(products) > 0 {
		writeProductToExcel(products)
	}
	// 存成 Excel
	if len(centers) > 0 {
		writeCenterToExcel(centers)
	}
}
