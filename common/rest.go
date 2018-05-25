package common

import "github.com/bitly/go-simplejson"

func BuildCommonJson() *simplejson.Json {
	js := simplejson.New()
	js.Set("version", 1)
	js.Set("success", true)
	return js
}

