package main

import (
  "os"
  "bufio"
  "fmt"
  "strings"
  "sync"
  "path/filepath"
)


type lineInfo struct {
  lineNo int
  lineText string
}


type searchResult struct {
  filepath string
  lineInfos []lineInfo
}

var wg sync.WaitGroup
var searchResults = []searchResult{}

func main() {
  if len(os.Args) != 3 {
    fmt.Println("Usage : go_grep word filepath")
    os.Exit(1)
  }

  filepaths, err := filepath.Glob(os.Args[2])
  if err != nil {
    fmt.Println("Error while listing filepaths")
    os.Exit(1)
  }

  // initialize goroutines according to found files
  wg.Add(len(filepaths))

  for _, v := range filepaths {
    go gatherResults(v)
  }

  wg.Wait()

  totalFinds := 0
  totalFiles := 0
  for _, v := range searchResults {
    if len(v.lineInfos) == 0 {
      continue
    }

    fmt.Println("=================================================")
    fmt.Println(v.filepath)
    fmt.Println("-------------------------------------------------")
    for _, w := range v.lineInfos {
      fmt.Printf("Line %d : %s\n", w.lineNo, w.lineText)
    }
    fmt.Println()

    totalFinds += len(v.lineInfos)
    totalFiles += 1
  }

  fmt.Println("-------------------------------------------------")
  fmt.Printf("Found total %d matching lines from %d files\n", totalFinds, totalFiles)
  fmt.Println("-------------------------------------------------")
  
}

func gatherResults(filepath string) {
  newSearchResult := searchResult{filepath:filepath, lineInfos:readFile(filepath)}
  searchResults = append(searchResults, newSearchResult)
  wg.Done()
}

func readFile(filepath string) []lineInfo {
  fileReader, _ := os.Open(filepath)
  defer fileReader.Close()

  fileScanner := bufio.NewScanner(fileReader)

  results := []lineInfo{}

  idx := 1
  for {
    if fileScanner.Scan() == true {
      curLine := fileScanner.Text()

      if strings.Contains(curLine, os.Args[1]) {
        newLineInfo := lineInfo{lineNo:idx, lineText:curLine}
        results = append(results, newLineInfo)
      }
    } else {
      break
    }

    idx++
  }

  return results
}

