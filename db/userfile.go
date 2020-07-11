package db

import (
	mydb "./mysql"
	"database/sql"
	"fmt"
	"time"
)

type UserFile struct {
	UserName    string
	FileHash    string
	FileName    string
	FileSize    int64
	UploadAt    string
	LastUpdated string
}

func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user_file (`user_name`,`file_sha1`,`file_name`," +
			"`file_size`,`upload_at`) values (?,?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		return false
	}
	return true
}
