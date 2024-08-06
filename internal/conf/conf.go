package conf

type Config struct {
	AiBot struct {
		Secret  string
		Appid   string
		Setting struct {
			Name    string
			Logo    string
			Desc    string
			Guiding []struct {
				Text string
				Type string
			}
		}
	}
}
