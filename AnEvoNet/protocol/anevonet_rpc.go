// This file is automatically generated. Do not modify.

package anevonet_rpc

import (
	"strconv"
	"fmt"
	"Common"
)

type P2PDNA *Common.P2PDNA

const VERSION string = "0.1.0"

type Transport int32

var(
	TransportSCTP = Transport(3)
	TransportTCP = Transport(0)
	TransportWEBRTC = Transport(2)
	TransportWEBSOCKET = Transport(1)
	TransportByName = map[string]Transport{
		"Transport.SCTP": TransportSCTP,
		"Transport.TCP": TransportTCP,
		"Transport.WEBRTC": TransportWEBRTC,
		"Transport.WEBSOCKET": TransportWEBSOCKET,
	}
	TransportByValue = map[Transport]string{
		TransportSCTP: "Transport.SCTP",
		TransportTCP: "Transport.TCP",
		TransportWEBRTC: "Transport.WEBRTC",
		TransportWEBSOCKET: "Transport.WEBSOCKET",
	}
)

func (e Transport) String() string {
	name := TransportByValue[e]
	if name == "" {
		name = fmt.Sprintf("Unknown enum value Transport(%d)", e)
	}
	return name
}

func (e Transport) MarshalJSON() ([]byte, error) {
	name := TransportByValue[e]
	if name == "" {
		name = strconv.Itoa(int(e))
	}
	return []byte("\""+name+"\""), nil
}

func (e *Transport) UnmarshalJSON(b []byte) error {
	st := string(b)
	if st[0] == '"' {
		*e = Transport(TransportByName[st[1:len(st)-1]])
		return nil
	}
	i, err := strconv.Atoi(st)
	*e = Transport(i)
	return err
}

type BootstrapNetworkRes struct {
	Success bool `thrift:"1,required" json:"Success"`
}

type BootstrapRes struct {
	Peers []*Common.Peer `thrift:"1,required" json:"Peers"`
}

type Gene struct {
	Value int32 `thrift:"1,required" json:"value"`
}

type Peer struct {
	ID string `thrift:"1,required" json:"ID"`
	IP string `thrift:"2,required" json:"IP"`
	Port int32 `thrift:"3,required" json:"Port"`
}

type Timestamp struct {
	Sec int64 `thrift:"1,required" json:"sec"`
}

type ConnectionReq struct {
	Target *Common.Peer `thrift:"1,required" json:"Target"`
	Module string `thrift:"2,required" json:"Module"`
}

type ConnectionRes struct {
	Socket string `thrift:"1,required" json:"Socket"`
}

type Module struct {
	Name string `thrift:"1,required" json:"Name"`
	DNA *Common.P2PDNA `thrift:"2,required" json:"DNA"`
}

type RegisterRes struct {
	Socket string `thrift:"1,required" json:"Socket"`
	DNA *Common.P2PDNA `thrift:"2,required" json:"DNA"`
	ID *Common.Peer `thrift:"3,required" json:"ID"`
}

type Service struct {
	Name string `thrift:"1,required" json:"Name"`
}

type StatusRes struct {
	ID *Common.Peer `thrift:"1,required" json:"ID"`
	OnlinePeers int32 `thrift:"2,required" json:"OnlinePeers"`
}

type RPCClient interface {
	Call(method string, request interface{}, response interface{}) error
}

type InternalRpc interface {
	BootstrapAlgorithm() (*BootstrapRes, error)
	BootstrapNetwork(P *Common.Peer) (bool, error)
	RegisterModule(M *Module) (*RegisterRes, error)
	RegisterService(S *Service) (bool, error)
	RequestConnection(Req *ConnectionReq) (*ConnectionRes, error)
	ShutdownConnection(Req *ConnectionRes) error
	Status() (*StatusRes, error)
}

type InternalRpcServer struct {
	Implementation InternalRpc
}

func (s *InternalRpcServer) BootstrapAlgorithm(req *InternalRpcBootstrapAlgorithmRequest, res *InternalRpcBootstrapAlgorithmResponse) error {
	val, err := s.Implementation.BootstrapAlgorithm()
	res.Value = val
	return err
}

func (s *InternalRpcServer) BootstrapNetwork(req *InternalRpcBootstrapNetworkRequest, res *InternalRpcBootstrapNetworkResponse) error {
	val, err := s.Implementation.BootstrapNetwork(req.P)
	res.Value = val
	return err
}

func (s *InternalRpcServer) RegisterModule(req *InternalRpcRegisterModuleRequest, res *InternalRpcRegisterModuleResponse) error {
	val, err := s.Implementation.RegisterModule(req.M)
	res.Value = val
	return err
}

func (s *InternalRpcServer) RegisterService(req *InternalRpcRegisterServiceRequest, res *InternalRpcRegisterServiceResponse) error {
	val, err := s.Implementation.RegisterService(req.S)
	res.Value = val
	return err
}

func (s *InternalRpcServer) RequestConnection(req *InternalRpcRequestConnectionRequest, res *InternalRpcRequestConnectionResponse) error {
	val, err := s.Implementation.RequestConnection(req.Req)
	res.Value = val
	return err
}

func (s *InternalRpcServer) ShutdownConnection(req *InternalRpcShutdownConnectionRequest, res *InternalRpcShutdownConnectionResponse) error {
	err := s.Implementation.ShutdownConnection(req.Req)
	return err
}

func (s *InternalRpcServer) Status(req *InternalRpcStatusRequest, res *InternalRpcStatusResponse) error {
	val, err := s.Implementation.Status()
	res.Value = val
	return err
}

type InternalRpcBootstrapAlgorithmRequest struct {
}

type InternalRpcBootstrapAlgorithmResponse struct {
	Value *BootstrapRes `thrift:"0,required" json:"value"`
}

type InternalRpcBootstrapNetworkRequest struct {
	P *Common.Peer `thrift:"1,required" json:"p"`
}

type InternalRpcBootstrapNetworkResponse struct {
	Value bool `thrift:"0,required" json:"value"`
}

type InternalRpcRegisterModuleRequest struct {
	M *Module `thrift:"1,required" json:"m"`
}

type InternalRpcRegisterModuleResponse struct {
	Value *RegisterRes `thrift:"0,required" json:"value"`
}

type InternalRpcRegisterServiceRequest struct {
	S *Service `thrift:"1,required" json:"s"`
}

type InternalRpcRegisterServiceResponse struct {
	Value bool `thrift:"0,required" json:"value"`
}

type InternalRpcRequestConnectionRequest struct {
	Req *ConnectionReq `thrift:"1,required" json:"req"`
}

type InternalRpcRequestConnectionResponse struct {
	Value *ConnectionRes `thrift:"0,required" json:"value"`
}

type InternalRpcShutdownConnectionRequest struct {
	Req *ConnectionRes `thrift:"1,required" json:"req"`
}

type InternalRpcShutdownConnectionResponse struct {
}

type InternalRpcStatusRequest struct {
}

type InternalRpcStatusResponse struct {
	Value *StatusRes `thrift:"0,required" json:"value"`
}

type InternalRpcClient struct {
	Client RPCClient
}

func (s *InternalRpcClient) BootstrapAlgorithm() (*BootstrapRes, error) {
	req := &InternalRpcBootstrapAlgorithmRequest{
	}
	res := &InternalRpcBootstrapAlgorithmResponse{}
	err := s.Client.Call("BootstrapAlgorithm", req, res)
	return res.Value, err
}

func (s *InternalRpcClient) BootstrapNetwork(P *Common.Peer) (bool, error) {
	req := &InternalRpcBootstrapNetworkRequest{
		P: P,
	}
	res := &InternalRpcBootstrapNetworkResponse{}
	err := s.Client.Call("BootstrapNetwork", req, res)
	return res.Value, err
}

func (s *InternalRpcClient) RegisterModule(M *Module) (*RegisterRes, error) {
	req := &InternalRpcRegisterModuleRequest{
		M: M,
	}
	res := &InternalRpcRegisterModuleResponse{}
	err := s.Client.Call("RegisterModule", req, res)
	return res.Value, err
}

func (s *InternalRpcClient) RegisterService(S *Service) (bool, error) {
	req := &InternalRpcRegisterServiceRequest{
		S: S,
	}
	res := &InternalRpcRegisterServiceResponse{}
	err := s.Client.Call("RegisterService", req, res)
	return res.Value, err
}

func (s *InternalRpcClient) RequestConnection(Req *ConnectionReq) (*ConnectionRes, error) {
	req := &InternalRpcRequestConnectionRequest{
		Req: Req,
	}
	res := &InternalRpcRequestConnectionResponse{}
	err := s.Client.Call("RequestConnection", req, res)
	return res.Value, err
}

func (s *InternalRpcClient) ShutdownConnection(Req *ConnectionRes) error {
	req := &InternalRpcShutdownConnectionRequest{
		Req: Req,
	}
	res := &InternalRpcShutdownConnectionResponse{}
	err := s.Client.Call("ShutdownConnection", req, res)
	return err
}

func (s *InternalRpcClient) Status() (*StatusRes, error) {
	req := &InternalRpcStatusRequest{
	}
	res := &InternalRpcStatusResponse{}
	err := s.Client.Call("Status", req, res)
	return res.Value, err
}
