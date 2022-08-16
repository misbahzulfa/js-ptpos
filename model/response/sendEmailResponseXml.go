package response

import "encoding/xml"

type SendEmailResponseXml struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Body    struct {
		Text              string `xml:",chardata"`
		SendEmailResponse struct {
			Text   string `xml:",chardata"`
			Ns     string `xml:"ns,attr"`
			Return struct {
				Text      string `xml:",chardata"`
				Ax21      string `xml:"ax21,attr"`
				Xsi       string `xml:"xsi,attr"`
				Type      string `xml:"type,attr"`
				AddValues struct {
					Text string `xml:",chardata"`
					Nil  string `xml:"nil,attr"`
				} `xml:"addValues"`
				Kode string `xml:"kode"`
				Msg  string `xml:"msg"`
				Ret  string `xml:"ret"`
			} `xml:"return"`
		} `xml:"sendEmailResponse"`
	} `xml:"Body"`
}
