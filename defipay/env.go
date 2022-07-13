package defipay

type Env struct {
	Host   string
	PubKey string
}

func Sandbox() Env {
	return Env{Host: "https://api-test.defipay.biz/api-service", PubKey: "0347833d34c31dbe2f44e4300c05615b0ef5927eb16311ec85265fad47f3574f8f"}
	//return Env{Host: "http://api-test.defipay.biz/api-service", PubKey: "03412208e920ba78d97c33c1476db11506cb6d4b3fc218a8207ca523c8d392a3f0"}
	//return Env{Host: "http://api-test.defipay.biz/api-service", PubKey: "02a17fffb024cce6220ddf91b40711dc15fd8f830e23f6160c6a4eac8bc0eba820"}
}

func Prod() Env {
	return Env{Host: "https://api.defipay.biz/api-service", PubKey: "02adb46f0c10b5ec51d0df2183a812fdf7b330ef2c948e36cdb479f1af73a22753"}
}
