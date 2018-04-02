package smgo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

var ErrUnsupportedEncoding = errors.New("Unsupported encoding")

// Parse parses the GO source code from the src io.ReadSeeker and returns a declarations tree *smgo.File.
func Parse(src io.Reader, encoding string) (*File, error) {
	srcBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading src")
	}
	encoding = strings.ToUpper(encoding)
	if encoding != "UTF-8" {
		return nil, ErrUnsupportedEncoding
	}

	fset := token.NewFileSet()
	srcAST, err := parser.ParseFile(fset, "", srcBytes, parser.ParseComments)
	if err != nil {
		file := &File{
			LocationSpan: LocationSpan{
				Start: Location{1, 0},
				End:   Location{1, 0},
			},
			FooterSpan: RuneSpan{0, -1},
			ParsingErrors: []*ParsingError{
				{
					Location: Location{1, 0},
					Message:  err.Error(),
				},
			},
		}
		return file, nil
	}

	fv := &fileVisitor{
		FileSet: fset,
	}
	ast.Walk(fv, srcAST)

	err = fixBlockBoundaries(fv.File, fv.Blocks, srcBytes)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading fixing boundaries")
	}

	return fv.File, nil
}

type fileVisitor struct {
	FileSet *token.FileSet
	File    *File
	Blocks  []block
}

func (v *fileVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.File:
		v.File = createFile(v.FileSet, n)
		v.Blocks = append(v.Blocks, block{
			Type:      nodeBlock,
			Node:      v.File.Nodes[0],
			Container: nil,
		})
		return v
	case *ast.FuncDecl:
		funcNode := createFunc(v.FileSet, n)
		v.File.Nodes = append(v.File.Nodes, funcNode)
		v.Blocks = append(v.Blocks, block{
			Type:      nodeBlock,
			Node:      funcNode,
			Container: nil,
		})
		return nil
	default:
		return nil
	}
}

func createFile(fset *token.FileSet, n *ast.File) *File {
	return &File{
		LocationSpan: LocationSpan{
			Start: Location{
				Line:   fset.Position(n.Pos()).Line,
				Column: fset.Position(n.Pos()).Column,
			},
			End: Location{
				Line:   fset.Position(n.End()).Line,
				Column: fset.Position(n.End()).Column,
			},
		},
		FooterSpan: RuneSpan{
			Start: 0,
			End:   -1,
		},
		Nodes: []*Node{
			{
				Type: PackageNode,
				Name: n.Name.Name,
				LocationSpan: LocationSpan{
					Start: Location{
						Line:   fset.Position(n.Package).Line,
						Column: fset.Position(n.Package).Column,
					},
					End: Location{
						Line:   fset.Position(n.Name.Pos()).Line,
						Column: fset.Position(n.Name.End()).Column,
					},
				},
				Span: RuneSpan{
					Start: fset.Position(n.Package).Offset,
					End:   fset.Position(n.Name.End()).Offset,
				},
			},
		},
	}
}

func createFunc(fset *token.FileSet, n *ast.FuncDecl) *Node {
	return &Node{
		Type: FunctionNode,
		Name: n.Name.Name,
		LocationSpan: LocationSpan{
			Start: Location{
				Line:   fset.Position(n.Pos()).Line,
				Column: fset.Position(n.Pos()).Column,
			},
			End: Location{
				Line:   fset.Position(n.End()).Line,
				Column: fset.Position(n.End()).Column,
			},
		},
		Span: RuneSpan{
			Start: fset.Position(n.Pos()).Offset,
			End:   fset.Position(n.End()).Offset,
		},
	}
}