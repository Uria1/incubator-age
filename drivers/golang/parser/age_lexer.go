// Code generated from Age.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 157,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 7, 11, 82, 10, 11, 12, 11, 14, 11, 85, 11, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 98, 10,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 108,
	10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 5, 18, 121, 10, 18, 3, 18, 3, 18, 3, 18, 6, 18, 126, 10,
	18, 13, 18, 14, 18, 127, 5, 18, 130, 10, 18, 3, 18, 5, 18, 133, 10, 18,
	3, 19, 3, 19, 3, 19, 7, 19, 138, 10, 19, 12, 19, 14, 19, 141, 11, 19, 5,
	19, 143, 10, 19, 3, 20, 3, 20, 5, 20, 147, 10, 20, 3, 20, 3, 20, 3, 21,
	6, 21, 152, 10, 21, 13, 21, 14, 21, 153, 3, 21, 3, 21, 2, 2, 22, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 2, 29, 2, 31, 2, 33, 2, 35, 15, 37, 2, 39, 2, 41, 16, 3, 2,
	10, 10, 2, 36, 36, 49, 49, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116,
	118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 2, 33, 36, 36, 94, 94, 3,
	2, 50, 59, 3, 2, 51, 59, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 162, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 45, 3, 2, 2, 2, 7, 47, 3, 2, 2, 2, 9,
	49, 3, 2, 2, 2, 11, 51, 3, 2, 2, 2, 13, 53, 3, 2, 2, 2, 15, 55, 3, 2, 2,
	2, 17, 64, 3, 2, 2, 2, 19, 71, 3, 2, 2, 2, 21, 78, 3, 2, 2, 2, 23, 97,
	3, 2, 2, 2, 25, 99, 3, 2, 2, 2, 27, 104, 3, 2, 2, 2, 29, 109, 3, 2, 2,
	2, 31, 115, 3, 2, 2, 2, 33, 117, 3, 2, 2, 2, 35, 120, 3, 2, 2, 2, 37, 142,
	3, 2, 2, 2, 39, 144, 3, 2, 2, 2, 41, 151, 3, 2, 2, 2, 43, 44, 7, 93, 2,
	2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 46, 2, 2, 46, 6, 3, 2, 2, 2, 47, 48, 7,
	95, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50, 7, 125, 2, 2, 50, 10, 3, 2, 2, 2,
	51, 52, 7, 127, 2, 2, 52, 12, 3, 2, 2, 2, 53, 54, 7, 60, 2, 2, 54, 14,
	3, 2, 2, 2, 55, 56, 7, 60, 2, 2, 56, 57, 7, 60, 2, 2, 57, 58, 7, 120, 2,
	2, 58, 59, 7, 103, 2, 2, 59, 60, 7, 116, 2, 2, 60, 61, 7, 118, 2, 2, 61,
	62, 7, 103, 2, 2, 62, 63, 7, 122, 2, 2, 63, 16, 3, 2, 2, 2, 64, 65, 7,
	60, 2, 2, 65, 66, 7, 60, 2, 2, 66, 67, 7, 103, 2, 2, 67, 68, 7, 102, 2,
	2, 68, 69, 7, 105, 2, 2, 69, 70, 7, 103, 2, 2, 70, 18, 3, 2, 2, 2, 71,
	72, 7, 60, 2, 2, 72, 73, 7, 60, 2, 2, 73, 74, 7, 114, 2, 2, 74, 75, 7,
	99, 2, 2, 75, 76, 7, 118, 2, 2, 76, 77, 7, 106, 2, 2, 77, 20, 3, 2, 2,
	2, 78, 83, 7, 36, 2, 2, 79, 82, 5, 27, 14, 2, 80, 82, 5, 33, 17, 2, 81,
	79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2,
	2, 83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87,
	7, 36, 2, 2, 87, 22, 3, 2, 2, 2, 88, 89, 7, 118, 2, 2, 89, 90, 7, 116,
	2, 2, 90, 91, 7, 119, 2, 2, 91, 98, 7, 103, 2, 2, 92, 93, 7, 104, 2, 2,
	93, 94, 7, 99, 2, 2, 94, 95, 7, 110, 2, 2, 95, 96, 7, 117, 2, 2, 96, 98,
	7, 103, 2, 2, 97, 88, 3, 2, 2, 2, 97, 92, 3, 2, 2, 2, 98, 24, 3, 2, 2,
	2, 99, 100, 7, 112, 2, 2, 100, 101, 7, 119, 2, 2, 101, 102, 7, 110, 2,
	2, 102, 103, 7, 110, 2, 2, 103, 26, 3, 2, 2, 2, 104, 107, 7, 94, 2, 2,
	105, 108, 9, 2, 2, 2, 106, 108, 5, 29, 15, 2, 107, 105, 3, 2, 2, 2, 107,
	106, 3, 2, 2, 2, 108, 28, 3, 2, 2, 2, 109, 110, 7, 119, 2, 2, 110, 111,
	5, 31, 16, 2, 111, 112, 5, 31, 16, 2, 112, 113, 5, 31, 16, 2, 113, 114,
	5, 31, 16, 2, 114, 30, 3, 2, 2, 2, 115, 116, 9, 3, 2, 2, 116, 32, 3, 2,
	2, 2, 117, 118, 10, 4, 2, 2, 118, 34, 3, 2, 2, 2, 119, 121, 7, 47, 2, 2,
	120, 119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122,
	129, 5, 37, 19, 2, 123, 125, 7, 48, 2, 2, 124, 126, 9, 5, 2, 2, 125, 124,
	3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2,
	2, 2, 128, 130, 3, 2, 2, 2, 129, 123, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2,
	130, 132, 3, 2, 2, 2, 131, 133, 5, 39, 20, 2, 132, 131, 3, 2, 2, 2, 132,
	133, 3, 2, 2, 2, 133, 36, 3, 2, 2, 2, 134, 143, 7, 50, 2, 2, 135, 139,
	9, 6, 2, 2, 136, 138, 9, 5, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141, 3, 2,
	2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 143, 3, 2, 2, 2,
	141, 139, 3, 2, 2, 2, 142, 134, 3, 2, 2, 2, 142, 135, 3, 2, 2, 2, 143,
	38, 3, 2, 2, 2, 144, 146, 9, 7, 2, 2, 145, 147, 9, 8, 2, 2, 146, 145, 3,
	2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 5, 37, 19,
	2, 149, 40, 3, 2, 2, 2, 150, 152, 9, 9, 2, 2, 151, 150, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155,
	3, 2, 2, 2, 155, 156, 8, 21, 2, 2, 156, 42, 3, 2, 2, 2, 15, 2, 81, 83,
	97, 107, 120, 127, 129, 132, 139, 142, 146, 153, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "','", "']'", "'{'", "'}'", "':'", "'::vertex'", "'::edge'",
	"'::path'", "", "", "'null'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "KW_VERTEX", "KW_EDGE", "KW_PATH", "STRING",
	"BOOL", "NULL", "NUMBER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "KW_VERTEX", "KW_EDGE",
	"KW_PATH", "STRING", "BOOL", "NULL", "ESC", "UNICODE", "HEX", "SAFECODEPOINT",
	"NUMBER", "INT", "EXP", "WS",
}

type AgeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewAgeLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *AgeLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAgeLexer(input antlr.CharStream) *AgeLexer {
	l := new(AgeLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Age.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AgeLexer tokens.
const (
	AgeLexerT__0      = 1
	AgeLexerT__1      = 2
	AgeLexerT__2      = 3
	AgeLexerT__3      = 4
	AgeLexerT__4      = 5
	AgeLexerT__5      = 6
	AgeLexerKW_VERTEX = 7
	AgeLexerKW_EDGE   = 8
	AgeLexerKW_PATH   = 9
	AgeLexerSTRING    = 10
	AgeLexerBOOL      = 11
	AgeLexerNULL      = 12
	AgeLexerNUMBER    = 13
	AgeLexerWS        = 14
)
