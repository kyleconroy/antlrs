// Generated from JSON.g4 by ANTLR 4.6.

package json // JSON
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JSONParser.
type JSONVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JSONParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by JSONParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by JSONParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by JSONParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by JSONParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
