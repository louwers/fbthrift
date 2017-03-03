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

type Animal int64
const (
  Animal_DOG Animal = 1
  Animal_CAT Animal = 2
  Animal_TARANTULA Animal = 3
)

func (p Animal) String() string {
  switch p {
  case Animal_DOG: return "DOG"
  case Animal_CAT: return "CAT"
  case Animal_TARANTULA: return "TARANTULA"
  }
  return "<UNSET>"
}

func AnimalFromString(s string) (Animal, error) {
  switch s {
  case "DOG": return Animal_DOG, nil 
  case "CAT": return Animal_CAT, nil 
  case "TARANTULA": return Animal_TARANTULA, nil 
  }
  return Animal(0), fmt.Errorf("not a valid Animal string")
}


func AnimalPtr(v Animal) *Animal { return &v }

type PersonID int64

func PersonIDPtr(v PersonID) *PersonID { return &v }

// Attributes:
//  - Red
//  - Green
//  - Blue
//  - Alpha
type Color struct {
  Red float64 `thrift:"red,1" json:"red"`
  Green float64 `thrift:"green,2" json:"green"`
  Blue float64 `thrift:"blue,3" json:"blue"`
  Alpha float64 `thrift:"alpha,4" json:"alpha"`
}

func NewColor() *Color {
  return &Color{}
}


func (p *Color) GetRed() float64 {
  return p.Red
}

func (p *Color) GetGreen() float64 {
  return p.Green
}

func (p *Color) GetBlue() float64 {
  return p.Blue
}

func (p *Color) GetAlpha() float64 {
  return p.Alpha
}
func (p *Color) Read(iprot thrift.TProtocol) error {
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
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
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

func (p *Color)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Red = v
}
  return nil
}

func (p *Color)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Green = v
}
  return nil
}

func (p *Color)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Blue = v
}
  return nil
}

func (p *Color)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadDouble(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Alpha = v
}
  return nil
}

func (p *Color) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Color"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := p.writeField4(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Color) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("red", thrift.DOUBLE, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:red: ", p), err) }
  if err := oprot.WriteDouble(float64(p.Red)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.red (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:red: ", p), err) }
  return err
}

func (p *Color) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("green", thrift.DOUBLE, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:green: ", p), err) }
  if err := oprot.WriteDouble(float64(p.Green)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.green (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:green: ", p), err) }
  return err
}

func (p *Color) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("blue", thrift.DOUBLE, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:blue: ", p), err) }
  if err := oprot.WriteDouble(float64(p.Blue)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.blue (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:blue: ", p), err) }
  return err
}

func (p *Color) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("alpha", thrift.DOUBLE, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:alpha: ", p), err) }
  if err := oprot.WriteDouble(float64(p.Alpha)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.alpha (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:alpha: ", p), err) }
  return err
}

func (p *Color) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Color(%+v)", *p)
}

// Attributes:
//  - Color
//  - LicensePlate
//  - Description
//  - Name
type Vehicle struct {
  Color *Color `thrift:"color,1" json:"color"`
  LicensePlate *string `thrift:"licensePlate,2" json:"licensePlate,omitempty"`
  Description *string `thrift:"description,3" json:"description,omitempty"`
  Name *string `thrift:"name,4" json:"name,omitempty"`
}

func NewVehicle() *Vehicle {
  return &Vehicle{}
}

var Vehicle_Color_DEFAULT *Color
func (p *Vehicle) GetColor() *Color {
  if !p.IsSetColor() {
    return Vehicle_Color_DEFAULT
  }
return p.Color
}
var Vehicle_LicensePlate_DEFAULT string
func (p *Vehicle) GetLicensePlate() string {
  if !p.IsSetLicensePlate() {
    return Vehicle_LicensePlate_DEFAULT
  }
return *p.LicensePlate
}
var Vehicle_Description_DEFAULT string
func (p *Vehicle) GetDescription() string {
  if !p.IsSetDescription() {
    return Vehicle_Description_DEFAULT
  }
return *p.Description
}
var Vehicle_Name_DEFAULT string
func (p *Vehicle) GetName() string {
  if !p.IsSetName() {
    return Vehicle_Name_DEFAULT
  }
return *p.Name
}
func (p *Vehicle) IsSetColor() bool {
  return p.Color != nil
}

func (p *Vehicle) IsSetLicensePlate() bool {
  return p.LicensePlate != nil
}

func (p *Vehicle) IsSetDescription() bool {
  return p.Description != nil
}

func (p *Vehicle) IsSetName() bool {
  return p.Name != nil
}

func (p *Vehicle) Read(iprot thrift.TProtocol) error {
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
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
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

func (p *Vehicle)  ReadField1(iprot thrift.TProtocol) error {
  p.Color = &Color{}
  if err := p.Color.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Color), err)
  }
  return nil
}

func (p *Vehicle)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.LicensePlate = &v
}
  return nil
}

func (p *Vehicle)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Description = &v
}
  return nil
}

func (p *Vehicle)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Name = &v
}
  return nil
}

func (p *Vehicle) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Vehicle"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := p.writeField4(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Vehicle) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("color", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:color: ", p), err) }
  if err := p.Color.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Color), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:color: ", p), err) }
  return err
}

func (p *Vehicle) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetLicensePlate() {
    if err := oprot.WriteFieldBegin("licensePlate", thrift.STRING, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:licensePlate: ", p), err) }
    if err := oprot.WriteString(string(*p.LicensePlate)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.licensePlate (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:licensePlate: ", p), err) }
  }
  return err
}

func (p *Vehicle) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetDescription() {
    if err := oprot.WriteFieldBegin("description", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:description: ", p), err) }
    if err := oprot.WriteString(string(*p.Description)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.description (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:description: ", p), err) }
  }
  return err
}

func (p *Vehicle) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetName() {
    if err := oprot.WriteFieldBegin("name", thrift.STRING, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:name: ", p), err) }
    if err := oprot.WriteString(string(*p.Name)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.name (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:name: ", p), err) }
  }
  return err
}

func (p *Vehicle) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Vehicle(%+v)", *p)
}

// Attributes:
//  - Id
//  - Name
//  - Age
//  - Address
//  - FavoriteColor
//  - Friends
//  - BestFriend
//  - PetNames
//  - AfraidOfAnimal
//  - Vehicles
type Person struct {
  Id PersonID `thrift:"id,1" json:"id"`
  Name string `thrift:"name,2" json:"name"`
  Age *int16 `thrift:"age,3" json:"age,omitempty"`
  Address *string `thrift:"address,4" json:"address,omitempty"`
  FavoriteColor *Color `thrift:"favoriteColor,5" json:"favoriteColor,omitempty"`
  Friends map[int64]bool `thrift:"friends,6" json:"friends,omitempty"`
  BestFriend *PersonID `thrift:"bestFriend,7" json:"bestFriend,omitempty"`
  PetNames map[Animal]string `thrift:"petNames,8" json:"petNames,omitempty"`
  AfraidOfAnimal *Animal `thrift:"afraidOfAnimal,9" json:"afraidOfAnimal,omitempty"`
  Vehicles []*Vehicle `thrift:"vehicles,10" json:"vehicles,omitempty"`
}

func NewPerson() *Person {
  return &Person{}
}


func (p *Person) GetId() PersonID {
  return p.Id
}

func (p *Person) GetName() string {
  return p.Name
}
var Person_Age_DEFAULT int16
func (p *Person) GetAge() int16 {
  if !p.IsSetAge() {
    return Person_Age_DEFAULT
  }
return *p.Age
}
var Person_Address_DEFAULT string
func (p *Person) GetAddress() string {
  if !p.IsSetAddress() {
    return Person_Address_DEFAULT
  }
return *p.Address
}
var Person_FavoriteColor_DEFAULT *Color
func (p *Person) GetFavoriteColor() *Color {
  if !p.IsSetFavoriteColor() {
    return Person_FavoriteColor_DEFAULT
  }
return p.FavoriteColor
}
var Person_Friends_DEFAULT map[int64]bool

func (p *Person) GetFriends() map[int64]bool {
  return p.Friends
}
var Person_BestFriend_DEFAULT PersonID
func (p *Person) GetBestFriend() PersonID {
  if !p.IsSetBestFriend() {
    return Person_BestFriend_DEFAULT
  }
return *p.BestFriend
}
var Person_PetNames_DEFAULT map[Animal]string

func (p *Person) GetPetNames() map[Animal]string {
  return p.PetNames
}
var Person_AfraidOfAnimal_DEFAULT Animal
func (p *Person) GetAfraidOfAnimal() Animal {
  if !p.IsSetAfraidOfAnimal() {
    return Person_AfraidOfAnimal_DEFAULT
  }
return *p.AfraidOfAnimal
}
var Person_Vehicles_DEFAULT []*Vehicle

func (p *Person) GetVehicles() []*Vehicle {
  return p.Vehicles
}
func (p *Person) IsSetAge() bool {
  return p.Age != nil
}

func (p *Person) IsSetAddress() bool {
  return p.Address != nil
}

func (p *Person) IsSetFavoriteColor() bool {
  return p.FavoriteColor != nil
}

func (p *Person) IsSetFriends() bool {
  return p.Friends != nil
}

func (p *Person) IsSetBestFriend() bool {
  return p.BestFriend != nil
}

func (p *Person) IsSetPetNames() bool {
  return p.PetNames != nil
}

func (p *Person) IsSetAfraidOfAnimal() bool {
  return p.AfraidOfAnimal != nil
}

func (p *Person) IsSetVehicles() bool {
  return p.Vehicles != nil
}

func (p *Person) Read(iprot thrift.TProtocol) error {
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
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    case 6:
      if err := p.ReadField6(iprot); err != nil {
        return err
      }
    case 7:
      if err := p.ReadField7(iprot); err != nil {
        return err
      }
    case 8:
      if err := p.ReadField8(iprot); err != nil {
        return err
      }
    case 9:
      if err := p.ReadField9(iprot); err != nil {
        return err
      }
    case 10:
      if err := p.ReadField10(iprot); err != nil {
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

func (p *Person)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := PersonID(v)
  p.Id = temp
}
  return nil
}

func (p *Person)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *Person)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI16(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Age = &v
}
  return nil
}

func (p *Person)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Address = &v
}
  return nil
}

func (p *Person)  ReadField5(iprot thrift.TProtocol) error {
  p.FavoriteColor = &Color{}
  if err := p.FavoriteColor.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.FavoriteColor), err)
  }
  return nil
}

func (p *Person)  ReadField6(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadSetBegin()
  if err != nil {
    return thrift.PrependError("error reading set begin: ", err)
  }
  tSet := make(map[int64]bool, size)
  p.Friends =  tSet
  for i := 0; i < size; i ++ {
var _elem0 int64
    if v, err := iprot.ReadI64(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.Friends[_elem0] = true
  }
  if err := iprot.ReadSetEnd(); err != nil {
    return thrift.PrependError("error reading set end: ", err)
  }
  return nil
}

func (p *Person)  ReadField7(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  temp := PersonID(v)
  p.BestFriend = &temp
}
  return nil
}

func (p *Person)  ReadField8(iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin()
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(map[Animal]string, size)
  p.PetNames =  tMap
  for i := 0; i < size; i ++ {
var _key1 Animal
    if v, err := iprot.ReadI32(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    temp := Animal(v)
    _key1 = temp
}
var _val2 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val2 = v
}
    p.PetNames[_key1] = _val2
  }
  if err := iprot.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *Person)  ReadField9(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 9: ", err)
} else {
  temp := Animal(v)
  p.AfraidOfAnimal = &temp
}
  return nil
}

func (p *Person)  ReadField10(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*Vehicle, 0, size)
  p.Vehicles =  tSlice
  for i := 0; i < size; i ++ {
    _elem3 := &Vehicle{}
    if err := _elem3.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem3), err)
    }
    p.Vehicles = append(p.Vehicles, _elem3)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *Person) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Person"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := p.writeField4(oprot); err != nil { return err }
  if err := p.writeField5(oprot); err != nil { return err }
  if err := p.writeField6(oprot); err != nil { return err }
  if err := p.writeField7(oprot); err != nil { return err }
  if err := p.writeField8(oprot); err != nil { return err }
  if err := p.writeField9(oprot); err != nil { return err }
  if err := p.writeField10(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Person) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteI64(int64(p.Id)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *Person) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err) }
  return err
}

func (p *Person) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetAge() {
    if err := oprot.WriteFieldBegin("age", thrift.I16, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:age: ", p), err) }
    if err := oprot.WriteI16(int16(*p.Age)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.age (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:age: ", p), err) }
  }
  return err
}

func (p *Person) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetAddress() {
    if err := oprot.WriteFieldBegin("address", thrift.STRING, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:address: ", p), err) }
    if err := oprot.WriteString(string(*p.Address)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.address (4) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:address: ", p), err) }
  }
  return err
}

func (p *Person) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetFavoriteColor() {
    if err := oprot.WriteFieldBegin("favoriteColor", thrift.STRUCT, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:favoriteColor: ", p), err) }
    if err := p.FavoriteColor.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.FavoriteColor), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:favoriteColor: ", p), err) }
  }
  return err
}

func (p *Person) writeField6(oprot thrift.TProtocol) (err error) {
  if p.IsSetFriends() {
    if err := oprot.WriteFieldBegin("friends", thrift.SET, 6); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:friends: ", p), err) }
    if err := oprot.WriteSetBegin(thrift.I64, len(p.Friends)); err != nil {
      return thrift.PrependError("error writing set begin: ", err)
    }
    for v, _ := range p.Friends {
      if err := oprot.WriteI64(int64(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteSetEnd(); err != nil {
      return thrift.PrependError("error writing set end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 6:friends: ", p), err) }
  }
  return err
}

func (p *Person) writeField7(oprot thrift.TProtocol) (err error) {
  if p.IsSetBestFriend() {
    if err := oprot.WriteFieldBegin("bestFriend", thrift.I64, 7); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:bestFriend: ", p), err) }
    if err := oprot.WriteI64(int64(*p.BestFriend)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.bestFriend (7) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 7:bestFriend: ", p), err) }
  }
  return err
}

func (p *Person) writeField8(oprot thrift.TProtocol) (err error) {
  if p.IsSetPetNames() {
    if err := oprot.WriteFieldBegin("petNames", thrift.MAP, 8); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:petNames: ", p), err) }
    if err := oprot.WriteMapBegin(thrift.I32, thrift.STRING, len(p.PetNames)); err != nil {
      return thrift.PrependError("error writing map begin: ", err)
    }
    for k, v := range p.PetNames {
      if err := oprot.WriteI32(int32(k)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteMapEnd(); err != nil {
      return thrift.PrependError("error writing map end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 8:petNames: ", p), err) }
  }
  return err
}

func (p *Person) writeField9(oprot thrift.TProtocol) (err error) {
  if p.IsSetAfraidOfAnimal() {
    if err := oprot.WriteFieldBegin("afraidOfAnimal", thrift.I32, 9); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:afraidOfAnimal: ", p), err) }
    if err := oprot.WriteI32(int32(*p.AfraidOfAnimal)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.afraidOfAnimal (9) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 9:afraidOfAnimal: ", p), err) }
  }
  return err
}

func (p *Person) writeField10(oprot thrift.TProtocol) (err error) {
  if p.IsSetVehicles() {
    if err := oprot.WriteFieldBegin("vehicles", thrift.LIST, 10); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:vehicles: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Vehicles)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.Vehicles {
      if err := v.Write(oprot); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
      }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 10:vehicles: ", p), err) }
  }
  return err
}

func (p *Person) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Person(%+v)", *p)
}

