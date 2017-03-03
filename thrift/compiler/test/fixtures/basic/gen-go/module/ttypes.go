// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package module

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int;

type MyEnum int64
const (
  MyEnum_MyValue1 MyEnum = 0
  MyEnum_MyValue2 MyEnum = 1
)

func (p MyEnum) String() string {
  switch p {
  case MyEnum_MyValue1: return "MyValue1"
  case MyEnum_MyValue2: return "MyValue2"
  }
  return "<UNSET>"
}

func MyEnumFromString(s string) (MyEnum, error) {
  switch s {
  case "MyValue1": return MyEnum_MyValue1, nil 
  case "MyValue2": return MyEnum_MyValue2, nil 
  }
  return MyEnum(0), fmt.Errorf("not a valid MyEnum string")
}


func MyEnumPtr(v MyEnum) *MyEnum { return &v }

// Attributes:
//  - MyIntField
//  - MyStringField
type MyStruct struct {
  MyIntField int64 `thrift:"MyIntField,1" json:"MyIntField"`
  MyStringField string `thrift:"MyStringField,2" json:"MyStringField"`
}

func NewMyStruct() *MyStruct {
  return &MyStruct{}
}


func (p *MyStruct) GetMyIntField() int64 {
  return p.MyIntField
}

func (p *MyStruct) GetMyStringField() string {
  return p.MyStringField
}
func (p *MyStruct) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyStruct)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.MyIntField = v
}
  return nil
}

func (p *MyStruct)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.MyStringField = v
}
  return nil
}

func (p *MyStruct) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("MyStruct"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyStruct) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("MyIntField", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:MyIntField: ", p), err) }
  if err := oprot.WriteI64(int64(p.MyIntField)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.MyIntField (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:MyIntField: ", p), err) }
  return err
}

func (p *MyStruct) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("MyStringField", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:MyStringField: ", p), err) }
  if err := oprot.WriteString(string(p.MyStringField)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.MyStringField (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:MyStringField: ", p), err) }
  return err
}

func (p *MyStruct) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyStruct(%+v)", *p)
}

