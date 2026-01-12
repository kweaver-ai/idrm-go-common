package impl

import (
	"strconv"
	"strings"
)

// parsePrecision 解析虚拟化引擎返回的字段长度精度
func parsePrecision(oriType string) (*int, *int) {
	olen := len(oriType)
	start := strings.Index(oriType, "(")
	end := strings.Index(oriType, ")")
	if start == -1 || end == -1 || start >= olen || end >= olen {
		return nil, nil
	}
	pstr := oriType[start+1 : end]
	ps := strings.Split(pstr, ",")
	if len(ps) != 2 {
		length, _ := strconv.Atoi(pstr)
		if length > 0 {
			return &length, nil
		}
		return nil, nil
	}
	var lenPtr, prePtr *int
	length, err1 := strconv.Atoi(strings.TrimSpace(ps[0]))
	precision, err2 := strconv.Atoi(strings.TrimSpace(ps[1]))
	if err1 == nil {
		lenPtr = &length
	}
	if err2 == nil {
		prePtr = &precision
	}
	return lenPtr, prePtr
}
