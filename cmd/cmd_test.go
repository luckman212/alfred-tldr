package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/konoui/go-alfred"
	"github.com/mattn/go-shellwords"
)

func testWorkflowOutput(t *testing.T, outWantData, outGotData, errWantData, errGotData []byte) {
	t.Helper()
	if diff := alfred.DiffScriptFilter(outWantData, outGotData); diff != "" {
		t.Errorf("+want -got\n%+v", diff)
	}

	if string(errWantData) != string(errGotData) {
		t.Errorf("Workflow want: %+v\n, got: %+v\n", string(errWantData), string(errGotData))
	}
}

func testCLIOutput(t *testing.T, outWantData, outGotData, errWantData, errGotData []byte) {
	t.Helper()

	want := string(outWantData)
	got := string(outGotData)
	if want != got {
		t.Errorf("CLI want: %+v\n, got: %+v\n", want, got)
	}

	want = string(errWantData)
	got = string(errGotData)
	if want != got {
		t.Errorf("CLI want: %+v\n, got: %+v\n", want, got)
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		description string
		expectErr   bool
		filepath    string
		command     string
		tldrMaxAge  time.Duration
		errMsg      string
	}{
		{
			description: "text output. lsof",
			expectErr:   false,
			command:     "lsof --update",
			filepath:    "./test_output_lsof.txt",
			tldrMaxAge:  tldrMaxAge,
		},
		{
			description: "text output. sub command git checkout",
			expectErr:   false,
			command:     "git checkout",
			filepath:    "./test_output_git-checkout.txt",
			tldrMaxAge:  tldrMaxAge,
		},
		{
			description: "text output. expired cache",
			expectErr:   false,
			command:     "lsof",
			filepath:    "./test_output_lsof.txt",
			errMsg:      "more than a week passed, should update tldr using --update\n",
			tldrMaxAge:  0 * time.Hour,
		},
		{
			description: "text output. page not found",
			expectErr:   false,
			command:     "lsoff",
			filepath:    "./test_output_no_page.txt",
			errMsg:      "This page doesn't exist yet!\nSubmit new pages here: https://github.com/tldr-pages/tldr\n",
			tldrMaxAge:  tldrMaxAge,
		},
		{
			description: "alfred workflow. lsof",
			expectErr:   false,
			command:     "lsof --update --workflow",
			filepath:    "./test_output_lsof.json",
			tldrMaxAge:  tldrMaxAge,
		},
		{
			description: "alfred workflow. sub command git checkout",
			expectErr:   false,
			command:     "git checkout --workflow",
			filepath:    "./test_output_git-checkout.json",
			tldrMaxAge:  tldrMaxAge,
		},
		{
			description: "alfred workflow. fuzzy search",
			expectErr:   false,
			command:     " gitchec --workflow --fuzzy",
			filepath:    "./test_output_git-checkout_with_fuzzy.json",
			tldrMaxAge:  tldrMaxAge,
		},
		{
			description: "alfred workflow. show no error when cache expired",
			expectErr:   false,
			command:     "lsof --workflow",
			filepath:    "./test_output_lsof.json",
			tldrMaxAge:  0 * time.Hour,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			// set cache max age
			tldrMaxAge = tt.tldrMaxAge

			wantData, err := ioutil.ReadFile(tt.filepath)
			if err != nil {
				t.Fatal(err)
			}

			outBuf, errBuf := new(bytes.Buffer), new(bytes.Buffer)
			outStream, errStream = outBuf, errBuf
			cmdArgs, err := shellwords.Parse(tt.command)
			if err != nil {
				t.Fatalf("args parse error: %+v", err)
			}
			rootCmd := NewRootCmd()
			rootCmd.SetOutput(outStream)
			rootCmd.SetArgs(cmdArgs)

			err = rootCmd.Execute()
			if tt.expectErr && err == nil {
				t.Errorf("expect error happens, but got response")
			}

			if !tt.expectErr && err != nil {
				t.Errorf("unexpected error got: %+v", err)
			}

			outGotData := outBuf.Bytes()
			errGotData := errBuf.Bytes()
			// switch test
			if strings.Contains(tt.command, "--workflow") || strings.Contains(tt.command, "-w") {
				testWorkflowOutput(t, wantData, outGotData, []byte(tt.errMsg), errGotData)
			} else {
				testCLIOutput(t, wantData, outGotData, []byte(tt.errMsg), errGotData)
			}
		})
	}
}
