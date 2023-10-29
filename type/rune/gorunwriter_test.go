package gorunewriter_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/marrbor/goutil/closer"
	"github.com/marrbor/goutil/type/rune"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	NoBreakSpace = "\u00A0"
	WaveDash     = "\u301C"
)

func TestRuneWriter_Write(t *testing.T) {
	const fileName = "out.csv"
	_, err := os.Stat(fileName)
	if err != nil {
		// remove file
		assert.NoError(t, os.Remove(fileName))
	}

	file, err := os.Create(fileName)
	assert.NoError(t, err)
	defer closer.Close(file)

	writer := csv.NewWriter(&gorunewriter.RuneWriter{Writer: transform.NewWriter(file, japanese.ShiftJIS.NewEncoder())})
	writer.UseCRLF = true

	header := []string{
		"header1",
		"header2",
	}
	body := []string{
		fmt.Sprintf("あ%sい", NoBreakSpace),
		fmt.Sprintf("十時 %s 十二時", WaveDash),
	}

	err = writer.Write(header)
	assert.NoError(t, err)

	err = writer.Write(body)
	assert.NoError(t, err)

	writer.Flush()
	err = writer.Error()
	assert.NoError(t, err)
}
