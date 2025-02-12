// Code generated by counterfeiter. DO NOT EDIT.
package builderfakes

import (
	"sync"

	"github.com/pivotal/kpack/pkg/cnb"
	"github.com/pivotal/kpack/pkg/reconciler/v1alpha1/builder"
	"github.com/pivotal/kpack/pkg/registry"
)

type FakeMetadataRetriever struct {
	GetBuilderImageStub        func(registry.ImageRef) (cnb.BuilderImage, error)
	getBuilderImageMutex       sync.RWMutex
	getBuilderImageArgsForCall []struct {
		arg1 registry.ImageRef
	}
	getBuilderImageReturns struct {
		result1 cnb.BuilderImage
		result2 error
	}
	getBuilderImageReturnsOnCall map[int]struct {
		result1 cnb.BuilderImage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetadataRetriever) GetBuilderImage(arg1 registry.ImageRef) (cnb.BuilderImage, error) {
	fake.getBuilderImageMutex.Lock()
	ret, specificReturn := fake.getBuilderImageReturnsOnCall[len(fake.getBuilderImageArgsForCall)]
	fake.getBuilderImageArgsForCall = append(fake.getBuilderImageArgsForCall, struct {
		arg1 registry.ImageRef
	}{arg1})
	fake.recordInvocation("GetBuilderImage", []interface{}{arg1})
	fake.getBuilderImageMutex.Unlock()
	if fake.GetBuilderImageStub != nil {
		return fake.GetBuilderImageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBuilderImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMetadataRetriever) GetBuilderImageCallCount() int {
	fake.getBuilderImageMutex.RLock()
	defer fake.getBuilderImageMutex.RUnlock()
	return len(fake.getBuilderImageArgsForCall)
}

func (fake *FakeMetadataRetriever) GetBuilderImageCalls(stub func(registry.ImageRef) (cnb.BuilderImage, error)) {
	fake.getBuilderImageMutex.Lock()
	defer fake.getBuilderImageMutex.Unlock()
	fake.GetBuilderImageStub = stub
}

func (fake *FakeMetadataRetriever) GetBuilderImageArgsForCall(i int) registry.ImageRef {
	fake.getBuilderImageMutex.RLock()
	defer fake.getBuilderImageMutex.RUnlock()
	argsForCall := fake.getBuilderImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMetadataRetriever) GetBuilderImageReturns(result1 cnb.BuilderImage, result2 error) {
	fake.getBuilderImageMutex.Lock()
	defer fake.getBuilderImageMutex.Unlock()
	fake.GetBuilderImageStub = nil
	fake.getBuilderImageReturns = struct {
		result1 cnb.BuilderImage
		result2 error
	}{result1, result2}
}

func (fake *FakeMetadataRetriever) GetBuilderImageReturnsOnCall(i int, result1 cnb.BuilderImage, result2 error) {
	fake.getBuilderImageMutex.Lock()
	defer fake.getBuilderImageMutex.Unlock()
	fake.GetBuilderImageStub = nil
	if fake.getBuilderImageReturnsOnCall == nil {
		fake.getBuilderImageReturnsOnCall = make(map[int]struct {
			result1 cnb.BuilderImage
			result2 error
		})
	}
	fake.getBuilderImageReturnsOnCall[i] = struct {
		result1 cnb.BuilderImage
		result2 error
	}{result1, result2}
}

func (fake *FakeMetadataRetriever) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBuilderImageMutex.RLock()
	defer fake.getBuilderImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMetadataRetriever) recordInvocation(key string, args []interface{}) {
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

var _ builder.MetadataRetriever = new(FakeMetadataRetriever)
