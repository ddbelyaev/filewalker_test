package main

import (
	"log"
	"os"
	"path/filepath"
)

// [TODO] Insert comment
func main() {
	WalkFuncSync()
}

// [TODO] Insert comment
var skipDirMap map[string]bool = map[string]bool{
	".git":   true,
	"vendor": true,
}

// [TODO] Insert comment
const lel string = "1"

// [TODO] Insert comment
const lul = "1"

// [TODO] Insert comment
var kek int

// [TODO] Insert comment
var lol string

// this comment should not be here

// [TODO] Insert comment
// [SOLVED] skips directory
// WalkFuncSync does lalalala
func WalkFuncSync() []string {
	dir := "/home/d/go/src/github.com/ddbelyaev/codoc/"
	files := make([]string, 0)
	visit := func(path string, f os.FileInfo, err error) error {
		log.Println(f.Name())
		// if _, ok := skipDirMap[filepath.Base(f.Name())]; f.IsDir() && ok {
		// 	return filepath.SkipDir
		// }

		if f.IsDir() && path != dir {
			return nil
		}
		// if f.Mode().IsRegular() && filepath.Ext(path) == ".go" {
		// 	fmt.Printf("Visited: %s File name: %s Size: %d bytes\n",
		// 		path, f.Name(), f.Size())
		// 	files = append(files, path)
		// 	return nil
		// }
		return nil
	}

	filepath.Walk(dir, visit)
	return files
}

// [TODO] Insert comment
type kekStruct struct {
	attrName string
}

// [TODO] Insert comment
func (k *kekStruct) lol() error {
	return nil
}

// [TODO] Insert comment
func (*kekStruct) kek() error {
	return nil
}
