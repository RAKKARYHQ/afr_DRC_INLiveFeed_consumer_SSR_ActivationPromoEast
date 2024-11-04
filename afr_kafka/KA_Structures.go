package afr_kafka

import (
	"time"
)

type KafkaTopic struct {
	Key         string `bson:"Key" json:"Key"`
	Name        string `bson:"Name" json:"Name"`
	Description string `bson:"Description" json:"Description"`
}

type DB_SSR struct {
	Transid string    `bson:"Transid" json:"Transid"`
	Subid   string    `bson:"Subid" json:"Subid"`
	Time    time.Time `bson:"Time" json:"Time"`
	Type    string    `bson:"Type" json:"Type"`
	Accnum  string    `bson:"Accnum" json:"Accnum"`
	Acctype string    `bson:"Acctype" json:"Acctype"`
	Cos     string    `bson:"Cos" json:"Cos"`
	State   string    `bson:"State" json:"State"`
	Opid    string    `bson:"Opid" json:"Opid"`
	Cardnum string    `bson:"Cardnum" json:"Cardnum"`
	Strtbal float64   `bson:"Strtbal" json:"Strtbal"`
	Endbal  float64   `bson:"Endbal" json:"Endbal"`
	Strtft  float64   `bson:"Strtft" json:"Strtft"`
	Endft   float64   `bson:"Endft" json:"Endft"`
	Strtfm  float64   `bson:"Strtfm" json:"Strtfm"`
	Endfm   float64   `bson:"Endfm" json:"Endfm"`
	Strtfs  float64   `bson:"Strtfs" json:"Strtfs"`
	Endfs   float64   `bson:"Endfs" json:"Endfs"`
	Strtfd  float64   `bson:"Strtfd" json:"Strtfd"`
	Endfd   float64   `bson:"Endfd" json:"Endfd"`
	Tokbun  string    `bson:"Tokbun" json:"Tokbun"`
	Amount  float64   `bson:"Amount" json:"Amount"`
}

type DB_CDRV struct {
	sessid string `bson:"sessid" json:"sessid"`
	Subid  string `bson:"Subid" json:"Subid"`
	Callid string `bson:"Callid" json:"Callid"`
	Trmid  string `bson:"Trmid" json:"Trmid"`
	Portid string `bson:"Portid" json:"Portid"`
	Appid  string `bson:"Appid" json:"Appid"`
	//Date        string  `bson:"Date" json:"Date"`
	Time    time.Time `bson:"Time" json:"Time"`
	Tzr     string    `bson:"Tzr" json:"Tzr"`
	Tzl     string    `bson:"Tzl" json:"Tzl"`
	Calldur float64   `bson:"Calldur" json:"Calldur"`
	//Chdate      string  `bson:"Chdate" json:"Chdate"`
	Chtime      time.Time `bson:"Chtime" json:"Chtime"`
	Chdur       float64   `bson:"Chdur" json:"Chdur"`
	Term        string    `bson:"Term" json:"Term"`
	Status      string    `bson:"Status" json:"Status"`
	Type        string    `bson:"Type" json:"Type"`
	Calling     string    `bson:"Calling" json:"Calling"`
	Accnum      string    `bson:"Accnum" json:"Accnum"`
	Origin      string    `bson:"Origin" json:"Origin"`
	Called      string    `bson:"Called" json:"Called"`
	Cparty      string    `bson:"Cparty" json:"Cparty"`
	Follow      string    `bson:"Follow" json:"Follow"`
	Cos         string    `bson:"Cos" json:"Cos"`
	Credused    float64   `bson:"Credused" json:"Credused"`
	Region      string    `bson:"Region" json:"Region"`
	rgname      string    `bson:"rgname" json:"rgname"`
	Currency    string    `bson:"Currency" json:"Currency"`
	Strtbal     float64   `bson:"Strtbal" json:"Strtbal"`
	Tcredused   float64   `bson:"Tcredused" json:"Tcredused"`
	Debitax     float64   `bson:"Debitax" json:"Debitax"`
	Debisurch   float64   `bson:"Debisurch" json:"Debisurch"`
	Credirebt   float64   `bson:"Credirebt" json:"Credirebt"`
	Debirate1   float64   `bson:"Debirate1" json:"Debirate1"`
	Debirate2   float64   `bson:"Debirate2" json:"Debirate2"`
	Option      string    `bson:"Option" json:"Option"`
	Table       string    `bson:"Table" json:"Table"`
	Plan        string    `bson:"Plan" json:"Plan"`
	Campaign    string    `bson:"Campaign" json:"Campaign"`
	Tcounter    float64   `bson:"Tcounter" json:"Tcounter"`
	Mcounter    float64   `bson:"Mcounter" json:"Mcounter"`
	Seq         float64   `bson:"Seq" json:"Seq"`
	Interim     float64   `bson:"Interim" json:"Interim"`
	Interim_dur float64   `bson:"Interim_dur" json:"Interim_dur"`
	Token       string    `bson:"Token" json:"Token"`
}

type DB_CDRS struct {
	Subid                float64   `bson:"Subid" json:"Subid"`
	Appid                string    `bson:"Appid" json:"Appid"`
	Threadid             float64   `bson:"Threadid" json:"Threadid"`
	Origin               string    `bson:"Origin" json:"Origin"`
	Chtime               time.Time `bson:"Chtime" json:"Chtime"`
	Status               float64   `bson:"Status" json:"Status"`
	Term                 float64   `bson:"Term" json:"Term"`
	Currency             string    `bson:"Currency" json:"Currency"`
	Type                 float64   `bson:"Type" json:"Type"`
	Calling              string    `bson:"Calling" json:"Calling"`
	Cparty               string    `bson:"Cparty" json:"Cparty"`
	Accnum               string    `bson:"Accnum" json:"Accnum"`
	Orig_sms_type        float64   `bson:"Orig_sms_type" json:"Orig_sms_type"`
	Original_dest        string    `bson:"Original_dest" json:"Original_dest"`
	Original_origin      string    `bson:"Original_origin" json:"Original_origin"`
	Debirate1            float64   `bson:"Debirate1" json:"Debirate1"`
	Realdebit            float64   `bson:"Realdebit" json:"Realdebit"`
	Credused             float64   `bson:"Credused" json:"Credused"`
	Cos                  float64   `bson:"Cos" json:"Cos"`
	Strtbal              float64   `bson:"Strtbal" json:"Strtbal"`
	Freesms              float64   `bson:"Freesms" json:"Freesms"`
	Tuc                  float64   `bson:"Tuc" json:"Tuc"`
	Fandftype            string    `bson:"Fandftype" json:"Fandftype"`
	Credirebt            float64   `bson:"Credirebt" json:"Credirebt"`
	Option               float64   `bson:"Option" json:"Option"`
	Length               float64   `bson:"Length" json:"Length"`
	Encoding             float64   `bson:"Encoding" json:"Encoding"`
	Udhi                 string    `bson:"Udhi" json:"Udhi"`
	Validity_date        string    `bson:"Validity_date" json:"Validity_date"`
	Validity_time        string    `bson:"Validity_time" json:"Validity_time"`
	Delivery_attempts    float64   `bson:"Delivery_attempts" json:"Delivery_attempts"`
	Ref_num              float64   `bson:"Ref_num" json:"Ref_num"`
	Total_seg            float64   `bson:"Total_seg" json:"Total_seg"`
	Current_seg          float64   `bson:"Current_seg" json:"Current_seg"`
	Origmsc              string    `bson:"Origmsc" json:"Origmsc"`
	Destmsc              string    `bson:"Destmsc" json:"Destmsc"`
	Srrflag              string    `bson:"Srrflag" json:"Srrflag"`
	Dfcause              float64   `bson:"Dfcause" json:"Dfcause"`
	Mtroute              string    `bson:"Mtroute" json:"Mtroute"`
	Fwddest              string    `bson:"Fwddest" json:"Fwddest"`
	Esme_hostname        string    `bson:"Esme_hostname" json:"Esme_hostname"`
	Source_imsi          string    `bson:"Source_imsi" json:"Source_imsi"`
	Dest_imsi            string    `bson:"Dest_imsi" json:"Dest_imsi"`
	VIMSI                string    `bson:"VIMSI" json:"VIMSI"`
	Orig_SCA             string    `bson:"Orig_SCA" json:"Orig_SCA"`
	Receipt_id           float64   `bson:"Receipt_id" json:"Receipt_id"`
	Source_esme_hostname string    `bson:"Source_esme_hostname" json:"Source_esme_hostname"`
	Rateplan             string    `bson:"Rateplan" json:"Rateplan"`
	Campaign             string    `bson:"Campaign" json:"Campaign"`
	Moneyctr             float64   `bson:"Moneyctr" json:"Moneyctr"`
	Smsctr               float64   `bson:"Smsctr" json:"Smsctr"`
	Seq                  float64   `bson:"Seq" json:"Seq"`
	Interim              float64   `bson:"Interim" json:"Interim"`
	Token                string    `bson:"Token" json:"Token"`
}

type DB_DDR struct {
	Subid       string    `bson:"Subid" json:"Subid"`
	Accnum      string    `bson:"Accnum" json:"Accnum"`
	Acctype     float64   `bson:"Acctype" json:"Acctype"`
	Sessid      string    `bson:"Sessid" json:"Sessid"`
	Origin      string    `bson:"Origin" json:"Origin"`
	Time        time.Time `bson:"Time" json:"Time"`
	Dur         float64   `bson:"Dur" json:"Dur"`
	Units       float64   `bson:"Units" json:"Units"`
	Cost        float64   `bson:"Cost" json:"Cost"`
	Currency    string    `bson:"Currency" json:"Currency"`
	Strtbal     float64   `bson:"Strtbal" json:"Strtbal"`
	Credused    float64   `bson:"Credused" json:"Credused"`
	Dataused    float64   `bson:"Dataused" json:"Dataused"`
	Type        float64   `bson:"Type" json:"Type"`
	Cos         float64   `bson:"Cos" json:"Cos"`
	Rating      float64   `bson:"Rating" json:"Rating"`
	Disconnect  float64   `bson:"Disconnect" json:"Disconnect"`
	Term        float64   `bson:"Term" json:"Term"`
	Partial     float64   `bson:"Partial" json:"Partial"`
	Orealm      string    `bson:"Orealm" json:"Orealm"`
	Ohost       string    `bson:"Ohost" json:"Ohost"`
	Dhost       string    `bson:"Dhost" json:"Dhost"`
	Sgsn        string    `bson:"Sgsn" json:"Sgsn"`
	Ggsn        string    `bson:"Ggsn" json:"Ggsn"`
	Chrgid      string    `bson:"Chrgid" json:"Chrgid"`
	Pdp         string    `bson:"Pdp" json:"Pdp"`
	Apn         string    `bson:"Apn" json:"Apn"`
	Roam        float64   `bson:"Roam" json:"Roam"`
	Csunits     string    `bson:"Csunits" json:"Csunits"`
	Mktcampaign string    `bson:"Mktcampaign" json:"Mktcampaign"`
	Mcounter    float64   `bson:"Mcounter" json:"Mcounter"`
	Dcounter    float64   `bson:"Dcounter" json:"Dcounter"`
	Ucounter    float64   `bson:"Ucounter" json:"Ucounter"`
	Campaign    string    `bson:"Campaign" json:"Campaign"`
	Credtok     string    `bson:"Credtok" json:"Credtok"`
	Strttime    time.Time `bson:"Strttime" json:"Strttime"`
	Endtime     time.Time `bson:"Endtime" json:"Endtime"`
	Amount      float64   `bson:"Amount" json:"Amount"`
	Totrech     string    `bson:"Totrech" json:"Totrech"`
	Totuse      string    `bson:"Totuse" json:"Totuse"`
	Qos         string    `bson:"Qos" json:"Qos"`
	Application string    `bson:"Application" json:"Application"`
	Source      string    `bson:"Source" json:"Source"`
}

type DB_EDR struct {
	Msisdn     string    `bson:"Msisdn" json:"Msisdn"`
	Tok_sub_id string    `bson:"Tok_sub_id" json:"Tok_sub_id"`
	Event      string    `bson:"Event" json:"Event"`
	Service    string    `bson:"Service" json:"Service"`
	Edrtype    string    `bson:"Edrtype" json:"Edrtype"`
	Token      string    `bson:"Token" json:"Token"`
	Sessid     string    `bson:"Sessid" json:"Sessid"`
	Origin     string    `bson:"Origin" json:"Origin"`
	Time       time.Time `bson:"Time" json:"Time"`
	Threshold  float64   `bson:"Threshold" json:"Threshold"`
	Granted    float64   `bson:"Granted" json:"Granted"`
	Current    float64   `bson:"Current" json:"Current"`
}
