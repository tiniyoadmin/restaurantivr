package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/labstack/echo"
)

type Number struct {
	Text            string `xml:",chardata"`
	SendDigits      string `xml:"sendDigits,attr,omitempty"`
	SendOnPreanswer string `xml:"sendOnPreanswer,attr,omitempty"`
}

type User struct {
	Text            string `xml:",chardata"`
	SendDigits      string `xml:"sendDigits,attr,omitempty"`
	SendOnPreanswer string `xml:"sendOnPreanswer,attr,omitempty"`
	SipHeaders      string `xml:"sipHeaders,attr,omitempty"`
}

type Sip struct {
	Text            string `xml:",chardata"`
	SendDigits      string `xml:"sendDigits,attr,omitempty"`
	SendOnPreanswer string `xml:"sendOnPreanswer,attr,omitempty"`
	SipHeaders      string `xml:"sipHeaders,attr,omitempty"`
}

type Play struct {
	XMLName        xml.Name `xml:"Play"`
	Text           string   `xml:",chardata"`
	Loop           string   `xml:"loop,attr,omitempty"`
	CallbackURL    string   `xml:"callback_url,attr,omitempty"`
	CallbackMethod string   `xml:"callback_method,attr,omitempty"`
}

type Say struct {
	XMLName  xml.Name `xml:"Say"`
	Text     string   `xml:",chardata"`
	Loop     string   `xml:"loop,attr,omitempty"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
}

type Redirect struct {
	XMLName xml.Name `xml:"Redirect"`
	Text    string   `xml:",chardata"`
	Method  string   `xml:"method,attr,omitempty"`
}

type Dial struct {
	Text           string  `xml:",chardata"`
	Number         *Number `xml:"Number,omitempty"`
	User           *User   `xml:"User,omitempty"`
	Sip            *Sip    `xml:"Sip,omitempty"`
	RingTone       string  `xml:"ringTone,attr,omitempty"`
	Record         string  `xml:"record,attr,omitempty"`
	AnswerOnBridge bool    `xml:"answerOnBridge,attr,omitempty"`
	CallerId       string  `xml:"callerId,attr,omitempty"`
	Timeout        string  `xml:"timeout,attr,omitempty"`
	Action         string  `xml:"action,attr,omitempty"`
}

type Pause struct {
	Text   string `xml:",chardata"`
	Length int    `xml:"length,attr,omitempty"`
}

type Reject struct {
	Text   string `xml:",chardata"`
	Reason string `xml:"reason,attr,omitempty"`
}

type Gather struct {
	XMLName xml.Name `xml:"Gather"`
	Text    string   `xml:",chardata"`
	Say     *Say     `xml:"Say"`
	Play    *Play    `xml:"Play"`

	Action              string `xml:"action,attr,omitempty"`
	Method              string `xml:"method,attr,omitempty"`
	FinishOnKey         string `xml:"finishOnKey,attr,omitempty"`
	NumDigits           string `xml:"numDigits,attr,omitempty"`
	Timeout             string `xml:"timeout,attr,omitempty"`
	ActionOnEmptyResult string `xml:"actionOnEmptyResult,attr,omitempty"`
}

type Hangup struct {
	XMLName xml.Name `xml:"Hangup"`
	Text    string   `xml:",chardata"`
}

type Response struct {
	XMLName  xml.Name  `xml:"Response"`
	Text     string    `xml:",chardata"`
	Redirect *Redirect `xml:"Redirect,omitempty"`
	Reject   *Reject   `xml:"Reject"`
	Gather   *Gather   `xml:"Gather,omitempty"`
	Play     *Play     `xml:"Play,omitempty"`
	Pause    *Pause    `xml:"Pause,omitempty"`
	Say      *Say      `xml:"Say,omitempty"`
	Dial     *Dial     `xml:"Dial,omitempty"`
	Hangup   *Hangup   `xml:Hangup,omitempty`
}

type StatusCallback struct {
	CallSid       string `json:"CallSid" form:"CallSid" query:"CallSid"`
	AccountSid    string `json:"AccountSid" form:"AccountSid" query:"AccountSid"`
	From          string `json:"From" form:"From" query:"From"`
	To            string `json:"To" form:"To" query:"To"`
	CallStatus    string `json:"CallStatus" form:"CallStatus" query:"CallStatus"`
	ApiVersion    string `json:"ApiVersion" form:"ApiVersion" query:"ApiVersion"`
	Direction     string `json:"Direction" form:"Direction" query:"Direction"`
	ForwardedFrom string `json:"ForwardedFrom" form:"ForwardedFrom" query:"ForwardedFrom"`
	CallerName    string `json:"CallerName" form:"CallerName" query:"CallerName"`
	ParentCallSid string `json:"ParentCallSid" form:"ParentCallSid" query:"ParentCallSid"`

	CallDuration      string `json:"CallDuration,omitempty" form:"CallDuration" query:"CallDuration"`
	SipResponseCode   string `json:"SipResponseCode,omitempty" form:"SipResponseCode" query:"SipResponseCode"`
	RecordingUrl      string `json:"RecordingUrl,omitempty" form:"RecordingUrl" query:"RecordingUrl"`
	RecordingSid      string `json:"RecordingSid,omitempty" form:"RecordingSid" query:"RecordingSid"`
	RecordingDuration string `json:"RecordingDuration,omitempty" form:"RecordingDuration" query:"RecordingDuration"`
	Timestamp         string `json:"Timestamp,omitempty" form:"Timestamp" query:"Timestamp"`
	CallbackSource    string `json:"CallbackSource,omitempty" form:"CallbackSource" query:"CallbackSource"`
	SequenceNumber    string `json:"SequenceNumber,omitempty" form:"SequenceNumber" query:"SequenceNumber"`
	Digits            string `json:"Digits,omitempty" form:"Digits" query:"Digits"`
}

type DialActionStatusCallback struct {
	DialCallStatus   string `json:"DialCallStatus" form:"DialCallStatus" query:"DialCallStatus"`
	DialCallSid      string `json:"DialCallSid" form:"DialCallSid" query:"DialCallSid"`
	DialCallDuration string `json:"DialCallDuration" form:"DialCallDuration" query:"DialCallDuration"`
	RecordingUrl     string `json:"RecordingUrl" form:"RecordingUrl" query:"RecordingUrl"`
	CallSid          string `json:"CallSid" form:"CallSid" query:"CallSid"`
	AccountSid       string `json:"AccountSid" form:"AccountSid" query:"AccountSid"`
	From             string `json:"From" form:"From" query:"From"`
	To               string `json:"To" form:"To" query:"To"`
	CallStatus       string `json:"CallStatus" form:"CallStatus" query:"CallStatus"`
	ApiVersion       string `json:"ApiVersion" form:"ApiVersion" query:"ApiVersion"`
	Direction        string `json:"Direction" form:"Direction" query:"Direction"`
	ForwardedFrom    string `json:"ForwardedFrom" form:"ForwardedFrom" query:"ForwardedFrom"`
	CallerName       string `json:"CallerName" form:"CallerName" query:"CallerName"`
	ParentCallSid    string `json:"ParentCallSid" form:"ParentCallSid" query:"ParentCallSid"`
}

// https://66da3a82c5f1.ngrok.io/TiniyoApplications/MainRestaurantMenu

func main() {
	e := echo.New()
	numberRestMap := new(PhonenumberMap)

	e.GET("/v1/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!!!")
	})

	e.POST("/TiniyoApplications/DirectCall", func(c echo.Context) error {
		resp := &Response{}
		resp.Text = ""
		u := StatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := &Response{}
			rejectresp.Reject = &Reject{
				Reason: "rejected",
			}
			return c.XML(http.StatusOK, resp)
		}
		if strings.HasPrefix(u.To, "sip") {
			resp.Dial = &Dial{
				AnswerOnBridge: true,
				Sip: &Sip{
					Text: u.To,
				},
			}
		} else {
			resp.Dial = &Dial{
				AnswerOnBridge: true,
				Number: &Number{
					Text: u.To,
				},
			}
		}
		return c.XML(http.StatusOK, resp)
	})

	// from number : customer number
	// to number : inbound tiniyo number.
	e.GET("/TiniyoApplications/MainRestaurantMenu", func(c echo.Context) error {
		u := StatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := &Response{}
			rejectresp.Reject = &Reject{
				Reason: "rejected",
			}
			return c.XML(http.StatusOK, rejectresp)
		}
		ivrRest := new(RestaurentIVR)
		numberRestMap.StoreNumberInstance(u.From, ivrRest)
		resp := ivrRest.GetMainMenuResponse()
		return c.XML(http.StatusOK, resp)
	})

	e.POST("/TiniyoApplications/DtmfReceived", func(c echo.Context) error {
		u := StatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}

		fmt.Println("Dtmf Digit : ", u.Digits, u.From)

		lastChar := u.Digits[len(u.Digits)-1:]

		if lastChar == "#" {
			u.Digits = u.Digits[0 : len(u.Digits)-1]
		}

		ivrRest := numberRestMap.GetNumberInstance(u.From)

		if ivrRest == nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}

		resp := ivrRest.ProcessDTMFDigits(u.Digits)

		return c.XML(http.StatusOK, resp)
	})

	e.GET("/TiniyoApplications/ReceptionIVR", func(c echo.Context) error {
		resp := &Response{}
		resp.Play = &Play{
			Text: "https://tiniyo.s3-ap-southeast-1.amazonaws.com/public/WelcomeMessage.mp3",
		}
		resp.Dial = &Dial{
			Timeout:  "30",
			CallerId: "913366236661",
			Number: &Number{
				Text: "+919903333376",
			},
			Action: "https://tiniyo.dev/TiniyoApplications/ReceptionIVRCB",
		}
		return c.XML(http.StatusOK, resp)
	})

	e.POST("/TiniyoApplications/ReceptionIVRCB", func(c echo.Context) error {
		u := DialActionStatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}
		resp := &Response{}
		if u.DialCallStatus == "busy" ||
			u.DialCallStatus == "no-answer" ||
			u.DialCallStatus == "failed" ||
			u.DialCallStatus == "canceled" {
			resp.Dial = &Dial{
				Timeout:  "45",
				CallerId: "913366236661",
				Number: &Number{
					Text: "+919614478482",
				},
				Action: "https://tiniyo.dev/TiniyoApplications/ReceptionIVRCB2",
			}
		}
		return c.XML(http.StatusOK, resp)
	})

	e.POST("/TiniyoApplications/ReceptionIVRCB2", func(c echo.Context) error {
		u := DialActionStatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}
		resp := &Response{}
		if u.DialCallStatus == "busy" ||
			u.DialCallStatus == "no-answer" ||
			u.DialCallStatus == "failed" ||
			u.DialCallStatus == "canceled" {
			// Send Message to both numbers.
		}
		return c.XML(http.StatusOK, resp)
	})

	e.GET("/TiniyoApplications/KolkataMixtapeWelcome", func(c echo.Context) error {
		resp := &Response{}
		resp.Text = ""
		resp.Gather = &Gather{
			Action:      "https://tiniyo.dev/TiniyoApplications/MixtapeDtmfReceived",
			NumDigits:   "1",
			FinishOnKey: "#",
			Method:      "POST",
		}
		resp.Gather.Play = &Play{
			Text: "https://tiniyo.s3-ap-southeast-1.amazonaws.com/public/KolkataMixtapeWelcome.mp3",
		}
		resp.Say = &Say{
			Text: "We didn't receive any input. Goodbye!",
		}
		return c.XML(http.StatusOK, resp)
	})

	e.POST("/TiniyoApplications/MixtapeDtmfReceived", func(c echo.Context) error {
		resp := &Response{}
		resp.Play = &Play{
			Text: "https://tiniyo.s3-ap-southeast-1.amazonaws.com/public/KolkataMixtapeHold.mp3",
		}
		resp.Dial = &Dial{
			Timeout:  "25",
			CallerId: "913366236660",
			RingTone: "https://tiniyo.s3-ap-southeast-1.amazonaws.com/public/KolkataMixtapeWelcome.mp3",
			Number: &Number{
				Text: "+919967609476",
			},
			Action: "https://tiniyo.dev/TiniyoApplications/KolkataMixtapeCB",
		}
		return c.XML(http.StatusOK, resp)
	})

	e.POST("/TiniyoApplications/KolkataMixtapeCB", func(c echo.Context) error {
		u := DialActionStatusCallback{}
		err := c.Bind(&u)
		if err != nil {
			rejectresp := GetRejectedResponse()
			return c.XML(http.StatusOK, rejectresp)
		}
		resp := &Response{}
		if u.DialCallStatus == "busy" ||
			u.DialCallStatus == "no-answer" ||
			u.DialCallStatus == "failed" ||
			u.DialCallStatus == "canceled" {
			resp.Dial = &Dial{
				Timeout:  "45",
				CallerId: "913366236660",
				Number: &Number{
					Text: "+917976055614",
				},
			}
		}
		return c.XML(http.StatusOK, resp)
	})

	port := strconv.Itoa(beego.BConfig.Listen.HTTPPort)

	e.Logger.Fatal(e.Start(":" + port))
}
