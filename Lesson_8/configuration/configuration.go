package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/namsral/flag"
	"net/url"
)

type Config struct {
	UseCache bool
	Url      string
}

var (
	useCache    = flag.Bool("use-cache", true, "Использовать кэш (true/false)?")
	externalUrl = flag.String("external-url", "", "Внешний URL в полном формате")
)

func Load(envfiles ...string) (config *Config, error error) {
	if envfiles != nil {
		if err := godotenv.Load(envfiles...); err != nil {
			return nil, fmt.Errorf("произошла ошибка парсинга файла окружения")
		}
	}

	flag.Parse()

	if *externalUrl != "" && !isUrl(*externalUrl) {
		return nil, fmt.Errorf("необходимо ввести валидный Url")
	}

	return &Config{UseCache: *useCache, Url: *externalUrl}, nil
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
