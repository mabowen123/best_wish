package until

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func JoinDomain(base, sub string) string {
	// 如果sub已经有/，则不需要添加/
	if strings.HasPrefix(sub, "/") {
		return base + sub
	}
	// 如果sub没有/，则需要添加/
	return base + "/" + sub
}

func ToBeInt64(data interface{}) int64 {
	// 处理 ID 字段
	var idAsInt int
	switch v := data.(type) {
	case string:
		// 字符串类型，转换为整数
		var err error
		idAsInt, err = strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting ID to int:", err)
			return 0
		}
		return int64(idAsInt)
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case float64:
		return int64(v)
	case float32:
		return int64(v)
	case int64:
		return v
	}
	return 0
}

func ReplaceAllCharAndEmojiToBlank(s string, specialChars []string) string {
	for _, char := range specialChars {
		s = strings.ReplaceAll(s, char, "")
	}
	reg, _ := regexp.Compile("[\U0001F600-\U0001F64F\U0001F300-\U0001F5FF\U0001F680-\U0001F6FF\U0001F1E0-\U0001F1FF]+")
	s = reg.ReplaceAllString(s, "")
	return s
}
