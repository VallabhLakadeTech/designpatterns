package main

import "sync"

type DBConnection struct {
	instanceCount int
}

var (
	dbConnection        DBConnection
	runOnce             sync.Once
	GlobalInstanceCount int
)

func GetDBInstance() DBConnection {
	runOnce.Do(func() {
		GlobalInstanceCount = GlobalInstanceCount + 1
		dbConnection = DBConnection{instanceCount: GlobalInstanceCount}
	})
	return dbConnection
}
