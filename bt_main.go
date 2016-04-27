package main

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "bytes"
    "io"
    "strconv"
)

func main() {
    url := "http://dl.frostwire.com/torrents/audio/music/Dave_Doobie_Aaron_Doobie_Duke_Sims_FrostClick_FrostWire_MP3_March_09_2016.torrent"

    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer response.Body.Close()
    
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(body))
}

func parseMetainfo(s string) map[string]string {
    m := make(map[string]string)
    
    i := 1

    for i < len(s){
        k,v,in,e := eatPair( i, s )
        if e != nil {
            panic(e.Error())
        }
       
        m[k] = v
        i = in
    }
   
    return m
}

func eatPair(i int, s string) (string,string,int,error) {
    s0,i,e := eatString(i,s)
    if e != nil {
        return "","",0,e
    }

    s1,i,e := eatString(i,s)
    if e != nil {
        return "","",0,e
    }

    return s0, s1, i, nil
}

func eatString(i int, s string) (string,int,error) {
    v, i, e := eatInt(i, s)
    if e != nil {
        return "",-1,e
    }
    
    i, e = eatColon(i, s)
    if e != nil {
        return "",-1, e
    }

    so := s[i:i+v]

    return so, i+v, nil
}

func eatColon(i int, s string) (int, error) {
    if string(s[i]) != ":" {
        return 0,fmt.Errorf("expected %q at %d, but got %q", ":", i, string(s[i]))
    }
    return i+1, nil
}

func eatInt(i int, s string) (int, int, error) {
    b := new(bytes.Buffer)
    
    if _, e := strconv.Atoi(string(s[i])); e != nil {
        return 0,0,fmt.Errorf("expected numeric rune at %d, but got %q", i, string(s[i]))
    }
    
    for {
        if _, e := strconv.Atoi(string(s[i])); e == nil {
            io.WriteString(b, string(s[i]))
            i++
        } else {
            break
        }
    }
    
    pi,e := strconv.Atoi(b.String())
    if e != nil {
        return 0,0,e
    }

    return pi, i, e
}
