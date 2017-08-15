// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insonmnia.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskStatusReply_Status int32

const (
	TaskStatusReply_UNKNOWN  TaskStatusReply_Status = 0
	TaskStatusReply_SPOOLING TaskStatusReply_Status = 1
	TaskStatusReply_SPAWNING TaskStatusReply_Status = 2
	TaskStatusReply_RUNNING  TaskStatusReply_Status = 3
	TaskStatusReply_FINISHED TaskStatusReply_Status = 4
	TaskStatusReply_BROKEN   TaskStatusReply_Status = 5
)

var TaskStatusReply_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPOOLING",
	2: "SPAWNING",
	3: "RUNNING",
	4: "FINISHED",
	5: "BROKEN",
}
var TaskStatusReply_Status_value = map[string]int32{
	"UNKNOWN":  0,
	"SPOOLING": 1,
	"SPAWNING": 2,
	"RUNNING":  3,
	"FINISHED": 4,
	"BROKEN":   5,
}

func (x TaskStatusReply_Status) String() string {
	return proto.EnumName(TaskStatusReply_Status_name, int32(x))
}
func (TaskStatusReply_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{6, 0} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type PingReply struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PingReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type InfoReply struct {
	Stats map[string]*InfoReplyStats `protobuf:"bytes,1,rep,name=Stats" json:"Stats,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *InfoReply) Reset()                    { *m = InfoReply{} }
func (m *InfoReply) String() string            { return proto.CompactTextString(m) }
func (*InfoReply) ProtoMessage()               {}
func (*InfoReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *InfoReply) GetStats() map[string]*InfoReplyStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type InfoReplyStats struct {
	CPU    *InfoReplyStatsCpu    `protobuf:"bytes,1,opt,name=CPU" json:"CPU,omitempty"`
	Memory *InfoReplyStatsMemory `protobuf:"bytes,2,opt,name=Memory" json:"Memory,omitempty"`
}

func (m *InfoReplyStats) Reset()                    { *m = InfoReplyStats{} }
func (m *InfoReplyStats) String() string            { return proto.CompactTextString(m) }
func (*InfoReplyStats) ProtoMessage()               {}
func (*InfoReplyStats) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

func (m *InfoReplyStats) GetCPU() *InfoReplyStatsCpu {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *InfoReplyStats) GetMemory() *InfoReplyStatsMemory {
	if m != nil {
		return m.Memory
	}
	return nil
}

type InfoReplyStatsCpu struct {
	TotalUsage uint64 `protobuf:"varint,1,opt,name=totalUsage" json:"totalUsage,omitempty"`
}

func (m *InfoReplyStatsCpu) Reset()                    { *m = InfoReplyStatsCpu{} }
func (m *InfoReplyStatsCpu) String() string            { return proto.CompactTextString(m) }
func (*InfoReplyStatsCpu) ProtoMessage()               {}
func (*InfoReplyStatsCpu) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0, 0} }

func (m *InfoReplyStatsCpu) GetTotalUsage() uint64 {
	if m != nil {
		return m.TotalUsage
	}
	return 0
}

type InfoReplyStatsMemory struct {
	MaxUsage uint64 `protobuf:"varint,1,opt,name=maxUsage" json:"maxUsage,omitempty"`
}

func (m *InfoReplyStatsMemory) Reset()                    { *m = InfoReplyStatsMemory{} }
func (m *InfoReplyStatsMemory) String() string            { return proto.CompactTextString(m) }
func (*InfoReplyStatsMemory) ProtoMessage()               {}
func (*InfoReplyStatsMemory) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0, 1} }

func (m *InfoReplyStatsMemory) GetMaxUsage() uint64 {
	if m != nil {
		return m.MaxUsage
	}
	return 0
}

type StopTaskRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StopTaskRequest) Reset()                    { *m = StopTaskRequest{} }
func (m *StopTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*StopTaskRequest) ProtoMessage()               {}
func (*StopTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *StopTaskRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StopTaskReply struct {
}

func (m *StopTaskReply) Reset()                    { *m = StopTaskReply{} }
func (m *StopTaskReply) String() string            { return proto.CompactTextString(m) }
func (*StopTaskReply) ProtoMessage()               {}
func (*StopTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type TaskStatusRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TaskStatusRequest) Reset()                    { *m = TaskStatusRequest{} }
func (m *TaskStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskStatusRequest) ProtoMessage()               {}
func (*TaskStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *TaskStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TaskStatusReply struct {
	Status TaskStatusReply_Status `protobuf:"varint,1,opt,name=status,enum=sonm.TaskStatusReply_Status" json:"status,omitempty"`
}

func (m *TaskStatusReply) Reset()                    { *m = TaskStatusReply{} }
func (m *TaskStatusReply) String() string            { return proto.CompactTextString(m) }
func (*TaskStatusReply) ProtoMessage()               {}
func (*TaskStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *TaskStatusReply) GetStatus() TaskStatusReply_Status {
	if m != nil {
		return m.Status
	}
	return TaskStatusReply_UNKNOWN
}

type StatusMapReply struct {
	Statuses map[string]*TaskStatusReply `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *StatusMapReply) Reset()                    { *m = StatusMapReply{} }
func (m *StatusMapReply) String() string            { return proto.CompactTextString(m) }
func (*StatusMapReply) ProtoMessage()               {}
func (*StatusMapReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *StatusMapReply) GetStatuses() map[string]*TaskStatusReply {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type ContainerResources struct {
	Memory   int64 `protobuf:"varint,1,opt,name=memory" json:"memory,omitempty"`
	NanoCPUs int64 `protobuf:"varint,2,opt,name=nanoCPUs" json:"nanoCPUs,omitempty"`
}

func (m *ContainerResources) Reset()                    { *m = ContainerResources{} }
func (m *ContainerResources) String() string            { return proto.CompactTextString(m) }
func (*ContainerResources) ProtoMessage()               {}
func (*ContainerResources) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ContainerResources) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ContainerResources) GetNanoCPUs() int64 {
	if m != nil {
		return m.NanoCPUs
	}
	return 0
}

type ContainerRestartPolicy struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MaximumRetryCount uint32 `protobuf:"varint,2,opt,name=maximumRetryCount" json:"maximumRetryCount,omitempty"`
}

func (m *ContainerRestartPolicy) Reset()                    { *m = ContainerRestartPolicy{} }
func (m *ContainerRestartPolicy) String() string            { return proto.CompactTextString(m) }
func (*ContainerRestartPolicy) ProtoMessage()               {}
func (*ContainerRestartPolicy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *ContainerRestartPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRestartPolicy) GetMaximumRetryCount() uint32 {
	if m != nil {
		return m.MaximumRetryCount
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "sonm.PingRequest")
	proto.RegisterType((*PingReply)(nil), "sonm.PingReply")
	proto.RegisterType((*InfoReply)(nil), "sonm.InfoReply")
	proto.RegisterType((*InfoReplyStats)(nil), "sonm.InfoReply.stats")
	proto.RegisterType((*InfoReplyStatsCpu)(nil), "sonm.InfoReply.stats.cpu")
	proto.RegisterType((*InfoReplyStatsMemory)(nil), "sonm.InfoReply.stats.memory")
	proto.RegisterType((*StopTaskRequest)(nil), "sonm.StopTaskRequest")
	proto.RegisterType((*StopTaskReply)(nil), "sonm.StopTaskReply")
	proto.RegisterType((*TaskStatusRequest)(nil), "sonm.TaskStatusRequest")
	proto.RegisterType((*TaskStatusReply)(nil), "sonm.TaskStatusReply")
	proto.RegisterType((*StatusMapReply)(nil), "sonm.StatusMapReply")
	proto.RegisterType((*ContainerResources)(nil), "sonm.ContainerResources")
	proto.RegisterType((*ContainerRestartPolicy)(nil), "sonm.ContainerRestartPolicy")
	proto.RegisterEnum("sonm.TaskStatusReply_Status", TaskStatusReply_Status_name, TaskStatusReply_Status_value)
}

func init() { proto.RegisterFile("insonmnia.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0x4d, 0xb2, 0x1b, 0x77, 0xcf, 0xba, 0x1f, 0x1d, 0xb0, 0xd4, 0x20, 0xa2, 0xa9, 0x42,
	0xa1, 0x12, 0xa4, 0xf6, 0x42, 0xbc, 0x10, 0x74, 0x5d, 0xed, 0x52, 0x9b, 0x5d, 0x66, 0x0d, 0x15,
	0xef, 0xc6, 0xed, 0x58, 0x42, 0x93, 0x99, 0x98, 0x99, 0x48, 0xf3, 0x38, 0x82, 0xef, 0xe0, 0x2b,
	0xf8, 0x58, 0x32, 0x1f, 0xbb, 0x4d, 0xd7, 0x0a, 0xde, 0xcd, 0x3f, 0xe7, 0x77, 0xfe, 0x73, 0x72,
	0xce, 0x19, 0x18, 0xa6, 0x4c, 0x70, 0x96, 0xb3, 0x94, 0x44, 0x45, 0xc9, 0x25, 0x47, 0x2d, 0x25,
	0xc3, 0x3e, 0xf4, 0xe6, 0x29, 0x3b, 0xc7, 0xf4, 0x5b, 0x45, 0x85, 0x0c, 0x77, 0xa1, 0x6b, 0x64,
	0x91, 0xd5, 0x68, 0x1b, 0x7c, 0x21, 0x89, 0xac, 0xc4, 0x8e, 0xf3, 0xd0, 0xd9, 0xeb, 0x62, 0xab,
	0xc2, 0xdf, 0x2e, 0x74, 0xa7, 0xec, 0x2b, 0x37, 0xd4, 0x33, 0x68, 0x2f, 0x24, 0x91, 0x0a, 0xf2,
	0xf6, 0x7a, 0x07, 0x41, 0xa4, 0x7c, 0xa3, 0x75, 0x3c, 0xd2, 0xc1, 0x09, 0x93, 0x65, 0x8d, 0x0d,
	0x18, 0xfc, 0x72, 0xa0, 0xad, 0xac, 0x04, 0xda, 0x07, 0x6f, 0x3c, 0x4f, 0xb4, 0x7d, 0xef, 0xe0,
	0xde, 0x66, 0xa6, 0x66, 0xa2, 0x65, 0x51, 0x61, 0x45, 0xa1, 0x43, 0xf0, 0x4f, 0x68, 0xce, 0xcb,
	0x7a, 0xc7, 0xd5, 0xfc, 0xfd, 0x9b, 0xf9, 0x5c, 0x33, 0xd8, 0xb2, 0xc1, 0x13, 0xf0, 0x96, 0x45,
	0x85, 0x1e, 0x00, 0x48, 0x2e, 0x49, 0x96, 0x08, 0x72, 0x4e, 0xf5, 0x85, 0x2d, 0xdc, 0xf8, 0x12,
	0x3c, 0x06, 0xdf, 0x24, 0xa2, 0x00, 0x3a, 0x39, 0xb9, 0x6c, 0x72, 0x6b, 0x1d, 0xcc, 0x00, 0xae,
	0x7e, 0x07, 0x8d, 0xc0, 0xbb, 0xa0, 0xb5, 0x6d, 0x8e, 0x3a, 0xa2, 0x7d, 0x68, 0x7f, 0x27, 0x59,
	0x45, 0x6d, 0x85, 0x77, 0x6f, 0xac, 0x10, 0x1b, 0xe6, 0xa5, 0xfb, 0xc2, 0x09, 0x1f, 0xc1, 0x70,
	0x21, 0x79, 0xf1, 0x91, 0x88, 0x0b, 0x3b, 0x02, 0x34, 0x00, 0x37, 0x3d, 0xb3, 0xa6, 0x6e, 0x7a,
	0x16, 0x0e, 0xa1, 0x7f, 0x85, 0x14, 0x59, 0x1d, 0xee, 0xc2, 0x96, 0x12, 0x0b, 0x3d, 0x8c, 0x7f,
	0x65, 0xfd, 0x70, 0x60, 0xd8, 0xa4, 0xd4, 0xa4, 0x0e, 0xaf, 0xcd, 0x73, 0xb0, 0x6a, 0xe0, 0x06,
	0x16, 0xd9, 0xf3, 0x6a, 0xda, 0x9f, 0xc0, 0x37, 0x5f, 0x50, 0x0f, 0x6e, 0x27, 0xf1, 0x71, 0x3c,
	0x3b, 0x8d, 0x47, 0xb7, 0xd0, 0x1d, 0xe8, 0x2c, 0xe6, 0xb3, 0xd9, 0x87, 0x69, 0xfc, 0x7e, 0xe4,
	0x18, 0xf5, 0xfa, 0x34, 0x56, 0xca, 0x55, 0x20, 0x4e, 0x62, 0x2d, 0x3c, 0x15, 0x7a, 0x37, 0x8d,
	0xa7, 0x8b, 0xa3, 0xc9, 0xdb, 0x51, 0x0b, 0x01, 0xf8, 0x6f, 0xf0, 0xec, 0x78, 0x12, 0x8f, 0xda,
	0xe1, 0x4f, 0x07, 0x06, 0xc6, 0xfa, 0x84, 0x14, 0xa6, 0xc4, 0x57, 0xd0, 0x31, 0xd7, 0xd2, 0xd5,
	0x3e, 0x85, 0xa6, 0xc8, 0xeb, 0x9c, 0x95, 0xd4, 0xee, 0xd5, 0x3a, 0x27, 0xc0, 0xaa, 0x59, 0x8d,
	0xd0, 0x7f, 0xcf, 0x68, 0xa3, 0x09, 0xcd, 0x19, 0x1d, 0x01, 0x1a, 0x73, 0x26, 0x49, 0xca, 0x68,
	0x89, 0xa9, 0xe0, 0x55, 0xb9, 0xa4, 0x42, 0x3d, 0x0e, 0xb3, 0x30, 0xda, 0xdb, 0xc3, 0x8d, 0xf5,
	0x61, 0x84, 0xf1, 0xf1, 0x3c, 0x11, 0xfa, 0x06, 0x0f, 0xaf, 0x75, 0xf8, 0x19, 0xb6, 0x9b, 0x4e,
	0x92, 0x94, 0x72, 0xce, 0xb3, 0x74, 0x59, 0x23, 0x04, 0x2d, 0x46, 0x72, 0x6a, 0xeb, 0xd4, 0x67,
	0xf4, 0x14, 0xb6, 0x72, 0x72, 0x99, 0xe6, 0x55, 0x8e, 0xa9, 0x2c, 0xeb, 0x31, 0xaf, 0x98, 0xd4,
	0x96, 0x7d, 0xfc, 0x77, 0xe0, 0x8b, 0xaf, 0x5f, 0xf5, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x00, 0x64, 0x3a, 0x5b, 0xe8, 0x03, 0x00, 0x00,
}