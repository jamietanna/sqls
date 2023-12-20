package formatter

import (
	"github.com/sqls-server/sqls/ast"
	"github.com/sqls-server/sqls/token"
)

func unshift(slice []ast.Node, node ...ast.Node) []ast.Node {
	return append(node, slice...)
}

var whitespaceNode = ast.NewItem(&token.Token{
	Kind:  token.Whitespace,
	Value: " ",
})

func whiteSpaceNodes(num int) []ast.Node {
	res := make([]ast.Node, num)
	for i := 0; i < num; i++ {
		res[i] = whitespaceNode
	}
	return res
}

var linebreakNode = &ast.LineBreak{
	Toks: []ast.Node{linebreakItem},
}

var linebreakItem = ast.NewItem(&token.Token{
	Kind:  token.Whitespace,
	Value: "\n",
})

var tabItem = ast.NewItem(&token.Token{
	Kind:  token.Whitespace,
	Value: "\t",
})

var periodItem = ast.NewItem(&token.Token{
	Kind:  token.Period,
	Value: ".",
})

var lparenItem = ast.NewItem(&token.Token{
	Kind:  token.LParen,
	Value: "(",
})

var rparenItem = ast.NewItem(&token.Token{
	Kind:  token.RParen,
	Value: ")",
})

var commaItem = ast.NewItem(&token.Token{
	Kind:  token.Comma,
	Value: ",",
})
