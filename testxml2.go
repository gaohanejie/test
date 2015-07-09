package main

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"log"

	"code.google.com/p/mahonia"
)

// 客户登陆[800101]
type Request_800101 struct {
	XMLName xml.Name    `xml:"request"`
	Head    RequestHead `xml:"head"`
	Body    RequestBody `xml:"body"`
}
type RequestHead struct {
	H_exch_code string `xml:"h_exch_code"`
	H_bank_no   string `xml:"h_bank_no"`
	H_user_id   string `xml:"h_user_id"`
	H_branch_id string `xml:"h_branch_id"`
	H_fact_date string `xml:"h_fact_date"`
	H_fact_time string `xml:"h_fact_time"`
	H_exch_date string `xml:"h_exch_date"`
	H_serial_no string `xml:"h_serial_no"`
	H_rsp_code  string `xml:"h_rsp_code"`
	H_rsp_msg   string `xml:"h_rsp_msg"`
}
type RequestBody struct {
	Record []ReqRecord_800101 `xml:"record"`
}

// 客户登陆[800101]请求报文体
type ReqRecord_800101 struct {
	user_pwd          string `xml:"h_rsp_msg"`
	login_ip          string `xml:"login_ip"`
	login_server_code string `xml:"login_server_code"`
}
type Result struct {
	XMLName xml.Name `xml:"persons"`
	Persons []Person `xml:"person"`
}
type Person struct {
	Name      string   `xml:"name,attr"`
	Age       int      `xml:"age,attr"`
	Career    string   `xml:"career"`
	Interests []string `xml:"interests>interest"`
}

func main() {
	content, err := ioutil.ReadFile("studygolang.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
	log.Println(result.Persons[0].Name)

	//input := "<?xml version=\"1.0\" encoding=\"GBK\"?><request><head><h_exch_code>800101</h_exch_code><h_bank_no>1111</h_bank_no><h_user_id>1038738897</h_user_id><h_branch_id>B00008211</h_branch_id><h_fact_date>20110321</h_fact_date><h_fact_time>16:28:30</h_fact_time><h_exch_date>20130929</h_exch_date><h_serial_no>123456</h_serial_no><h_rsp_code>hj123545</h_rsp_code><h_rsp_msg>ar</h_rsp_msg></head><body><record><user_pwd>54f9b3396fe28c208d525db21588965c</user_pwd></record></body></request>"
	//inputReader := strings.NewReader(input)
	content, err = ioutil.ReadFile("request.xml")
	decoder := mahonia.NewDecoder("gb18030")
	r := bufio.NewReader(decoder.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(content)
	var request Request_800101
	err = xml.Unmarshal(content, &request)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(request)
}
