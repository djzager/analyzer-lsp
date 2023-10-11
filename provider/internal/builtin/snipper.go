package builtin

import (
	"fmt"
	"os"
	"strings"

	"github.com/konveyor/analyzer-lsp/engine"
	"go.lsp.dev/uri"
)

var _ engine.CodeSnip = &builtinProvider{}

func (p *builtinProvider) GetCodeSnip(u uri.URI, loc engine.Location) (string, error) {
	if !strings.Contains(string(u), uri.FileScheme) {
		return "", fmt.Errorf("invalid file uri")
	}

	readFile, err := os.Open(u.Filename())
	if err != nil {
		p.log.V(5).Error(err, "Unable to read file")
		return "", err
	}
	defer readFile.Close()

	return "", nil
}
