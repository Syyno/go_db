package database

import (
	"mydb/internal/database/compute"
	"mydb/internal/database/storage"
	"mydb/internal/domain/operations"

	"go.uber.org/zap"
)

type Database struct {
	logger   *zap.Logger
	computer *compute.Computer
	storage  *storage.Storage
}

func New(computer *compute.Computer, storage *storage.Storage, logger *zap.Logger) *Database {
	return &Database{
		logger:   logger,
		computer: computer,
		storage:  storage,
	}
}

func (db *Database) Handle(command string) (string, error) {
	res, err := db.computer.Handle(command)
	if err != nil {
		return "", err
	}

	var result string
	switch res.OpId {
	case operations.Get:
		{
			var ok bool
			result, ok = db.storage.Get(res.Args[1])
			if !ok {
				result = "KEY NOT FOUND"
			}
		}
	case operations.Del:
		{
			ok := db.storage.Delete(res.Args[1])
			if ok {
				result = "OK"
			} else {
				result = "KEY NOT FOUND"
			}
		}
	case operations.Set:
		{
			db.storage.Save(res.Args[1], res.Args[2])
			result = "OK"
		}
	}

	return result, nil
}
