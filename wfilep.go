package wfilep

import (
	"bytes"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
)

// ファイルが存在するか
func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// 拡張子を返す
//  ex) dir\test.txt
//  ->  txt
func Extension(p string) string {
	e := path.Ext(p)
	// ピリオド以降取得(.txt → txt)
	if i := strings.Index(e, "."); i > -1 {
		e = e[i+1:]
	}
	return e
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

// ファイル名を返す
//  ex) dir\test.txt
//  ->  test.txt
func Filename(p string) string {
	i := strings.LastIndex(p, "\\")
	return p[i+1 : len(p)]
}

// タイトルを返す
//  ex) dir\test.txt <br>
//  ->  test
func Title(p string) string {
	f := Filename(p)
	i := strings.LastIndex(f, ".")
	if -1 == i {
		return ""
	}
	t := f[0:i]
	return t
}

// パスの最後に区切り文字を追加する
func AddBackSlash(p string) string {
	var b bytes.Buffer
	b.WriteString(p)
	if p[len(p)-1:] != "\\" {
		// stringの連結より高速
		b.WriteString("\\")
	}
	return b.String()
}

// 短いパス名を取得する
func ShortName(path string) (string, error) {
	p := syscall.StringToUTF16(path)
	b := p
	n, err := syscall.GetShortPathName(&p[0], &b[0], uint32(len(b)))
	if err != nil {
		return "", err
	}
	if n > uint32(len(b)) {
		b = make([]uint16, n)
		n, err = syscall.GetShortPathName(&p[0], &b[0], uint32(len(b)))
		if err != nil {
			return "", err
		}
	}
	return syscall.UTF16ToString(b), nil
}

// ディレクトリかどうか
func IsDir(p string) (bool, error) {
	f, err := os.Stat(p)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}
