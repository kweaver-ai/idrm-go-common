package v1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// WhereSQL 生成 Where 子句
func (d *Detail) WhereSQL() (sql string, err error) {
	if d.RowFilters == nil {
		return
	}
	if len(d.RowFilters.Where) == 0 {
		return
	}

	whereArgs := make([]string, 0, len(d.RowFilters.Where))
	for i := range d.RowFilters.Where {
		var wherePreGroupFmt string
		for j := range d.RowFilters.Where[i].Member {
			var opAndValueSQL string
			opAndValueSQL, err = whereOPAndValueFmt(
				escape(d.RowFilters.Where[i].Member[j].NameEn),
				d.RowFilters.Where[i].Member[j].Operator,
				d.RowFilters.Where[i].Member[j].Value,
				d.RowFilters.Where[i].Member[j].DataType)
			if err != nil {
				return
			}
			if wherePreGroupFmt != "" {
				wherePreGroupFmt = wherePreGroupFmt + " " + string(d.RowFilters.Where[i].Relation) + " " + opAndValueSQL
			} else {
				wherePreGroupFmt = opAndValueSQL
			}
		}
		wherePreGroupFmt = "(" + wherePreGroupFmt + ")"
		whereArgs = append(whereArgs, wherePreGroupFmt)
	}
	var whereRelation string
	if d.RowFilters.WhereRelation != "" {
		whereRelation = fmt.Sprintf(` %s `, d.RowFilters.WhereRelation)
	} else {
		whereRelation = " AND "
	}
	sql = strings.Join(whereArgs, whereRelation)
	return
}

func whereOPAndValueFmt(name string, op Operator, value string, dataType DataType) (whereBackendSql string, err error) {
	special := strings.NewReplacer(`\`, `\\\\`, `'`, `\'`, `%`, `\%`, `_`, `\_`)
	switch op {
	case "<", "<=", ">", ">=":
		if _, err = strconv.ParseFloat(value, 64); err != nil {
			return whereBackendSql, errors.New("where conf invalid")
		}
		whereBackendSql = fmt.Sprintf("%s %s %s", name, op, value)
	case "=", "<>":
		if dataType == DataTypeNumber {
			if _, err = strconv.ParseFloat(value, 64); err != nil {
				return whereBackendSql, errors.New("where conf invalid")
			}
			whereBackendSql = fmt.Sprintf("%s %s %s", name, op, value)
		} else if dataType == DataTypeChar {
			whereBackendSql = fmt.Sprintf("%s %s '%s'", name, op, value)
		} else {
			return "", errors.New("523 where op not allowed")
		}
	case "null":
		whereBackendSql = fmt.Sprintf("%s IS NULL", name)
	case "not null":
		whereBackendSql = fmt.Sprintf("%s IS NOT NULL", name)
	case "include":
		if dataType == DataTypeChar {
			value = special.Replace(value)
			whereBackendSql = fmt.Sprintf("%s LIKE '%s'", name, "%"+value+"%")
		} else {
			return "", errors.New("534 where op not allowed")
		}
	case "not include":
		if dataType == DataTypeChar {
			value = special.Replace(value)
			whereBackendSql = fmt.Sprintf("%s NOT LIKE '%s'", name, "%"+value+"%")
		} else {
			return "", errors.New("541 where op not allowed")
		}
	case "prefix":
		if dataType == DataTypeChar {
			value = special.Replace(value)
			whereBackendSql = fmt.Sprintf("%s LIKE '%s'", name, value+"%")
		} else {
			return "", errors.New("548 where op not allowed")
		}
	case "not prefix":
		if dataType == DataTypeChar {
			value = special.Replace(value)
			whereBackendSql = fmt.Sprintf("%s NOT LIKE '%s'", name, value+"%")
		} else {
			return "", errors.New("555 where op not allowed")
		}
	case "in list":
		valueList := strings.Split(value, ",")
		for i := range valueList {
			if dataType == DataTypeChar {
				valueList[i] = "'" + valueList[i] + "'"
			}
		}
		value = strings.Join(valueList, ",")
		whereBackendSql = fmt.Sprintf("%s IN %s", name, "("+value+")")
	case "belong":
		valueList := strings.Split(value, ",")
		for i := range valueList {
			if dataType == DataTypeChar {
				valueList[i] = "'" + valueList[i] + "'"
			}
		}
		value = strings.Join(valueList, ",")
		whereBackendSql = fmt.Sprintf("%s IN %s", name, "("+value+")")
	case "true":
		whereBackendSql = fmt.Sprintf("%s = true", name)
	case "false":
		whereBackendSql = fmt.Sprintf("%s = false", name)
	case "before":
		valueList := strings.Split(value, " ")
		whereBackendSql = fmt.Sprintf(`%s >= DATE_add('%s', -%s, CURRENT_TIMESTAMP AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Shanghai') AND %s <= CURRENT_TIMESTAMP AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Shanghai'`, name, valueList[1], valueList[0], name)
	case "current":
		if value == "%Y" || value == "%Y-%m" || value == "%Y-%m-%d" || value == "%Y-%m-%d %H" || value == "%Y-%m-%d %H:%i" || value == "%x-%v" {
			whereBackendSql = fmt.Sprintf("DATE_FORMAT(%s, '%s') = DATE_FORMAT(CURRENT_TIMESTAMP AT TIME ZONE 'UTC' AT TIME ZONE 'Asia/Shanghai', '%s')", name, value, value)
		} else {
			return "", errors.New("586 where op not allowed")
		}
	case "between":
		valueList := strings.Split(value, ",")
		whereBackendSql = fmt.Sprintf("%s BETWEEN DATE_TRUNC('minute', CAST('%s' AS TIMESTAMP)) AND DATE_TRUNC('minute', CAST('%s' AS TIMESTAMP))", name, valueList[0], valueList[1])
	default:
		return "", errors.New("592 where op not allowed")
	}
	return
}

// quote 转义字段名称
func escape(s string) string {
	s = strings.Replace(s, "\"", "\"\"", -1)
	// 虚拟化引擎要求字段名称使用英文双引号 "" 转义，避免与关键字冲突
	s = fmt.Sprintf(`"%s"`, s)
	return s
}
