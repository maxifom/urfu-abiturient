package main

import (
	"context"
	"sync"
	"time"

	"urfu-abiturient-api/ent"
)

type Dictionary struct {
	Basis   []string `json:"basis"`
	Form    []string `json:"form"`
	Program []string `json:"program"`
	Type    []string `json:"type"`
	Status  []string `json:"status"`
}

type DictionaryCache struct {
	dict        Dictionary
	lastUpdated time.Time
	m           sync.Mutex
}

func NewDictionaryCache(dict Dictionary) *DictionaryCache {
	return &DictionaryCache{dict: dict, lastUpdated: time.Now()}
}

func (c *DictionaryCache) Get() (Dictionary, bool) {
	c.m.Lock()
	defer c.m.Unlock()

	return c.dict, time.Since(c.lastUpdated) > 5*time.Minute
}

func (c *DictionaryCache) Update(dict Dictionary) {
	c.m.Lock()
	defer c.m.Unlock()

	c.dict = dict
}

func GetDictFromDB(client *ent.Client) Dictionary {
	basis := client.AbiturientEntry.Query().Order(ent.Asc("basis")).GroupBy("basis").StringsX(context.Background())
	form := client.AbiturientEntry.Query().Order(ent.Asc("form")).GroupBy("form").StringsX(context.Background())
	program := client.AbiturientEntry.Query().Order(ent.Asc("program")).GroupBy("program").StringsX(context.Background())
	types := client.AbiturientEntry.Query().Order(ent.Asc("type")).GroupBy("type").StringsX(context.Background())
	status := client.AbiturientEntry.Query().Order(ent.Asc("status")).GroupBy("status").StringsX(context.Background())

	dict := Dictionary{
		Basis:   basis,
		Form:    form,
		Program: program,
		Type:    types,
		Status:  status,
	}

	return dict
}
