// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/providers (interfaces: Provider,DatacenterConfig,MachineConfig)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	bootstrapper "github.com/aws/eks-anywhere/pkg/bootstrapper"
	cluster "github.com/aws/eks-anywhere/pkg/cluster"
	providers "github.com/aws/eks-anywhere/pkg/providers"
	types "github.com/aws/eks-anywhere/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// BootstrapClusterOpts mocks base method.
func (m *MockProvider) BootstrapClusterOpts(arg0 *cluster.Spec) ([]bootstrapper.BootstrapClusterOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapClusterOpts", arg0)
	ret0, _ := ret[0].([]bootstrapper.BootstrapClusterOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BootstrapClusterOpts indicates an expected call of BootstrapClusterOpts.
func (mr *MockProviderMockRecorder) BootstrapClusterOpts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapClusterOpts", reflect.TypeOf((*MockProvider)(nil).BootstrapClusterOpts), arg0)
}

// ChangeDiff mocks base method.
func (m *MockProvider) ChangeDiff(arg0, arg1 *cluster.Spec) *types.ComponentChangeDiff {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDiff", arg0, arg1)
	ret0, _ := ret[0].(*types.ComponentChangeDiff)
	return ret0
}

// ChangeDiff indicates an expected call of ChangeDiff.
func (mr *MockProviderMockRecorder) ChangeDiff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDiff", reflect.TypeOf((*MockProvider)(nil).ChangeDiff), arg0, arg1)
}

// DatacenterConfig mocks base method.
func (m *MockProvider) DatacenterConfig(arg0 *cluster.Spec) providers.DatacenterConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatacenterConfig", arg0)
	ret0, _ := ret[0].(providers.DatacenterConfig)
	return ret0
}

// DatacenterConfig indicates an expected call of DatacenterConfig.
func (mr *MockProviderMockRecorder) DatacenterConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatacenterConfig", reflect.TypeOf((*MockProvider)(nil).DatacenterConfig), arg0)
}

// DatacenterResourceType mocks base method.
func (m *MockProvider) DatacenterResourceType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatacenterResourceType")
	ret0, _ := ret[0].(string)
	return ret0
}

// DatacenterResourceType indicates an expected call of DatacenterResourceType.
func (mr *MockProviderMockRecorder) DatacenterResourceType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatacenterResourceType", reflect.TypeOf((*MockProvider)(nil).DatacenterResourceType))
}

// DeleteResources mocks base method.
func (m *MockProvider) DeleteResources(arg0 context.Context, arg1 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResources indicates an expected call of DeleteResources.
func (mr *MockProviderMockRecorder) DeleteResources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResources", reflect.TypeOf((*MockProvider)(nil).DeleteResources), arg0, arg1)
}

// EnvMap mocks base method.
func (m *MockProvider) EnvMap(arg0 *cluster.Spec) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvMap", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvMap indicates an expected call of EnvMap.
func (mr *MockProviderMockRecorder) EnvMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvMap", reflect.TypeOf((*MockProvider)(nil).EnvMap), arg0)
}

// GenerateCAPISpecForCreate mocks base method.
func (m *MockProvider) GenerateCAPISpecForCreate(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) ([]byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCAPISpecForCreate", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateCAPISpecForCreate indicates an expected call of GenerateCAPISpecForCreate.
func (mr *MockProviderMockRecorder) GenerateCAPISpecForCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCAPISpecForCreate", reflect.TypeOf((*MockProvider)(nil).GenerateCAPISpecForCreate), arg0, arg1, arg2)
}

// GenerateCAPISpecForUpgrade mocks base method.
func (m *MockProvider) GenerateCAPISpecForUpgrade(arg0 context.Context, arg1, arg2 *types.Cluster, arg3, arg4 *cluster.Spec) ([]byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCAPISpecForUpgrade", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateCAPISpecForUpgrade indicates an expected call of GenerateCAPISpecForUpgrade.
func (mr *MockProviderMockRecorder) GenerateCAPISpecForUpgrade(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCAPISpecForUpgrade", reflect.TypeOf((*MockProvider)(nil).GenerateCAPISpecForUpgrade), arg0, arg1, arg2, arg3, arg4)
}

// GenerateStorageClass mocks base method.
func (m *MockProvider) GenerateStorageClass() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateStorageClass")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GenerateStorageClass indicates an expected call of GenerateStorageClass.
func (mr *MockProviderMockRecorder) GenerateStorageClass() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateStorageClass", reflect.TypeOf((*MockProvider)(nil).GenerateStorageClass))
}

// GetDeployments mocks base method.
func (m *MockProvider) GetDeployments() map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployments")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// GetDeployments indicates an expected call of GetDeployments.
func (mr *MockProviderMockRecorder) GetDeployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockProvider)(nil).GetDeployments))
}

// GetInfrastructureBundle mocks base method.
func (m *MockProvider) GetInfrastructureBundle(arg0 *cluster.Spec) *types.InfrastructureBundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfrastructureBundle", arg0)
	ret0, _ := ret[0].(*types.InfrastructureBundle)
	return ret0
}

// GetInfrastructureBundle indicates an expected call of GetInfrastructureBundle.
func (mr *MockProviderMockRecorder) GetInfrastructureBundle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfrastructureBundle", reflect.TypeOf((*MockProvider)(nil).GetInfrastructureBundle), arg0)
}

// InstallCustomProviderComponents mocks base method.
func (m *MockProvider) InstallCustomProviderComponents(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallCustomProviderComponents", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallCustomProviderComponents indicates an expected call of InstallCustomProviderComponents.
func (mr *MockProviderMockRecorder) InstallCustomProviderComponents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallCustomProviderComponents", reflect.TypeOf((*MockProvider)(nil).InstallCustomProviderComponents), arg0, arg1)
}

// MachineConfigs mocks base method.
func (m *MockProvider) MachineConfigs(arg0 *cluster.Spec) []providers.MachineConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineConfigs", arg0)
	ret0, _ := ret[0].([]providers.MachineConfig)
	return ret0
}

// MachineConfigs indicates an expected call of MachineConfigs.
func (mr *MockProviderMockRecorder) MachineConfigs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineConfigs", reflect.TypeOf((*MockProvider)(nil).MachineConfigs), arg0)
}

// MachineResourceType mocks base method.
func (m *MockProvider) MachineResourceType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineResourceType")
	ret0, _ := ret[0].(string)
	return ret0
}

// MachineResourceType indicates an expected call of MachineResourceType.
func (mr *MockProviderMockRecorder) MachineResourceType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineResourceType", reflect.TypeOf((*MockProvider)(nil).MachineResourceType))
}

// Name mocks base method.
func (m *MockProvider) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockProviderMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockProvider)(nil).Name))
}

// PostBootstrapDeleteForUpgrade mocks base method.
func (m *MockProvider) PostBootstrapDeleteForUpgrade(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBootstrapDeleteForUpgrade", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostBootstrapDeleteForUpgrade indicates an expected call of PostBootstrapDeleteForUpgrade.
func (mr *MockProviderMockRecorder) PostBootstrapDeleteForUpgrade(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBootstrapDeleteForUpgrade", reflect.TypeOf((*MockProvider)(nil).PostBootstrapDeleteForUpgrade), arg0)
}

// PostBootstrapSetup mocks base method.
func (m *MockProvider) PostBootstrapSetup(arg0 context.Context, arg1 *v1alpha1.Cluster, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBootstrapSetup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostBootstrapSetup indicates an expected call of PostBootstrapSetup.
func (mr *MockProviderMockRecorder) PostBootstrapSetup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBootstrapSetup", reflect.TypeOf((*MockProvider)(nil).PostBootstrapSetup), arg0, arg1, arg2)
}

// PostBootstrapSetupUpgrade mocks base method.
func (m *MockProvider) PostBootstrapSetupUpgrade(arg0 context.Context, arg1 *v1alpha1.Cluster, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBootstrapSetupUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostBootstrapSetupUpgrade indicates an expected call of PostBootstrapSetupUpgrade.
func (mr *MockProviderMockRecorder) PostBootstrapSetupUpgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBootstrapSetupUpgrade", reflect.TypeOf((*MockProvider)(nil).PostBootstrapSetupUpgrade), arg0, arg1, arg2)
}

// PostClusterDeleteValidate mocks base method.
func (m *MockProvider) PostClusterDeleteValidate(arg0 context.Context, arg1 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostClusterDeleteValidate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostClusterDeleteValidate indicates an expected call of PostClusterDeleteValidate.
func (mr *MockProviderMockRecorder) PostClusterDeleteValidate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostClusterDeleteValidate", reflect.TypeOf((*MockProvider)(nil).PostClusterDeleteValidate), arg0, arg1)
}

// PostMoveManagementToBootstrap mocks base method.
func (m *MockProvider) PostMoveManagementToBootstrap(arg0 context.Context, arg1 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMoveManagementToBootstrap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMoveManagementToBootstrap indicates an expected call of PostMoveManagementToBootstrap.
func (mr *MockProviderMockRecorder) PostMoveManagementToBootstrap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMoveManagementToBootstrap", reflect.TypeOf((*MockProvider)(nil).PostMoveManagementToBootstrap), arg0, arg1)
}

// PostWorkloadInit mocks base method.
func (m *MockProvider) PostWorkloadInit(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostWorkloadInit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostWorkloadInit indicates an expected call of PostWorkloadInit.
func (mr *MockProviderMockRecorder) PostWorkloadInit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostWorkloadInit", reflect.TypeOf((*MockProvider)(nil).PostWorkloadInit), arg0, arg1, arg2)
}

// PreCAPIInstallOnBootstrap mocks base method.
func (m *MockProvider) PreCAPIInstallOnBootstrap(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreCAPIInstallOnBootstrap", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreCAPIInstallOnBootstrap indicates an expected call of PreCAPIInstallOnBootstrap.
func (mr *MockProviderMockRecorder) PreCAPIInstallOnBootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreCAPIInstallOnBootstrap", reflect.TypeOf((*MockProvider)(nil).PreCAPIInstallOnBootstrap), arg0, arg1, arg2)
}

// RunPostControlPlaneUpgrade mocks base method.
func (m *MockProvider) RunPostControlPlaneUpgrade(arg0 context.Context, arg1, arg2 *cluster.Spec, arg3, arg4 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunPostControlPlaneUpgrade", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunPostControlPlaneUpgrade indicates an expected call of RunPostControlPlaneUpgrade.
func (mr *MockProviderMockRecorder) RunPostControlPlaneUpgrade(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunPostControlPlaneUpgrade", reflect.TypeOf((*MockProvider)(nil).RunPostControlPlaneUpgrade), arg0, arg1, arg2, arg3, arg4)
}

// SetupAndValidateCreateCluster mocks base method.
func (m *MockProvider) SetupAndValidateCreateCluster(arg0 context.Context, arg1 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupAndValidateCreateCluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupAndValidateCreateCluster indicates an expected call of SetupAndValidateCreateCluster.
func (mr *MockProviderMockRecorder) SetupAndValidateCreateCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupAndValidateCreateCluster", reflect.TypeOf((*MockProvider)(nil).SetupAndValidateCreateCluster), arg0, arg1)
}

// SetupAndValidateDeleteCluster mocks base method.
func (m *MockProvider) SetupAndValidateDeleteCluster(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupAndValidateDeleteCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupAndValidateDeleteCluster indicates an expected call of SetupAndValidateDeleteCluster.
func (mr *MockProviderMockRecorder) SetupAndValidateDeleteCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupAndValidateDeleteCluster", reflect.TypeOf((*MockProvider)(nil).SetupAndValidateDeleteCluster), arg0, arg1, arg2)
}

// SetupAndValidateUpgradeCluster mocks base method.
func (m *MockProvider) SetupAndValidateUpgradeCluster(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupAndValidateUpgradeCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupAndValidateUpgradeCluster indicates an expected call of SetupAndValidateUpgradeCluster.
func (mr *MockProviderMockRecorder) SetupAndValidateUpgradeCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupAndValidateUpgradeCluster", reflect.TypeOf((*MockProvider)(nil).SetupAndValidateUpgradeCluster), arg0, arg1, arg2, arg3)
}

// UpdateKubeConfig mocks base method.
func (m *MockProvider) UpdateKubeConfig(arg0 *[]byte, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKubeConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateKubeConfig indicates an expected call of UpdateKubeConfig.
func (mr *MockProviderMockRecorder) UpdateKubeConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKubeConfig", reflect.TypeOf((*MockProvider)(nil).UpdateKubeConfig), arg0, arg1)
}

// UpdateSecrets mocks base method.
func (m *MockProvider) UpdateSecrets(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecrets indicates an expected call of UpdateSecrets.
func (mr *MockProviderMockRecorder) UpdateSecrets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecrets", reflect.TypeOf((*MockProvider)(nil).UpdateSecrets), arg0, arg1, arg2)
}

// UpgradeNeeded mocks base method.
func (m *MockProvider) UpgradeNeeded(arg0 context.Context, arg1, arg2 *cluster.Spec, arg3 *types.Cluster) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeNeeded", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeNeeded indicates an expected call of UpgradeNeeded.
func (mr *MockProviderMockRecorder) UpgradeNeeded(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeNeeded", reflect.TypeOf((*MockProvider)(nil).UpgradeNeeded), arg0, arg1, arg2, arg3)
}

// ValidateNewSpec mocks base method.
func (m *MockProvider) ValidateNewSpec(arg0 context.Context, arg1 *types.Cluster, arg2 *cluster.Spec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateNewSpec", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateNewSpec indicates an expected call of ValidateNewSpec.
func (mr *MockProviderMockRecorder) ValidateNewSpec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateNewSpec", reflect.TypeOf((*MockProvider)(nil).ValidateNewSpec), arg0, arg1, arg2)
}

// Version mocks base method.
func (m *MockProvider) Version(arg0 *cluster.Spec) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockProviderMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockProvider)(nil).Version), arg0)
}

// MockDatacenterConfig is a mock of DatacenterConfig interface.
type MockDatacenterConfig struct {
	ctrl     *gomock.Controller
	recorder *MockDatacenterConfigMockRecorder
}

// MockDatacenterConfigMockRecorder is the mock recorder for MockDatacenterConfig.
type MockDatacenterConfigMockRecorder struct {
	mock *MockDatacenterConfig
}

// NewMockDatacenterConfig creates a new mock instance.
func NewMockDatacenterConfig(ctrl *gomock.Controller) *MockDatacenterConfig {
	mock := &MockDatacenterConfig{ctrl: ctrl}
	mock.recorder = &MockDatacenterConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatacenterConfig) EXPECT() *MockDatacenterConfigMockRecorder {
	return m.recorder
}

// ClearPauseAnnotation mocks base method.
func (m *MockDatacenterConfig) ClearPauseAnnotation() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearPauseAnnotation")
}

// ClearPauseAnnotation indicates an expected call of ClearPauseAnnotation.
func (mr *MockDatacenterConfigMockRecorder) ClearPauseAnnotation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearPauseAnnotation", reflect.TypeOf((*MockDatacenterConfig)(nil).ClearPauseAnnotation))
}

// Kind mocks base method.
func (m *MockDatacenterConfig) Kind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockDatacenterConfigMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockDatacenterConfig)(nil).Kind))
}

// Marshallable mocks base method.
func (m *MockDatacenterConfig) Marshallable() v1alpha1.Marshallable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshallable")
	ret0, _ := ret[0].(v1alpha1.Marshallable)
	return ret0
}

// Marshallable indicates an expected call of Marshallable.
func (mr *MockDatacenterConfigMockRecorder) Marshallable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshallable", reflect.TypeOf((*MockDatacenterConfig)(nil).Marshallable))
}

// PauseReconcile mocks base method.
func (m *MockDatacenterConfig) PauseReconcile() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PauseReconcile")
}

// PauseReconcile indicates an expected call of PauseReconcile.
func (mr *MockDatacenterConfigMockRecorder) PauseReconcile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseReconcile", reflect.TypeOf((*MockDatacenterConfig)(nil).PauseReconcile))
}

// MockMachineConfig is a mock of MachineConfig interface.
type MockMachineConfig struct {
	ctrl     *gomock.Controller
	recorder *MockMachineConfigMockRecorder
}

// MockMachineConfigMockRecorder is the mock recorder for MockMachineConfig.
type MockMachineConfigMockRecorder struct {
	mock *MockMachineConfig
}

// NewMockMachineConfig creates a new mock instance.
func NewMockMachineConfig(ctrl *gomock.Controller) *MockMachineConfig {
	mock := &MockMachineConfig{ctrl: ctrl}
	mock.recorder = &MockMachineConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineConfig) EXPECT() *MockMachineConfigMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockMachineConfig) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockMachineConfigMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockMachineConfig)(nil).GetName))
}

// GetNamespace mocks base method.
func (m *MockMachineConfig) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockMachineConfigMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockMachineConfig)(nil).GetNamespace))
}

// Marshallable mocks base method.
func (m *MockMachineConfig) Marshallable() v1alpha1.Marshallable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshallable")
	ret0, _ := ret[0].(v1alpha1.Marshallable)
	return ret0
}

// Marshallable indicates an expected call of Marshallable.
func (mr *MockMachineConfigMockRecorder) Marshallable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshallable", reflect.TypeOf((*MockMachineConfig)(nil).Marshallable))
}

// OSFamily mocks base method.
func (m *MockMachineConfig) OSFamily() v1alpha1.OSFamily {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OSFamily")
	ret0, _ := ret[0].(v1alpha1.OSFamily)
	return ret0
}

// OSFamily indicates an expected call of OSFamily.
func (mr *MockMachineConfigMockRecorder) OSFamily() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OSFamily", reflect.TypeOf((*MockMachineConfig)(nil).OSFamily))
}
