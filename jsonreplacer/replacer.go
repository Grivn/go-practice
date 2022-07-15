package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func NewReplacer(replaceFunc func(interface{}) interface{}, options ...Option) IReplacer {
	r := new(replacer)
	for _, v := range options {
		v.Apply(r)
	}
	r.replacer = replaceFunc
	return r
}

type Option interface {
	Apply(*replacer)
}

type OptionFunc func(*replacer)

func (f OptionFunc) Apply(r *replacer) {
	f(r)
}

func WithMaxDepth(n int) Option {
	return OptionFunc(func(r *replacer) {
		r.maxDepth = n
	})
}

func WithKind(k reflect.Kind) Option {
	return OptionFunc(func(r *replacer) {
		r.kind = k
	})
}

func WithKeyword(k string) Option {
	return OptionFunc(func(r *replacer) {
		if r.keyword == nil {
			r.keyword = make([]string, 0, 1)
		}
		r.keyword = append(r.keyword, k)
	})

}

func WithKeywords(k []string) Option {
	return OptionFunc(func(r *replacer) {
		if r.keyword == nil {
			r.keyword = make([]string, 0, 1)
		}
		r.keyword = append(r.keyword, k...)
	})
}

type IReplacer interface {
	Replacing([]byte) ([]byte, error)
}

type replacer struct {
	maxDepth int                           // 最大递归替换深度，如果为0，表示不限制递归深度
	kind     reflect.Kind                  // 匹配的字段类型，如果未reflect.Invalid，表示不关心类型
	keyword  []string                      // 匹配的Key名，如果未指定即为nil，则仅做类型匹配
	replacer func(interface{}) interface{} // 替换器

	currDepth int // 当前递归深度
}

func (r replacer) Replacing(content []byte) ([]byte, error) {
	// 校验是否合法json
	if !json.Valid(content) {
		return nil, errors.New("invalid json")
	}

	// 反序列化json，兼容root为对象或数组的场景
	var holder interface{}
	if content[0] == '{' {
		holder = map[string]interface{}(nil)
	} else if content[0] == '[' {
		holder = []interface{}(nil)
	}

	err := json.Unmarshal(content, &holder)
	if err != nil {
		return nil, err
	}

	// 遍历/替换过程
	return json.Marshal(r.replacingElements("", holder))
}

func (r replacer) deepEnough() bool {
	// 如果maxDepth为0，表示不限制递归深度，总返回false
	// 如果maxDepth大于0，且当前深度没有达到maxDepth，则返回false
	// 否则返回true
	return r.maxDepth != 0 && r.currDepth > r.maxDepth
}

func (r replacer) needReplace(key string) bool {
	// 如果使用者没有指定keyword，那么总是返回true，表示不需要区分key
	if r.keyword == nil || len(r.keyword) == 0 {
		return true
	}
	// 查看当前key是否是使用者指定的key集合中的一个
	for _, v := range r.keyword {
		if v == key {
			return true
		}
	}
	return false
}

func (r replacer) convertArbitraryNumber(n interface{}, expect reflect.Kind) interface{} {
	// 获取以float64表达的json Number类型值
	rn, ok := n.(float64)
	if !ok {
		return n
	}

	// 将以float64表达的Number转换为目标类型，目标类型由使用者指定
	switch expect {
	case reflect.Float32:
		return float32(rn) // precision lost, cut off
	case reflect.Int:
		return int(rn) // precision lost
	case reflect.Int8:
		return int8(rn) // cut off
	case reflect.Int16:
		return int16(rn) // cut off
	case reflect.Int32:
		return int32(rn) // cut off
	case reflect.Int64:
		return int64(rn) // precision lost
	case reflect.Uint:
		return uint(rn) // cut off, sign lost
	case reflect.Uint8:
		return uint8(rn) // cut off, sign lost
	case reflect.Uint16:
		return uint16(rn) // cut off, sign lost
	case reflect.Uint32:
		return uint32(rn) // cut off, sign lost
	case reflect.Uint64:
		return uint64(rn) // precision lost, sign lost
	}
	return rn
}

func (r replacer) replacingElements(key string, value interface{}) interface{} {
	// 递归深度判断
	fmt.Println("-----")
	fmt.Println("key ", key)
	fmt.Println("val ", value)
	fmt.Println("     ")
	if r.deepEnough() {
		return value
	}

	// 反射获取value的类型信息
	vt := reflect.TypeOf(value)
	k := vt.Kind()

	fmt.Println("value type: ", vt)
	fmt.Println("value kind: ", k)
	fmt.Println("r kind: ", r.kind)

	// 根据kind分别处理
	switch k {
	case r.kind: // 是期望的值类型
		if r.needReplace(key) {
			// 当前key是目标key集合中的内容
			return r.replacer(value)
		}
	case reflect.Slice:
		// 获得slice的value
		val := reflect.ValueOf(value)
		r.currDepth++ // increace depth
		for i := 0; i < val.Len(); i++ {
			// 遍历slice
			elem := val.Index(i)
			if elem.CanInterface() {
				// 将成员递归处理
				relem := r.replacingElements(key, elem.Interface())
				if elem.CanSet() {
					elem.Set(reflect.ValueOf(relem))
				}
			}
		}
	case reflect.Map:
		val, ok := value.(map[string]interface{})
		if ok {
			r.currDepth++ // increace depth
			for mkey, mval := range val {
				// 将成员递归处理
				val[mkey] = r.replacingElements(mkey, mval)
			}
		}
	default:
		if k == reflect.Float64 && (r.kind >= reflect.Int && r.kind <= reflect.Float32) {
			// for arbitrary numberic types, adjust underlying type to expected
			if r.needReplace(key) {
				// 这里调用convertArbitraryNumber对float64进行了转换，确保使用者在对value进行断言时
				// 可以获得自己在replacer.kind成员上指定的类型值
				return r.replacer(r.convertArbitraryNumber(value, r.kind))
			}
		} else if r.kind == reflect.Invalid {
			// 对于任意类型的值，不判断value的类型
			if r.needReplace(key) {
				return r.replacer(value)
			}
		}
	}
	return value
}
