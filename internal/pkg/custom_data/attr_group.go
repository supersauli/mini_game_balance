package custom_data

import "encoding/json"

type AttrGroup struct {
	prop     map[string]*Value
	notice   []Notice
	isNotice bool
}

func (g *AttrGroup) Load(data []byte) error {
	var valList []*Value
	err := json.Unmarshal(data, &valList)
	if err != nil {
		return err
	}
	return nil
}

func (g *AttrGroup) Serialize() (data []byte, err error) {

	var valList []*Value
	for _, v := range g.prop {
		valList = append(valList, v)
	}

	data, err = json.Marshal(valList)
	if err != nil {
		return
	}
	return
}

func (g *AttrGroup) AddNotice(notice Notice) {
	g.isNotice = true
	g.notice = append(g.notice, notice)
}

func (g *AttrGroup) SetStrVal(name string, val string) {
	//
	if v, ok := g.prop[name]; ok {
		if g.isNotice {
			var oldVal Value
			oldVal.Copy(v)
			v.SetStrVal(val)
			var newVal Value
			newVal.Copy(v)
			for _, v := range g.notice {
				v.send(&oldVal, &newVal)
			}
		} else {
			v.SetStrVal(val)
		}
	}
}

func (g *AttrGroup) GetStrVal(name string) string {
	if v, ok := g.prop[name]; ok {
		return v.GetStrVal()
	}
	return ""
}

func (g *AttrGroup) SetIntVal(name string, val int) {
	if v, ok := g.prop[name]; ok {
		if g.isNotice {
			var oldVal Value
			oldVal.Copy(v)
			v.SetIntVal(val)
			var newVal Value
			newVal.Copy(v)
			for _, v := range g.notice {
				v.send(&oldVal, &newVal)
			}
		} else {
			v.SetIntVal(val)
		}

	}
}

func (g *AttrGroup) GetIntVal(name string) int {
	if v, ok := g.prop[name]; ok {
		return v.GetIntVal()
	}
	return 0
}

func (g *AttrGroup) IncIntVal(name string, val int) {
	if v, ok := g.prop[name]; ok {
		if g.isNotice {
			var oldVal Value
			oldVal.Copy(v)
			v.IncIntVal(val)
			var newVal Value
			newVal.Copy(v)
			for _, v := range g.notice {
				v.send(&oldVal, &newVal)
			}
		} else {
			v.IncIntVal(val)
		}
	}
}

func (g *AttrGroup) SetInt64Val(name string, val int64) {
	if v, ok := g.prop[name]; ok {
		if g.isNotice {
			var oldVal Value
			oldVal.Copy(v)
			v.SetInt64Val(val)
			var newVal Value
			newVal.Copy(v)
			for _, v := range g.notice {
				v.send(&oldVal, &newVal)
			}
		} else {
			v.SetInt64Val(val)
		}

	}
}

func (g *AttrGroup) GetInt64Val(name string) int64 {
	if v, ok := g.prop[name]; ok {
		return v.GetInt64Val()
	}
	return 0

}

func (g *AttrGroup) IncInt64Val(name string, val int64) {
	if v, ok := g.prop[name]; ok {
		if g.isNotice {
			var oldVal Value
			oldVal.Copy(v)
			v.IncInt64Val(val)
			var newVal Value
			newVal.Copy(v)
			for _, v := range g.notice {
				v.send(&oldVal, &newVal)
			}
		} else {
			v.IncInt64Val(val)

		}
	}
}

func (g *AttrGroup) SetFloat32Val(name string, val float32) {
	if v, ok := g.prop[name]; ok {
		if g.isNotice {
			var oldVal Value
			oldVal.Copy(v)
			v.SetFloat32Val(val)
			var newVal Value
			newVal.Copy(v)
			for _, v := range g.notice {
				v.send(&oldVal, &newVal)
			}
		} else {
			v.SetFloat32Val(val)
		}
	}
}

func (g *AttrGroup) GetFloat32Val(name string) float32 {
	if v, ok := g.prop[name]; ok {
		return v.GetFloat32Val()
	}
	return 0
}

func (g *AttrGroup) SetFloat64Val(name string, val float64) {
	if v, ok := g.prop[name]; ok {
		if g.isNotice {
			var oldVal Value
			oldVal.Copy(v)
			v.SetFloat64Val(val)
			var newVal Value
			newVal.Copy(v)
			for _, v := range g.notice {
				v.send(&oldVal, &newVal)
			}
		} else {
			v.SetFloat64Val(val)
		}
	}
}

func (g *AttrGroup) GetFloat64Val(name string) float64 {
	if v, ok := g.prop[name]; ok {
		return v.GetFloat64Val()
	}
	return 0
}
