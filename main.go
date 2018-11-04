package main

import (
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "io"
  "log"
  "net/http"
  "os"
  "time"

  "github.com/davecgh/go-spew/spew"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
)

// data model (struct) of each block in the blockchain
// each block will contain data to be written to the blockchain
// Index - position of data record in the blockchain
// Timestamp - determined when the data is written
// BPM - beats per min
// Hash - SHA256 identifier of this data record
// PrevHash - SHA256 identifier of the previous data record

type Block struct {
  Index     int
  Timestamp string
  BPM       int
  Hash      string
  PrevHash  string
}

var Blockchain []Block
