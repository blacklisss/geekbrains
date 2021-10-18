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
		{"configuration.env", &Config{UseCache: true, Url: "http://yandex.ru"}, nil},
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

func TestLoad2(t *testing.T) {
	var err error
	var config *Config

	var tests = []struct {
		useCache    bool
		externalUrl string
		want        *Config
		error       error
	}{
		{true, "http://ya.ru", &Config{UseCache: true, Url: "http://ya.ru"}, nil},
		{false, "http://ya1.ru", &Config{UseCache: false, Url: "http://ya1.ru"}, nil},
		{false, "ya.ru", nil, fmt.Errorf("необходимо ввести валидный Url")},
	}

	for _, test := range tests {
		useCache = &test.useCache
		externalUrl = &test.externalUrl
		config, err = Load()
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
