package hcat

import "apache/hadoop/hive"

type HCatTable struct {
  hive.Table
}

type HCatStorageDescriptor struct {
  hive.StorageDescriptor
}

type HCatFieldSchema struct {
  hive.FieldSchema
}

type HCatSerDeInfo struct {
  hive.SerDeInfo
}

type HCatOrder struct {
  hive.Order
}

type HCatPrincipalPrivilegeSet struct {
  hive.PrincipalPrivilegeSet
}

type HCatPrivilegeGrantInfo struct {
  hive.PrivilegeGrantInfo
}