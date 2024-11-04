package afr_kafka

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////////////////////////
// Functions to generate OTP PINs and authentication
// //////////////////////////////////////////////////////////////////////////////////////////////////
const (
	letterAndDigitsBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 52 possibilities
	digitsBytes          = "0123456789"
	letterIdxBits        = 6                    // 6 bits to represent 64 possibilities / indexes
	letterIdxMask        = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func GenerateAPIAuthenticationKey(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterAndDigitsBytes) {
			result[i] = letterAndDigitsBytes[idx]
			i++
		}
	}
	return string(result)
}

func GenerateOTP(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(digitsBytes) {
			result[i] = digitsBytes[idx]
			i++
		}
	}
	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////

func Ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}

func GetTimeStamp() (YYYYMMDD string, HHMISS string) {
	CurrentDateTime := time.Now()
	yyyy, mm, dd := CurrentDateTime.Date()
	hr, mi, ss := CurrentDateTime.Clock()

	str_yyyy := strconv.Itoa(yyyy)
	str_mm := strconv.Itoa(int(mm))
	if mm < 10 {
		str_mm = "0" + str_mm
	}
	str_dd := strconv.Itoa(dd)
	if dd < 10 {
		str_dd = "0" + str_dd
	}
	YYYYMMDD = str_yyyy + str_mm + str_dd

	str_hr := strconv.Itoa(int(hr))
	if hr < 10 {
		str_hr = "0" + str_hr
	}
	str_mi := strconv.Itoa(int(mi))
	if mi < 10 {
		str_mi = "0" + str_mi
	}
	str_ss := strconv.Itoa(int(ss))
	if ss < 10 {
		str_ss = "0" + str_ss
	}
	HHMISS = str_hr + str_mi + str_ss
	return YYYYMMDD, HHMISS
}

func GetTimeParts(_Date time.Time) (YYYY string, MM string, MMM string, DD string, DD_Ordinal string, HH string, MI string) {
	yyyy, mm, dd := _Date.Date()
	hr, mi, _ := _Date.Clock()

	YYYY = strconv.Itoa(yyyy)
	MM = strconv.Itoa(int(mm))
	if mm < 10 {
		MM = "0" + MM
	}
	MMM = mm.String()[:3]

	DD = strconv.Itoa(dd)
	if dd < 10 {
		DD = "0" + DD
	}
	DD_Ordinal = Ordinal(dd)
	HH = strconv.Itoa(int(hr))
	if hr < 10 {
		HH = "0" + HH
	}
	MI = strconv.Itoa(int(mi))
	if mi < 10 {
		MI = "0" + MI
	}
	return YYYY, MM, MMM, DD, DD_Ordinal, HH, MI
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////
// //Hashing Functions///////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////////////////////////
var Hash_Secret_Key string = "A3k%a2l&$&&34Fo1~~2Fo|003j|`j%&*hlksalkdj7|~jb343!2e09df{12$$^1u)U(*@("

func Hashing(input string, secretKey string) (Output string) {
	md5 := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()
	io.WriteString(md5, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
	//sha_512_256 := sha512.Sum512_256([]byte(input))
	hmac512 := hmac.New(sha512.New, []byte(secretKey))
	hmac512.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(hmac512.Sum(nil))

	//fmt.Printf("md5:\t\t%x\n", md5.Sum(nil))
	//fmt.Printf("sha256:\t\t%x\n", sha_256.Sum(nil))
	//fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))
	//fmt.Printf("sha512_256:\t%x\n", sha_512_256)
	//fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
}

// ////////////////////////////////////////////////////////
// Function to get the source IP of the request
// ////////////////////////////////////////////////////////
func GetRequestIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}
	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("no valid ip found")
}

// ////////////////////////////////////////////////////////////////////////
// ////////////////Decimal precision & Comma Separator/////////////////////
// ////////////////////////////////////////////////////////////////////////
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func Commaf(v float64) string {
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	comma := []byte{','}

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////
// /////SEND SMS////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////////////////////////////////////////////////////////////////////////
// func SendSMS(Sender string, target string, SMSText string) (_rErr error) {
// 	log.Println("Sending SMS: Sender (" + Sender + "), Target (" + target + "), text (" + SMSText + ") ")
// 	//	url := "http://10.250.8.50:15403/?systemid=EMIS_OA&password=r5qtRVIV1a&Originator=" + Sender + "&dest_addr=" + target + "&msg_text=" + url.QueryEscape(SMSText) + "%C2%AEistered_delivery=0&ston=5&snpi=0&dton=1&dnpi=1&encoding=1"
// 	url := "http://" + Configuration.SMPP_HTTP.IP + ":" + Configuration.SMPP_HTTP.Port + "/?systemid=" + Configuration.SMPP.Login + "&password=" + Configuration.SMPP.Password + "&Originator=" + Sender + "&dest_addr=" + target + "&msg_text=" + url.QueryEscape(SMSText) + "&ston=5&snpi=0&dton=1&dnpi=1&encoding=1"
// 	method := "GET"
// 	req, err := http.NewRequest(method, url, nil)
// 	if err != nil {
// 		log.Println("Error sending SMS: ", err)
// 		return err
// 	}
// 	//client := &http.Client{}
// 	client := &http.Client{
// 		Timeout: 15 * time.Second, //is SMSC not reachable request will time out after 15 sec
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println("Error sending SMS: ", err)
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode == 200 {
// 		return nil
// 	} else {
// 		err := errors.New("error sending SMS")
// 		log.Println("Error sending SMS: ", err)
// 		return err
// 	}
// }

// SendEmail : Generic Send Email using smtp
//
//	Usage Example:
//		SendEmail("vantage@africell.com", []string{"gkfoury@africell.com"}, []string{}, []string{}, "Test Email subject",
//			"Hello <b> George</b> <br> This is an email with attachement", []string{"/home/jouj/Pictures/Screenshot_20200824_154032.png"}, []string{})
//
//		SendEmail("vantage@africell.com", []string{"gkfoury@africell.com"}, []string{}, []string{}, "Test Email subject",
//			"Hello <b> George</b> <br> This is an email with Inline attachement<br> <img src=\"cid:Screenshot_20200824_154032.png\"/>", []string{}, []string{"/home/jouj/Pictures/Screenshot_20200824_154032.png"})
// func SendEmail(from string, to []string, cc []string, bcc []string, subject string, htmlBody string, attachments []string, inlineImgs []string) (err error) {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", from)
// 	m.SetHeader("To", to...)
// 	if len(cc) > 0 {
// 		m.SetHeader("Cc", cc...)
// 	}
// 	if len(bcc) > 0 {
// 		m.SetHeader("Bcc", bcc...)
// 	}
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", htmlBody)
// 	if len(attachments) > 0 {
// 		for _, s := range attachments {
// 			m.Attach(s)
// 		}
// 	}
// 	if len(inlineImgs) > 0 {
// 		for _, s := range inlineImgs {
// 			m.Embed(s)
// 		}
// 	}
// 	d := gomail.NewDialer(Configuration.SMTP.IP, Configuration.SMTP.Port, Configuration.SMTP.Login, Configuration.SMTP.Password)

// 	if err := d.DialAndSend(m); err != nil {
// 		log.Println(err)
// 	}
// 	return
// }
