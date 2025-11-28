package base

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
)

type H = map[string]any

var Config struct {
	Addr       string         `config:"addr"`
	Token      string         `config:"token"`
	Cookie     string         `config:"music.cookie"`
	NeteaseAPI string         `config:"music.netease"`
	QQAPI      string         `config:"music.qq"`
	Pgsql      string         `config:"pgsql"`
	Debug      bool           `config:"debug"`
	Persist    []PersistHouse `config:"persist"`
}

type PersistHouse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Password string `json:"password"`
}

func envKeyFromTag(tag string) string {
	// "music.cookie" => "ALISTEN_MUSIC_COOKIE"
	key := strings.ToUpper(strings.ReplaceAll(tag, ".", "_"))
	return "ALISTEN_" + key
}

func InitConfig() {
	var (
		v                     = reflect.ValueOf(&Config).Elem()
		t                     = v.Type()
		stringType            = reflect.TypeOf("")
		boolType              = reflect.TypeOf(true)
		slicePersistHouseType = reflect.TypeOf([]PersistHouse{})
	)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("config")
		if tag == "" {
			continue
		}

		envKey := envKeyFromTag(tag)
		raw := os.Getenv(envKey)
		if raw == "" {
			continue
		}

		switch field.Type {
		case stringType:
			v.Field(i).SetString(raw)

		case boolType:
			if raw == "true" || raw == "1" {
				v.Field(i).SetBool(true)
			} else {
				v.Field(i).SetBool(false)
			}

		case slicePersistHouseType:
			var houses []PersistHouse
			if err := json.Unmarshal([]byte(raw), &houses); err == nil {
				v.Field(i).Set(reflect.ValueOf(houses))
			}

		default:
			panic("unsupported type: " + field.Type.String())
		}
	}
}
