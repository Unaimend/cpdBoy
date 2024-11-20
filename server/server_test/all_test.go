package server

import "testing"
import "github.com/Unaimend/cpdBoy/server"



func TestQuoteAndJoin(t *testing.T) {
  result := server.QuoteAndJoin("a,b,c")
  expected := "\"a\",\"b\",\"c\""
  
  if result != expected {
    t.Errorf("Is: %s; want %s", result, expected)
  }
}
