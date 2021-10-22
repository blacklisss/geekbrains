package configuration

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	var err error
	var config *Config

	var tests = []struct {
		envFile string
		want    *Config
		error   error
	}{
		{"config.json", &Config{UseCache: true, Url: "https://gb.ru"}, nil},
		{"config.yaml", &Config{UseCache: true, Url: "https://google.com"}, nil},
		{"config.env", &Config{UseCache: true, Url: "http://yandex.ru"}, nil},
		{"configuration1.env", nil, fmt.Errorf("произошла ошибка парсинга файла окружения")},
	}

	for _, test := range tests {
		config, err = Load(test.envFile)
		if err != nil {
			if err.Error() != test.error.Error() {
				t.Errorf("Ожидалось: «%v». Пришло: «%v»", test.error, err)
			}
		} else {
			if *config != *test.want {
				t.Errorf("Требуется: «%v». Пришло: «%v»", test.want, config)
			}
		}
	}
}
