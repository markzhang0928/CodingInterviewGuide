package main

import (
	"sync"
	"testing"
)

func Test_sync_map(t *testing.T) {
	var sMp sync.Map
	sMp.Store("key1", "value1")
	v, ok := sMp.Load("key1")
	if !ok {
		t.Error("key1 not found")
		return
	} else if v.(string) != "value1" {
		t.Errorf("expected value: %s, actual: %v", "value1", v)
	}
	k := "key2"
	sMp.Store("key2", "value2")
	sMp.Store("key3", "value3")
	sMp.Range(func(key, value any) bool {
		t.Logf("key: %v, value: %+v", key, value)
		return key.(string) != k
	})
	//str, _ := v.(string)
	//t.Errorf("value: %+v", str)
	//t.Logf("value: %v", v)
	// sMp.Store("key2", "value2")
	// v, ok = sMp.Load("key2")
	// if !ok {
	// 	t.Error("key2 not found")
	// }
}
