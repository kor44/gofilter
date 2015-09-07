package gofilter

import (
	"bytes"
	"net"
	"regexp"
	"strings"
)

type node interface {
	Apply(map[string]interface{}) bool
}

type node2 interface {
	Apply(map[string]interface{}) bool
	FieldName() string
	applyOne(v interface{}) bool
}

// AND
type nodeAnd struct {
	left  node
	right node
}

func (n *nodeAnd) Apply(p map[string]interface{}) bool {
	return (n.left.Apply(p) && n.right.Apply(p))
}

// OR
type nodeOr struct {
	left  node
	right node
}

func (n *nodeOr) Apply(p map[string]interface{}) bool {
	return n.left.Apply(p) || n.right.Apply(p)
}

// NOT
type nodeNot struct {
	right node
}

func (n *nodeNot) Apply(p map[string]interface{}) bool {
	return !n.right.Apply(p)
}

// EXIST
type nodeExist struct {
	fieldName string
}

func (n *nodeExist) Apply(p map[string]interface{}) bool {
	_, ok := p[n.fieldName]
	return ok
}

// EQUAL
type nodeEq struct {
	fieldName string
	value     interface{}
}

func (n *nodeEq) Apply(p map[string]interface{}) bool {
	return applyRange(p, n)
}

func (n *nodeEq) applyOne(v interface{}) bool {
	switch n.value.(type) {
	case []byte:
		if x, ok := v.([]byte); ok {
			return bytes.Equal(x, n.value.([]byte))
		} else if x, ok := v.(string); ok {
			return bytes.Equal([]byte(x), n.value.([]byte))
		}

	case net.IP:
		if x, ok := v.(net.IP); ok {
			return x.Equal(n.value.(net.IP))
		}
	case *net.IPNet:
		if x, ok := v.(net.IP); ok {
			return n.value.(*net.IPNet).Contains(x)
		}
	case net.HardwareAddr:
		if x, ok := v.(net.HardwareAddr); ok {
			return bytes.Equal(n.value.([]byte), x)
		}

	case uint:
		if x, ok := v.(uint); ok {
			return x == n.value.(uint)
		}
	case uint8:
		if x, ok := v.(uint8); ok {
			return x == n.value.(uint8)
		}
	case uint16:
		if x, ok := v.(uint16); ok {
			return x == n.value.(uint16)
		}
	case uint32:
		if x, ok := v.(uint32); ok {
			return x == n.value.(uint32)
		}
	case uint64:
		if x, ok := v.(uint64); ok {
			return x == n.value.(uint64)
		}
	case int:
		if x, ok := v.(int); ok {
			return x == n.value.(int)
		}
	case int8:
		if x, ok := v.(int8); ok {
			return x == n.value.(int8)
		}
	case int16:
		if x, ok := v.(int16); ok {
			return x == n.value.(int16)
		}
	case int32:
		if x, ok := v.(int32); ok {
			return x == n.value.(int32)
		}
	case int64:
		if x, ok := v.(int64); ok {
			return x == n.value.(int64)
		}

	case float32:
		if x, ok := v.(float32); ok {
			return x == n.value.(float32)
		}
	case float64:
		if x, ok := v.(float64); ok {
			return x == n.value.(float64)
		}

	case string:
		if x, ok := v.(string); ok {
			return x == n.value.(string)
		}
	}

	return false // было v == n.value
}

func (n *nodeEq) FieldName() string {
	return n.fieldName
}

// TEST_GT
type nodeGt struct {
	fieldName string
	value     interface{}
}

func (n *nodeGt) Apply(p map[string]interface{}) bool {
	return applyRange(p, n)
}

func (n *nodeGt) applyOne(v interface{}) bool {
	switch n.value.(type) {
	case []byte:
		if x, ok := v.([]byte); ok {
			if bytes.Compare(x, n.value.([]byte)) == 1 {
				return true
			} else {
				return false
			}
		} else if x, ok := v.(string); ok {
			if bytes.Compare([]byte(x), n.value.([]byte)) == 1 {
				return true
			} else {
				return false
			}
		}

	case net.IP:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x.To16()), []byte(n.value.(net.IP))) == 1 {
				return true
			} else {
				return false
			}
		}
	case *net.IPNet:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x.To16()), []byte(n.value.(*net.IPNet).IP)) == 1 {
				return true
			} else {
				return false
			}
		}
	case net.HardwareAddr:
		if x, ok := v.(net.HardwareAddr); ok {
			if bytes.Compare([]byte(x), []byte(n.value.(net.HardwareAddr))) == 1 {
				return true
			} else {
				return false
			}
		}

	case uint:
		if x, ok := v.(uint); ok {
			return x > n.value.(uint)
		}
	case uint8:
		if x, ok := v.(uint8); ok {
			return x > n.value.(uint8)
		}
	case uint16:
		if x, ok := v.(uint16); ok {
			return x > n.value.(uint16)
		}
	case uint32:
		if x, ok := v.(uint32); ok {
			return x > n.value.(uint32)
		}
	case uint64:
		if x, ok := v.(uint64); ok {
			return x > n.value.(uint64)
		}
	case int:
		if x, ok := v.(int); ok {
			return x > n.value.(int)
		}
	case int8:
		if x, ok := v.(int8); ok {
			return x > n.value.(int8)
		}
	case int16:
		if x, ok := v.(int16); ok {
			return x > n.value.(int16)
		}
	case int32:
		if x, ok := v.(int32); ok {
			return x > n.value.(int32)
		}
	case int64:
		if x, ok := v.(int64); ok {
			return x > n.value.(int64)
		}
	case float32:
		if x, ok := v.(float32); ok {
			return x > n.value.(float32)
		}
	case float64:
		if x, ok := v.(float64); ok {
			return x > n.value.(float64)
		}
	case string:
		if x, ok := v.(string); ok {
			return x > n.value.(string)
		}
	}
	return false
}

func (n *nodeGt) FieldName() string {
	return n.fieldName
}

// TEST_GE
type nodeGe struct {
	fieldName string
	value     interface{}
}

func (n *nodeGe) Apply(p map[string]interface{}) bool {
	return applyRange(p, n)
}

func (n *nodeGe) applyOne(v interface{}) bool {
	switch n.value.(type) {
	case []byte:
		if x, ok := v.([]byte); ok {
			if bytes.Compare(x, n.value.([]byte)) >= 0 {
				return true
			} else {
				return false
			}
		} else if x, ok := v.(string); ok {
			if bytes.Compare([]byte(x), n.value.([]byte)) >= 0 {
				return true
			} else {
				return false
			}
		}

	case net.IP:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x.To16()), []byte(n.value.(net.IP))) >= 0 {
				return true
			} else {
				return false
			}
		}
	case *net.IPNet:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x.To16()), []byte(n.value.(*net.IPNet).IP)) >= 0 {
				return true
			} else {
				return false
			}
		}
	case net.HardwareAddr:
		if x, ok := v.(net.HardwareAddr); ok {
			if bytes.Compare([]byte(x), []byte(n.value.(net.HardwareAddr))) >= 0 {
				return true
			} else {
				return false
			}
		}

	case int:
		if x, ok := v.(int); ok {
			return x >= n.value.(int)
		}
	case int8:
		if x, ok := v.(int8); ok {
			return x >= n.value.(int8)
		}
	case int16:
		if x, ok := v.(int16); ok {
			return x >= n.value.(int16)
		}
	case int32:
		if x, ok := v.(int32); ok {
			return x >= n.value.(int32)
		}
	case int64:
		if x, ok := v.(int64); ok {
			return x >= n.value.(int64)
		}
	case uint:
		if x, ok := v.(uint); ok {
			return x >= n.value.(uint)
		}
	case uint8:
		if x, ok := v.(uint8); ok {
			return x >= n.value.(uint8)
		}
	case uint16:
		if x, ok := v.(uint16); ok {
			return x >= n.value.(uint16)
		}
	case uint32:
		if x, ok := v.(uint32); ok {
			return x >= n.value.(uint32)
		}
	case uint64:
		if x, ok := v.(uint64); ok {
			return x >= n.value.(uint64)
		}
	case float32:
		if x, ok := v.(float32); ok {
			return x >= n.value.(float32)
		}
	case float64:
		if x, ok := v.(float64); ok {
			return x >= n.value.(float64)
		}
	case string:
		if x, ok := v.(string); ok {
			return x >= n.value.(string)
		}
	}
	return false
}

func (n *nodeGe) FieldName() string {
	return n.fieldName
}

// TEST_LT
type nodeLt struct {
	fieldName string
	value     interface{}
}

func (n *nodeLt) Apply(p map[string]interface{}) bool {
	return applyRange(p, n)
}

func (n *nodeLt) applyOne(v interface{}) bool {
	switch n.value.(type) {
	case []byte:
		if x, ok := v.([]byte); ok {
			if bytes.Compare(x, n.value.([]byte)) == -1 {
				return true
			} else {
				return false
			}
		} else if x, ok := v.(string); ok {
			if bytes.Compare([]byte(x), n.value.([]byte)) == -1 {
				return true
			} else {
				return false
			}
		}

	case net.IP:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x.To16()), []byte(n.value.(net.IP))) == -1 {
				return true
			} else {
				return false
			}
		}
	case *net.IPNet:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x.To16()), []byte(n.value.(*net.IPNet).IP)) == -1 {
				return true
			} else {
				return false
			}
		}
	case net.HardwareAddr:
		if x, ok := v.(net.HardwareAddr); ok {
			if bytes.Compare([]byte(x), []byte(n.value.(net.HardwareAddr))) == -1 {
				return true
			} else {
				return false
			}
		}

	case uint:
		if x, ok := v.(uint); ok {
			return x < n.value.(uint)
		}
	case uint8:
		if x, ok := v.(uint8); ok {
			return x < n.value.(uint8)
		}
	case uint16:
		if x, ok := v.(uint16); ok {
			return x < n.value.(uint16)
		}
	case uint32:
		if x, ok := v.(uint32); ok {
			return x < n.value.(uint32)
		}
	case uint64:
		if x, ok := v.(uint64); ok {
			return x < n.value.(uint64)
		}
	case int:
		if x, ok := v.(int); ok {
			return x < n.value.(int)
		}
	case int8:
		if x, ok := v.(int8); ok {
			return x < n.value.(int8)
		}
	case int16:
		if x, ok := v.(int16); ok {
			return x < n.value.(int16)
		}
	case int32:
		if x, ok := v.(int32); ok {
			return x < n.value.(int32)
		}
	case int64:
		if x, ok := v.(int64); ok {
			return x < n.value.(int64)
		}
	case float32:
		if x, ok := v.(float32); ok {
			return x < n.value.(float32)
		}
	case float64:
		if x, ok := v.(float64); ok {
			return x < n.value.(float64)
		}
	case string:
		if x, ok := v.(string); ok {
			return x < n.value.(string)
		}
	}
	return false
}

func (n *nodeLt) FieldName() string {
	return n.fieldName
}

// TEST_LE
type nodeLe struct {
	fieldName string
	value     interface{}
}

func (n *nodeLe) Apply(p map[string]interface{}) bool {
	return applyRange(p, n)
}

func (n *nodeLe) applyOne(v interface{}) bool {
	switch n.value.(type) {
	case []byte:
		if x, ok := v.([]byte); ok {
			if bytes.Compare(x, n.value.([]byte)) <= 0 {
				return true
			} else {
				return false
			}
		} else if x, ok := v.(string); ok {
			if bytes.Compare([]byte(x), n.value.([]byte)) <= 0 {
				return true
			} else {
				return false
			}
		}

	case net.IP:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x), []byte(n.value.(net.IP))) <= 0 {
				return true
			} else {
				return false
			}
		}
	case *net.IPNet:
		if x, ok := v.(net.IP); ok {
			if bytes.Compare([]byte(x), []byte(n.value.(*net.IPNet).IP)) <= 0 {
				return true
			} else {
				return false
			}
		}
	case net.HardwareAddr:
		if x, ok := v.(net.HardwareAddr); ok {
			if bytes.Compare([]byte(x), []byte(n.value.(net.HardwareAddr))) <= 0 {
				return true
			} else {
				return false
			}
		}

	case uint:
		if x, ok := v.(uint); ok {
			return x <= n.value.(uint)
		}
	case uint8:
		if x, ok := v.(uint8); ok {
			return x <= n.value.(uint8)
		}
	case uint16:
		if x, ok := v.(uint16); ok {
			return x <= n.value.(uint16)
		}
	case uint32:
		if x, ok := v.(uint32); ok {
			return x <= n.value.(uint32)
		}
	case uint64:
		if x, ok := v.(uint64); ok {
			return x <= n.value.(uint64)
		}
	case int:
		if x, ok := v.(int); ok {
			return x <= n.value.(int)
		}
	case int8:
		if x, ok := v.(int8); ok {
			return x <= n.value.(int8)
		}
	case int16:
		if x, ok := v.(int16); ok {
			return x <= n.value.(int16)
		}
	case int32:
		if x, ok := v.(int32); ok {
			return x <= n.value.(int32)
		}
	case int64:
		if x, ok := v.(int64); ok {
			return x <= n.value.(int64)
		}
	case float32:
		if x, ok := v.(float32); ok {
			return x <= n.value.(float32)
		}
	case float64:
		if x, ok := v.(float64); ok {
			return x <= n.value.(float64)
		}
	case string:
		if x, ok := v.(string); ok {
			return x <= n.value.(string)
		}
	}
	return false
}

func (n *nodeLe) FieldName() string {
	return n.fieldName
}

// TEST_CONTAINS
type nodeContains struct {
	fieldName string
	value     interface{}
}

func (n *nodeContains) Apply(p map[string]interface{}) bool {
	return applyRange(p, n)
}

func (n *nodeContains) applyOne(v interface{}) bool {
	switch n.value.(type) {
	case string:
		if x, ok := v.(string); ok {
			return strings.Contains(x, n.value.(string))
		}
	case []byte:
		node_value := n.value.([]byte)
		switch v.(type) {
		case []byte:
			return bytes.Contains(v.([]byte), node_value)
		case string:
			return bytes.Contains([]byte(v.(string)), node_value)
		/*case net.IP:
			return bytes.Contains([]byte(v.(net.IP)), node_value)
		case net.IPNet:
			return bytes.Contains([]byte(v.(net.IPNet).IP), node_value)*/
		case net.HardwareAddr:
			return bytes.Contains([]byte(v.(net.HardwareAddr)), node_value)
		}
	}
	return false
}

func (n *nodeContains) FieldName() string {
	return n.fieldName
}

// TEST_MATCHES
type nodeMatch struct {
	fieldName string
	reg_expr  *regexp.Regexp
}

func (n *nodeMatch) Apply(p map[string]interface{}) bool {
	return applyRange(p, n)
}

func (n *nodeMatch) applyOne(v interface{}) bool {

	switch v.(type) {
	case string:
		return n.reg_expr.MatchString(v.(string))
	}
	return false
}

func (n *nodeMatch) FieldName() string {
	return n.fieldName
}

// invoke apply for slice of values
func applyRange(p map[string]interface{}, n node2) bool {
	v, ok := p[n.FieldName()]
	if !ok {
		return false
	}

	switch v.(type) {
	case []bool:
		for _, x := range v.([]bool) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []uint:
		for _, x := range v.([]uint) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []uint8:
		for _, x := range v.([]uint8) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []uint16:
		for _, x := range v.([]uint16) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []uint32:
		for _, x := range v.([]uint32) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []uint64:
		for _, x := range v.([]uint64) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []int:
		for _, x := range v.([]int) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []int8:
		for _, x := range v.([]int8) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []int16:
		for _, x := range v.([]int16) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []int32:
		for _, x := range v.([]int32) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []int64:
		for _, x := range v.([]int64) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []float32:
		for _, x := range v.([]float32) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []float64:
		for _, x := range v.([]float64) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []string:
		for _, x := range v.([]string) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []net.IP:
		for _, x := range v.([]net.IP) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	case []net.HardwareAddr:
		for _, x := range v.([]net.HardwareAddr) {
			if n.applyOne(x) {
				return true
			}
		}
		return false
	default:
		return n.applyOne(v)
	}
	return false
}
