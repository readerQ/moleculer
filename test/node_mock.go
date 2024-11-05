package test

import "time"

type NodeMock struct {
	UpdateResult          bool
	ID                    string
	IncreaseSequenceCalls int
	HeartBeatCalls        int
	ExportAsMapResult     map[string]interface{}
	IsAvailableResult     bool
	IsExpiredResult       bool
	PublishCalls          int
}

func (node *NodeMock) Update(id string, info map[string]interface{}) (bool, []map[string]interface{}) {
	return node.UpdateResult, []map[string]interface{}{}
}

func (node *NodeMock) Unavailable() {
	node.IsAvailableResult = false
}
func (node *NodeMock) Available() {
	node.IsAvailableResult = true
}

func (node *NodeMock) GetID() string {
	return node.ID
}

func (node *NodeMock) IncreaseSequence() {
	node.IncreaseSequenceCalls++
}

func (node *NodeMock) ExportAsMap() map[string]interface{} {
	return node.ExportAsMapResult
}
func (node *NodeMock) IsAvailable() bool {
	return node.IsAvailableResult
}
func (node *NodeMock) HeartBeat(heartbeat map[string]interface{}) {
	node.HeartBeatCalls++
}
func (node *NodeMock) IsExpired(timeout time.Duration) bool {
	return node.IsExpiredResult
}
func (node *NodeMock) Publish(service map[string]interface{}) {
	node.PublishCalls++
}

func (node *NodeMock) GetCpu() int64 {
	return 0
}
func (node *NodeMock) GetCpuSequence() int64 {
	return 0
}

func (node *NodeMock) GetHost() string {
	return "localhost"
}
func (node *NodeMock) GetHostname() string {
	return "localhost"
}
func (node *NodeMock) GetIpList() []string {
	return []string{"127.0.0.1"}
}
func (node *NodeMock) GetPort() int {
	return 5103
}
func (node *NodeMock) GetSequence() int64 {
	return int64(node.IncreaseSequenceCalls)
}
func (node *NodeMock) GetUdpAddress() string {
	return ""
}
func (node *NodeMock) IsLocal() bool {
	return true
}
func (node *NodeMock) UpdateInfo(info map[string]interface{}) []map[string]interface{} {
	return nil
}
func (node *NodeMock) UpdateMetrics() {

}
