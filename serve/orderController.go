package serve

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"newPay/defipay"
	"newPay/useshop"
	"strconv"
	"time"
)

type Records struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	DisplayName        string `json:"displayName"`
	ShortName          string `json:"shortName"`
	LogoUrl            string `json:"logoUrl"`
	ChainAssertId      string `json:"chainAssertId"`
	ChainAssertDecimal string `json:"chainAssertDecimal"`
}

type OrderController struct {
}

func QueryCurreny(orderCurrency string) (err error, r *Records) {
	client := GetClient()
	query, apiError := client.TokenQuery("1", "10")
	if apiError != nil {
		panic(apiError.Message)
	}

	mJson, err := query.Array()
	if err != nil {
		fmt.Errorf("mjson err is %s", err.Error())
		return
	}
	currencySlice := make([]*Records, 0)
	for i, _ := range mJson {
		currencyItem := new(Records)
		item := query.GetIndex(i)
		itemJson, _ := item.Encode()
		json.Unmarshal(itemJson, currencyItem)
		currencySlice = append(currencySlice, currencyItem)
	}

	for _, item := range currencySlice {
		if item.Name == orderCurrency {
			r = item
			break
		}
	}
	if r == nil {
		err = fmt.Errorf("order currency nonexistent")
		return
	}
	return
}

func GetClient() defipay.Client {
	var client = defipay.Client{
		Signer: defipay.LocalSigner{
			PrivateKey: GetConf(ThreePriKey),
		},
		Env:   defipay.Sandbox(),
		Debug: false,
	}
	return client
}

func RequestDefiPay(notifyUrl string, returnUrl string, memberTransNo string, amount string, currency string, tokenIds string) (res *simplejson.Json, err *defipay.ApiError) {
	client := GetClient()
	res, err = client.CreateOrder(notifyUrl, returnUrl, memberTransNo, amount, currency, tokenIds)
	if err != nil {
		return
	}
	return
}

func (ctl OrderController) Create(c *gin.Context) {
	//解密
	client := c.PostForm("clientId")
	orderData := c.PostForm("orderData")
	signature := c.PostForm("signature")

	key := GetConf(ClientKey)
	md5key := useshop.GetSha512(key)

	err, res := useshop.StrDecrypt(orderData, md5key)
	if err != nil {
		fmt.Errorf("err is %s", err.Error())
		return
	}
	out, err := simplejson.NewJson([]byte(res))
	if err != nil {
		fmt.Errorf("newJson err is %s", err.Error())
		return
	}

	apiOrderData := new(useshop.OrderData)
	apiOrderData.OrderNo, _ = out.Get("orderNo").String()
	apiOrderData.OrderCurrency, _ = out.Get("orderCurrency").String()
	apiOrderData.OrderAmount, _ = out.Get("orderAmount").String()
	apiOrderData.FirstName, _ = out.Get("firstName").String()
	apiOrderData.LastName, _ = out.Get("lastName").String()
	apiOrderData.Email, _ = out.Get("email").String()
	apiOrderData.Phone, _ = out.Get("phone").String()
	apiOrderData.Country, _ = out.Get("country").String()
	apiOrderData.State, _ = out.Get("state").String()
	apiOrderData.City, _ = out.Get("city").String()
	apiOrderData.Address, _ = out.Get("address").String()
	apiOrderData.Zip, _ = out.Get("zip").String()
	apiOrderData.IsMobile, _ = out.Get("isMobile").Int()
	apiOrderData.Ip, _ = out.Get("ip").String()
	apiOrderData.Timestamp, _ = out.Get("timestamp").String()
	apiOrderData.UrlCallback, _ = out.Get("urlCallback").String()
	apiOrderData.UrlComplete, _ = out.Get("urlComplete").String()

	goodsDetails, _ := out.Get("goodsDetail").Array()
	currencySlice := make([]*useshop.GoodsDeatail, 0)
	for i, _ := range goodsDetails {
		currencyItem := new(useshop.GoodsDeatail)
		item := out.Get("goodsDetail").GetIndex(i)
		itemJson, _ := item.Encode()
		json.Unmarshal(itemJson, currencyItem)
		currencySlice = append(currencySlice, currencyItem)
	}
	apiOrderData.GoodsDetail = currencySlice

	orderTotal, _ := out.Get("orderTotal").Array()
	orderTotalSlice := make([]*useshop.OrderTotal, 0)
	for i, _ := range orderTotal {
		orderTotalItem := new(useshop.OrderTotal)
		item := out.Get("orderTotal").GetIndex(i)
		itemJson, _ := item.Encode()
		json.Unmarshal(itemJson, orderTotalItem)
		orderTotalSlice = append(orderTotalSlice, orderTotalItem)
	}
	apiOrderData.OrderTotal = orderTotalSlice

	//判断币种
	err, r := QueryCurreny(apiOrderData.OrderCurrency)
	if err != nil {
		fmt.Errorf("unmarshal err is %s", err.Error())
		return
	}
	//精度
	rate, err := strconv.Atoi(r.ChainAssertDecimal)
	if err != nil {
		fmt.Errorf("strconv err is %s,%d", err.Error(), rate)
		return
	}
	//amount := strconv.FormatFloat(float64(apiOrderData.OrderAmount), 'f', rate, 32)
	amount := apiOrderData.OrderAmount
	amountToken := r.Id
	mapJson := useshop.StructToMapJson(apiOrderData)
	mapJson["clientId"] = client
	sortStr, err := useshop.MapKeySort(mapJson)
	if err != nil {
		fmt.Errorf("sort Map err is %s", err.Error())
		return
	}
	designa := useshop.GetSha256(sortStr, key)
	if designa != signature {
		fmt.Println("signature fail", client, orderData, signature)
		return
	}
	//请求第三方支付
	rsp, errJson := RequestDefiPay(GetConf(NotifyUrl), GetConf(ReturnUrl), apiOrderData.OrderNo, amount, apiOrderData.OrderCurrency, strconv.Itoa(amountToken))
	if errJson != nil || !errJson.Success {
		fmt.Errorf("requestDefiPay fail %s", errJson.Message)
		return
	}
	fmt.Println(rsp)
	transNoJson := rsp.Get("transNo")
	transInfo, err := transNoJson.Encode()
	if err != nil {
		fmt.Errorf("transNo fail %s", errJson.Message)
		return
	}
	c.String(200, string(transInfo))
}

func DispathReturn(callbackReturn *useshop.CallbackReturnParams, signature string) (result string, err error) {
	rqs := url.Values{
		"clientId":  {callbackReturn.ClientId},
		"orderNo":   {callbackReturn.OrderNo},
		"reference": {callbackReturn.Reference},
		"result":    {callbackReturn.Result},
		"message":   {callbackReturn.Message},
		"timestamp": {callbackReturn.Timestamp},
		"signature": {signature},
	}

	rsp, err := http.PostForm(GetConf(LocalUrl), rqs)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	return string(body), nil
}

func queryOrder(transNo string) bool {
	client := GetClient()
	result, apiError := client.QueryOrder(transNo)
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return false
	}
	str, _ := result.Encode()
	fmt.Println("queryOrder：")
	fmt.Println(string(str))
	return true
}

func (ctl OrderController) DispatchReturn(c *gin.Context) {
	//验签
	stime := c.GetHeader("BIZ-TIMESTAMP")
	signture := c.GetHeader("BIZ-RESP-SIGNATURE")
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Errorf("err is %s", err.Error())
		return
	}
	content := string(body) + "|" + stime
	client := GetClient()
	if client.VerifyEcc(content, signture) {
		json, _ := simplejson.NewJson(body)
		state, _ := json.Get("state").Int()
		if state == 200 {
			transNo, _ := json.Get("transNo").String()
			if queryOrder(transNo) {
				//异步调用独立站
				memberTransNo, _ := json.Get("memberTransNo").String()
				//回调
				getReturn := &useshop.CallbackReturnParams{
					ClientId:  GetConf(ClientId),
					OrderNo:   memberTransNo,
					Reference: transNo,
					Result:    "completed",
					Message:   "",
					Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
				}

				returnDs := useshop.StructToMapJson(getReturn)
				rSort, err := useshop.MapKeySort(returnDs)
				if err != nil {
					fmt.Println("Error Marshal >>>>>>>>")
					fmt.Errorf("err is %s", err.Error())
					return
				}
				signature := useshop.GetSha256(rSort, GetConf(ClientKey))
				getRes, err := DispathReturn(getReturn, signature)
				if err != nil {
					fmt.Errorf("sort Map err is %s", err.Error())
					return
				}
				log.Println("Notify return >>>>>>>>")
				log.Println(getRes)
			}
		}
	}
	c.String(200, "ok")
}
