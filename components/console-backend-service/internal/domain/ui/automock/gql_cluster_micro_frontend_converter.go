// Code generated by mockery v1.0.0
package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/kyma-project/kyma/common/microfrontend-client/pkg/apis/ui/v1alpha1"

// gqlClusterMicroFrontendConverter is an autogenerated mock type for the gqlClusterMicroFrontendConverter type
type gqlClusterMicroFrontendConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlClusterMicroFrontendConverter) ToGQL(in *v1alpha1.ClusterMicroFrontend) (*gqlschema.ClusterMicroFrontend, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ClusterMicroFrontend
	if rf, ok := ret.Get(0).(func(*v1alpha1.ClusterMicroFrontend) *gqlschema.ClusterMicroFrontend); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ClusterMicroFrontend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.ClusterMicroFrontend) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlClusterMicroFrontendConverter) ToGQLs(in []*v1alpha1.ClusterMicroFrontend) ([]gqlschema.ClusterMicroFrontend, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ClusterMicroFrontend
	if rf, ok := ret.Get(0).(func([]*v1alpha1.ClusterMicroFrontend) []gqlschema.ClusterMicroFrontend); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ClusterMicroFrontend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1alpha1.ClusterMicroFrontend) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
