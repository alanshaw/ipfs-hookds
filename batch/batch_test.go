package batch

import (
	"bytes"
	"testing"

	"github.com/ipfs/go-datastore"
)

func TestIsBatch(t *testing.T) {
	ds := datastore.NewMapDatastore()
	defer ds.Close()

	bch, err := ds.Batch()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	// ensure it implements datastore.Batch
	var bds datastore.Batch = NewBatch(bch)

	err = bds.Put(datastore.NewKey("test"), []byte("test"))
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = bds.Commit()
	if err != nil {
		t.Fatal("unexpected error", err)
	}
}

func TestBatchHookPut(t *testing.T) {
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
		afterHookCalled = true
		return err
	}

	ds := datastore.NewMapDatastore()
	defer ds.Close()

	bch, err := ds.Batch()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	hbh := NewBatch(bch, WithBeforePut(onBeforePut), WithAfterPut(onAfterPut))

	err = hbh.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = hbh.Commit()
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

func TestBatchHookDelete(t *testing.T) {
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
	defer ds.Close()

	err := ds.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	bch, err := ds.Batch()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	hbh := NewBatch(bch, WithBeforeDelete(onBeforeDelete), WithAfterDelete(onAfterDelete))

	err = hbh.Delete(key)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = hbh.Commit()
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

func TestBatchHookCommit(t *testing.T) {
	beforeHookCalled := false
	afterHookCalled := false

	key := datastore.NewKey("test")
	value := []byte("test")

	onBeforeCommit := func() {
		beforeHookCalled = true
	}

	onAfterCommit := func(err error) error {
		afterHookCalled = true
		return err
	}

	ds := datastore.NewMapDatastore()
	defer ds.Close()

	bch, err := ds.Batch()
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	hbh := NewBatch(bch, WithBeforeCommit(onBeforeCommit), WithAfterCommit(onAfterCommit))

	err = hbh.Put(key, value)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	err = hbh.Commit()
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
