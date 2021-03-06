import logging
import os
import shutil
import subprocess
import xml.etree.ElementTree

from string import Template


def clone():
    if not os.path.exists("_build/grammars"):
        logging.info("Cloning https://github.com/antlr/grammars-v4...")
        subprocess.check_call([
            "git",
            "clone",
            "https://github.com/antlr/grammars-v4",
            "_build/grammars",
        ])

test_template = Template("""package $package

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
			lexer := New${lexer}Lexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := New${parser}Parser(stream)
			p.AddErrorListener(TestingErrorListener{t: t})
			p.${entrypoint}()
		})
	}
}""")

SUPPORTED_GRAMMARS = set([
    "json",
    "tinyc",
    "gml",
    "csv",
])

def upperfirst(x):
    return x[0].upper() + x[1:]

def main():
    logging.basicConfig(level=logging.INFO)

    # Download jar
    clone()

    for pl in os.listdir("_build/grammars"):
        if pl not in SUPPORTED_GRAMMARS:
            continue
        fullpath = os.path.join("_build", "grammars", pl)
        pkg = pl.replace("-", "")
        if os.path.isfile(fullpath):
            continue
        if os.path.exists(pkg):
            continue

        for grammar in os.listdir(fullpath):
            if not grammar.endswith(".g4"):
                continue
            logging.info("Building {}".format(pkg))
            
            # Create the directory
            os.makedirs(pkg)

            # Generate the code
            subprocess.call([
                "java",
                "-jar", "../../antlr-4.6-complete.jar",
                "-package", pkg,
                "-visitor",
                "-o", os.path.join("..", "..", "..", pkg),
                "-Dlanguage=Go",
                grammar,
            ], cwd=fullpath)

        examples = os.path.join(fullpath, "examples")
        if not os.path.exists(examples):
            continue

        # Copy over the examples
        shutil.copytree(
            examples,
            os.path.join(pkg,"testdata"),
        )

        # Generate the test file
        testfile = os.path.join(pkg, "{}_parser_test.go".format(pkg))

        # Find the entry point
        pomfile = os.path.join(fullpath, "pom.xml")
        pom = xml.etree.ElementTree.parse(pomfile).getroot()
        entry_point = pom.find(".//{http://maven.apache.org/POM/4.0.0}entryPoint")
        grammar_name = pom.find(".//{http://maven.apache.org/POM/4.0.0}grammarName")

        with open(testfile,'w') as f:
            f.write(test_template.substitute(
                package=pkg,
                lexer=grammar_name.text,
                parser=grammar_name.text,
                entrypoint=upperfirst(entry_point.text),
            ))

if __name__ == "__main__":
    main()
