package custom_data

import "go.uber.org/zap"

const (
	VALUE_TYPE_EMPTY = iota
	VALUE_TYPE_INT
	VALUE_TYPE_INT64
	VALUE_TYPE_STRING
	VALUE_TYPE_FLOAT32
	VALUE_TYPE_FLOAT64
)

type Value struct {
	Type int32       `json:"type"`
	Val  interface{} `json:"val"`
	Name string      `json:"name"`
	Desc string      `json:"desc,omitempty"`
}

func (v *Value) Copy(old *Value) {
	v.Val = old.Val
	v.Type = old.Type
	v.Name = old.Name
	v.Desc = old.Desc
}

func (v *Value) SetStrVal(val string) {
	if v.Type != VALUE_TYPE_EMPTY && v.Type != VALUE_TYPE_STRING {
		zap.L().Error("set string value type error", zap.String("name", v.Name), zap.Any("val", val), zap.Any("type", v.Type))
		return
	}
	v.Val = val
	v.Type = VALUE_TYPE_STRING
}
func (v *Value) GetStrVal() string {
	if v.Type != VALUE_TYPE_STRING {
		zap.L().Error("get string value type error", zap.String("name", v.Name), zap.Any("val", v.Val), zap.Any("type", v.Type))
		return ""
	}
	return v.Val.(string)
}

func (v *Value) SetIntVal(val int) {
	if v.Type != VALUE_TYPE_EMPTY && v.Type != VALUE_TYPE_INT {
		zap.L().Error("set int value type error", zap.String("name", v.Name), zap.Any("val", val), zap.Any("type", v.Type))
		return
	}
	v.Val = val
	v.Type = VALUE_TYPE_INT
}
func (v *Value) GetIntVal() int {
	if v.Type != VALUE_TYPE_INT {
		zap.L().Error("get int value type error", zap.String("name", v.Name), zap.Any("val", v.Val), zap.Any("type", v.Type))
		return 0
	}
	return v.Val.(int)
}

func (v *Value) IncIntVal(inc int) int {

	if v.Type != VALUE_TYPE_INT {
		zap.L().Error("inc int value type error", zap.String("name", v.Name), zap.Any("val", inc), zap.Any("type", v.Type))
		return 0
	}
	v.Val = v.Val.(int) + inc
	return v.Val.(int)
}

func (v *Value) SetInt64Val(val int64) {
	if v.Type != VALUE_TYPE_EMPTY && v.Type != VALUE_TYPE_INT64 {
		zap.L().Error("set int64 value type error", zap.String("name", v.Name), zap.Any("val", val), zap.Any("type", v.Type))
		return
	}
	v.Val = val
	v.Type = VALUE_TYPE_INT64
}

func (v *Value) GetInt64Val() int64 {
	if v.Type != VALUE_TYPE_INT64 {
		zap.L().Error("get int64 value type error", zap.String("name", v.Name), zap.Any("val", v.Val), zap.Any("type", v.Type))
		return 0
	}
	return v.Val.(int64)
}

func (v *Value) IncInt64Val(inc int64) int64 {

	if v.Type != VALUE_TYPE_INT64 {
		zap.L().Error("inc int64 value type error", zap.String("name", v.Name), zap.Any("val", inc), zap.Any("type", v.Type))
		return 0
	}

	v.Val = v.Val.(int64) + inc
	return v.Val.(int64)
}

func (v *Value) SetFloat32Val(val float32) {
	if v.Type != VALUE_TYPE_EMPTY && v.Type != VALUE_TYPE_FLOAT32 {
		zap.L().Error("set float32 value type error", zap.String("name", v.Name), zap.Any("val", val), zap.Any("type", v.Type))
		return
	}
	v.Val = val
	v.Type = VALUE_TYPE_FLOAT32
}

func (v *Value) GetFloat32Val() float32 {
	if v.Type != VALUE_TYPE_FLOAT32 {
		zap.L().Error("get float32 value type error", zap.String("name", v.Name), zap.Any("val", v.Val), zap.Any("type", v.Type))
		return 0
	}
	return v.Val.(float32)

}

func (v *Value) SetFloat64Val(val float64) {
	if v.Type != VALUE_TYPE_EMPTY && v.Type != VALUE_TYPE_FLOAT64 {
		zap.L().Error("set float64 value type error", zap.String("name", v.Name), zap.Any("val", val), zap.Any("type", v.Type))
		return

	}
	v.Val = val
	v.Type = VALUE_TYPE_FLOAT64
}

func (v *Value) GetFloat64Val() float64 {
	if v.Type != VALUE_TYPE_FLOAT64 {
		zap.L().Error("get float64 value type error", zap.String("name", v.Name), zap.Any("val", v.Val), zap.Any("type", v.Type))
		return 0
	}
	return v.Val.(float64)

}

func (v *Value) SetDesc(desc string) {
	v.Desc = desc
}
func (v *Value) GetType() int32 {
	return v.Type
}
