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
	"strconv"
)

type Config struct {
	UseCache bool   `json:"use_cache" yaml:"use_cache"`
	Url      string `json:"url" yaml:"url"`
}

func Load(configFile string) (config *Config, err error) {
	config = NewConfig()
	switch filepath.Ext(configFile) {
	case ".json":
		if err = LoadJsonConfig(&configFile, config); err != nil {
			return
		}
	case ".yaml":
		if err = LoadYamlConfig(&configFile, config); err != nil {
			return
		}
	case ".env":
		if err = LoadEnvConfig(&configFile, config); err != nil {
			return
		}
	default:
		return nil, fmt.Errorf("invalid format of configuration file")
	}

	if len(os.Args) > 1 && filepath.Ext(os.Args[0]) != ".test" { // тут не нашел лучшего решения исключить тесты из вызова. Буду благодарен, если подскажете варианты
		if err = LoadFromArgs(config); err != nil {
			return
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

	if !validateUrl(&config.Url) {
		return fmt.Errorf("необходимо ввести валидный Url")
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

	if !validateUrl(&config.Url) {
		return fmt.Errorf("необходимо ввести валидный Url")
	}

	return nil
}

func LoadEnvConfig(configFile *string, config *Config) error {
	if err := godotenv.Load(*configFile); err != nil {
		return fmt.Errorf("произошла ошибка парсинга файла окружения")
	}

	useCacheTmp := os.Getenv("USE_CACHE")
	useCacheBool, err := strconv.ParseBool(useCacheTmp)
	if err != nil {
		return fmt.Errorf("failed convert to bool useCache variable")
	}
	config.UseCache = useCacheBool

	config.Url = os.Getenv("EXTERNAL_URL")
	if !validateUrl(&config.Url) {
		return fmt.Errorf("необходимо ввести валидный Url")
	}

	return nil
}

func LoadFromArgs(config *Config) error {
	var (
		useCache    = flag.Bool("use-cache", true, "Использовать кэш (true/false)?")
		externalUrl = flag.String("external-url", "", "Внешний URL в полном формате")
	)

	flag.Parse()

	config.UseCache = *useCache
	config.Url = *externalUrl

	if !validateUrl(&config.Url) {
		return fmt.Errorf("необходимо ввести валидный Url")
	}

	return nil
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateUrl(url *string) bool {
	if *url != "" && !isUrl(*url) {
		return false
	}

	return true
}

func NewConfig() *Config {
	return &Config{}
}
