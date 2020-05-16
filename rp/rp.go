package rp

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

// OInfo ...
type OInfo struct {
	Name  string
	IsDir bool
}

// Test ...
func Test() {
	fmt.Println("Test")
}

// Test2 ...
func Test2() {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("starting dir: ", wd)
}

// Test3 ...
func Test3() {
	wd, err := os.Getwd()

	files, err := ioutil.ReadDir(wd)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {

		} else {
			fmt.Println(f.Name())
		}
	}
}

// ReadFromPlace ...
func ReadFromPlace(dirname string) {

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Printf("-----> %s is a directory\n", f.Name())
			ReadFromPlace(f.Name())
		} else {
			fmt.Println(f.Name())
		}
	}

}

// Test5 ...
func Test5(root string) {
	var files []string
	err := filepath.Walk(root, func(path string, into os.FileInfo, err error) error {
		fmt.Println(path)
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	/*
		for _, file := range files {
			fmt.Println(file)
		}*/
}

// Algorithm1 ...
func Algorithm1(root string) ([]OInfo, error) {
	var o []OInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		i := OInfo{
			Name:  path,
			IsDir: MyIsDir(path),
		}
		o = append(o, i)
		return nil
	})

	return o, err
}

// RenameWrapper ...
func RenameWrapper(oldName, newName string) error {
	return os.Rename(oldName, newName)
}

// FindAndReplace ...
func FindAndReplace(root, oldName, newName string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if !MyIsDir(path) {
			base := filepath.Base(path)
			if base == oldName {
				RenameWrapper(path, filepath.Join(filepath.Dir(path), newName))
			}
		}
		return nil
	})

}

// MyIsDir ...
func MyIsDir(name string) bool {
	fmt.Println(name)
	fi, err := os.Stat(name)
	
	if err != nil {
		panic(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false

	default:
		panic("not valid")
	}
}

// Test6 ...
func Test6(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, f := range files {

		full := filepath.Join(path, f.Name())

		if f.IsDir() {
			Test6(full)
		}
		fmt.Println(full)
	}
}

// FindAndRename2 ...
func FindAndRename2(root, targetName, newName string) {
	files, err := ioutil.ReadDir(root)

	if err != nil {
		panic(err)
	}

	for _, f := range files {
		full := filepath.Join(root, f.Name())

		if f.IsDir() {
			FindAndRename2(full, targetName, newName)
		} else {
			if f.Name() == targetName {
				newTarget := filepath.Join(root, newName)

				if fileExists(newTarget) {
					fmt.Printf("dir contains %s\n", newName)
					continue
				}

				fmt.Printf("renaming %s to %s\n", full, newTarget)

				os.Rename(full, newTarget)
			}
		}
	}
}

// Exists ...
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil 
	}
	return true, err
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}


// Fix ...
// func Fix(s string) string {
// 	usr, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}


// 	return usr.HomeDir + s[1:]
// }