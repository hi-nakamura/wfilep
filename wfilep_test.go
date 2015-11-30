package wfilep_test

import (
	. "github.com/hi-nakamura/wfilep"
	"testing"
)

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

func TestExt(t *testing.T) {
	// 拡張子が存在する
	if "txt" != Extension("sample\\testfile.txt") {
		t.Errorf("拡張子が一致しない")
	}

	// 拡張子が存在しない
	if "" != Extension("sample\\testfile") {
		t.Errorf("拡張子が空以外となる")
	}

	// ファイルが存在しない
	if "" == Extension("sample\\testfile.txt2") {
		t.Errorf("ファイル存在しない場合の動作が不正")
	}
}

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

func TestFilename(t *testing.T) {
	// ファイル指定
	if filename := Filename("sample\\testfile.txt"); "testfile.txt" != filename {
		t.Errorf("ファイル名が取得されない [%s]", filename)
	}

	// ディレクトリ指定
	if filename := Filename("sample"); "sample" != filename {
		t.Errorf("ディレクトリ名が取得されない [%s]", filename)
	}
}

func TestTitle(t *testing.T) {
	// ファイル指定
	if title := Title("sample\\testfile.txt"); title != "testfile" {
		t.Errorf("タイトルが取得されない [%s]", title)
	}

	// ディレクトリ指定
	if title := Title("sample"); title != "" {
		t.Errorf("ディレクトリ名指定で空にならない [%s]", title)
	}
}

func TestAddBackSlash(t *testing.T) {
	// バックスラッシュ必要
	if dir := AddBackSlash("sample"); dir != "sample\\" {
		t.Errorf("バックスラッシュ必要時の動作不正 [%s]", dir)
	}

	// バックスラッシュ不要
	if dir := AddBackSlash("sample\\"); dir != "sample\\" {
		t.Errorf("バックスラッシュ不要時の動作不正 [%s]", dir)
	}
}

func TestShortName(t *testing.T) {
	{
		// 存在しないパス
		p, err := ShortName("sample\\testfile2.txt")
		if p != "" {
			t.Errorf("存在しないパスで名前が取得される [%s]", p)
		}
		if err == nil {
			t.Errorf("存在しないパスでエラーが発生しない")
		}
	}

	{
		// 存在するパス(短いパス)
		p, err := ShortName("sample\\testfile.txt")
		if p != "sample\\testfile.txt" {
			t.Errorf("短いパスで正しく取得されない [%s]", p)
		}
		if err != nil {
			t.Errorf("短いパスでエラーが発生 [%s]", err)
		}
	}

	{
		// 存在するパス(長いパス)
		p, err := ShortName("sample\\aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.txt")
		if p != "sample\\AAAAAA~1.TXT" {
			t.Errorf("長いパスで正しく取得されない [%s]", p)
		}
		if err != nil {
			t.Errorf("長いパスでエラーが発生 [%s]", err)
		}
	}
}

func TestIsDir(t *testing.T) {
	{
		// 存在しないディレクトリ
		b, err := IsDir("sample2")
		if b {
			t.Errorf("存在しないディレクトリの結果が不正")
		}
		if err == nil {
			t.Errorf("存在しないディレクトリでエラー発生しない")
		}
	}

	{
		// 存在しないファイルパス
		b, err := IsDir("sample\\testfile2.txt")
		if b {
			t.Errorf("存在しないファイルパスの結果が不正")
		}
		if err == nil {
			t.Errorf("存在しないファイルパスでエラー発生しない")
		}
	}

	{
		// 存在するディレクトリ
		b, err := IsDir("sample")
		if !b {
			t.Errorf("存在するディレクトリの結果が不正")
		}
		if err != nil {
			t.Errorf("存在するディレクトリでエラー発生 [%s]", err)
		}
	}

	{
		// 存在するファイルパス
		b, err := IsDir("sample\\testfile.txt")
		if b {
			t.Errorf("存在するファイルパスの結果が不正")
		}
		if err != nil {
			t.Errorf("存在するファイルパスでエラー発生 [%s]", err)
		}
	}
}
