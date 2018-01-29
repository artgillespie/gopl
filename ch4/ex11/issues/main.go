package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func editWithUsersTextEditor(s string) (string, error) {
	// write out the string we want the user to edit to a tmp file
	f, err := ioutil.TempFile("", "issues-")
	if err != nil {
		return "", err
	}
	tmpPath := f.Name()
	_, err = f.WriteString(s)
	f.Close()
	if err != nil {
		return "", err
	}

	// get the editor, fail if the env var isn't set
	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		return "", fmt.Errorf("No EDITOR environment variable set")
	}
	// get the full path to the user's text editor
	fullEditorPath, err := exec.LookPath(editor)
	if err != nil {
		return "", err
	}

	// launch the user's editor
	cmd := exec.Command(fullEditorPath, tmpPath)
	// even if we're not using these, they have to be connected in order for this to work
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	// open the tmp file and read the user's edits back in
	f, err = os.Open(tmpPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	// remove the tmp file
	err = os.Remove(tmpPath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func main() {
	s, err := editWithUsersTextEditor("Hello, Text Editor!")
	if err != nil {
		panic(fmt.Sprintf("Couldn't edit with text editor: %s", err))
	}
	fmt.Println("After Edit:")
	fmt.Println(s)
}
