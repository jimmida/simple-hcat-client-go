package hcat

import (
  "fmt"
  "apache/hadoop/hive"
)

func TableInfo(table *hive.Table) {
  fmt.Println("TableName       : ", table.GetTableName()       )
  fmt.Println("DbName          : ", table.GetDbName()          )
  fmt.Println("Owner           : ", table.GetOwner()           )
  fmt.Println("CreateTime      : ", table.GetCreateTime()      )
  fmt.Println("LastAccessTime  : ", table.GetLastAccessTime()  )
  fmt.Println("Retention       : ", table.GetRetention()       )
  fmt.Println("Storage Descr   : ")
  SdInfo(table.GetSd())
  fmt.Println("PartitionKeys   : ", table.GetPartitionKeys()   )
  fmt.Println("Parameters      : ", table.GetParameters()      )
  fmt.Println("ViewOriginalText: ", table.GetViewOriginalText())
  fmt.Println("ViewExpandedText: ", table.GetViewExpandedText())
  fmt.Println("TableType       : ", table.GetTableType()       )
  fmt.Println("Privileges      : ", table.GetPrivileges()      )
  fmt.Println("--------------------------")
}

func SdInfo(sd *hive.StorageDescriptor) {
  fmt.Println("\tCols        : ")
  ColsInfo(sd.GetCols())
  fmt.Println("\tLocation    : ", sd.GetLocation()    )
  fmt.Println("\tInputFormat : ", sd.GetInputFormat() )
  fmt.Println("\tOutputFormat: ", sd.GetOutputFormat())
  fmt.Println("\tCompressed  : ", sd.GetCompressed()  )
  fmt.Println("\tNumBuckets  : ", sd.GetNumBuckets()  )
  fmt.Println("\tSerdeInfo   : ")
  SerDeInfo(sd.GetSerdeInfo())
  fmt.Println("\tBucketCols  : ", sd.GetBucketCols()  )
  fmt.Println("\tSortCols    : ", sd.GetSortCols()    )
  fmt.Println("\tParameters  : ", sd.GetParameters()  )
}

func ColsInfo(cols []*hive.FieldSchema) {
  for _, col := range cols {
    fmt.Printf("\t\tName: %v; TypeA1: %v; Comment: %v\n",
      col.GetName(), col.GetTypeA1(), col.GetComment() )
  }
}

func SerDeInfo(serde *hive.SerDeInfo) {
  fmt.Println("\t\tName            : ", serde.GetName            () )
  fmt.Println("\t\tSerializationLib: ", serde.GetSerializationLib() )
  fmt.Println("\t\tParameters      : ", serde.GetParameters      () )
}