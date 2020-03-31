package hook

import (
	"bytes"
	"testing"

	"github.com/alanshaw/ipfs-hookds/opts"
	"github.com/ipfs/go-datastore"
)

func TestIsBatching(t *testing.T) {
	// ensure it implements datastore.Batching
	var bds datastore.Batching = NewBatching(datastore.NewMapDatastore())
	bds.Close()
}

func TestBatchingHookBatch(t *testing.T) {
	beforeHookCalled := false
	afterHookCalled := false

	key := datastore.NewKey("test")
	value := []byte("test")

	onBeforeBatch := func() {
		beforeHookCalled = true
	}

	onAfterBatch := func(bch datastore.Batch, err error) (datastore.Batch, error) {
		afterHookCalled = true
		return bch, err
	}

	ds := datastore.NewMapDatastore()
	bds := NewBatching(ds, opts.OnBeforeBatch(onBeforeBatch), opts.OnAfterBatch(onAfterBatch))
	defer bds.Close()

	bch, err := bds.Batch()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = bch.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = bch.Commit()
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

func TestBatchingHookPut(t *testing.T) {
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
	bds := NewBatching(ds, opts.OnBeforePut(onBeforePut), opts.OnAfterPut(onAfterPut))
	defer bds.Close()

	err := bds.Put(key, value)
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

func TestBatchingHookGet(t *testing.T) {
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
	bds := NewBatching(ds, opts.OnBeforeGet(onBeforeGet), opts.OnAfterGet(onAfterGet))
	defer bds.Close()

	err := bds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	v, err := bds.Get(key)
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

func TestBatchingHookDelete(t *testing.T) {
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
		afterHookCalled = true
		return err
	}

	ds := datastore.NewMapDatastore()
	bds := NewBatching(ds, opts.OnBeforeDelete(onBeforeDelete), opts.OnAfterDelete(onAfterDelete))
	defer bds.Close()

	err := bds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = bds.Delete(key)
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

func TestBatchingHookHas(t *testing.T) {
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
		if !exists {
			t.Fatal("expected key to exist")
		}
		afterHookCalled = true
		return exists, err
	}

	ds := datastore.NewMapDatastore()
	bds := NewBatching(ds, opts.OnBeforeHas(onBeforeHas), opts.OnAfterHas(onAfterHas))
	defer bds.Close()

	err := bds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	exists, err := bds.Has(key)
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

func TestBatchingHookGetSize(t *testing.T) {
	key := datastore.NewKey("test")
	value := []byte("test")

	ds := datastore.NewMapDatastore()
	bds := NewBatching(ds)
	defer bds.Close()

	err := bds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	size, err := bds.GetSize(key)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if size != len(value) {
		t.Fatal("incorrect size")
	}
}
