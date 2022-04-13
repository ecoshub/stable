package main

import (
	"fmt"
	"stable/field"
	"stable/style"
	"stable/table"
)

func printStructArray() {

	persons := []*Person{
		{Name: "Ruby Cohen", Age: 30, Height: 1.80, Male: true},
		{Name: "Bethany Parsons", Age: 29, Height: 1.58},
		{Name: "Ronnie Rodriguez", Age: 28, Height: 1.78, Male: true},
		{Name: "Rosa Daniels", Age: 31, Height: 1.80, Male: true},
	}

	t, err := table.ToTable(persons)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t)
}

func printStruct() {
	p := &Person{
		Name:   "Ruby Cohen",
		Age:    31,
		Height: 1.8,
		Male:   true,
	}
	t, err := table.ToTable(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
}

func test() {
	table := table.New("Benchmark of Hashing")

	table.AddFields(
		"File",
		"File Size (KB)",
		"Chunk Size (KB)",
		"Time (ms)",
	)

	table.AddFieldWithOptions("Proportion (%)", &field.Options{
		Format:     "%0.2f",
		Alignement: style.AlignementCenter,
	})

	table.Row("/var/log/system/test/crontab.log", 12.515, 14.0, "32", 0.223)
	table.Row("/var/log/system/test/monit.log", 85.521, 43.32, "322", 0.742)
	table.Row("/var/log/system/test/logrotate.log", 96.57, 65.123, "31112", 0.321)
	table.Row("/var/log/system/test/docker-deamon.log", 13.3511, 34.01, "3652", 0.895)

	fmt.Println(table)

}

var (
	values = []interface{}{
		map[string]interface{}{
			"name":   "Ruby Cohen",
			"age":    31,
			"height": 1.8,
			"male":   true,
		},

		[]map[string]interface{}{
			{
				"name":   "Bethany Parsons",
				"age":    31,
				"height": 1.6,
				"male":   true,
			},
			{
				"name":   "Ronnie Rodriguez",
				"age":    30,
				"height": 1.7,
				"male":   true,
			},
			{
				"name":   "Rosa Daniels",
				"age":    33,
				"height": 1.7,
				"male":   false,
			},
		},

		[]byte(`{
		"encoding": "en-US,en;q=0.8",
		"host": "headers.jsontest.com",
		"accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3",
		"accept": "text/html",
		"time": "03:53:25 AM",
		"epoch": 1362196405309,
		"date": "03-02-2013"
	}`),

		[]byte(`[
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"},
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"},
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"},
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"},
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"},
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"},
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"},
	{"accept": "en-US,en;q=0.8","accept-charset": "ISO-8859-1,utf-8;q=0.7,*;q=0.3","accept": "text/html","time": "03:53:25 AM","epoch": 1362196405309,"date": "03-02-2013"}
	]`),

		[]struct {
			File       string
			FileSize   float64
			ChunkSize  float64
			Time       string
			Proportion float64
		}{
			{
				File:       "/var/log/system/test/crontab.log",
				FileSize:   12.515,
				ChunkSize:  14.0,
				Time:       "32",
				Proportion: 0.223,
			},
			{
				File:       "/var/log/system/test/monit.log",
				FileSize:   85.521,
				ChunkSize:  43.32,
				Time:       "322",
				Proportion: 0.742,
			},
			{
				File:       "/var/log/system/test/logrotate.log",
				FileSize:   96.57,
				ChunkSize:  65.123,
				Time:       "31112",
				Proportion: 0.321,
			},
			{
				File:       "/var/log/system/test/docker-deamon.log",
				FileSize:   13.3511,
				ChunkSize:  34.01,
				Time:       "3652",
				Proportion: 0.895,
			},
		},
	}
)

func anonymus() {
	for i, v := range values {
		t, err := table.ToTable(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		t.SetStyle(style.BorderStylePrintableLine)
		t.SetCaption(fmt.Sprintf("%d. table", i+1))
		fmt.Println(t)
	}
}
