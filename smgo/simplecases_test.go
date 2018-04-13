package smgo_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jriquelme/SemanticMergeGO/smgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newLocationSpan(startLine, startColumn, endLine, endColumn int) smgo.LocationSpan {
	return smgo.LocationSpan{
		Start: smgo.Location{
			Line:   startLine,
			Column: startColumn,
		},
		End: smgo.Location{
			Line:   endLine,
			Column: endColumn,
		},
	}
}

func TestParseEmpty(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		smgo.PrintBlocks = true
	}

	src := bytes.NewReader([]byte{})
	file, err := smgo.Parse(src, "UTF-8")
	assert.NotNil(t, file)
	assert.Nil(t, err)

	assert.Equal(t, &smgo.File{
		LocationSpan: newLocationSpan(1, 0, 1, 0),
		FooterSpan:   smgo.RuneSpan{0, -1},
		Containers:   nil,
		Nodes:        nil,
		ParsingErrors: []*smgo.ParsingError{
			{
				Location: smgo.Location{1, 0},
				Message:  "1:1: expected 'package', found 'EOF'",
			},
		},
	}, file)
	if t.Failed() {
		spew.Dump(t.Name(), file)
	}
}

func TestParseSimpleConst(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		smgo.PrintBlocks = true
	}

	simpleConst, err := os.Open("testdata/simple_const.go")
	require.Nil(t, err)
	defer simpleConst.Close()

	file, err := smgo.Parse(simpleConst, "UTF-8")
	assert.NotNil(t, file)
	assert.Nil(t, err)

	assert.Equal(t, &smgo.File{
		LocationSpan: newLocationSpan(1, 0, 5, 25),
		FooterSpan:   smgo.RuneSpan{0, -1},
		Containers:   nil,
		Nodes: []*smgo.Node{
			{
				Type:         smgo.PackageNode,
				Name:         "simpleconst",
				LocationSpan: newLocationSpan(1, 0, 1, 20),
				Span:         smgo.RuneSpan{0, 19},
			},
			{
				Type:         smgo.ConstNode,
				Name:         "N",
				LocationSpan: newLocationSpan(2, 0, 3, 12),
				Span:         smgo.RuneSpan{20, 32},
			},
			{
				Type:         smgo.ConstNode,
				Name:         "Name",
				LocationSpan: newLocationSpan(4, 0, 5, 25),
				Span:         smgo.RuneSpan{33, 58},
			},
		},
		ParsingErrors: nil,
	}, file)
	if t.Failed() {
		spew.Dump(t.Name(), file)
	}
}

func TestParseSimpleFunc(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		smgo.PrintBlocks = true
	}

	simpleFunc, err := os.Open("testdata/simple_func.go")
	require.Nil(t, err)
	defer simpleFunc.Close()

	file, err := smgo.Parse(simpleFunc, "UTF-8")
	assert.NotNil(t, file)
	assert.Nil(t, err)

	assert.Equal(t, &smgo.File{
		LocationSpan: newLocationSpan(1, 0, 5, 2),
		FooterSpan:   smgo.RuneSpan{0, -1},
		Containers:   nil,
		Nodes: []*smgo.Node{
			{
				Type:         smgo.PackageNode,
				Name:         "simplefunc",
				LocationSpan: newLocationSpan(1, 0, 1, 19),
				Span:         smgo.RuneSpan{0, 18},
			},
			{
				Type:         smgo.FunctionNode,
				Name:         "Hi",
				LocationSpan: newLocationSpan(2, 0, 5, 2),
				Span:         smgo.RuneSpan{19, 47},
			},
		},
		ParsingErrors: nil,
	}, file)
	if t.Failed() {
		spew.Dump(t.Name(), file)
	}
}

func TestParseSimpleImport(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		smgo.PrintBlocks = true
	}

	simpleImport, err := os.Open("testdata/simple_import.go_src")
	require.Nil(t, err)
	defer simpleImport.Close()

	file, err := smgo.Parse(simpleImport, "UTF-8")
	assert.NotNil(t, file)
	assert.Nil(t, err)

	assert.Equal(t, &smgo.File{
		LocationSpan: newLocationSpan(1, 0, 3, 13),
		FooterSpan:   smgo.RuneSpan{0, -1},
		Containers:   nil,
		Nodes: []*smgo.Node{
			{
				Type:         smgo.PackageNode,
				Name:         "simpleimport",
				LocationSpan: newLocationSpan(1, 0, 1, 21),
				Span:         smgo.RuneSpan{0, 20},
			},
			{
				Type:         smgo.ImportNode,
				Name:         "fmt",
				LocationSpan: newLocationSpan(2, 0, 3, 13),
				Span:         smgo.RuneSpan{21, 34},
			},
		},
		ParsingErrors: nil,
	}, file)
	if t.Failed() {
		spew.Dump(t.Name(), file)
	}

}

func TestParseSimpleStruct(t *testing.T) {
	t.Parallel()
	if testing.Verbose() {
		smgo.PrintBlocks = true
	}

	simpleStruct, err := os.Open("testdata/simple_struct.go")
	require.Nil(t, err)
	defer simpleStruct.Close()

	file, err := smgo.Parse(simpleStruct, "UTF-8")
	assert.NotNil(t, file)
	assert.Nil(t, err)

	assert.Equal(t, &smgo.File{
		LocationSpan: newLocationSpan(1, 0, 9, 2),
		FooterSpan:   smgo.RuneSpan{0, -1},
		Containers: []*smgo.Container{
			{
				Type:         smgo.StructContainer,
				Name:         "Person",
				LocationSpan: newLocationSpan(2, 0, 5, 2),
				HeaderSpan:   smgo.RuneSpan{21, 42},
				FooterSpan:   smgo.RuneSpan{56, 57},
				Containers:   nil,
				Nodes: []*smgo.Node{
					{
						Type:         smgo.FieldNode,
						Name:         "Name",
						LocationSpan: newLocationSpan(4, 0, 4, 13),
						Span:         smgo.RuneSpan{43, 55},
					},
				},
			},
		},
		Nodes: []*smgo.Node{
			{
				Type:         smgo.PackageNode,
				Name:         "simplestruct",
				LocationSpan: newLocationSpan(1, 0, 1, 21),
				Span:         smgo.RuneSpan{0, 20},
			},
			{
				Type:         smgo.FunctionNode,
				Name:         "SayHi",
				LocationSpan: newLocationSpan(6, 0, 9, 2),
				Span:         smgo.RuneSpan{58, 115},
			},
		},
		ParsingErrors: nil,
	}, file)
	if t.Failed() {
		spew.Dump(t.Name(), file)
	}

}
