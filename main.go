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

// this calculateHash function concatenates Index, Timestamp, BPM,
// prevHash of the block provided as argument and returns the sha256
// hash as string

func calculateHash(block Block) string {
  record := string(block.Index) + block.Timestamp + string(block.BPM) + block.Prevhash
  h := sha256.New()
  h.write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}

// this generates a new block if supplied the previous block

func generateBlock(oldBlock Block, BPM int) (Block, error) {

  var newBlock Blockchain
  t := time.Now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.BPM = BPM
  newBlock.PrevHash = oldBlock.Hash
  newBlock.Hash = calculateHash(newBlock)

  return newBlock, nil
}
