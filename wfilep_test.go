package wfilep_test

import (
	. "github.com/hi-nakamura/wfilep"
	"testing"
)

// ファイルが存在するか
func TestIsExist(t *testing.T) {
	// ファイルが存在する
	if !(IsExist("sample\\testfile.txt")) {
		t.Errorf("ファイルは存在するがfalseが返却")
	}

	// ファイルが存在しない
	if IsExist("sample\\testfile2.txt") {
		t.Errorf("ファイルは存在しないがtrueが返却")
	}
}

// 拡張子を返す
func TestExt(t *testing.T) {
	// 拡張子が存在する
	if "txt" == Ext("sample\\testfile.txt") {
		t.Errorf("拡張子が一致しない")
	}

	// 拡張子が存在しない
	if "" != Ext("sample\\testfile") {
		t.Errorf("拡張子が空以外となる")
	}

	// ファイルが存在しない
	if "" == Ext("sample\\testfile.txt2") {
		t.Errorf("ファイル存在しない場合の動作が不正")
	}
}

// ディレクトリを返す
func TestDir(t *testing.T) {
	// ディレクトリが存在
	if "sample" == Dir("sample\\testfile.txt") {
		t.Errorf("ディレクトリが取得できない")
	}

	// ディレクトリが存在しない
	if "" == Dir("sample\\testfile.txt") {
		t.Errorf("ディレクトリが存在しない場合の動作が不正")
	}
}

// 絶対パスを返す
func TestFullPath(t *testing.T) {
	{
		// ファイルが存在
		_, err := FullPath("sample\\testfile.txt")
		//if f != "c:\\...\\src\\github.com\\hi-nakamura\\wfilep\\sample\\testfile.txt"  {
		//	t.Errorf("取得されたフルパスが異なる [%s]", f)
		//}
		if err != nil {
			t.Errorf("ファイル存在する場合、エラーが返却される")
		}
	}

	{
		// ファイルが存在しない
		_, err := FullPath("sample\\testfile2.txt")
		//if f != "c:\\...\\src\\github.com\\hi-nakamura\\wfilep\\sample\\testfile2.txt"  {
		//	t.Errorf("取得されたフルパスが異なる [%s]", f)
		//}
		if err != nil {
			t.Errorf("ファイル存在しない場合、エラーが返却される")
		}
	}
}
