package hcat

// Parameters:
//  - DbName
func (h *HCatClient) HGetAllTables(db_name string) (r []string, err error) {
  return h.GetAllTables(db_name)
}

// Parameters:
//  - Dbname
//  - TblName
func (h *HCatClient) HGetTable(dbname string, tbl_name string) (r *HCatTable, err error) {
  table, err := h.GetTable(dbname, tbl_name)
  return &HCatTable{ *table }, err
}