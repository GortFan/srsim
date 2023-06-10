// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/simimpact/srsim/pkg/engine (interfaces: Engine)

// Package mock is a generated GoMock package.
package mock

import (
	rand "math/rand"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	event "github.com/simimpact/srsim/pkg/engine/event"
	info "github.com/simimpact/srsim/pkg/engine/info"
	key "github.com/simimpact/srsim/pkg/key"
	model "github.com/simimpact/srsim/pkg/model"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// AddModifier mocks base method.
func (m *MockEngine) AddModifier(arg0 key.TargetID, arg1 info.Modifier) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddModifier", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddModifier indicates an expected call of AddModifier.
func (mr *MockEngineMockRecorder) AddModifier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddModifier", reflect.TypeOf((*MockEngine)(nil).AddModifier), arg0, arg1)
}

// AddShield mocks base method.
func (m *MockEngine) AddShield() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddShield")
}

// AddShield indicates an expected call of AddShield.
func (mr *MockEngineMockRecorder) AddShield() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddShield", reflect.TypeOf((*MockEngine)(nil).AddShield))
}

// AddTarget mocks base method.
func (m *MockEngine) AddTarget() key.TargetID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTarget")
	ret0, _ := ret[0].(key.TargetID)
	return ret0
}

// AddTarget indicates an expected call of AddTarget.
func (mr *MockEngineMockRecorder) AddTarget() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTarget", reflect.TypeOf((*MockEngine)(nil).AddTarget))
}

// AdjacentTo mocks base method.
func (m *MockEngine) AdjacentTo(arg0 key.TargetID) []key.TargetID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdjacentTo", arg0)
	ret0, _ := ret[0].([]key.TargetID)
	return ret0
}

// AdjacentTo indicates an expected call of AdjacentTo.
func (mr *MockEngineMockRecorder) AdjacentTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdjacentTo", reflect.TypeOf((*MockEngine)(nil).AdjacentTo), arg0)
}

// Attack mocks base method.
func (m *MockEngine) Attack(arg0 info.Attack) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Attack", arg0)
}

// Attack indicates an expected call of Attack.
func (mr *MockEngineMockRecorder) Attack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attack", reflect.TypeOf((*MockEngine)(nil).Attack), arg0)
}

// CharacterInfo mocks base method.
func (m *MockEngine) CharacterInfo(arg0 key.TargetID) (info.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharacterInfo", arg0)
	ret0, _ := ret[0].(info.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharacterInfo indicates an expected call of CharacterInfo.
func (mr *MockEngineMockRecorder) CharacterInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharacterInfo", reflect.TypeOf((*MockEngine)(nil).CharacterInfo), arg0)
}

// Characters mocks base method.
func (m *MockEngine) Characters() []key.TargetID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Characters")
	ret0, _ := ret[0].([]key.TargetID)
	return ret0
}

// Characters indicates an expected call of Characters.
func (mr *MockEngineMockRecorder) Characters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Characters", reflect.TypeOf((*MockEngine)(nil).Characters))
}

// Enemies mocks base method.
func (m *MockEngine) Enemies() []key.TargetID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enemies")
	ret0, _ := ret[0].([]key.TargetID)
	return ret0
}

// Enemies indicates an expected call of Enemies.
func (mr *MockEngineMockRecorder) Enemies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enemies", reflect.TypeOf((*MockEngine)(nil).Enemies))
}

// EnemyInfo mocks base method.
func (m *MockEngine) EnemyInfo(arg0 key.TargetID) (info.Enemy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnemyInfo", arg0)
	ret0, _ := ret[0].(info.Enemy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnemyInfo indicates an expected call of EnemyInfo.
func (mr *MockEngineMockRecorder) EnemyInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnemyInfo", reflect.TypeOf((*MockEngine)(nil).EnemyInfo), arg0)
}

// Events mocks base method.
func (m *MockEngine) Events() *event.System {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events")
	ret0, _ := ret[0].(*event.System)
	return ret0
}

// Events indicates an expected call of Events.
func (mr *MockEngineMockRecorder) Events() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockEngine)(nil).Events))
}

// ExtendModifierCount mocks base method.
func (m *MockEngine) ExtendModifierCount(arg0 key.TargetID, arg1 key.Modifier, arg2 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExtendModifierCount", arg0, arg1, arg2)
}

// ExtendModifierCount indicates an expected call of ExtendModifierCount.
func (mr *MockEngineMockRecorder) ExtendModifierCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendModifierCount", reflect.TypeOf((*MockEngine)(nil).ExtendModifierCount), arg0, arg1, arg2)
}

// ExtendModifierDuration mocks base method.
func (m *MockEngine) ExtendModifierDuration(arg0 key.TargetID, arg1 key.Modifier, arg2 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExtendModifierDuration", arg0, arg1, arg2)
}

// ExtendModifierDuration indicates an expected call of ExtendModifierDuration.
func (mr *MockEngineMockRecorder) ExtendModifierDuration(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendModifierDuration", reflect.TypeOf((*MockEngine)(nil).ExtendModifierDuration), arg0, arg1, arg2)
}

// HasBehaviorFlag mocks base method.
func (m *MockEngine) HasBehaviorFlag(arg0 key.TargetID, arg1 ...model.BehaviorFlag) bool {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasBehaviorFlag", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBehaviorFlag indicates an expected call of HasBehaviorFlag.
func (mr *MockEngineMockRecorder) HasBehaviorFlag(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBehaviorFlag", reflect.TypeOf((*MockEngine)(nil).HasBehaviorFlag), varargs...)
}

// HasModifier mocks base method.
func (m *MockEngine) HasModifier(arg0 key.TargetID, arg1 key.Modifier) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasModifier", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasModifier indicates an expected call of HasModifier.
func (mr *MockEngineMockRecorder) HasModifier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasModifier", reflect.TypeOf((*MockEngine)(nil).HasModifier), arg0, arg1)
}

// Heal mocks base method.
func (m *MockEngine) Heal(arg0 info.Heal) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Heal", arg0)
}

// Heal indicates an expected call of Heal.
func (mr *MockEngineMockRecorder) Heal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heal", reflect.TypeOf((*MockEngine)(nil).Heal), arg0)
}

// InsertAbility mocks base method.
func (m *MockEngine) InsertAbility(arg0 info.Insert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertAbility", arg0)
}

// InsertAbility indicates an expected call of InsertAbility.
func (mr *MockEngineMockRecorder) InsertAbility(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAbility", reflect.TypeOf((*MockEngine)(nil).InsertAbility), arg0)
}

// InsertAction mocks base method.
func (m *MockEngine) InsertAction(arg0 key.TargetID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAction indicates an expected call of InsertAction.
func (mr *MockEngineMockRecorder) InsertAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAction", reflect.TypeOf((*MockEngine)(nil).InsertAction), arg0)
}

// IsCharacter mocks base method.
func (m *MockEngine) IsCharacter(arg0 key.TargetID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCharacter", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCharacter indicates an expected call of IsCharacter.
func (mr *MockEngineMockRecorder) IsCharacter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCharacter", reflect.TypeOf((*MockEngine)(nil).IsCharacter), arg0)
}

// IsEnemy mocks base method.
func (m *MockEngine) IsEnemy(arg0 key.TargetID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnemy", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEnemy indicates an expected call of IsEnemy.
func (mr *MockEngineMockRecorder) IsEnemy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnemy", reflect.TypeOf((*MockEngine)(nil).IsEnemy), arg0)
}

// IsValid mocks base method.
func (m *MockEngine) IsValid(arg0 key.TargetID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid.
func (mr *MockEngineMockRecorder) IsValid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockEngine)(nil).IsValid), arg0)
}

// ModifierCount mocks base method.
func (m *MockEngine) ModifierCount(arg0 key.TargetID, arg1 model.StatusType) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifierCount", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// ModifierCount indicates an expected call of ModifierCount.
func (mr *MockEngineMockRecorder) ModifierCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifierCount", reflect.TypeOf((*MockEngine)(nil).ModifierCount), arg0, arg1)
}

// ModifyCurrentGaugeCost mocks base method.
func (m *MockEngine) ModifyCurrentGaugeCost(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ModifyCurrentGaugeCost", arg0)
}

// ModifyCurrentGaugeCost indicates an expected call of ModifyCurrentGaugeCost.
func (mr *MockEngineMockRecorder) ModifyCurrentGaugeCost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyCurrentGaugeCost", reflect.TypeOf((*MockEngine)(nil).ModifyCurrentGaugeCost), arg0)
}

// ModifyEnergy mocks base method.
func (m *MockEngine) ModifyEnergy(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyEnergy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyEnergy indicates an expected call of ModifyEnergy.
func (mr *MockEngineMockRecorder) ModifyEnergy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyEnergy", reflect.TypeOf((*MockEngine)(nil).ModifyEnergy), arg0, arg1)
}

// ModifyEnergyFixed mocks base method.
func (m *MockEngine) ModifyEnergyFixed(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyEnergyFixed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyEnergyFixed indicates an expected call of ModifyEnergyFixed.
func (mr *MockEngineMockRecorder) ModifyEnergyFixed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyEnergyFixed", reflect.TypeOf((*MockEngine)(nil).ModifyEnergyFixed), arg0, arg1)
}

// ModifyGaugeAV mocks base method.
func (m *MockEngine) ModifyGaugeAV(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyGaugeAV", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyGaugeAV indicates an expected call of ModifyGaugeAV.
func (mr *MockEngineMockRecorder) ModifyGaugeAV(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyGaugeAV", reflect.TypeOf((*MockEngine)(nil).ModifyGaugeAV), arg0, arg1)
}

// ModifyGaugeNormalized mocks base method.
func (m *MockEngine) ModifyGaugeNormalized(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyGaugeNormalized", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyGaugeNormalized indicates an expected call of ModifyGaugeNormalized.
func (mr *MockEngineMockRecorder) ModifyGaugeNormalized(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyGaugeNormalized", reflect.TypeOf((*MockEngine)(nil).ModifyGaugeNormalized), arg0, arg1)
}

// ModifyHPByAmount mocks base method.
func (m *MockEngine) ModifyHPByAmount(arg0, arg1 key.TargetID, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyHPByAmount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyHPByAmount indicates an expected call of ModifyHPByAmount.
func (mr *MockEngineMockRecorder) ModifyHPByAmount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyHPByAmount", reflect.TypeOf((*MockEngine)(nil).ModifyHPByAmount), arg0, arg1, arg2)
}

// ModifyHPByRatio mocks base method.
func (m *MockEngine) ModifyHPByRatio(arg0, arg1 key.TargetID, arg2 info.ModifyHPByRatio) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyHPByRatio", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyHPByRatio indicates an expected call of ModifyHPByRatio.
func (mr *MockEngineMockRecorder) ModifyHPByRatio(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyHPByRatio", reflect.TypeOf((*MockEngine)(nil).ModifyHPByRatio), arg0, arg1, arg2)
}

// ModifyStance mocks base method.
func (m *MockEngine) ModifyStance(arg0, arg1 key.TargetID, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyStance", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyStance indicates an expected call of ModifyStance.
func (mr *MockEngineMockRecorder) ModifyStance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyStance", reflect.TypeOf((*MockEngine)(nil).ModifyStance), arg0, arg1, arg2)
}

// Neutrals mocks base method.
func (m *MockEngine) Neutrals() []key.TargetID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Neutrals")
	ret0, _ := ret[0].([]key.TargetID)
	return ret0
}

// Neutrals indicates an expected call of Neutrals.
func (mr *MockEngineMockRecorder) Neutrals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Neutrals", reflect.TypeOf((*MockEngine)(nil).Neutrals))
}

// Rand mocks base method.
func (m *MockEngine) Rand() *rand.Rand {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rand")
	ret0, _ := ret[0].(*rand.Rand)
	return ret0
}

// Rand indicates an expected call of Rand.
func (mr *MockEngineMockRecorder) Rand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rand", reflect.TypeOf((*MockEngine)(nil).Rand))
}

// RemoveModifier mocks base method.
func (m *MockEngine) RemoveModifier(arg0 key.TargetID, arg1 key.Modifier) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveModifier", arg0, arg1)
}

// RemoveModifier indicates an expected call of RemoveModifier.
func (mr *MockEngineMockRecorder) RemoveModifier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveModifier", reflect.TypeOf((*MockEngine)(nil).RemoveModifier), arg0, arg1)
}

// RemoveModifierFromSource mocks base method.
func (m *MockEngine) RemoveModifierFromSource(arg0, arg1 key.TargetID, arg2 key.Modifier) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveModifierFromSource", arg0, arg1, arg2)
}

// RemoveModifierFromSource indicates an expected call of RemoveModifierFromSource.
func (mr *MockEngineMockRecorder) RemoveModifierFromSource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveModifierFromSource", reflect.TypeOf((*MockEngine)(nil).RemoveModifierFromSource), arg0, arg1, arg2)
}

// RemoveShield mocks base method.
func (m *MockEngine) RemoveShield() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveShield")
}

// RemoveShield indicates an expected call of RemoveShield.
func (mr *MockEngineMockRecorder) RemoveShield() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveShield", reflect.TypeOf((*MockEngine)(nil).RemoveShield))
}

// SetCurrentGaugeCost mocks base method.
func (m *MockEngine) SetCurrentGaugeCost(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentGaugeCost", arg0)
}

// SetCurrentGaugeCost indicates an expected call of SetCurrentGaugeCost.
func (mr *MockEngineMockRecorder) SetCurrentGaugeCost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentGaugeCost", reflect.TypeOf((*MockEngine)(nil).SetCurrentGaugeCost), arg0)
}

// SetGauge mocks base method.
func (m *MockEngine) SetGauge(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGauge", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGauge indicates an expected call of SetGauge.
func (mr *MockEngineMockRecorder) SetGauge(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGauge", reflect.TypeOf((*MockEngine)(nil).SetGauge), arg0, arg1)
}

// SetHP mocks base method.
func (m *MockEngine) SetHP(arg0, arg1 key.TargetID, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHP", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHP indicates an expected call of SetHP.
func (mr *MockEngineMockRecorder) SetHP(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHP", reflect.TypeOf((*MockEngine)(nil).SetHP), arg0, arg1, arg2)
}

// Stats mocks base method.
func (m *MockEngine) Stats(arg0 key.TargetID) *info.Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats", arg0)
	ret0, _ := ret[0].(*info.Stats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockEngineMockRecorder) Stats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockEngine)(nil).Stats), arg0)
}
