package multiplication

import (
	"errors"
	"fmt"
)

const (
	Int  string = "int"
	Doub string = "double"
)

type Number struct {
	t string
	v interface{}
}

func NewNumber(n interface{}) (Number, error) {
	switch n.(type) {
	case int:
		return Number{Int, int64(n.(int))}, nil
	case int8:
		return Number{Int, int64(n.(int8))}, nil
	case int16:
		return Number{Int, int64(n.(int16))}, nil
	case int32:
		return Number{Int, int64(n.(int32))}, nil
	case int64:
		return Number{Int, n}, nil
	case float32:
		return Number{Doub, float64(n.(float32))}, nil
	case float64:
		return Number{Doub, n}, nil
	default:
		return Number{}, errors.New("type doesn't currently supported")
	}
}

func (n *Number) Mul(num Number) (Number, error) {
	if n.t != num.t {
		return Number{}, errors.New("type mismatch")
	}
	switch n.t {
	case Int:
		f := n.v.(int64)
		s := num.v.(int64)
		return Number{
			t: Int,
			v: f * s,
		}, nil
	case Doub:
		f := n.v.(float64)
		s := num.v.(float64)
		return Number{
			t: Doub,
			v: f * s,
		}, nil
	default:
		return Number{}, fmt.Errorf("type %v doesn't supported", n.t)
	}
}

func (n *Number) Add(num Number) (Number, error) {
	if n.t != num.t {
		return Number{}, errors.New("type mismatch")
	}
	switch n.t {
	case Int:
		f := n.v.(int64)
		s := num.v.(int64)
		return Number{
			t: Int,
			v: f + s,
		}, nil
	case Doub:
		f := n.v.(float64)
		s := num.v.(float64)
		return Number{
			t: Doub,
			v: f + s,
		}, nil
	default:
		return Number{}, fmt.Errorf("type %v doesn't supported", n.t)
	}
}
