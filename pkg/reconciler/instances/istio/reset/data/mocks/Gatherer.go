// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	data "github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/reset/data"
	kubernetes "k8s.io/client-go/kubernetes"

	mock "github.com/stretchr/testify/mock"

	retry "github.com/avast/retry-go"

	v1 "k8s.io/api/core/v1"
)

// Gatherer is an autogenerated mock type for the Gatherer type
type Gatherer struct {
	mock.Mock
}

// GetAllPods provides a mock function with given fields: kubeClient, retryOpts
func (_m *Gatherer) GetAllPods(kubeClient kubernetes.Interface, retryOpts []retry.Option) (*v1.PodList, error) {
	ret := _m.Called(kubeClient, retryOpts)

	var r0 *v1.PodList
	if rf, ok := ret.Get(0).(func(kubernetes.Interface, []retry.Option) *v1.PodList); ok {
		r0 = rf(kubeClient, retryOpts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PodList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(kubernetes.Interface, []retry.Option) error); ok {
		r1 = rf(kubeClient, retryOpts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPodsWithDifferentImage provides a mock function with given fields: inputPodsList, image
func (_m *Gatherer) GetPodsWithDifferentImage(inputPodsList v1.PodList, image data.ExpectedImage) v1.PodList {
	ret := _m.Called(inputPodsList, image)

	var r0 v1.PodList
	if rf, ok := ret.Get(0).(func(v1.PodList, data.ExpectedImage) v1.PodList); ok {
		r0 = rf(inputPodsList, image)
	} else {
		r0 = ret.Get(0).(v1.PodList)
	}

	return r0
}