package repl

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/chzyer/readline"
)

type Repl struct {
	Instructions []string
}

// Wait starts the repl and blocks
func (r *Repl) Wait() error {

	rl, err := readline.NewEx(&readline.Config{
		Prompt: "go > ",
	})
	if err != nil {
		return err
	}
	defer rl.Close()
	b := bytes.Buffer{}
outer:
	for {
		line, err := rl.Readline()
		if err != nil {
			return err
		}

		if len(line) == 0 {
			continue
		}

		if line[0] == '.' {
			switch line[1:] {
			case "exit":
				break outer
			case "clear":
				if len(r.Instructions) > 0 {
					r.Instructions = r.Instructions[:0]
					slog.Info("Clearing previous instructions")
				}
				continue
			default:
				line = "println(" + line[1:] + ")"
			}
		}

		if line[len(line)-1] != ';' {
			rl.SetPrompt(">>>> ")
			b.WriteString(line)
			b.WriteRune('\n')
			continue
		}

		if b.Len() > 0 {
			b.WriteString(line)
			line = b.String()
			b.Reset()
			rl.SetPrompt("go > ")
		}

		r.Instructions = append(r.Instructions, line[:len(line)-1])
		if err := r.Exec(); err != nil {
			slog.Error("Failed to run, clearing previous instructions", "err", err)
			r.Instructions = r.Instructions[:0]
		}
	}

	return nil
}

// Exec executes all previously input instructions
func (r *Repl) Exec() error {
	file, err := os.CreateTemp("", "repl_*.go")
	if err != nil {
		return err
	}
	if err := r.codeGen(file); err != nil {
		return err
	}
	return r.compileAndRun(file.Name())
}

func (r *Repl) codeGen(w io.Writer) error {
	buf := &bytes.Buffer{}
	imports := make([]string, 0, 16)
	lines := make([]string, 0, 16)
	for _, line := range r.Instructions {
		// handle multiline input
		for _, subLine := range strings.Split(line, "\n") {
			if strings.Contains(subLine, "import") {
				imports = append(imports, subLine)
				continue
			}
			lines = append(lines, subLine)
		}
	}
	buf.WriteString("package main;")
	for _, i := range imports {
		buf.WriteString(i)
		buf.WriteRune(';')
	}
	buf.WriteString("func main(){")
	for _, l := range lines {
		buf.WriteString(l)
		buf.WriteRune(';')
	}
	buf.WriteRune('}')
	slog.Debug("Generated go code", "code", buf.String())
	_, err := buf.WriteTo(w)
	return err
}

func (r *Repl) compileAndRun(path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	args := []string{"run", path}
	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
