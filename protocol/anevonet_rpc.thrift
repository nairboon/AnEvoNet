namespace js services
namespace go anevonet_rpc

include "Common.thrift"

struct Module {
  1: string name,
  2: Common.DNA dna,
  3: i32 passive_port
}

struct ConnectionReq {
 1: Common.Peer target,
}

struct ConnectionRes {
 1: string Socket,
}

struct RegisterRes {
 1: string Socket,
 2: Common.DNA DNA,
}

service local_rpc {
  RegisterRes RegisterModule(1:Module m),
  ConnectionRes RequestConnection(1:ConnectionReq m),
}
