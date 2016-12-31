// Generated from tinyc.g4 by ANTLR 4.6.

package tinyc // tinyc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetinycListener is a complete listener for a parse tree produced by tinycParser.
type BasetinycListener struct{}

var _ tinycListener = &BasetinycListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetinycListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetinycListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetinycListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetinycListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasetinycListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasetinycListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasetinycListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasetinycListener) ExitStatement(ctx *StatementContext) {}

// EnterParen_expr is called when production paren_expr is entered.
func (s *BasetinycListener) EnterParen_expr(ctx *Paren_exprContext) {}

// ExitParen_expr is called when production paren_expr is exited.
func (s *BasetinycListener) ExitParen_expr(ctx *Paren_exprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasetinycListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasetinycListener) ExitExpr(ctx *ExprContext) {}

// EnterTest is called when production test is entered.
func (s *BasetinycListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasetinycListener) ExitTest(ctx *TestContext) {}

// EnterSum is called when production sum is entered.
func (s *BasetinycListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production sum is exited.
func (s *BasetinycListener) ExitSum(ctx *SumContext) {}

// EnterTerm is called when production term is entered.
func (s *BasetinycListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasetinycListener) ExitTerm(ctx *TermContext) {}

// EnterId is called when production id is entered.
func (s *BasetinycListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BasetinycListener) ExitId(ctx *IdContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasetinycListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasetinycListener) ExitInteger(ctx *IntegerContext) {}
