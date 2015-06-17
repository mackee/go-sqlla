package genbase

import (
	"testing"
)

func TestParserParsePackageDir(t *testing.T) {
	p := &Parser{}
	pInfo, err := p.ParsePackageDir("./misc/fixture/a")
	if err != nil {
		t.Fatal(err)
	}

	if len(pInfo.Files) != 1 {
		t.Fatalf("unexpected: %d", len(pInfo.Files))
	}
}

func TestParserParsePackageFiles(t *testing.T) {
	p := &Parser{}
	pInfo, err := p.ParsePackageFiles([]string{"./misc/fixture/a/model.go"})
	if err != nil {
		t.Fatal(err)
	}

	if len(pInfo.Files) != 1 {
		t.Fatalf("unexpected: %d", len(pInfo.Files))
	}
}

func TestPackageInfoTypeInfos(t *testing.T) {
	p := &Parser{}
	pInfo, err := p.ParsePackageDir("./misc/fixture/a")
	if err != nil {
		t.Fatal(err)
	}

	tis := pInfo.TypeInfos()
	if len(tis) != 3 {
		for _, ti := range tis {
			t.Log(ti.Name())
		}
		t.Fatalf("unexpected: %d", len(tis))
	}
}

func TestPackageInfoCollectTaggedTypeInfos(t *testing.T) {
	p := &Parser{}
	pInfo, err := p.ParsePackageDir("./misc/fixture/a")
	if err != nil {
		t.Fatal(err)
	}

	tis := pInfo.CollectTaggedTypeInfos("+test")
	if len(tis) != 3 {
		for _, ti := range tis {
			t.Log(ti.Name())
		}
		t.Fatalf("unexpected: %d", len(tis))
	}
}

func TestPackageInfoCollectTypeInfos(t *testing.T) {
	p := &Parser{}
	pInfo, err := p.ParsePackageDir("./misc/fixture/a")
	if err != nil {
		t.Fatal(err)
	}

	tis := pInfo.CollectTypeInfos([]string{"C"})
	if len(tis) != 1 {
		for _, ti := range tis {
			t.Log(ti.Name())
		}
		t.Fatalf("unexpected: %d", len(tis))
	}
}
