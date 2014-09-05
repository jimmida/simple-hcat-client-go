package hcat

import (
  "git.apache.org/thrift.git/lib/go/thrift"
  "apache/hadoop/hive"
)

type HCatClient struct {
  // HCat *hive.ThriftHiveMetastoreClient
  hive.ThriftHiveMetastoreClient
  Socket *thrift.TSocket
}

func (h *HCatClient) Open() error {
  if err := h.Socket.Open(); err != nil {
   return err
  }
  return nil
}

func (h *HCatClient) Close() {
  h.Socket.Close()
}

func NewHCatClient(addr string) (*HCatClient, error) {
  transportFactory := thrift.NewTTransportFactory()
  protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
  socket, err := thrift.NewTSocket(addr)
  if err != nil {
    return nil, err
  }

  transport := transportFactory.GetTransport(socket)

  mstore := hive.NewThriftHiveMetastoreClientFactory(transport, protocolFactory)
  // return &HCatClient{ mstore, socket }, nil
  return &HCatClient{ *mstore, socket }, nil
}

