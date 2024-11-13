package main

type InfoStruct struct {
	ConfigFileType string
}

type KvStruct struct {
	Field         string
	FieldType     string
	FieldOriginal string
}

type Sortable []KvStruct

func (a Sortable) Len() int           { return len(a) }
func (a Sortable) Less(i, j int) bool { return a[i].Field < a[j].Field }
func (a Sortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
