// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hellothrift

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type HelloThrift interface {
  // Parameters:
  //  - Username
  SayHello(ctx context.Context, username string) (r string, err error)
  // Parameters:
  //  - Username
  Say(ctx context.Context, username string) (r string, err error)
}

type HelloThriftClient struct {
  c thrift.TClient
}

// Deprecated: Use NewHelloThrift instead
func NewHelloThriftClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HelloThriftClient {
  return &HelloThriftClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

// Deprecated: Use NewHelloThrift instead
func NewHelloThriftClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HelloThriftClient {
  return &HelloThriftClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewHelloThriftClient(c thrift.TClient) *HelloThriftClient {
  return &HelloThriftClient{
    c: c,
  }
}

// Parameters:
//  - Username
func (p *HelloThriftClient) SayHello(ctx context.Context, username string) (r string, err error) {
  var _args0 HelloThriftSayHelloArgs
  _args0.Username = username
  var _result1 HelloThriftSayHelloResult
  if err = p.c.Call(ctx, "SayHello", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

// Parameters:
//  - Username
func (p *HelloThriftClient) Say(ctx context.Context, username string) (r string, err error) {
  var _args2 HelloThriftSayArgs
  _args2.Username = username
  var _result3 HelloThriftSayResult
  if err = p.c.Call(ctx, "Say", &_args2, &_result3); err != nil {
    return
  }
  return _result3.GetSuccess(), nil
}

type HelloThriftProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler HelloThrift
}

func (p *HelloThriftProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *HelloThriftProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *HelloThriftProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewHelloThriftProcessor(handler HelloThrift) *HelloThriftProcessor {

  self4 := &HelloThriftProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["SayHello"] = &helloThriftProcessorSayHello{handler:handler}
  self4.processorMap["Say"] = &helloThriftProcessorSay{handler:handler}
return self4
}

func (p *HelloThriftProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x5.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x5

}

type helloThriftProcessorSayHello struct {
  handler HelloThrift
}

func (p *helloThriftProcessorSayHello) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := HelloThriftSayHelloArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("SayHello", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := HelloThriftSayHelloResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.SayHello(ctx, args.Username); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing SayHello: " + err2.Error())
    oprot.WriteMessageBegin("SayHello", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("SayHello", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type helloThriftProcessorSay struct {
  handler HelloThrift
}

func (p *helloThriftProcessorSay) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := HelloThriftSayArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("Say", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := HelloThriftSayResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Say(ctx, args.Username); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Say: " + err2.Error())
    oprot.WriteMessageBegin("Say", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("Say", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Username
type HelloThriftSayHelloArgs struct {
  Username string `thrift:"username,1" db:"username" json:"username"`
}

func NewHelloThriftSayHelloArgs() *HelloThriftSayHelloArgs {
  return &HelloThriftSayHelloArgs{}
}


func (p *HelloThriftSayHelloArgs) GetUsername() string {
  return p.Username
}
func (p *HelloThriftSayHelloArgs) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *HelloThriftSayHelloArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Username = v
}
  return nil
}

func (p *HelloThriftSayHelloArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SayHello_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloThriftSayHelloArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("username", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:username: ", p), err) }
  if err := oprot.WriteString(string(p.Username)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.username (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:username: ", p), err) }
  return err
}

func (p *HelloThriftSayHelloArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloThriftSayHelloArgs(%+v)", *p)
}

// Attributes:
//  - Success
type HelloThriftSayHelloResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewHelloThriftSayHelloResult() *HelloThriftSayHelloResult {
  return &HelloThriftSayHelloResult{}
}

var HelloThriftSayHelloResult_Success_DEFAULT string
func (p *HelloThriftSayHelloResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return HelloThriftSayHelloResult_Success_DEFAULT
  }
return *p.Success
}
func (p *HelloThriftSayHelloResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *HelloThriftSayHelloResult) Read(iprot thrift.TProtocol) error {
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
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *HelloThriftSayHelloResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *HelloThriftSayHelloResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("SayHello_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloThriftSayHelloResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *HelloThriftSayHelloResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloThriftSayHelloResult(%+v)", *p)
}

// Attributes:
//  - Username
type HelloThriftSayArgs struct {
  Username string `thrift:"username,1" db:"username" json:"username"`
}

func NewHelloThriftSayArgs() *HelloThriftSayArgs {
  return &HelloThriftSayArgs{}
}


func (p *HelloThriftSayArgs) GetUsername() string {
  return p.Username
}
func (p *HelloThriftSayArgs) Read(iprot thrift.TProtocol) error {
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
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *HelloThriftSayArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Username = v
}
  return nil
}

func (p *HelloThriftSayArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Say_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloThriftSayArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("username", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:username: ", p), err) }
  if err := oprot.WriteString(string(p.Username)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.username (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:username: ", p), err) }
  return err
}

func (p *HelloThriftSayArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloThriftSayArgs(%+v)", *p)
}

// Attributes:
//  - Success
type HelloThriftSayResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewHelloThriftSayResult() *HelloThriftSayResult {
  return &HelloThriftSayResult{}
}

var HelloThriftSayResult_Success_DEFAULT string
func (p *HelloThriftSayResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return HelloThriftSayResult_Success_DEFAULT
  }
return *p.Success
}
func (p *HelloThriftSayResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *HelloThriftSayResult) Read(iprot thrift.TProtocol) error {
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
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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

func (p *HelloThriftSayResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *HelloThriftSayResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Say_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *HelloThriftSayResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *HelloThriftSayResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HelloThriftSayResult(%+v)", *p)
}


