package main

import (
  "fmt"
  "reflect"
  "hcat"
)

func main() {
  addr := "10.103.219.33:9083"
  hcatClient, err := hcat.NewHCatClient(addr)
  if err != nil {
    fmt.Println("Error creating HCatClient:", err)
    return
  }

  if err := hcatClient.Open(); err != nil {
    fmt.Println("Error opening HCatClient connection:", err)
    return
  }

  tables, err := hcatClient.HGetAllTables("default")
  fmt.Println(tables)

  hcatTable, err := hcatClient.HGetTable("default", "employees")
  hcat.TableInfo(&hcatTable.Table)
  fmt.Println(hcatTable)
  fmt.Println(reflect.TypeOf(hcatTable).Kind())
  fmt.Println(*hcatTable.Sd.StoredAsSubDirectories)

  hcatClient.Close()
}
