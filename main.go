package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"notes/Notes"
	"os"
	"strings"
)

const (
    fileName = ".notes.json"
)

func main(){
    
    add := flag.Bool("add", false, "add a new note")
    remove := flag.Int("r", 0, "remove an existing note")
    show := flag.Bool("show", false , "show list of notes")

    flag.Parse()

    notes := &Notes.List{}

     err := notes.Load(fileName)
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(0)
    } 

    switch {
        case *add :
            note, err := getInput(os.Stdin, flag.Args()...)
            if err != nil {
                fmt.Fprintln(os.Stderr, err.Error())
                os.Exit(1)
            }
            notes.Add(note)
            err = notes.Store(fileName)
            if err != nil {
                fmt.Fprintln(os.Stderr, err.Error())
                os.Exit(1)
            } 

        case *remove > 0:
            err := notes.Delete(*remove)
            if err != nil {
                fmt.Fprintln(os.Stderr, err.Error())
                os.Exit(1)
            } 
        
            err = notes.Store(fileName)
            if err != nil {
                fmt.Fprintln(os.Stderr, err.Error())
                os.Exit(1)
            }
        case *show: 
            notes.Print()
        default:
            fmt.Fprintln(os.Stdout, "invalid command")
            os.Exit(0)

    }
}

func getInput(r io.Reader, s ...string) (string, error){
   
    if len(s) > 0{
        return strings.Join(s, " "), nil
    }
    scanner := bufio.NewScanner(r)
    scanner.Scan()
    if err := scanner.Err(); err != nil {
    
        return "",err
    }
    text := scanner.Text()
    if len(text) == 0 {
        return "", errors.New("empty note is not allowed, please insert a note jackass")

    }
    return text, nil
}
