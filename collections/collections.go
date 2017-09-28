package collections

type Collection interface {
	Add(interface{}) (bool, error)
	AddAll(Collection) (bool, error)
	Clear()
	Contains(interface{}) (bool, error)
	Remove(interface{}) (bool, error)
	Size() uint
	String() string
}
