package hook

import (
	"bytes"
	"testing"

	"github.com/ipfs/go-datastore"
)

func TestIsDatastore(t *testing.T) {
	// ensure it implements datastore.Datastore
	var hds datastore.Datastore = NewDatastore(datastore.NewMapDatastore())
	hds.Close()
}

func TestHookPut(t *testing.T) {
	beforeHookCalled := false
	afterHookCalled := false

	key := datastore.NewKey("test")
	value := []byte("test")

	onBeforePut := func(k datastore.Key, v []byte) (datastore.Key, []byte) {
		if k != key {
			t.Fatal("incorrect key")
		}
		if bytes.Compare(v, value) != 0 {
			t.Fatal("incorrect value")
		}
		beforeHookCalled = true
		return k, v
	}

	onAfterPut := func(k datastore.Key, v []byte, err error) error {
		if k != key {
			t.Fatal("incorrect key")
		}
		if bytes.Compare(v, value) != 0 {
			t.Fatal("incorrect value")
		}
		afterHookCalled = true
		return err
	}

	ds := datastore.NewMapDatastore()
	hds := NewDatastore(ds, WithBeforePut(onBeforePut), WithAfterPut(onAfterPut))
	defer hds.Close()

	err := hds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if !beforeHookCalled {
		t.Fatal("before hook not called")
	}

	if !afterHookCalled {
		t.Fatal("after hook not called")
	}
}

func TestHookGet(t *testing.T) {
	beforeHookCalled := false
	afterHookCalled := false

	key := datastore.NewKey("test")
	value := []byte("test")

	onBeforeGet := func(k datastore.Key) datastore.Key {
		if k != key {
			t.Fatal("incorrect key")
		}
		beforeHookCalled = true
		return k
	}

	onAfterGet := func(k datastore.Key, v []byte, err error) ([]byte, error) {
		if k != key {
			t.Fatal("incorrect key")
		}
		if bytes.Compare(v, value) != 0 {
			t.Fatal("incorrect value")
		}
		afterHookCalled = true
		return v, err
	}

	ds := datastore.NewMapDatastore()
	hds := NewDatastore(ds, WithBeforeGet(onBeforeGet), WithAfterGet(onAfterGet))
	defer hds.Close()

	err := hds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	v, err := hds.Get(key)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if bytes.Compare(v, value) != 0 {
		t.Fatal("incorrect value")
	}

	if !beforeHookCalled {
		t.Fatal("before hook not called")
	}

	if !afterHookCalled {
		t.Fatal("after hook not called")
	}
}

func TestHookDelete(t *testing.T) {
	beforeHookCalled := false
	afterHookCalled := false

	key := datastore.NewKey("test")
	value := []byte("test")

	onBeforeDelete := func(k datastore.Key) datastore.Key {
		if k != key {
			t.Fatal("incorrect key")
		}
		beforeHookCalled = true
		return k
	}

	onAfterDelete := func(k datastore.Key, err error) error {
		if k != key {
			t.Fatal("incorrect key")
		}
		afterHookCalled = true
		return err
	}

	ds := datastore.NewMapDatastore()
	hds := NewDatastore(ds, WithBeforeDelete(onBeforeDelete), WithAfterDelete(onAfterDelete))
	defer hds.Close()

	err := hds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = hds.Delete(key)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if !beforeHookCalled {
		t.Fatal("before hook not called")
	}

	if !afterHookCalled {
		t.Fatal("after hook not called")
	}
}

func TestHookHas(t *testing.T) {
	beforeHookCalled := false
	afterHookCalled := false

	key := datastore.NewKey("test")
	value := []byte("test")

	onBeforeHas := func(k datastore.Key) datastore.Key {
		if k != key {
			t.Fatal("incorrect key")
		}
		beforeHookCalled = true
		return k
	}

	onAfterHas := func(k datastore.Key, exists bool, err error) (bool, error) {
		if k != key {
			t.Fatal("incorrect key")
		}
		afterHookCalled = true
		return exists, err
	}

	ds := datastore.NewMapDatastore()
	hds := NewDatastore(ds, WithBeforeHas(onBeforeHas), WithAfterHas(onAfterHas))
	defer hds.Close()

	err := hds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	exists, err := hds.Has(key)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if !exists {
		t.Fatal("expected key to exist")
	}

	if !beforeHookCalled {
		t.Fatal("before hook not called")
	}

	if !afterHookCalled {
		t.Fatal("after hook not called")
	}
}

func TestHookGetSize(t *testing.T) {
	key := datastore.NewKey("test")
	value := []byte("test")

	ds := datastore.NewMapDatastore()
	hds := NewDatastore(ds)
	defer hds.Close()

	err := hds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	size, err := hds.GetSize(key)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if size != len(value) {
		t.Fatal("incorrect size")
	}
}
