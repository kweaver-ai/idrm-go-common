package iter

//常用的迭代方法
//减少日常的for循环代码的书写量

import "fmt"

func isEmpty[R any](r R) bool {
	var empty R
	return fmt.Sprintf("%v", r) == fmt.Sprintf("%v", empty)
}

// Gen 从数组ds中遍历，通过方法，返回每个符合条件的值，再组成数组[]R返回
// 如果不想某个值返回，那就直接返回R的空值形式，空值形式如下：
// var s string
// var s int
// var s map[string]string
// 等等
func Gen[R any, D any](ds []D, f func(D) R) []R {
	rs := make([]R, 0)
	for _, d := range ds {
		if r := f(d); !isEmpty(r) {
			rs = append(rs, r)
		}
	}
	return rs
}

func defaultKeyFunc[T any](t T) string {
	return fmt.Sprintf("%v", t)
}

// Unique 切片去重, fs是key的方法，不填默认是字符串方法，方便调用
func Unique[T any](ds []T, fs ...func(T) string) []T {
	var set = map[string]bool{}
	var res = make([]T, 0, len(ds))

	kf := defaultKeyFunc[T]
	if len(fs) > 0 {
		kf = fs[0]
	}

	for _, v := range ds {
		key := kf(v)
		if !set[key] && key != "" {
			res = append(res, v)
			set[key] = true
		}
	}
	return res
}

// GenMap 将数组ts，转换为，map[string]T的map，Key是方法
func GenMap[D, S any](ds []S, key func(S) (string, D)) map[string]D {
	r := make(map[string]D)
	for i := range ds {
		k, v := key(ds[i])
		if k != "" {
			r[k] = v
		}
	}
	return r
}

// StringMap 将数组ts，转换为，map[string]T的map，Key是方法
func StringMap[T any](ts []T, key func(T) string) map[string]T {
	r := make(map[string]T)
	for i := range ts {
		if k := key(ts[i]); k != "" {
			r[k] = ts[i]
		}
	}
	return r
}

// Tree 根据结果生成树状结构
// ds: 源数组
// id: 用于在数组中查找某个元素的方法
// pid: 获取父节点的方法
// append: 追加方法，如何节点追加到树的方法
func Tree[D any](ds []D, id func(D) string, pid func(D) string, append func(D, D)) D {
	dict := StringMap[D](ds, id)
	var root D
	for i := range ds {
		parentID := pid(ds[i])
		//获取root节点
		if parentID == "" {
			root = ds[i]
			continue
		}
		parent, ok := dict[parentID]
		if ok {
			append(parent, ds[i])
		} else if isEmpty(root) {
			root = ds[i]
		}
	}
	return root
}
