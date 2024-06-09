package iface

const (
	TK_SELECT  = "SELECT"
	TK_UPDATE  = "UPDATE"
	TK_DELETE  = "DELETE"
	TK_INSERT  = "INSERT"
	TK_WHERE   = "WHERE"
	TK_AND     = "AND"
	TK_OR      = "OR"
	TK_BETWEEN = "BETWEEN"
	TK_EQ      = "="
	TK_GT      = ">"
	TK_LT      = "<"
	TK_NE      = "<>"
)

type Token struct {
	value string
}

func NewToken(value string) Token {
	return Token{
		value: value,
	}
}

func (t *Token) Value() string {
	return t.value
}
