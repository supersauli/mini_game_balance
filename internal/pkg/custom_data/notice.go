package custom_data

type noticeCBFunc func(*Value,*Value)
type Notice struct {

	f []noticeCBFunc
}

func (n *Notice) send(old,new *Value) {
 for _,f := range n.f{
	 f(old,new)
 }

}
func (n*Notice) AddCB(f noticeCBFunc){
	n.f = append(n.f,f)
}
