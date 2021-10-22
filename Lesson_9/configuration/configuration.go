package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/namsral/flag"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
	"path/filepath"
)

type Config struct {
	UseCache bool   `json:"use_cache" yaml:"use_cache"`
	Url      string `json:"url" yaml:"url"`
}

var (
	useCache    = flag.Bool("use-cache", true, "Использовать кэш (true/false)?")
	externalUrl = flag.String("external-url", "", "Внешний URL в полном формате")
)

func Load(configFile string) (config *Config, error error) {
	config = &Config{}
	switch filepath.Ext(configFile) {
	case ".json":
		if error = LoadJsonConfig(&configFile, config); error != nil {
			return
		}
	case ".yaml":
		if error = LoadYamlConfig(&configFile, config); error != nil {
			return
		}
	case ".env":
		if error = LoadEnvConfig(&configFile, config); error != nil {
			return
		}
	}

	if len(os.Args) > 1 {
		flag.Parse()
		if !validateUrl(externalUrl) {
			return nil, fmt.Errorf("необходимо ввести валидный Url")
		}
		config.UseCache = *useCache
		if *externalUrl != "" {
			config.Url = *externalUrl
		}
	}

	return
}

func LoadJsonConfig(configFile *string, config *Config) error {
	contents, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Error read config file: %s\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(contents, config)
	if err != nil {
		return fmt.Errorf("invalid json: %s\n", err)
	}

	return nil
}

func LoadYamlConfig(configFile *string, config *Config) error {
	contents, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Error read config file: %s\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(contents, config)
	if err != nil {
		return fmt.Errorf("invalid yaml: %s\n", err)
	}

	return nil
}

func LoadEnvConfig(configFile *string, config *Config) error {
	if err := godotenv.Load(*configFile); err != nil {
		return fmt.Errorf("произошла ошибка парсинга файла окружения")
	}

	if !validateUrl(externalUrl) {
		return fmt.Errorf("необходимо ввести валидный Url")
	}

	flag.Parse()

	config.UseCache = *useCache
	config.Url = *externalUrl

	return nil
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateUrl(url *string) bool {
	if *externalUrl != "" && !isUrl(*externalUrl) {
		return false
	}

	return true
}
