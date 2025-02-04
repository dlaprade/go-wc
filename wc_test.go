package main 

import (
  "testing"
  "os"
)

func TestProcessLength(t *testing.T) {
  file, err := os.Open("test.txt")
  if err != nil {
    
  }
  defer file.Close()

  data := Data{}

  t.Cleanup( resetFlags )
  lengthFlag = true

  Process(file, &data)

  if data.line != 79 {
    t.Errorf("Expected line length: %d but received: %d", 79, data.line)
  }
}

func TestProcessBytes(t *testing.T) {
  file, err := os.Open("test.txt")
  if err != nil {
    
  }
  defer file.Close()

  data := Data{}

  bytesFlag = true

  t.Cleanup( resetFlags )
  Process(file, &data)

  if data.bytes != 342190 {
    t.Errorf("Expected bytes: %d but received: %d", 342190, data.bytes)
  }
}

func TestProcessLines(t *testing.T) {
  file, err := os.Open("test.txt")
  if err != nil {
    
  }
  defer file.Close()

  data := Data{}

  t.Cleanup( resetFlags )
  linesFlag = true

  Process(file, &data)

  if data.lines != 7145 {
    t.Errorf("Expected lines: %d but received: %d", 7145, data.lines)
  }
}

func TestProcessCharacters(t *testing.T) {
  file, err := os.Open("test.txt")
  if err != nil {
    
  }
  defer file.Close()

  data := Data{}

  t.Cleanup( resetFlags )
  multibyteFlag = true

  Process(file, &data)

  if data.characters != 339292 {
    t.Errorf("Expected characters: %d but received: %d", 339292, data.characters) 
  }
}

func TestProcessWords(t *testing.T) {
  file, err := os.Open("test.txt")
  if err != nil {
    
  }
  defer file.Close()

  data := Data{}

  t.Cleanup( resetFlags )
  wordFlag = true

  Process(file, &data)

  if data.words != 58164 {
    t.Errorf("Expected words: %d but received: %d", 58164, data.words)
  }
}

func TestProcessFailure(t *testing.T) {
  file, err := os.Open("test.txt")
  if err != nil {
    
  }
  defer file.Close()

  data := Data{}

  t.Cleanup( resetFlags )

  Process(file, &data)

  if data.lines != 0 {
    t.Errorf("Expected lines: %d but received: %d", 0, data.lines)
  }
  if data.bytes != 0 {
    t.Errorf("Expected bytes: %d but received: %d", 0, data.bytes)
  }
  if data.characters != 0 {
    t.Errorf("Expected characters: %d but received: %d", 0, data.characters)
  }
  if data.line != 0 {
    t.Errorf("Expected line: %d but received: %d", 0, data.line)
  }
  if data.words != 0 {
    t.Errorf("Expected words: %d but received: %d", 0, data.words)
  }
}

func resetFlags() {
  lengthFlag    = false
  bytesFlag     = false
  multibyteFlag = false
  linesFlag     = false
  wordFlag      = false
}
