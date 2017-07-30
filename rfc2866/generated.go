// Generated by radius-dict-gen. DO NOT EDIT.

package rfc2866

import (
	"net"
	"strconv"

	"layeh.com/radius"
)

var _ = radius.Type(0)
var _ = strconv.Itoa
var _ = net.ParseIP

const (
	AcctStatusType_Type     radius.Type = 40
	AcctDelayTime_Type      radius.Type = 41
	AcctInputOctets_Type    radius.Type = 42
	AcctOutputOctets_Type   radius.Type = 43
	AcctSessionID_Type      radius.Type = 44
	AcctAuthentic_Type      radius.Type = 45
	AcctSessionTime_Type    radius.Type = 46
	AcctInputPackets_Type   radius.Type = 47
	AcctOutputPackets_Type  radius.Type = 48
	AcctTerminateCause_Type radius.Type = 49
	AcctMultiSessionID_Type radius.Type = 50
	AcctLinkCount_Type      radius.Type = 51
)

type AcctStatusType uint32

const (
	AcctStatusType_Value_Start         AcctStatusType = 1
	AcctStatusType_Value_Stop          AcctStatusType = 2
	AcctStatusType_Value_InterimUpdate AcctStatusType = 3
	AcctStatusType_Value_AccountingOn  AcctStatusType = 7
	AcctStatusType_Value_AccountingOff AcctStatusType = 8
	AcctStatusType_Value_Failed        AcctStatusType = 15
)

func (a AcctStatusType) String() string {
	switch a {
	case AcctStatusType_Value_Start:
		return `Start`
	case AcctStatusType_Value_Stop:
		return `Stop`
	case AcctStatusType_Value_InterimUpdate:
		return `Interim-Update`
	case AcctStatusType_Value_AccountingOn:
		return `Accounting-On`
	case AcctStatusType_Value_AccountingOff:
		return `Accounting-Off`
	case AcctStatusType_Value_Failed:
		return `Failed`
	}
	return "AcctStatusType(" + strconv.Itoa(int(a)) + ")"
}

func AcctStatusType_Add(p *radius.Packet, value AcctStatusType) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctStatusType_Type, a)
}

func AcctStatusType_Get(p *radius.Packet) (value AcctStatusType) {
	value, _ = AcctStatusType_Lookup(p)
	return
}

func AcctStatusType_Gets(p *radius.Packet) (values []AcctStatusType, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctStatusType_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctStatusType(i))
	}
	return
}

func AcctStatusType_Lookup(p *radius.Packet) (value AcctStatusType, err error) {
	a, ok := p.Lookup(AcctStatusType_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctStatusType(i)
	return
}

func AcctStatusType_Set(p *radius.Packet, value AcctStatusType) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctStatusType_Type, a)
}

type AcctDelayTime uint32

func (a AcctDelayTime) String() string {
	return "AcctDelayTime(" + strconv.Itoa(int(a)) + ")"
}

func AcctDelayTime_Add(p *radius.Packet, value AcctDelayTime) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctDelayTime_Type, a)
}

func AcctDelayTime_Get(p *radius.Packet) (value AcctDelayTime) {
	value, _ = AcctDelayTime_Lookup(p)
	return
}

func AcctDelayTime_Gets(p *radius.Packet) (values []AcctDelayTime, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctDelayTime_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctDelayTime(i))
	}
	return
}

func AcctDelayTime_Lookup(p *radius.Packet) (value AcctDelayTime, err error) {
	a, ok := p.Lookup(AcctDelayTime_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctDelayTime(i)
	return
}

func AcctDelayTime_Set(p *radius.Packet, value AcctDelayTime) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctDelayTime_Type, a)
}

type AcctInputOctets uint32

func (a AcctInputOctets) String() string {
	return "AcctInputOctets(" + strconv.Itoa(int(a)) + ")"
}

func AcctInputOctets_Add(p *radius.Packet, value AcctInputOctets) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctInputOctets_Type, a)
}

func AcctInputOctets_Get(p *radius.Packet) (value AcctInputOctets) {
	value, _ = AcctInputOctets_Lookup(p)
	return
}

func AcctInputOctets_Gets(p *radius.Packet) (values []AcctInputOctets, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctInputOctets_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctInputOctets(i))
	}
	return
}

func AcctInputOctets_Lookup(p *radius.Packet) (value AcctInputOctets, err error) {
	a, ok := p.Lookup(AcctInputOctets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctInputOctets(i)
	return
}

func AcctInputOctets_Set(p *radius.Packet, value AcctInputOctets) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctInputOctets_Type, a)
}

type AcctOutputOctets uint32

func (a AcctOutputOctets) String() string {
	return "AcctOutputOctets(" + strconv.Itoa(int(a)) + ")"
}

func AcctOutputOctets_Add(p *radius.Packet, value AcctOutputOctets) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctOutputOctets_Type, a)
}

func AcctOutputOctets_Get(p *radius.Packet) (value AcctOutputOctets) {
	value, _ = AcctOutputOctets_Lookup(p)
	return
}

func AcctOutputOctets_Gets(p *radius.Packet) (values []AcctOutputOctets, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctOutputOctets_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctOutputOctets(i))
	}
	return
}

func AcctOutputOctets_Lookup(p *radius.Packet) (value AcctOutputOctets, err error) {
	a, ok := p.Lookup(AcctOutputOctets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctOutputOctets(i)
	return
}

func AcctOutputOctets_Set(p *radius.Packet, value AcctOutputOctets) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctOutputOctets_Type, a)
}

func AcctSessionID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute

	a, err = radius.NewBytes(value)

	if err != nil {
		return
	}
	p.Add(AcctSessionID_Type, a)
	return nil
}

func AcctSessionID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute

	a, err = radius.NewString(value)

	if err != nil {
		return
	}
	p.Add(AcctSessionID_Type, a)
	return nil
}

func AcctSessionID_Get(p *radius.Packet) (value []byte) {
	value, _ = AcctSessionID_Lookup(p)
	return
}

func AcctSessionID_GetString(p *radius.Packet) (value string) {
	return string(AcctSessionID_Get(p))
}

func AcctSessionID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[AcctSessionID_Type] {

		i = radius.Bytes(attr)

		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctSessionID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[AcctSessionID_Type] {

		i = radius.String(attr)

		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctSessionID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(AcctSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}

	value = radius.Bytes(a)

	return
}

func AcctSessionID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(AcctSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}

	value = radius.String(a)

	return
}

func AcctSessionID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute

	a, err = radius.NewBytes(value)

	if err != nil {
		return
	}
	p.Set(AcctSessionID_Type, a)
	return
}

func AcctSessionID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute

	a, err = radius.NewString(value)

	if err != nil {
		return
	}
	p.Set(AcctSessionID_Type, a)
	return
}

type AcctAuthentic uint32

const (
	AcctAuthentic_Value_RADIUS   AcctAuthentic = 1
	AcctAuthentic_Value_Local    AcctAuthentic = 2
	AcctAuthentic_Value_Remote   AcctAuthentic = 3
	AcctAuthentic_Value_Diameter AcctAuthentic = 4
)

func (a AcctAuthentic) String() string {
	switch a {
	case AcctAuthentic_Value_RADIUS:
		return `RADIUS`
	case AcctAuthentic_Value_Local:
		return `Local`
	case AcctAuthentic_Value_Remote:
		return `Remote`
	case AcctAuthentic_Value_Diameter:
		return `Diameter`
	}
	return "AcctAuthentic(" + strconv.Itoa(int(a)) + ")"
}

func AcctAuthentic_Add(p *radius.Packet, value AcctAuthentic) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctAuthentic_Type, a)
}

func AcctAuthentic_Get(p *radius.Packet) (value AcctAuthentic) {
	value, _ = AcctAuthentic_Lookup(p)
	return
}

func AcctAuthentic_Gets(p *radius.Packet) (values []AcctAuthentic, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctAuthentic_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctAuthentic(i))
	}
	return
}

func AcctAuthentic_Lookup(p *radius.Packet) (value AcctAuthentic, err error) {
	a, ok := p.Lookup(AcctAuthentic_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctAuthentic(i)
	return
}

func AcctAuthentic_Set(p *radius.Packet, value AcctAuthentic) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctAuthentic_Type, a)
}

type AcctSessionTime uint32

func (a AcctSessionTime) String() string {
	return "AcctSessionTime(" + strconv.Itoa(int(a)) + ")"
}

func AcctSessionTime_Add(p *radius.Packet, value AcctSessionTime) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctSessionTime_Type, a)
}

func AcctSessionTime_Get(p *radius.Packet) (value AcctSessionTime) {
	value, _ = AcctSessionTime_Lookup(p)
	return
}

func AcctSessionTime_Gets(p *radius.Packet) (values []AcctSessionTime, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctSessionTime_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctSessionTime(i))
	}
	return
}

func AcctSessionTime_Lookup(p *radius.Packet) (value AcctSessionTime, err error) {
	a, ok := p.Lookup(AcctSessionTime_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctSessionTime(i)
	return
}

func AcctSessionTime_Set(p *radius.Packet, value AcctSessionTime) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctSessionTime_Type, a)
}

type AcctInputPackets uint32

func (a AcctInputPackets) String() string {
	return "AcctInputPackets(" + strconv.Itoa(int(a)) + ")"
}

func AcctInputPackets_Add(p *radius.Packet, value AcctInputPackets) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctInputPackets_Type, a)
}

func AcctInputPackets_Get(p *radius.Packet) (value AcctInputPackets) {
	value, _ = AcctInputPackets_Lookup(p)
	return
}

func AcctInputPackets_Gets(p *radius.Packet) (values []AcctInputPackets, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctInputPackets_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctInputPackets(i))
	}
	return
}

func AcctInputPackets_Lookup(p *radius.Packet) (value AcctInputPackets, err error) {
	a, ok := p.Lookup(AcctInputPackets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctInputPackets(i)
	return
}

func AcctInputPackets_Set(p *radius.Packet, value AcctInputPackets) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctInputPackets_Type, a)
}

type AcctOutputPackets uint32

func (a AcctOutputPackets) String() string {
	return "AcctOutputPackets(" + strconv.Itoa(int(a)) + ")"
}

func AcctOutputPackets_Add(p *radius.Packet, value AcctOutputPackets) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctOutputPackets_Type, a)
}

func AcctOutputPackets_Get(p *radius.Packet) (value AcctOutputPackets) {
	value, _ = AcctOutputPackets_Lookup(p)
	return
}

func AcctOutputPackets_Gets(p *radius.Packet) (values []AcctOutputPackets, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctOutputPackets_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctOutputPackets(i))
	}
	return
}

func AcctOutputPackets_Lookup(p *radius.Packet) (value AcctOutputPackets, err error) {
	a, ok := p.Lookup(AcctOutputPackets_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctOutputPackets(i)
	return
}

func AcctOutputPackets_Set(p *radius.Packet, value AcctOutputPackets) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctOutputPackets_Type, a)
}

type AcctTerminateCause uint32

const (
	AcctTerminateCause_Value_UserRequest        AcctTerminateCause = 1
	AcctTerminateCause_Value_LostCarrier        AcctTerminateCause = 2
	AcctTerminateCause_Value_LostService        AcctTerminateCause = 3
	AcctTerminateCause_Value_IdleTimeout        AcctTerminateCause = 4
	AcctTerminateCause_Value_SessionTimeout     AcctTerminateCause = 5
	AcctTerminateCause_Value_AdminReset         AcctTerminateCause = 6
	AcctTerminateCause_Value_AdminReboot        AcctTerminateCause = 7
	AcctTerminateCause_Value_PortError          AcctTerminateCause = 8
	AcctTerminateCause_Value_NASError           AcctTerminateCause = 9
	AcctTerminateCause_Value_NASRequest         AcctTerminateCause = 10
	AcctTerminateCause_Value_NASReboot          AcctTerminateCause = 11
	AcctTerminateCause_Value_PortUnneeded       AcctTerminateCause = 12
	AcctTerminateCause_Value_PortPreempted      AcctTerminateCause = 13
	AcctTerminateCause_Value_PortSuspended      AcctTerminateCause = 14
	AcctTerminateCause_Value_ServiceUnavailable AcctTerminateCause = 15
	AcctTerminateCause_Value_Callback           AcctTerminateCause = 16
	AcctTerminateCause_Value_UserError          AcctTerminateCause = 17
	AcctTerminateCause_Value_HostRequest        AcctTerminateCause = 18
)

func (a AcctTerminateCause) String() string {
	switch a {
	case AcctTerminateCause_Value_UserRequest:
		return `User-Request`
	case AcctTerminateCause_Value_LostCarrier:
		return `Lost-Carrier`
	case AcctTerminateCause_Value_LostService:
		return `Lost-Service`
	case AcctTerminateCause_Value_IdleTimeout:
		return `Idle-Timeout`
	case AcctTerminateCause_Value_SessionTimeout:
		return `Session-Timeout`
	case AcctTerminateCause_Value_AdminReset:
		return `Admin-Reset`
	case AcctTerminateCause_Value_AdminReboot:
		return `Admin-Reboot`
	case AcctTerminateCause_Value_PortError:
		return `Port-Error`
	case AcctTerminateCause_Value_NASError:
		return `NAS-Error`
	case AcctTerminateCause_Value_NASRequest:
		return `NAS-Request`
	case AcctTerminateCause_Value_NASReboot:
		return `NAS-Reboot`
	case AcctTerminateCause_Value_PortUnneeded:
		return `Port-Unneeded`
	case AcctTerminateCause_Value_PortPreempted:
		return `Port-Preempted`
	case AcctTerminateCause_Value_PortSuspended:
		return `Port-Suspended`
	case AcctTerminateCause_Value_ServiceUnavailable:
		return `Service-Unavailable`
	case AcctTerminateCause_Value_Callback:
		return `Callback`
	case AcctTerminateCause_Value_UserError:
		return `User-Error`
	case AcctTerminateCause_Value_HostRequest:
		return `Host-Request`
	}
	return "AcctTerminateCause(" + strconv.Itoa(int(a)) + ")"
}

func AcctTerminateCause_Add(p *radius.Packet, value AcctTerminateCause) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctTerminateCause_Type, a)
}

func AcctTerminateCause_Get(p *radius.Packet) (value AcctTerminateCause) {
	value, _ = AcctTerminateCause_Lookup(p)
	return
}

func AcctTerminateCause_Gets(p *radius.Packet) (values []AcctTerminateCause, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctTerminateCause_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctTerminateCause(i))
	}
	return
}

func AcctTerminateCause_Lookup(p *radius.Packet) (value AcctTerminateCause, err error) {
	a, ok := p.Lookup(AcctTerminateCause_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctTerminateCause(i)
	return
}

func AcctTerminateCause_Set(p *radius.Packet, value AcctTerminateCause) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctTerminateCause_Type, a)
}

func AcctMultiSessionID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute

	a, err = radius.NewBytes(value)

	if err != nil {
		return
	}
	p.Add(AcctMultiSessionID_Type, a)
	return nil
}

func AcctMultiSessionID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute

	a, err = radius.NewString(value)

	if err != nil {
		return
	}
	p.Add(AcctMultiSessionID_Type, a)
	return nil
}

func AcctMultiSessionID_Get(p *radius.Packet) (value []byte) {
	value, _ = AcctMultiSessionID_Lookup(p)
	return
}

func AcctMultiSessionID_GetString(p *radius.Packet) (value string) {
	return string(AcctMultiSessionID_Get(p))
}

func AcctMultiSessionID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[AcctMultiSessionID_Type] {

		i = radius.Bytes(attr)

		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctMultiSessionID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[AcctMultiSessionID_Type] {

		i = radius.String(attr)

		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AcctMultiSessionID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(AcctMultiSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}

	value = radius.Bytes(a)

	return
}

func AcctMultiSessionID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(AcctMultiSessionID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}

	value = radius.String(a)

	return
}

func AcctMultiSessionID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute

	a, err = radius.NewBytes(value)

	if err != nil {
		return
	}
	p.Set(AcctMultiSessionID_Type, a)
	return
}

func AcctMultiSessionID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute

	a, err = radius.NewString(value)

	if err != nil {
		return
	}
	p.Set(AcctMultiSessionID_Type, a)
	return
}

type AcctLinkCount uint32

func (a AcctLinkCount) String() string {
	return "AcctLinkCount(" + strconv.Itoa(int(a)) + ")"
}

func AcctLinkCount_Add(p *radius.Packet, value AcctLinkCount) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctLinkCount_Type, a)
}

func AcctLinkCount_Get(p *radius.Packet) (value AcctLinkCount) {
	value, _ = AcctLinkCount_Lookup(p)
	return
}

func AcctLinkCount_Gets(p *radius.Packet) (values []AcctLinkCount, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctLinkCount_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctLinkCount(i))
	}
	return
}

func AcctLinkCount_Lookup(p *radius.Packet) (value AcctLinkCount, err error) {
	a, ok := p.Lookup(AcctLinkCount_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctLinkCount(i)
	return
}

func AcctLinkCount_Set(p *radius.Packet, value AcctLinkCount) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctLinkCount_Type, a)
}
