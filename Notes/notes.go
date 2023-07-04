package Notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)


type Note struct {
    Note string
    CreatedAt time.Time
}

//add and remove notes and display them

type List []Note

func (l *List) Load(fileName string) error {
    file, err := ioutil.ReadFile(fileName)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return nil
        }else {
            return err
        }

    }
    if len(file) == 0 {
       return  err
    }
    err = json.Unmarshal(file, l)
    if err != nil {
        return err
    }

    return nil
}

func (l *List) Store(fileName string) error {
    
    data, err := json.Marshal(l)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(fileName, data, 0644)
}

func (l *List) Add(note string) {
    
    newNote := Note{
        Note:note,
        CreatedAt: time.Now(),
    }
    *l = append(*l, newNote)
}

func (l *List) Delete(index int) error {
    n := *l
    if index < 0 || index > len(n) {
        return errors.New("Invalid index")
    }

    *l = append(n[:index-1], n[index:]...)
    return nil
}

func (l *List)  Print() {
    
    table := simpletable.New()

    table.Header = &simpletable.Header{
        Cells : []*simpletable.Cell{
            {Align: simpletable.AlignCenter, Text: "Notes"},
            {Align: simpletable.AlignCenter, Text: "CreateAt"},
        },
    }

    for _, row := range *l {
        r := []*simpletable.Cell{
            {Align: simpletable.AlignRight, Text:fmt.Sprintf("%s", row.Note)},
            {Align: simpletable.AlignRight, Text: row.CreatedAt.Format(time.RFC822)},
        }
        table.Body.Cells = append(table.Body.Cells, r)

    }
    table.SetStyle(simpletable.StyleCompactLite)
    fmt.Println(table.String())
}

