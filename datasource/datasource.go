package datasource

type DataBase interface {
	GetTable(tableName string) Table
}

type Table interface {
	Insert(model interface{}) string
	Query(opt map[string]interface{}) []interface{}
	Update(opt map[string]interface{},des map[string]interface{})
	Delete(opt map[string]interface{}) []interface{}
}

func GetDataBase(dbType string,dbName string) DataBase{
	if dbType == "mongoDB"{
		return NewMongoDB(dbName)
	}

	return NewMongoDB(dbName)
}