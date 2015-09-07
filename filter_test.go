package gofilter

import (
	"net"
	"testing"
)

func init() {
	//	filterDebug = 5
	//	filterErrorVerbose = true

	RegisterField("f_uint", FT_UINT)
	RegisterField("f_uint8", FT_UINT8)
	RegisterField("f_uint16", FT_UINT16)
	RegisterField("f_uint32", FT_UINT32)
	RegisterField("f_uint64", FT_UINT64)

	RegisterField("f_int", FT_INT)
	RegisterField("f_int8", FT_INT8)
	RegisterField("f_int16", FT_INT16)
	RegisterField("f_int32", FT_INT32)
	RegisterField("f_int64", FT_INT64)

	RegisterField("f_string", FT_STRING)

	RegisterField("f_ipv4", FT_IP)
	RegisterField("f_ipv6", FT_IP)

	RegisterField("ip.addr", FT_IP)
	RegisterField("ip.src", FT_IP)
	RegisterField("ip.dst", FT_IP)

	RegisterField("f_mac", FT_MAC)

	RegisterField("f_bool.1", FT_BOOL)
	RegisterField("f_bool.2", FT_BOOL)

	RegisterField("f_float32", FT_FLOAT32)
	RegisterField("f_float64", FT_FLOAT64)

	RegisterField("f_string.1", FT_STRING)
	RegisterField("f_string.2", FT_STRING)
	RegisterField("f_string.3", FT_STRING)

}

func TestFilterRegisterFilter(t *testing.T) {
	err := RegisterField("f_uint", FT_UINT)
	if err != ErrFieldExist {
		t.Error("Must be error: field is already registered")
	}
}

func TestFilterParseUnknownField(t *testing.T) {
	_, err := NewFilter("xxx && udp.srcport==12")
	if err == nil {
		t.Error("Must be error: Field with name \"xxx\" does not exists")
	}
}

func TestFilterFieldName(t *testing.T) {
	_, err := NewFilter("f_uint")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseUint(t *testing.T) {
	_, err := NewFilter("f_uint==4294967295 && f_uint8==255 && f_uint16==65535 && f_uint32==4294967295 && f_uint64==18446744073709551615")
	if err != nil {
		t.Error(err)
	}
	_, err = NewFilter("f_uint==0 && f_uint8==0 && f_uint16==0 && f_uint32==0 && f_uint64==0")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseInt(t *testing.T) {
	_, err := NewFilter("f_int==2147483647 && f_int8==127 && f_int16==32767 && f_int32==2147483647 && f_int64==9223372036854775807")
	if err != nil {
		t.Error(err)
	}
	_, err = NewFilter("f_int==-2147483648 && f_int8==-128 && f_int16==-32768 && f_int32==-2147483648 && f_int64==-9223372036854775808")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseString(t *testing.T) {
	_, err := NewFilter("f_string==\"text\" or f_string contains 12 && f_string==01:23:45:67:89:ab:AB:cd:ef")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseRegexp(t *testing.T) {
	_, err := NewFilter("f_string matches \"gl=se$\"")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseIP(t *testing.T) {
	_, err := NewFilter("f_ipv4 == 192.168.1.1 or f_ipv6==::1 or f_ipv6==2001:db8::1")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseMac(t *testing.T) {
	_, err := NewFilter("f_mac == 01:23:45:67:89:ab:cd:ef or f_mac == 0123.4567.89ab.cdef")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseBool(t *testing.T) {
	_, err := NewFilter("f_bool.1 == true or f_bool.2 != false")
	if err != nil {
		t.Error(err)
	}
}

func TestFilterParseFloat(t *testing.T) {
	_, err := NewFilter("f_float32 == 123.345 or f_float64 != 74123412341234.123412341243")
	if err != nil {
		t.Error(err)
	}
}

// Apply filters
func TestFilterApplyIntUint(t *testing.T) {
	f, err := NewFilter("f_int == 1 and f_int8 == 13")
	if err != nil {
		t.Fatal(err)
	}

	if !f.Apply(map[string]interface{}{"f_int": 1, "f_int8": int8(13)}) {
		t.Error("Packet must pass")
	}

	if f.Apply(map[string]interface{}{"f_int": 1, "f_int8": int8(14)}) {
		t.Error("Packet must not pass")
	}

	// field with multiple values
	f2, err := NewFilter("f_int != 2")
	if err != nil {
		t.Fatal(err)
	}
	if !f2.Apply(map[string]interface{}{"f_int": []int{1, 3, 4}}) {
		t.Error("Packet must pass")
	}

	if f2.Apply(map[string]interface{}{"f_int": []int{1, 2, 3, 4}}) {
		t.Error("Packet must not pass")
	}
}

func TestFilterApplyString(t *testing.T) {
	f, err := NewFilter("f_string.1 == \"1\" and f_string.2 == 47:45:54  and f_string.3 == abc123")
	if err != nil {
		t.Fatal(err)
	}
	if !f.Apply(map[string]interface{}{"f_string.1": "1", "f_string.2": "GET", "f_string.3": "abc123"}) {
		t.Error("Packet must pass")
	}

	if f.Apply(map[string]interface{}{"f_string.1": "2", "f_string.2": "GET", "f_string.3": "abc123"}) {
		t.Error("Packet must not pass")
	}

	f2, err := NewFilter("f_string.1 contains \"1\" and f_string.2 contains 47:45:54  and f_string.3 contains abc123")
	if err != nil {
		t.Fatal(err)
	}
	if !f2.Apply(map[string]interface{}{"f_string.1": "asdf1asdf", "f_string.2": "text - GET ---", "f_string.3": "asf fffabc123"}) {
		t.Error("Packet must pass")
	}

	if f2.Apply(map[string]interface{}{"f_string.1": "test234test", "f_string.2": "xxxxETyyy", "f_string.3": "abc125"}) {
		t.Error("Packet must not pass")
	}

}

func TestFilterApplyIP(t *testing.T) {
	f, err := NewFilter("f_ipv4 > 192.168.1.0 and f_ipv6 !=2001:db8::1")
	if err != nil {
		t.Fatal(err)
	}

	if !f.Apply(map[string]interface{}{"f_ipv4": net.ParseIP("192.168.100.1"), "f_ipv6": net.ParseIP("2001:db8::2")}) {
		t.Error("Packet must pass")
	}

	if f.Apply(map[string]interface{}{"f_ipv4": net.ParseIP("172.16.0.1"), "f_ipv6": net.ParseIP("2001:db8::1")}) {
		t.Error("Packet must not pass")
	}

	// field with multiple values
	f2, err := NewFilter("ip.addr != 2001:db8::1")
	if err != nil {
		t.Fatal(err)
	}
	if f2.Apply(map[string]interface{}{
		"ip.addr": []net.IP{
			net.ParseIP("192.168.100.1"),
			net.ParseIP("2001:db8::1")}},
	) {
		t.Error("Packet must not pass")
	}

	// compare CIDR
	f3, err := NewFilter("ip.src==192.168.0.0/16 and ip.dst==192.168.0.0/16")
	if err != nil {
		t.Fatal(err)
	}
	if !f3.Apply(map[string]interface{}{"ip.src": net.ParseIP("192.168.100.1"), "ip.dst": net.ParseIP("192.168.1.1")}) {
		t.Error("Packet must pass")
	}
	if f.Apply(map[string]interface{}{"ip.src": net.ParseIP("172.16.0.1"), "ip.dst": net.ParseIP("10.0.0.1")}) {
		t.Error("Packet must not pass")
	}

	f4, err := NewFilter("ip.src >= 192.168.0.0/16")
	if err != nil {
		t.Fatal(err)
	}

	_, ipnet, _ := net.ParseCIDR("192.168.0.12/16")
	if !f4.Apply(map[string]interface{}{"ip.src": ipnet.IP, "ip.dst": net.ParseIP("192.168.1.1")}) {

		t.Error("Packet must pass")
	}
	if f4.Apply(map[string]interface{}{"ip.src": net.ParseIP("172.16.0.1"), "ip.dst": net.ParseIP("10.0.0.1")}) {
		t.Error("Packet must not pass")
	}
}

func TestFilterApplyMac(t *testing.T) {
	// compare contains
	f, err := NewFilter("f_mac contains 23:45")
	if err != nil {
		t.Fatal(err)
	}
	h1, _ := net.ParseMAC("01:23:45:67:89:ab:cd:ef")
	if !f.Apply(map[string]interface{}{"f_mac": h1}) {
		t.Error("Packet must pass")
	}
	h2, _ := net.ParseMAC("01:23:44:67:89:ab:cd:ef1")
	if f.Apply(map[string]interface{}{"f_mac": h2}) {
		t.Error("Packet must not pass")
	}
}
