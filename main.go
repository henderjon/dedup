package main

import (
	"bufio"
	"log"
	"os"
	"time"
	"fmt"
	"flag"
	"crypto/sha1"
	"io/ioutil"
	"encoding/hex"
	"strings"
)

var help bool
var grep bool

func init() {
	flag.BoolVar(&help, "help", false, "line-by-line removes control characters from a file")
	flag.BoolVar(&grep, "grep", false, "removes 'line-number:' prefixes; assumes you used `find . | grep -n string | ...`")
	flag.Parse()
}

func main() {

	if help {
		flag.PrintDefaults()
		return
	}

	now := time.Now()

	hMap := make(map[string]string, 0)
	dup  := make(map[string][]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if grep {
			ind := strings.IndexByte(line, ':')
			if ind != -1 {
				line = line[ind+1:]
			}
		}

		bts, _ := ioutil.ReadFile(line)
		sum := sha1.Sum(bts)
		sumStr := hex.EncodeToString(sum[:])
		if v, ok := hMap[sumStr]; ok {
			dup[sumStr] = append(dup[sumStr], line)
			dup[sumStr] = append(dup[sumStr], v)
		}else{
			hMap[sumStr] = line
		}
	}

	for k, v := range dup {
		fmt.Println("sum", k)
		for _, v := range v {
			fmt.Println("  - ", v)
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Println(os.Stderr, "reading standard input:", err)
	}

	log.Println("found", len(dup), "set of duplicates across", len(hMap), "files in", time.Since(now).String())
}
