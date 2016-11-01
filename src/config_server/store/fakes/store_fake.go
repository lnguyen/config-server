// This file was generated by counterfeiter
package fakes

import (
	"config_server/store"
	"sync"
)

type FakeStore struct {
	PutStub        func(key string, value string) error
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		key   string
		value string
	}
	putReturns struct {
		result1 error
	}
	GetByNameStub        func(name string) (store.Configuration, error)
	getByNameMutex       sync.RWMutex
	getByNameArgsForCall []struct {
		name string
	}
	getByNameReturns struct {
		result1 store.Configuration
		result2 error
	}
	GetByIDStub        func(id string) (store.Configuration, error)
	getByIDMutex       sync.RWMutex
	getByIDArgsForCall []struct {
		id string
	}
	getByIDReturns struct {
		result1 store.Configuration
		result2 error
	}
	DeleteStub        func(key string) (bool, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		key string
	}
	deleteReturns struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) Put(key string, value string) error {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		key   string
		value string
	}{key, value})
	fake.recordInvocation("Put", []interface{}{key, value})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(key, value)
	} else {
		return fake.putReturns.result1
	}
}

func (fake *FakeStore) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeStore) PutArgsForCall(i int) (string, string) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].key, fake.putArgsForCall[i].value
}

func (fake *FakeStore) PutReturns(result1 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetByName(name string) (store.Configuration, error) {
	fake.getByNameMutex.Lock()
	fake.getByNameArgsForCall = append(fake.getByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetByName", []interface{}{name})
	fake.getByNameMutex.Unlock()
	if fake.GetByNameStub != nil {
		return fake.GetByNameStub(name)
	} else {
		return fake.getByNameReturns.result1, fake.getByNameReturns.result2
	}
}

func (fake *FakeStore) GetByNameCallCount() int {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	return len(fake.getByNameArgsForCall)
}

func (fake *FakeStore) GetByNameArgsForCall(i int) string {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	return fake.getByNameArgsForCall[i].name
}

func (fake *FakeStore) GetByNameReturns(result1 store.Configuration, result2 error) {
	fake.GetByNameStub = nil
	fake.getByNameReturns = struct {
		result1 store.Configuration
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetByID(id string) (store.Configuration, error) {
	fake.getByIDMutex.Lock()
	fake.getByIDArgsForCall = append(fake.getByIDArgsForCall, struct {
		id string
	}{id})
	fake.recordInvocation("GetByID", []interface{}{id})
	fake.getByIDMutex.Unlock()
	if fake.GetByIDStub != nil {
		return fake.GetByIDStub(id)
	} else {
		return fake.getByIDReturns.result1, fake.getByIDReturns.result2
	}
}

func (fake *FakeStore) GetByIDCallCount() int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return len(fake.getByIDArgsForCall)
}

func (fake *FakeStore) GetByIDArgsForCall(i int) string {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return fake.getByIDArgsForCall[i].id
}

func (fake *FakeStore) GetByIDReturns(result1 store.Configuration, result2 error) {
	fake.GetByIDStub = nil
	fake.getByIDReturns = struct {
		result1 store.Configuration
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Delete(key string) (bool, error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		key string
	}{key})
	fake.recordInvocation("Delete", []interface{}{key})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(key)
	} else {
		return fake.deleteReturns.result1, fake.deleteReturns.result2
	}
}

func (fake *FakeStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeStore) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].key
}

func (fake *FakeStore) DeleteReturns(result1 bool, result2 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ store.Store = new(FakeStore)
