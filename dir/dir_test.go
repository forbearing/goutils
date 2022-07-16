package dir

import (
	"os"
	"testing"
)

var ()

func TestDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal("create temporary directory failed: ", err)
	}
	dir2 := dir + "_copy"
	//dir := t.TempDir()
	file, err := os.CreateTemp(dir, "test")
	if err != nil {
		t.Fatal("create temporary file failed: ", err)
	}
	t.Cleanup(func() {
		os.RemoveAll(dir)
		os.RemoveAll(dir2)
	})

	t.Log(IsDir(dir))         // true
	t.Log(IsDir(file.Name())) // false

	os.CreateTemp(dir, "file1")
	os.CreateTemp(dir, "file2")
	os.CreateTemp(dir, "file3")
	os.MkdirTemp(dir, "sub1")
	os.MkdirTemp(dir, "sub2")
	os.MkdirTemp(dir, "sub3")

	dirList, err := GetAll(dir)
	checkErr(t, "GetAll", dirList, err)
	err = Copy(dir2, dir)
	checkErr(t, "Copy", nil, err)
}

func checkErr(t *testing.T, name string, val interface{}, err error) {
	if err != nil {
		t.Logf("%s failed: %v", name, val)
	} else {
		t.Logf("%s success: %v", name, val)
	}
}
