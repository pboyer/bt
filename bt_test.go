package main

import (
    "testing"
//    "fmt"
)

func TestEatString(t *testing.T){
    s, i, e := eatString(0,"8:12345678") 

    if e != nil {
        t.Fatalf(e.Error())
    }
   
    if s != "12345678" {
        t.Fatalf("expected %q, got %q", "12345678", s)
    }
    
    if i != 10 {
        t.Fatalf("expected 10, got %q", s)
    }
}

func TestEatPair(t *testing.T){
    s0, s1, i, e := eatPair(0,"3:foo6:foobar") 

    if e != nil {
        t.Fatalf(e.Error())
    }
   
    if s0 != "foo" {
        t.Fatalf("expected %q, got %q", "foo", s0)
    }
    
    if s1 != "foobar" {
        t.Fatalf("expected %q, got %q", "foobar", s1)
    }
    
    if i != 13 {
        t.Fatalf("expected 10, got %q", s1)
    }
}

func TestEatColon(t *testing.T){

}

func TestEatInt(t *testing.T){

}
