// Code generated from Excellent2.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gen // Excellent2
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 88, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 40, 10, 3, 12, 3, 14, 3, 43, 11, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 53, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 73, 10, 4, 12, 4, 14, 4, 76, 11, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 7, 6, 83, 10, 6, 12, 6, 14, 6, 86, 11, 6, 3, 6, 2, 4, 4,
	6, 7, 2, 4, 6, 8, 10, 2, 6, 3, 2, 11, 12, 3, 2, 9, 10, 3, 2, 16, 19, 3,
	2, 14, 15, 2, 100, 2, 12, 3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 52, 3, 2, 2,
	2, 8, 77, 3, 2, 2, 2, 10, 79, 3, 2, 2, 2, 12, 13, 5, 6, 4, 2, 13, 14, 7,
	2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 16, 8, 3, 1, 2, 16, 17, 5, 8, 5, 2, 17,
	19, 7, 4, 2, 2, 18, 20, 5, 10, 6, 2, 19, 18, 3, 2, 2, 2, 19, 20, 3, 2,
	2, 2, 20, 21, 3, 2, 2, 2, 21, 22, 7, 5, 2, 2, 22, 30, 3, 2, 2, 2, 23, 30,
	7, 26, 2, 2, 24, 30, 7, 21, 2, 2, 25, 30, 7, 22, 2, 2, 26, 30, 7, 23, 2,
	2, 27, 30, 7, 24, 2, 2, 28, 30, 7, 25, 2, 2, 29, 15, 3, 2, 2, 2, 29, 23,
	3, 2, 2, 2, 29, 24, 3, 2, 2, 2, 29, 25, 3, 2, 2, 2, 29, 26, 3, 2, 2, 2,
	29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 41, 3, 2, 2, 2, 31, 32, 12,
	10, 2, 2, 32, 33, 7, 8, 2, 2, 33, 40, 5, 4, 3, 11, 34, 35, 12, 9, 2, 2,
	35, 36, 7, 6, 2, 2, 36, 37, 5, 6, 4, 2, 37, 38, 7, 7, 2, 2, 38, 40, 3,
	2, 2, 2, 39, 31, 3, 2, 2, 2, 39, 34, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41,
	39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 5, 3, 2, 2, 2, 43, 41, 3, 2, 2,
	2, 44, 45, 8, 4, 1, 2, 45, 53, 5, 4, 3, 2, 46, 47, 7, 10, 2, 2, 47, 53,
	5, 6, 4, 10, 48, 49, 7, 4, 2, 2, 49, 50, 5, 6, 4, 2, 50, 51, 7, 5, 2, 2,
	51, 53, 3, 2, 2, 2, 52, 44, 3, 2, 2, 2, 52, 46, 3, 2, 2, 2, 52, 48, 3,
	2, 2, 2, 53, 74, 3, 2, 2, 2, 54, 55, 12, 9, 2, 2, 55, 56, 7, 13, 2, 2,
	56, 73, 5, 6, 4, 10, 57, 58, 12, 8, 2, 2, 58, 59, 9, 2, 2, 2, 59, 73, 5,
	6, 4, 9, 60, 61, 12, 7, 2, 2, 61, 62, 9, 3, 2, 2, 62, 73, 5, 6, 4, 8, 63,
	64, 12, 6, 2, 2, 64, 65, 9, 4, 2, 2, 65, 73, 5, 6, 4, 7, 66, 67, 12, 5,
	2, 2, 67, 68, 9, 5, 2, 2, 68, 73, 5, 6, 4, 6, 69, 70, 12, 4, 2, 2, 70,
	71, 7, 20, 2, 2, 71, 73, 5, 6, 4, 5, 72, 54, 3, 2, 2, 2, 72, 57, 3, 2,
	2, 2, 72, 60, 3, 2, 2, 2, 72, 63, 3, 2, 2, 2, 72, 66, 3, 2, 2, 2, 72, 69,
	3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2,
	75, 7, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 78, 7, 26, 2, 2, 78, 9, 3, 2,
	2, 2, 79, 84, 5, 6, 4, 2, 80, 81, 7, 3, 2, 2, 81, 83, 5, 6, 4, 2, 82, 80,
	3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2,
	85, 11, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 10, 19, 29, 39, 41, 52, 72, 74,
	84,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'('", "')'", "'['", "']'", "'.'", "'+'", "'-'", "'*'", "'/'",
	"'^'", "'='", "'!='", "'<='", "'<'", "'>='", "'>'", "'&'",
}
var symbolicNames = []string{
	"", "COMMA", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "DOT", "PLUS", "MINUS",
	"TIMES", "DIVIDE", "EXPONENT", "EQ", "NEQ", "LTE", "LT", "GTE", "GT", "AMPERSAND",
	"TEXT", "NUMBER", "TRUE", "FALSE", "NULL", "NAME", "WS", "ERROR",
}

var ruleNames = []string{
	"parse", "atom", "expression", "fnname", "parameters",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Excellent2Parser struct {
	*antlr.BaseParser
}

func NewExcellent2Parser(input antlr.TokenStream) *Excellent2Parser {
	this := new(Excellent2Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Excellent2.g4"

	return this
}

// Excellent2Parser tokens.
const (
	Excellent2ParserEOF       = antlr.TokenEOF
	Excellent2ParserCOMMA     = 1
	Excellent2ParserLPAREN    = 2
	Excellent2ParserRPAREN    = 3
	Excellent2ParserLBRACK    = 4
	Excellent2ParserRBRACK    = 5
	Excellent2ParserDOT       = 6
	Excellent2ParserPLUS      = 7
	Excellent2ParserMINUS     = 8
	Excellent2ParserTIMES     = 9
	Excellent2ParserDIVIDE    = 10
	Excellent2ParserEXPONENT  = 11
	Excellent2ParserEQ        = 12
	Excellent2ParserNEQ       = 13
	Excellent2ParserLTE       = 14
	Excellent2ParserLT        = 15
	Excellent2ParserGTE       = 16
	Excellent2ParserGT        = 17
	Excellent2ParserAMPERSAND = 18
	Excellent2ParserTEXT      = 19
	Excellent2ParserNUMBER    = 20
	Excellent2ParserTRUE      = 21
	Excellent2ParserFALSE     = 22
	Excellent2ParserNULL      = 23
	Excellent2ParserNAME      = 24
	Excellent2ParserWS        = 25
	Excellent2ParserERROR     = 26
)

// Excellent2Parser rules.
const (
	Excellent2ParserRULE_parse      = 0
	Excellent2ParserRULE_atom       = 1
	Excellent2ParserRULE_expression = 2
	Excellent2ParserRULE_fnname     = 3
	Excellent2ParserRULE_parameters = 4
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Excellent2ParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Excellent2ParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Excellent2Parser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Excellent2ParserRULE_parse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.expression(0)
	}
	{
		p.SetState(11)
		p.Match(Excellent2ParserEOF)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Excellent2ParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Excellent2ParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DotLookupContext struct {
	*AtomContext
}

func NewDotLookupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotLookupContext {
	var p = new(DotLookupContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *DotLookupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotLookupContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *DotLookupContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *DotLookupContext) DOT() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserDOT, 0)
}

func (s *DotLookupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterDotLookup(s)
	}
}

func (s *DotLookupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitDotLookup(s)
	}
}

func (s *DotLookupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitDotLookup(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	*AtomContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserNULL, 0)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	*AtomContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) Fnname() IFnnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFnnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFnnameContext)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserRPAREN, 0)
}

func (s *FunctionCallContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueContext struct {
	*AtomContext
}

func NewTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueContext {
	var p = new(TrueContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserTRUE, 0)
}

func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitTrue(s)
	}
}

func (s *TrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseContext struct {
	*AtomContext
}

func NewFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseContext {
	var p = new(FalseContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserFALSE, 0)
}

func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitFalse(s)
	}
}

func (s *FalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLookupContext struct {
	*AtomContext
}

func NewArrayLookupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLookupContext {
	var p = new(ArrayLookupContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ArrayLookupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLookupContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ArrayLookupContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserLBRACK, 0)
}

func (s *ArrayLookupContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLookupContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserRBRACK, 0)
}

func (s *ArrayLookupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterArrayLookup(s)
	}
}

func (s *ArrayLookupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitArrayLookup(s)
	}
}

func (s *ArrayLookupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitArrayLookup(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContextReferenceContext struct {
	*AtomContext
}

func NewContextReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContextReferenceContext {
	var p = new(ContextReferenceContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ContextReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextReferenceContext) NAME() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserNAME, 0)
}

func (s *ContextReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterContextReference(s)
	}
}

func (s *ContextReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitContextReference(s)
	}
}

func (s *ContextReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitContextReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type TextLiteralContext struct {
	*AtomContext
}

func NewTextLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextLiteralContext {
	var p = new(TextLiteralContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *TextLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextLiteralContext) TEXT() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserTEXT, 0)
}

func (s *TextLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterTextLiteral(s)
	}
}

func (s *TextLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitTextLiteral(s)
	}
}

func (s *TextLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitTextLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberLiteralContext struct {
	*AtomContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserNUMBER, 0)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Excellent2Parser) Atom() (localctx IAtomContext) {
	return p.atom(0)
}

func (p *Excellent2Parser) atom(_p int) (localctx IAtomContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAtomContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAtomContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, Excellent2ParserRULE_atom, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(14)
			p.Fnname()
		}
		{
			p.SetState(15)
			p.Match(Excellent2ParserLPAREN)
		}
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Excellent2ParserLPAREN)|(1<<Excellent2ParserMINUS)|(1<<Excellent2ParserTEXT)|(1<<Excellent2ParserNUMBER)|(1<<Excellent2ParserTRUE)|(1<<Excellent2ParserFALSE)|(1<<Excellent2ParserNULL)|(1<<Excellent2ParserNAME))) != 0 {
			{
				p.SetState(16)
				p.Parameters()
			}

		}
		{
			p.SetState(19)
			p.Match(Excellent2ParserRPAREN)
		}

	case 2:
		localctx = NewContextReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(21)
			p.Match(Excellent2ParserNAME)
		}

	case 3:
		localctx = NewTextLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(Excellent2ParserTEXT)
		}

	case 4:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(Excellent2ParserNUMBER)
		}

	case 5:
		localctx = NewTrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(Excellent2ParserTRUE)
		}

	case 6:
		localctx = NewFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(Excellent2ParserFALSE)
		}

	case 7:
		localctx = NewNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Match(Excellent2ParserNULL)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDotLookupContext(p, NewAtomContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_atom)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(30)
					p.Match(Excellent2ParserDOT)
				}
				{
					p.SetState(31)
					p.atom(9)
				}

			case 2:
				localctx = NewArrayLookupContext(p, NewAtomContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_atom)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(33)
					p.Match(Excellent2ParserLBRACK)
				}
				{
					p.SetState(34)
					p.expression(0)
				}
				{
					p.SetState(35)
					p.Match(Excellent2ParserRBRACK)
				}

			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Excellent2ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Excellent2ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParenthesesContext struct {
	*ExpressionContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserLPAREN, 0)
}

func (s *ParenthesesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserRPAREN, 0)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitParentheses(s)
	}
}

func (s *ParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegationContext struct {
	*ExpressionContext
}

func NewNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationContext {
	var p = new(NegationContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserMINUS, 0)
}

func (s *NegationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitNegation(s)
	}
}

func (s *NegationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitNegation(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ComparisonContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonContext) LTE() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserLTE, 0)
}

func (s *ComparisonContext) LT() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserLT, 0)
}

func (s *ComparisonContext) GTE() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserGTE, 0)
}

func (s *ComparisonContext) GT() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserGT, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConcatenationContext struct {
	*ExpressionContext
}

func NewConcatenationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatenationContext {
	var p = new(ConcatenationContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConcatenationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConcatenationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConcatenationContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserAMPERSAND, 0)
}

func (s *ConcatenationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterConcatenation(s)
	}
}

func (s *ConcatenationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitConcatenation(s)
	}
}

func (s *ConcatenationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitConcatenation(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicationOrDivisionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMultiplicationOrDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationOrDivisionContext {
	var p = new(MultiplicationOrDivisionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicationOrDivisionContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicationOrDivisionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicationOrDivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationOrDivisionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicationOrDivisionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplicationOrDivisionContext) TIMES() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserTIMES, 0)
}

func (s *MultiplicationOrDivisionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserDIVIDE, 0)
}

func (s *MultiplicationOrDivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterMultiplicationOrDivision(s)
	}
}

func (s *MultiplicationOrDivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitMultiplicationOrDivision(s)
	}
}

func (s *MultiplicationOrDivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitMultiplicationOrDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomReferenceContext struct {
	*ExpressionContext
}

func NewAtomReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomReferenceContext {
	var p = new(AtomReferenceContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AtomReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomReferenceContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterAtomReference(s)
	}
}

func (s *AtomReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitAtomReference(s)
	}
}

func (s *AtomReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitAtomReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditionOrSubtractionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAdditionOrSubtractionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionOrSubtractionContext {
	var p = new(AdditionOrSubtractionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AdditionOrSubtractionContext) GetOp() antlr.Token { return s.op }

func (s *AdditionOrSubtractionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditionOrSubtractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionOrSubtractionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AdditionOrSubtractionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AdditionOrSubtractionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserPLUS, 0)
}

func (s *AdditionOrSubtractionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserMINUS, 0)
}

func (s *AdditionOrSubtractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterAdditionOrSubtraction(s)
	}
}

func (s *AdditionOrSubtractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitAdditionOrSubtraction(s)
	}
}

func (s *AdditionOrSubtractionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitAdditionOrSubtraction(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualityContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityContext {
	var p = new(EqualityContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityContext) GetOp() antlr.Token { return s.op }

func (s *EqualityContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualityContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityContext) EQ() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserEQ, 0)
}

func (s *EqualityContext) NEQ() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserNEQ, 0)
}

func (s *EqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterEquality(s)
	}
}

func (s *EqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitEquality(s)
	}
}

func (s *EqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExponentContext struct {
	*ExpressionContext
}

func NewExponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentContext {
	var p = new(ExponentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExponentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExponentContext) EXPONENT() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserEXPONENT, 0)
}

func (s *ExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterExponent(s)
	}
}

func (s *ExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitExponent(s)
	}
}

func (s *ExponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitExponent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Excellent2Parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *Excellent2Parser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, Excellent2ParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Excellent2ParserTEXT, Excellent2ParserNUMBER, Excellent2ParserTRUE, Excellent2ParserFALSE, Excellent2ParserNULL, Excellent2ParserNAME:
		localctx = NewAtomReferenceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(43)
			p.atom(0)
		}

	case Excellent2ParserMINUS:
		localctx = NewNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(Excellent2ParserMINUS)
		}
		{
			p.SetState(45)
			p.expression(8)
		}

	case Excellent2ParserLPAREN:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(46)
			p.Match(Excellent2ParserLPAREN)
		}
		{
			p.SetState(47)
			p.expression(0)
		}
		{
			p.SetState(48)
			p.Match(Excellent2ParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_expression)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(53)
					p.Match(Excellent2ParserEXPONENT)
				}
				{
					p.SetState(54)
					p.expression(8)
				}

			case 2:
				localctx = NewMultiplicationOrDivisionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_expression)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(56)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicationOrDivisionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Excellent2ParserTIMES || _la == Excellent2ParserDIVIDE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicationOrDivisionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(57)
					p.expression(7)
				}

			case 3:
				localctx = NewAdditionOrSubtractionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_expression)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(59)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditionOrSubtractionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Excellent2ParserPLUS || _la == Excellent2ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditionOrSubtractionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(60)
					p.expression(6)
				}

			case 4:
				localctx = NewComparisonContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(62)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Excellent2ParserLTE)|(1<<Excellent2ParserLT)|(1<<Excellent2ParserGTE)|(1<<Excellent2ParserGT))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(63)
					p.expression(5)
				}

			case 5:
				localctx = NewEqualityContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_expression)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(65)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Excellent2ParserEQ || _la == Excellent2ParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(66)
					p.expression(4)
				}

			case 6:
				localctx = NewConcatenationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Excellent2ParserRULE_expression)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(68)
					p.Match(Excellent2ParserAMPERSAND)
				}
				{
					p.SetState(69)
					p.expression(3)
				}

			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IFnnameContext is an interface to support dynamic dispatch.
type IFnnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnnameContext differentiates from other interfaces.
	IsFnnameContext()
}

type FnnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnnameContext() *FnnameContext {
	var p = new(FnnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Excellent2ParserRULE_fnname
	return p
}

func (*FnnameContext) IsFnnameContext() {}

func NewFnnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnnameContext {
	var p = new(FnnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Excellent2ParserRULE_fnname

	return p
}

func (s *FnnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FnnameContext) NAME() antlr.TerminalNode {
	return s.GetToken(Excellent2ParserNAME, 0)
}

func (s *FnnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterFnname(s)
	}
}

func (s *FnnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitFnname(s)
	}
}

func (s *FnnameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitFnname(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Excellent2Parser) Fnname() (localctx IFnnameContext) {
	localctx = NewFnnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Excellent2ParserRULE_fnname)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(Excellent2ParserNAME)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Excellent2ParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Excellent2ParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) CopyFrom(ctx *ParametersContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionParametersContext struct {
	*ParametersContext
}

func NewFunctionParametersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	p.ParametersContext = NewEmptyParametersContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParametersContext))

	return p
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionParametersContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Excellent2ParserCOMMA)
}

func (s *FunctionParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Excellent2ParserCOMMA, i)
}

func (s *FunctionParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.EnterFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Excellent2Listener); ok {
		listenerT.ExitFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Excellent2Visitor:
		return t.VisitFunctionParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Excellent2Parser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Excellent2ParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFunctionParametersContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.expression(0)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Excellent2ParserCOMMA {
		{
			p.SetState(78)
			p.Match(Excellent2ParserCOMMA)
		}
		{
			p.SetState(79)
			p.expression(0)
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *Excellent2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *AtomContext = nil
		if localctx != nil {
			t = localctx.(*AtomContext)
		}
		return p.Atom_Sempred(t, predIndex)

	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Excellent2Parser) Atom_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Excellent2Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
