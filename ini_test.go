package iniparser

import (
	"testing"
)

func Test_LoadFile(t *testing.T) {
	ini, err := LoadFile("test.ini", "utf-8")
	if err != nil {
		t.Logf("err: %s", err)
		t.FailNow()
	}
	ip, ok := ini.GetString("proxy1", "ip")
	if !ok {
		t.FailNow()
	}
	port, ok := ini.GetInt("proxy1", "port")
	if !ok {
		t.FailNow()
	}
	weight, ok := ini.GetFloat("proxy1", "weight")
	if !ok {
		t.FailNow()
	}
	t.Logf("proxy1 ip: %s port: %d weight: %f", ip, port, weight)

	ip, _ = ini.GetString("proxy2", "ip")
	port, _ = ini.GetInt("proxy2", "port")
	weight, _ = ini.GetFloat("proxy2", "weight")
	t.Logf("proxy2 ip: %s port: %d weight: %f", ip, port, weight)
}
