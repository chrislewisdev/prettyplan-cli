package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"
)

func Prettify(value string) template.HTML {
	value = unescapeCharacters(value)

	if value == "<computed>" {
		return template.HTML("<em>&lt;computed&gt;</em>")
	} else if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
		return template.HTML("<em>" + value + "</em")
	} else if _, notNumber := strconv.ParseFloat(value, 32); json.Valid([]byte(value)) && notNumber != nil {
		var prettyJson bytes.Buffer
		json.Indent(&prettyJson, []byte(value), "", "  ")
		return template.HTML(fmt.Sprintf("<pre>%v</pre>", prettyJson.String()))
	} else {
		return template.HTML(value)
	}
}

func unescapeCharacters(value string) string {
	value = strings.Replace(value, `\n`, "\n", -1)
	value = strings.Replace(value, `\"`, "\"", -1)

	//This feels dumb - but, the above unescaping might unescape some \" sequences that need to be left alone
	//e.g. terraform properties inside json strings like ${terraform_property["index"]}
	//So this regex will let us re-escape any quotes inside a terraform property
	r := regexp.MustCompile(`(\${[^}]*?[^\\}])"([^}]*?})`)
	for r.MatchString(value) {
		value = r.ReplaceAllString(value, `$1\"$2`)
	}

	return value
}
