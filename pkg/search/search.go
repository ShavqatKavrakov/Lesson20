package search

import (
	"context"
	"math/rand"
	"strings"
	"sync"
	"time"
)

//Result описывает один результат поиска
type Result struct {
	//Фразу каторую искали
	Phrase string
	//Целиком вся строка, в которой нашли вхождение (без \n или \r\n в конце)
	Line string
	//Номер строки (начиная с 1), на каторой нашли вхождение
	LineNum int64
	//Номер позиции (начиная с 1), на каторой нашли вхождение
	ColNum int
}

//FoundPhrase ищет вхождение phrase в текстовых файлах file и возврашаеть результат поиска
func FoundPhrase(file string, phrase string) []Result {
	var result []Result
	colNum := 0
	lineNum := 0
	line := ""
	if file[len(file)-1] != 10 {
		s3 := string(append([]byte(file), "\n"...))
		file = s3
	}
	for _, elem := range file {
		if elem == 10 {
			lineNum++
			if colnum, ok := Found(line, phrase); ok {
				colNum = colnum
				Result := Result{
					Phrase:  phrase,
					Line:    line,
					LineNum: int64(lineNum),
					ColNum:  colNum,
				}
				result = append(result, Result)
			}
			line = ""
		}
		line += string(elem)
	}
	return result
}

//Found ищет phrase в линия и возврашает позиция phrase и проверить есть ли в линия phrase
func Found(Line string, phrase string) (int, bool) {
	colNum := 0
	prases := strings.Fields(Line)
	for _, prace := range prases {
		colNum++
		if prace == phrase {
			return colNum, true
		}
	}
	return 0, false
}

//All ищет все вхождения phrase в текстовых файлах files.
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	for _, file := range files {
		result := FoundPhrase(file, phrase)
		go func(ctx context.Context, result []Result, ch chan<- []Result) {
			defer wg.Done()
			want := rand.Intn(10)
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(want)):
				ch <- result
			}
		}(ctx, result, ch)
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	return ch

}
