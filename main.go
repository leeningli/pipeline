package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/olekukonko/tablewriter"
	"github.com/valyala/fasthttp"
)

type TestResp struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func main() {
	fmt.Println("Hello, leeingli.")
	dataTable := [][]string{
		[]string{"leeingli", "shenzhen", "1106"},
		[]string{"lichehuo", "shenzhen", "1207"},
		[]string{"wukong", "huaguoshan", "0000"}}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "City", "Bir"})
	table.AppendBulk(dataTable)
	table.Render()
	router := fasthttprouter.New()
	router.GET("/", testFunc)
	if err := fasthttp.ListenAndServe("0.0.0.0:12345", router.Handler); err != nil {
		fmt.Println("start fasthttp fail:", err.Error())
	}
}

func testFunc(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(200)
	data := TestResp{
		Name: "leeingli",
		Desc: "this is test for github&&wercker&&dockerhub",
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("struct to json is failed:", err)
	}
	ctx.WriteString(string(dataBytes))
}
