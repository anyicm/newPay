# newPay

newPay 是一個将useshop 扩展调用第三方支付defipay 的付款 回调的 gin框架 编写的中转程序

defipay 文档：`https://defipay.stoplight.io/docs/defipay-api-english-docs/n0ygelefq08x1-`
ueeshop 文档：`https://www.showdoc.com.cn/p/106d4c87bcb8a53d8490052b83b9d48d`

## 用法

### 所需要的环境 go 1.8+

#### 初始化

```
go mod tidy
```

#### 结构

```
conf
  -- conf.yml [配置文件]
defipay [官方包,可以去掉，使用官方] 
# 在orderController import "newPay/defipay" 替换为
# import "github.com/defipay/defipay-go-api/defipay"

serve
  -- orderController [主逻辑]
  -- serve  [调用gin框架路由注册]
useshop
  api-params  [独立站 请求数据结构]
  order [独立站签名，解签所需方法]
build.bat [windows 下打包后 多平台运行]
build.sh  [liunx 下打包后 多平台运行]
go.mod [所需依赖包]
main  [程序入口]
```

#### 转接流程

`/v1/order/create ` 客户端创建订单

```
用户下单 =》 独立站  =》 newPay =》 请求defiPay 检测开放的钱币种类，创建订单
```

`/v1/order/notify` defiPay 异步回调

```
defiPay请求  =》 newPay =》 请求独立站，返回成功和失败
```

