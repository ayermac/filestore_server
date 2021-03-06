package meta

import (
	"filestore_server/db"
)

// FileMeta: 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta: 新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// UpdateFileMetaDB: 新增/更新文件元信息到MySQL中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return db.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// GetFIleMetaDB: 从MySQL获取文件元信息
func GetFIleMetaDB(filesha1 string) (*FileMeta, error) {
	tfile, err := db.GetFileMeta(filesha1)
	if err != nil {
		return nil, err
	}

	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}

	return &fmeta, nil
}

// UpdateFileMeta: 通过sha1值i获取文件元信息
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// 删除元信息
func RemoveFileMeta(filesha1 string) {
	delete(fileMetas, filesha1)
}
