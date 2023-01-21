package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/joerdav/brev/lexer"
	"github.com/joerdav/brev/tokens"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Fprintf(out, prompt)
	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.NewLexer(strings.NewReader(line))
		for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%#v\n", tok)
		}
		fmt.Fprintf(out, prompt)
	}
}
