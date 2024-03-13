package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mt1976/rmg/config"
)

var C config.Config

type Rate struct {
	Bid        float64 `json:"bid,omitempty"`
	Ask        float64 `json:"ask,omitempty"`
	Owner      string  `json:"owner,omitempty"`
	RiskCentre string  `json:"riskcentre,omitempty"`
	DateTime   string  `json:"dTme,omitempty"`
	Amount     float64 `json:"amt,omitempty"`
	AmtCcy     string  `json:"amtCcy,omitempty"`
	Status     string  `json:"status,omitempty"`
	Category   string  `json:"category,omitempty"`
	ID         string  `json:"id,omitempty"`
	Source     string  `json:"src,omitempty"`
	StaleAfter string  `json:"staleat,omitempty"`
}

func (rt *Rate) GetBid() float64 {
	return rt.Bid
}

func (rt *Rate) GetAsk() float64 {
	return rt.Ask
}

func (rt *Rate) GetOwn() string {
	return rt.Owner
}

func (rt *Rate) GetRsk() string {
	return rt.RiskCentre
}

func (rt *Rate) GetDTmeString() string {
	return rt.DateTime
}

func (rt *Rate) GetDTmeTime() time.Time {
	fmt.Printf("C.DateTimeFormat: %v\n", config.Configuration.DateTimeFormat)
	fmt.Printf("rt.DateTime: %v\n", rt.DateTime)
	t, err := time.Parse(config.Configuration.DateTimeFormat, rt.DateTime)
	fmt.Printf("t: %v\n", t)
	if err != nil {
		fmt.Println("DateConversionError", err.Error())
	}
	return t
}

func (rt *Rate) GetAmtString() string {
	return strconv.FormatFloat(rt.Amount, 'f', 2, 64)
}

func (rt *Rate) GetAmt() float64 {
	return rt.Amount
}

func (rt *Rate) GetAmtCcy() string {
	return rt.AmtCcy
}

func (rt *Rate) GetSts() string {
	return rt.Status
}

func (rt *Rate) GetCat() string {
	return rt.Category
}

func (rt *Rate) GetID() string {
	return rt.ID
}

func (rt *Rate) GetSrc() string {
	return rt.Source
}

func (rt *Rate) SetBid(Bid string) *Rate {
	//strconv.ParseFloat(Bid, 64)
	arse, Err := strconv.ParseFloat(Bid, 64)
	if Err != nil {
		fmt.Println(Err)
	}
	rt.Bid = arse
	return rt
}

func (rt *Rate) SetBidFloat(Bid float64) *Rate {
	rt.Bid = Bid
	return rt
}

func (rt *Rate) SetAsk(Ask string) *Rate {
	arse, Err := strconv.ParseFloat(Ask, 64)
	if Err != nil {
		fmt.Println(Err)
	}
	rt.Ask = arse
	return rt
}

func (rt *Rate) SetAskFloat(Ask float64) *Rate {
	rt.Ask = Ask
	return rt
}

func (rt *Rate) SetOwn(Own string) *Rate {
	rt.Owner = Own
	return rt
}

func (rt *Rate) SetRsk(Rsk string) *Rate {
	rt.RiskCentre = Rsk
	return rt
}

func (rt *Rate) SetDTme(DTme string) *Rate {
	rt.DateTime = DTme
	return rt
}

func (rt *Rate) SetDTmeTime(DTme time.Time) *Rate {
	rt.DateTime = DTme.Format(C.DateTimeFormat)
	return rt
}

func (rt *Rate) SetAmtString(Amt string) *Rate {
	rt.Amount, _ = strconv.ParseFloat(Amt, 64)
	return rt
}

func (rt *Rate) SetAmt(Amt float64) *Rate {
	rt.Amount = Amt
	return rt
}

func (rt *Rate) SetAmtCcy(AmtCcy string) *Rate {
	rt.AmtCcy = AmtCcy
	return rt
}

func (rt *Rate) SetSts(Sts string) *Rate {
	rt.Status = Sts
	return rt
}

func (rt *Rate) SetCat(Cat string) *Rate {
	rt.Category = Cat
	return rt
}

func (rt *Rate) SetID(ID string) *Rate {
	rt.ID = ID
	return rt
}

func (rt *Rate) SetSrc(Src string) *Rate {
	rt.Source = Src
	return rt
}

func (rt *Rate) GetStaleAfter() time.Time {
	// parse to time
	xx, err := time.Parse(config.Configuration.DateTimeFormat, rt.StaleAfter)
	if err != nil {
		fmt.Println("Error parsing stale after time: ", err)
	}
	return xx
}

func (rt *Rate) GetStaleAfterString() string {
	return rt.StaleAfter
}

func (rt *Rate) SetStaleAfter(afterMS int) {
	fmt.Printf("rt.GetDTmeTime(): %v\n", rt.GetDTmeTime())
	fmt.Printf("afterMS: %v\n", afterMS)
	fmt.Printf("time.Millisecond: %v\n", time.Millisecond)
	fmt.Printf("config.Configuration.DateTimeFormat: %v\n", config.Configuration.DateTimeFormat)
	fmt.Printf("time.Duration(afterMS): %v\n", time.Duration(afterMS))
	fmt.Printf("(time.Duration(afterMS) * time.Millisecond): %v\n", (time.Duration(afterMS) * time.Millisecond))
	rt.StaleAfter = rt.GetDTmeTime().Add(time.Duration(afterMS) * time.Millisecond).Format(config.Configuration.DateTimeFormat)
}
