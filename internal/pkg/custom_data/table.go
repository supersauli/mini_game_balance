package custom_data

type Table struct {
	Name        string
	indexKey    map[string]struct{}
	indexUnique map[string]struct{}

	indexByKey    map[string][]*Row
	indexByUnique map[string]*Row
	Row           []*Row
	existUnique   bool
}

type Row struct {
	Attr     map[string]*Value
	UID      int64
	Activity bool
}

func (t *Table) AddRow(row *Row) {
	t.Row = append(t.Row, row)
}

func (t *Table) RemoveRow(row *Row) {
	//if t.existUnique {
	//	for _, v := range t.indexUnique {
	//		row.Attr[v]
	//			delete(t.indexByUnique, k)
	//		}
	//	}
	//}
	//for i, v := range t.Row {
	//	if v == row {
	//		t.Row = append(t.Row[:i], t.Row[i+1:]...)
	//	}
	//}
}

func (t *Table) FindRow(key string, val interface{}) *Row {

	return nil
}

func (t *Table) FindRows(key string, val interface{}) []*Row {

	return nil
}
