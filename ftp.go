package main

import (
	"log"
	"os"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"

	ldbauth "github.com/goftp/leveldb-auth"
	ldbperm "github.com/goftp/leveldb-perm"
	"github.com/syndtr/goleveldb/leveldb"
)

func generateDatabase() (auth *ldbauth.LDBAuth, perm *ldbperm.LDBPerm) {
	db, err := leveldb.OpenFile("./ftpdb", nil)
	if err != nil {
		log.Fatalf("FTP DB error %v", err)
	}

	auth = &ldbauth.LDBAuth{db}
	perm = ldbperm.NewLDBPerm(db, "root", "root", os.ModePerm)

	_, usererr := auth.GetUser("test")
	if usererr != nil {
		log.Printf("empty user %v", usererr)
		auth.AddUser("test", "test")
	}
	return
}

func main() {
	ftpdir := "./ftpdir"

	userdatabase, perm := generateDatabase()

	_, err := os.Lstat(ftpdir)
	if os.IsNotExist(err) {
		os.MkdirAll(ftpdir, os.ModePerm)
	}
	factory := &filedriver.FileDriverFactory{
		RootPath: ftpdir,
		Perm:     perm,
	}

	opt := &server.ServerOpts{
		Name:    "127.0.0.1",
		Factory: factory,
		Port:    2121,
		Auth:    userdatabase,
	}

	ftpServer := server.NewServer(opt)
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("FTP server error %v", err)
	}
}
