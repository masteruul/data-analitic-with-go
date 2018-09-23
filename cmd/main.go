package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kniren/gota/dataframe"
)

func main() {
	fmt.Printf("===READ CSV FILE GOLANG===")
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	file := "github.com/masteruul/data-analitic-with-go/storage/fifa_ranking.csv"
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	data := string(dat)
	df := dataframe.ReadCSV(strings.NewReader(data))
	fmt.Print(df)
}
