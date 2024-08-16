// Code generated by mockery 2.9.4. DO NOT EDIT.

package autocodesign

import (
	appstoreconnect "github.com/teamlapse/go-xcode/v2/autocodesign/devportalclient/appstoreconnect"
	mock "github.com/stretchr/testify/mock"
)

// MockLocalCodeSignAssetManager is an autogenerated mock type for the LocalCodeSignAssetManager type
type MockLocalCodeSignAssetManager struct {
	mock.Mock
}

// FindCodesignAssets provides a mock function with given fields: appLayout, distrType, certsByType, deviceIDs, minProfileDaysValid
func (_m *MockLocalCodeSignAssetManager) FindCodesignAssets(appLayout AppLayout, distrType DistributionType, certsByType map[appstoreconnect.CertificateType][]Certificate, deviceIDs []string, minProfileDaysValid int) (*AppCodesignAssets, *AppLayout, error) {
	ret := _m.Called(appLayout, distrType, certsByType, deviceIDs, minProfileDaysValid)

	var r0 *AppCodesignAssets
	if rf, ok := ret.Get(0).(func(AppLayout, DistributionType, map[appstoreconnect.CertificateType][]Certificate, []string, int) *AppCodesignAssets); ok {
		r0 = rf(appLayout, distrType, certsByType, deviceIDs, minProfileDaysValid)
	} else {
		if ret.Get(0) != nil {
			r0, ok = ret.Get(0).(*AppCodesignAssets)
			if !ok {
			}
		}
	}

	var r1 *AppLayout
	if rf, ok := ret.Get(1).(func(AppLayout, DistributionType, map[appstoreconnect.CertificateType][]Certificate, []string, int) *AppLayout); ok {
		r1 = rf(appLayout, distrType, certsByType, deviceIDs, minProfileDaysValid)
	} else {
		if ret.Get(1) != nil {
			r1, ok = ret.Get(1).(*AppLayout)
			if !ok {
			}
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(AppLayout, DistributionType, map[appstoreconnect.CertificateType][]Certificate, []string, int) error); ok {
		r2 = rf(appLayout, distrType, certsByType, deviceIDs, minProfileDaysValid)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
