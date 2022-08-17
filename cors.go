package goweb_cors

import (
	"github.com/pelletier/go-toml"
	"github.com/rs/cors"
	"log"
	"strconv"
)

func GetTree(tree *toml.Tree, key string) *toml.Tree {
	if value, ok := tree.Get(key).(*toml.Tree); ok {
		return value
	}
	return new(toml.Tree)
}

func GetStringArray(tree *toml.Tree, key string, values ...[]string) []string {
	strings := make([]string, 0)
	if array, ok := tree.Get(key).([]interface{}); ok {
		for _, value := range array {
			if str, ok := value.(string); ok {
				strings = append(strings, str)
			}
		}
	}
	if len(strings) == 0 && len(values) > 0 {
		return values[0]
	}
	return strings
}

func GetInt(tree *toml.Tree, key string, values ...int) int {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case int64:
			return int(value.(int64))
		case uint64:
			return int(value.(uint64))
		case float64:
			return int(value.(float64))
		case string:
			value, err := strconv.ParseInt(value.(string), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				return int(value)
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0
}

func GetBool(tree *toml.Tree, key string, values ...bool) bool {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case bool:
			return value.(bool)
		case string:
			value, err := strconv.ParseBool(value.(string))
			if err != nil {
				log.Println(err)
			} else {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return false
}

func IsLocal(Env string) bool {
	return Env == "local"
}

/**
* @func Cors
* @desc  Cors middleware
* @param config *toml.Tree ( get config from toml file, eg: config, err := toml.LoadFile(path) )
* @param Env string ( get env from toml file, eg: config = iris.TOML("./config/iris.toml")
config.Get("env").(string) )
* @return cors.Options
*/
func Cors(serverConfig *toml.Tree, EnvConfig string) *cors.Cors {
	config := GetTree(serverConfig, "cors")
	allowedOrigins := GetStringArray(config, "allowed_origins", []string{"*"})
	allowedMethods := GetStringArray(config, "allowed_methods", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := GetStringArray(config, "allowed_headers", []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"})
	exposedHeaders := GetStringArray(config, "exposed_headers", []string{"Content-Length", "Content-Encoding", "Content-Language", "Content-Type"})
	maxAge := GetInt(config, "max_age", 3600)
	allowCredentials := GetBool(config, "allow_credentials")
	optionsPassthrough := GetBool(config, "options_passthrough")
	debug := GetBool(config, "debug", IsLocal(EnvConfig))
	options := cors.Options{
		AllowedOrigins:     allowedOrigins,
		AllowedMethods:     allowedMethods,
		AllowedHeaders:     allowedHeaders,
		ExposedHeaders:     exposedHeaders,
		MaxAge:             maxAge,
		AllowCredentials:   allowCredentials,
		OptionsPassthrough: optionsPassthrough,
		Debug:              debug,
	}
	return cors.New(options)
}
