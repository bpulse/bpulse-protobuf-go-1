// Code generated by protoc-gen-go.
// source: CollectorMessageRS.proto
// DO NOT EDIT!

package collector

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PulsesRS_StatusType int32

const (
	PulsesRS_OK    PulsesRS_StatusType = 0
	PulsesRS_ERROR PulsesRS_StatusType = 1
)

var PulsesRS_StatusType_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}
var PulsesRS_StatusType_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x PulsesRS_StatusType) Enum() *PulsesRS_StatusType {
	p := new(PulsesRS_StatusType)
	*p = x
	return p
}
func (x PulsesRS_StatusType) String() string {
	return proto.EnumName(PulsesRS_StatusType_name, int32(x))
}
func (x *PulsesRS_StatusType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PulsesRS_StatusType_value, data, "PulsesRS_StatusType")
	if err != nil {
		return err
	}
	*x = PulsesRS_StatusType(value)
	return nil
}

type PulsesRS struct {
	RsTime           *int64               `protobuf:"varint,1,req,name=rsTime" json:"rsTime,omitempty"`
	Status           *PulsesRS_StatusType `protobuf:"varint,2,req,name=status,enum=collector.PulsesRS_StatusType,def=0" json:"status,omitempty"`
	Error            *string              `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PulsesRS) Reset()         { *m = PulsesRS{} }
func (m *PulsesRS) String() string { return proto.CompactTextString(m) }
func (*PulsesRS) ProtoMessage()    {}

const Default_PulsesRS_Status PulsesRS_StatusType = PulsesRS_OK

func (m *PulsesRS) GetRsTime() int64 {
	if m != nil && m.RsTime != nil {
		return *m.RsTime
	}
	return 0
}

func (m *PulsesRS) GetStatus() PulsesRS_StatusType {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_PulsesRS_Status
}

func (m *PulsesRS) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("collector.PulsesRS_StatusType", PulsesRS_StatusType_name, PulsesRS_StatusType_value)
}
