package useshop

type ApiParams struct {
	ClientId  string     `json:"clientId"`  //商户号
	OrderData *OrderData `json:"orderData"` //json格式订单数据
	//Signature string     `json:"signature"` //签名
}

type CallbackReturnParams struct {
	ClientId  string `json:"clientId"`  //商户号
	OrderNo   string `json:"orderNo"`   //交易订单号
	Reference string `json:"reference"` //交易流水号 (只属于此交易的唯一参考)
	Result    string `json:"result"`    //支付状态 (固定三个状态，completed：支付成功；pending：等待确认；denied：支付失败)
	Message   string `json:"message"`   //附加信息 (可以返回支付说明或者支付失败原因)
	Timestamp string `json:"timestamp"` //当前时间戳，如：1464693823 注意：需要string类型
}

type OrderData struct {
	OrderNo       string          `json:"orderNo"`       //交易订单号
	OrderCurrency string          `json:"orderCurrency"` //币种
	OrderAmount   string          `json:"orderAmount"`   //交易金额
	OrderTotal    []*OrderTotal   `json:"orderTotal"`    //价格明细
	FirstName     string          `json:"firstName"`     //收货地址-名
	LastName      string          `json:"lastName"`      //收货地址-姓
	Email         string          `json:"email"`         //收货地址-邮件地址
	Phone         string          `json:"phone"`         //收货地址-电话
	Country       string          `json:"country"`       //收货地址-国家，国家二位代码
	State         string          `json:"state"`         //收货地址-州/省
	City          string          `json:"city"`          //收货地址-城市
	Address       string          `json:"address"`       //收货地址-详细地址
	Zip           string          `json:"zip"`           //收货地址-邮政编码
	IsMobile      int             `json:"isMobile"`      //是否为手机版，1:手机版；0:PC版
	Ip            string          `json:"ip"`            //客户端ip
	Domain        string          `json:"domain"`        //发起支付的网站域名
	GoodsDetail   []*GoodsDeatail `json:"goodsDetail"`   //产品明细，最多传递20个产品纪录
	Timestamp     string          `json:"timestamp"`     //时间戳
	UrlCallback   string          `json:"urlCallback"`   //回调通知应异步发送到的 URL，固定值：http(s)://www.domain.com/payment/notify-url/ (请参考 3. 交易异步通知)
	UrlComplete   string          `json:"urlComplete"`   //成功完成付款流程后，客户必须重定向到的 URL，固定值：http(s)://www.domain.com/payment/return-url/ (请参考 4. 交易同步通知)

}

type OrderTotal struct {
	ItemTotal float32 `json:"itemTotal"` //产品价格
	Tax       float32 `json:"tax"`       //税费
	Shipping  float32 `json:"shipping"`  //运费
	Handling  float32 `json:"handling"`  //手续费
	Discount  float32 `json:"discount"`  //折扣
}

type GoodsDeatail struct {
	ProductName string  `json:"productName"` //产品名称
	Quantity    int     `json:"quantity"`    //产品数量
	Sku         string  `json:"sku"`         //产品SKU
	Price       float32 `json:"price"`       //产品单价
}
