package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

type LineInfo struct {
  lineNo int
  line   string
}

type FindInfo struct {
  filename string
  lines    []LineInfo
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage : $ go_grep word filepaths")
    return
  }

  word := os.Args[1]
  files := os.Args[2:]
  
  findInfos := []FindInfo{}
  for _, path := range files {
    findInfos = append(findInfos, ConFindWordInAllFiles(word, path)...)
  }

  for _, findInfo := range findInfos {
    if len(findInfo.lines) == 0 {
      continue
    }

    fmt.Println(findInfo.filename)
    fmt.Println("===========================================")
    for _, lineInfo := range findInfo.lines {
      fmt.Println("\n", lineInfo.lineNo, "\t", lineInfo.line)
    }
    fmt.Println("===========================================")
    fmt.Println()
  }
}

func GetFileList(path string) ([]string, error) {
  return filepath.Glob(path)
}

// func FindWordInAllFiles(word string, path string) []FindInfo {
//   findInfos := []FindInfo{}
//
//   filelist, err := GetFileList(path)
//   if err != nil {
//     fmt.Println("Check filepath. err:", err, "path:", path)
//     return findInfos
//   }
//
//   for _, filename := range filelist {
//     findInfos = append(findInfos, FindWordInFile(word, filename))
//   }
//
//   return findInfos
// }

// func FindWordInFile(word string, filename string) FindInfo {
//   findInfo := FindInfo{filename, []LineInfo{}}
//   file, err := os.Open(filename)
//   if err != nil {
//     fmt.Println("Unable to open file. ", filename)
//     return findInfo
//   }
//   defer file.Close()
//
//   lineNo := 1
//   scanner := bufio.NewScanner(file)
//   for scanner.Scan() {
//     line := scanner.Text()
//     if strings.Contains(line, word) {
//       findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
//     }
//     lineNo++
//   }
//
//   return findInfo
// }

func ConFindWordInAllFiles(word string, path string) []FindInfo {
  findInfos := []FindInfo{}

  filelist, err := GetFileList(path)
  if err != nil {
    fmt.Println("Check filepath. err:", err, "path:", path)
    return findInfos
  }

  ch := make(chan FindInfo)
  cnt := len(filelist)
  recvCnt := 0

  for _, filename := range filelist {
    go ConFindWordInFile(word, filename, ch)
  }

  for findInfo := range ch {
    findInfos = append(findInfos, findInfo)
    recvCnt++
    if recvCnt == cnt {
      break
    }
  }

  return findInfos
}

func ConFindWordInFile(word string, filename string, ch chan FindInfo) {
  findInfo := FindInfo{filename, []LineInfo{}}
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Unable to open file. ", filename)
    ch <- findInfo
    return
  }
  defer file.Close()

  lineNo := 1
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if strings.Contains(line, word) {
      findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
    }
    lineNo++
  }

  ch <- findInfo
}

