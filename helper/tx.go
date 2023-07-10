package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		PanicIfErrors(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfErrors(errCommit)
	}
}
