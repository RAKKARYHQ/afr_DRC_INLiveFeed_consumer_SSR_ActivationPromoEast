package afr_kafka

import "encoding/xml"

// //////////////////////////////////////
// /////SSR
// //////////////////////////////////////
type SSR_Message struct {
	SSR_Message xml.Name `xml:"message"`
	SSR         SSR      `xml:"ssr"`
}

type SSR struct {
	SSR     xml.Name `xml:"ssr"`
	Transid string   `xml:"transid"`
	Subid   string   `xml:"subid"`
	Time    string   `xml:"time"`
	Type    string   `xml:"type"`
	Accnum  string   `xml:"accnum"`
	Acctype string   `xml:"acctype"`
	Cos     string   `xml:"cos"`
	State   string   `xml:"state"`
	Opid    string   `xml:"opid"`
	Cardnum string   `xml:"cardnum"`
	Strtbal float64  `xml:"strtbal"`
	Endbal  float64  `xml:"endbal"`
	Strtft  float64  `xml:"strtft"`
	Endft   float64  `xml:"endft"`
	Strtfm  float64  `xml:"strtfm"`
	Endfm   float64  `xml:"endfm"`
	Strtfs  float64  `xml:"strtfs"`
	Endfs   float64  `xml:"endfs"`
	Strtfd  float64  `xml:"strtfd"`
	Endfd   float64  `xml:"endfd"`
	Tokbun  string   `xml:"tokbun"`
	Amount  float64  `xml:"amount"`
}

// //////////////////////////////////////
// /////CDRV
// //////////////////////////////////////
type CDRV_Message struct {
	CDRV_Message xml.Name `xml:"message"`
	CDRV         CDRV     `xml:"cdr"`
}

type CDRV struct {
	CDRV   xml.Name `xml:"cdr"`
	sessid string   `xml:"sessid"`
	Subid  string   `xml:"subid"`
	Callid string   `xml:"callid"`
	Trmid  string   `xml:"trmid"`
	Portid string   `xml:"portid"`
	Appid  string   `xml:"appid"`
	Date   string   `xml:"date"`
	Time   string   `xml:"time"`
	//Tzr         string   `xml:"tzr"`
	//Tzl         string   `xml:"tzl"`
	Calldur float64 `xml:"calldur"`
	Chdate  string  `xml:"chdate"`
	Chtime  string  `xml:"chtime"`
	Chdur   float64 `xml:"chdur"`
	Term    string  `xml:"term"`
	Status  string  `xml:"status"`
	Type    string  `xml:"type"`
	Calling string  `xml:"calling"`
	Accnum  string  `xml:"accnum"`
	Origin  string  `xml:"origin"`
	Called  string  `xml:"called"`
	Cparty  string  `xml:"cparty"`
	Follow  string  `xml:"follow"`
	Cos     string  `xml:"cos"`
	//	Credused    float64  `xml:"credused"`

	Region   string  `xml:"region"`
	rgname   string  `xml:"rgname"`
	Currency string  `xml:"currency"`
	Strtbal  float64 `xml:"strtbal"`
	//Tcredused   float64  `xml:"tcredused"`
	Debitax     float64 `xml:"debitax"`
	Debisurch   float64 `xml:"debisurch"`
	Credirebt   float64 `xml:"credirebt"`
	Debirate1   float64 `xml:"debirate1"`
	Debirate2   float64 `xml:"debirate2"`
	Option      string  `xml:"option"`
	Table       string  `xml:"table"`
	Plan        string  `xml:"plan"`
	Campaign    string  `xml:"campaign"`
	Token       string  `xml:"token"`
	Seq         float64 `xml:"seq"`
	Interim     float64 `xml:"interim"`
	Interim_dur float64 `xml:"interim_dur"`
}

// //////////////////////////////////////
// /////CDRS
// //////////////////////////////////////
type CDRS_Message struct {
	CDRS_Message xml.Name `xml:"message"`
	CDRS         CDRS     `xml:"cdr"`
}

type CDRS struct {
	CDRS                 xml.Name `xml:"cdr"`
	Subid                float64  `xml:"subid"`
	Appid                string   `xml:"appid"`
	Threadid             float64  `xml:"Threadid"`
	Origin               string   `xml:"Origin"`
	Date                 string   `xml:"Date"`
	Time                 string   `xml:"Time"`
	Chdate               string   `xml:"chdate"`
	Chtime               string   `xml:"chtime"`
	Status               float64  `xml:"status"`
	Term                 float64  `xml:"term"`
	Currency             string   `xml:"currency"`
	Type                 float64  `xml:"type"`
	Calling              string   `xml:"calling"`
	Cparty               string   `xml:"cparty"`
	Accnum               string   `xml:"accnum"`
	Orig_sms_type        float64  `xml:"orig_sms_type"`
	Original_dest        string   `xml:"original_dest"`
	Original_origin      string   `xml:"original_origin"`
	Debirate1            float64  `xml:"debirate1"`
	Realdebit            float64  `xml:"realdebit"`
	Credused             float64  `xml:"credused"`
	Cos                  float64  `xml:"cos"`
	Strtbal              float64  `xml:"strtbal"`
	Freesms              float64  `xml:"freesms"`
	Tuc                  float64  `xml:"tuc"`
	Fandftype            string   `xml:"fandftype"`
	Credirebt            float64  `xml:"credirebt"`
	Option               float64  `xml:"option"`
	Length               float64  `xml:"length"`
	Encoding             float64  `xml:"encoding"`
	Udhi                 string   `xml:"udhi"`
	Validity_date        string   `xml:"validity_date"`
	Validity_time        string   `xml:"validity_time"`
	Delivery_attempts    float64  `xml:"delivery_attempts"`
	Ref_num              float64  `xml:"ref_num"`
	Total_seg            float64  `xml:"total_seg"`
	Current_seg          float64  `xml:"current_seg"`
	Origmsc              string   `xml:"origmsc"`
	Destmsc              string   `xml:"destmsc"`
	Srrflag              string   `xml:"srrflag"`
	Dfcause              float64  `xml:"dfcause"`
	Mtroute              string   `xml:"mtroute"`
	Fwddest              string   `xml:"fwddest"`
	Esme_hostname        string   `xml:"esme_hostname"`
	Source_imsi          string   `xml:"source_imsi"`
	Dest_imsi            string   `xml:"dest_imsi"`
	VIMSI                string   `xml:"vIMSI"`
	Orig_SCA             string   `xml:"orig_SCA"`
	Receipt_id           float64  `xml:"Receipt_id"`
	Source_esme_hostname string   `xml:"source_esme_hostname"`
	Rateplan             string   `xml:"rateplan"`
	Campaign             string   `xml:"campaign"`
	Moneyctr             float64  `xml:"moneyctr"`
	Smsctr               float64  `xml:"smsctr"`
	Seq                  float64  `xml:"seq"`
	Interim              float64  `xml:"interim"`
	Token                string   `xml:"token"`
}

// //////////////////////////////////////
// /////DDR
// //////////////////////////////////////
type DDR_Message struct {
	DDR_Message xml.Name `xml:"message"`
	DDR         DDR      `xml:"ddr"`
}

type DDR struct {
	DDR         xml.Name `xml:"ddr"`
	Subid       string   `xml:"subid"`
	Accnum      string   `xml:"accnum"`
	Acctype     float64  `xml:"acctype"`
	Sessid      string   `xml:"sessid"`
	Origin      string   `xml:"origin"`
	Time        string   `xml:"time"`
	Dur         float64  `xml:"dur"`
	Units       float64  `xml:"units"`
	Cost        float64  `xml:"cost"`
	Currency    string   `xml:"currency"`
	Strtbal     float64  `xml:"strtbal"`
	Credused    float64  `xml:"credused"`
	Dataused    float64  `xml:"dataused"`
	Type        float64  `xml:"type"`
	Cos         float64  `xml:"cos"`
	Rating      float64  `xml:"rating"`
	Disconnect  float64  `xml:"disconnect"`
	Term        float64  `xml:"term"`
	Partial     float64  `xml:"partial"`
	Orealm      string   `xml:"orealm"`
	Ohost       string   `xml:"ohost"`
	Dhost       string   `xml:"dhost"`
	Sgsn        string   `xml:"sgsn"`
	Ggsn        string   `xml:"ggsn"`
	Chrgid      string   `xml:"chrgid"`
	Pdp         string   `xml:"pdp"`
	Apn         string   `xml:"apn"`
	Roam        float64  `xml:"roam"`
	Csunits     string   `xml:"csunits"`
	Mktcampaign string   `xml:"mktcampaign"`
	Mcounter    float64  `xml:"mcounter"`
	Dcounter    float64  `xml:"dcounter"`
	Ucounter    float64  `xml:"ucounter"`
	Campaign    string   `xml:"campaign"`
	Credtok     string   `xml:"credtok"`
	Strttime    string   `xml:"strttime"`
	Endtime     string   `xml:"endtime"`
	Amount      float64  `xml:"amount"`
	Totrech     string   `xml:"totrech"`
	Totuse      string   `xml:"totuse"`
	Qos         string   `xml:"qos"`
	Application string   `xml:"application"`
	Source      string   `xml:"source"`
}

// //////////////////////////////////////
// /////EDR
// //////////////////////////////////////
type EDR_Message struct {
	EDR_Message xml.Name `xml:"message"`
	EDR         EDR      `xml:"edr"`
}

type EDR struct {
	EDR        xml.Name `xml:"edr"`
	Msisdn     string   `xml:"msisdn"`
	Tok_sub_id string   `xml:"tok_sub_id"`
	Event      string   `xml:"event"`
	Service    string   `xml:"service"`
	Edrtype    string   `xml:"type"`
	Token      string   `xml:"token"`
	Sessid     string   `xml:"sessid"`
	Origin     string   `xml:"origin"`
	Time       string   `xml:"time"`
	Threshold  float64  `xml:"threshold"`
	Granted    float64  `xml:"granted"`
	Current    float64  `xml:"current"`
}
