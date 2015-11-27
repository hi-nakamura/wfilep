package wfilep

import (
	"os"
	"path"
	"path/filepath"
)

// ファイルが存在するか
func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// 拡張子を返す
//  ex) dir\test.txt
//  ->  txt
func Ext(p string) string {
	return path.Ext(p)
}

// ディレクトリを返す
//  ex1) dir1\dir2\test.txt
//  ->   dir1\dir2
func Dir(p string) string {
	return path.Dir(p)
}

// 絶対パスを返す
//  ex) test.txt <br>
//  ->  c:\\dir\test.txt
func FullPath(p string) (string, error) {
	return filepath.Abs(p)
}

/////////////////////////////////////////////////////
// 以下、未実装
/////////////////////////////////////////////////////

// ディレクトリかどうか
func IsDir(p string) bool {
	// TODO
	return false
}

// ファイル名を返す
//  ex) dir\test.txt
//  ->  test.txt
func Filename(p string) string {
	// TODO
	return ""
}

// タイトルを返す
//  ex) dir\test.txt <br>
//  ->  test
func Title(p string) string {
	// TODO
	return ""
}

// 短いパス名を取得する
func ShortName(p string) string {
	// TODO (::GetShortPathName)
	return ""
}

// パスの最後に区切り文字を追加する
func SddBackSlash(p string) string {
	// TODO 最後に「\\」を追加
	return ""
}
