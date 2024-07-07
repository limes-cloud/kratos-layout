package conf

type AuthInfo struct {
	AppId     string
	AppSecret string
}

type Config struct {
	DefaultUserAvatar string
	Auth              struct {
		YiBan AuthInfo
	}
}
