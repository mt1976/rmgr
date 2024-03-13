package model

import (
	"fmt"
	"strconv"
)

type Rate struct {
	Bid        float64 `json:"bid,omitempty"`
	Ask        float64 `json:"ask,omitempty"`
	Owner      string  `json:"owner,omitempty"`
	RiskCentre string  `json:"riskcentre,omitempty"`
	DateTime   string  `json:"dTme,omitempty"`
	Amount     string  `json:"amt,omitempty"`
	AmtCcy     string  `json:"amtCcy,omitempty"`
	Status     string  `json:"status,omitempty"`
	Category   string  `json:"category,omitempty"`
	ID         string  `json:"id,omitempty"`
	Source     string  `json:"src,omitempty"`
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

func (rt *Rate) GetDTme() string {
	return rt.DateTime
}

func (rt *Rate) GetAmt() string {
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

func (rt *Rate) SetAmt(Amt string) *Rate {
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
