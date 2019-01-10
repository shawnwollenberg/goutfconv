package convertutf8

import (
	"bytes"
	"io/ioutil"
	"log"
	"unsafe"

	"github.com/gogits/chardet"
	"github.com/rogpeppe/go-charset/charset"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ReadFileUTF16(filename string) ([]byte, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)
	decoded, err := ioutil.ReadAll(unicodeReader)
	return decoded, err
}

func ReturnData(filename string) (output string) {

	crst := "UTF-16"
	bt, _ := ioutil.ReadFile(filename)
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(bt)
	if err == nil {
		crst = result.Charset
	}

	if crst[0:3] == "UTF" {
		r, err := ReadFileUTF16(filename)
		if err != nil {
			log.Fatal(err)
		}
		x := string(r)
		return x

	} else {
		s, err := charset.NewReader(crst, bytes.NewReader(bt))
		if err != nil {
			log.Fatal(err)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(s)
		b := buf.Bytes()
		x := *(*string)(unsafe.Pointer(&b))
		return x
	}
}

func init() {

}
