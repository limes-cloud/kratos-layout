package conf

type AuthInfo struct {
	AppId     string
	AppSecret string
}

type Config struct {
	Auth struct {
		YiBan AuthInfo
	}
}
