package gml

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TestingErrorListener struct {
	*antlr.DefaultErrorListener
	t *testing.T
}

func (el TestingErrorListener) SyntaxError(r antlr.Recognizer, sym interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.t.Errorf("line %d:%d %s", line, column, msg)
}

func TestExamples(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		t.Run(file.Name(), func(t *testing.T) {
			input := antlr.NewFileStream(filepath.Join("testdata", file.Name()))
			lexer := NewgmlLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := NewgmlParser(stream)
			p.AddErrorListener(TestingErrorListener{t: t})
			p.Graph()
		})
	}
}
