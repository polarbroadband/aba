package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

type ZtpClient struct {
	XMLName xml.Name `xml:"ztpclient" json:"-"`          // json - ignore the field
	ID      string   `xml:"id,attr" json:"id,omitempty"` // json - ignore empty field
	Mask    int      `xml:"mask" json:"mask,omitempty"`
	Cluster string   `xml:"cluster" json:"cluster,omitempty"`
	Site    string   `xml:"site" json:"site,omitempty"`
	Device  string   `xml:"device" json:"device,omitempty"`
}

const ridd = `
<ztpclient id="YU90876A">
   <mask>30</mask>
   <device>Ethiopia</device>
</ztpclient>
`
const rid = `<ztpclient id="YU90876A"><mask>29</mask><device>ec01s05</device></ztpclient>`

func xmlToUrl() {
	fmt.Println("Hello, 世界")

	// splite the string with "" will split the each character
	k := ""
	l := "fhhf/skl/t"
	fmt.Println(strings.Split(l, k))

	// convert a xml object to URL query
	var ac ZtpClient
	if err := xml.Unmarshal([]byte(rid), &ac); err != nil {
		panic(err)
	}

	jq, err := json.Marshal(ac)
	if err != nil {
		fmt.Println(err)
	}
	var kvs map[string]interface{}
	json.Unmarshal(jq, &kvs)
	q := []string{}
	for k, v := range kvs {
		q = append(q, fmt.Sprintf("%s=%v", k, v))
	}
	url := "https://hhik:80443/oran/pq/sztp?" + strings.Join(q, "&")
	fmt.Println(url)
}
