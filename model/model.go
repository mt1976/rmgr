package model

type Welcome4 struct {
	Msg Msg `json:"msg"`
}

type Msg struct {
	Coll                         Coll   `json:"coll"`
	XmlnsXsi                     string `json:"_xmlns:xsi"`
	XsiNoNamespaceSchemaLocation string `json:"_xsi:noNamespaceSchemaLocation"`
}

type Coll struct {
	Rt []Rt `json:"rt"`
}

type Rt struct {
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
	Own    string `json:"own"`
	Rsk    string `json:"rsk"`
	DTme   string `json:"dTme"`
	Amt    string `json:"amt"`
	AmtCcy string `json:"amtCcy"`
	Sts    string `json:"sts"`
	Cat    string `json:"_cat"`
	ID     string `json:"_id"`
	Src    string `json:"_src"`
}

func (welcome4 *Welcome4) GetMsg() Msg {
	return welcome4.Msg
}

func (welcome4 *Welcome4) SetMsg(Msg Msg) *Welcome4 {
	welcome4.Msg = Msg
	return welcome4
}
func (msg *Msg) GetColl() Coll {
	return msg.Coll
}

func (msg *Msg) GetXmlnsXsi() string {
	return msg.XmlnsXsi
}

func (msg *Msg) GetXsiNoNamespaceSchemaLocation() string {
	return msg.XsiNoNamespaceSchemaLocation
}

func (msg *Msg) SetColl(Coll Coll) *Msg {
	msg.Coll = Coll
	return msg
}

func (msg *Msg) SetXmlnsXsi(XmlnsXsi string) *Msg {
	msg.XmlnsXsi = XmlnsXsi
	return msg
}

func (msg *Msg) SetXsiNoNamespaceSchemaLocation(XsiNoNamespaceSchemaLocation string) *Msg {
	msg.XsiNoNamespaceSchemaLocation = XsiNoNamespaceSchemaLocation
	return msg
}

// func (coll *Coll) GetRt() Rt {
// 	return coll.Rt
// }

//	func (coll *Coll) SetRt(Rt Rt) *Coll {
//		coll.Rt = Rt
//		return coll
//	}
func (rt *Rt) GetBid() string {
	return rt.Bid
}

func (rt *Rt) GetAsk() string {
	return rt.Ask
}

func (rt *Rt) GetOwn() string {
	return rt.Own
}

func (rt *Rt) GetRsk() string {
	return rt.Rsk
}

func (rt *Rt) GetDTme() string {
	return rt.DTme
}

func (rt *Rt) GetAmt() string {
	return rt.Amt
}

func (rt *Rt) GetAmtCcy() string {
	return rt.AmtCcy
}

func (rt *Rt) GetSts() string {
	return rt.Sts
}

func (rt *Rt) GetCat() string {
	return rt.Cat
}

func (rt *Rt) GetID() string {
	return rt.ID
}

func (rt *Rt) GetSrc() string {
	return rt.Src
}

func (rt *Rt) SetBid(Bid string) *Rt {
	rt.Bid = Bid
	return rt
}

func (rt *Rt) SetAsk(Ask string) *Rt {
	rt.Ask = Ask
	return rt
}

func (rt *Rt) SetOwn(Own string) *Rt {
	rt.Own = Own
	return rt
}

func (rt *Rt) SetRsk(Rsk string) *Rt {
	rt.Rsk = Rsk
	return rt
}

func (rt *Rt) SetDTme(DTme string) *Rt {
	rt.DTme = DTme
	return rt
}

func (rt *Rt) SetAmt(Amt string) *Rt {
	rt.Amt = Amt
	return rt
}

func (rt *Rt) SetAmtCcy(AmtCcy string) *Rt {
	rt.AmtCcy = AmtCcy
	return rt
}

func (rt *Rt) SetSts(Sts string) *Rt {
	rt.Sts = Sts
	return rt
}

func (rt *Rt) SetCat(Cat string) *Rt {
	rt.Cat = Cat
	return rt
}

func (rt *Rt) SetID(ID string) *Rt {
	rt.ID = ID
	return rt
}

func (rt *Rt) SetSrc(Src string) *Rt {
	rt.Src = Src
	return rt
}
