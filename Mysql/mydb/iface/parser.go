package iface

type Parser interface {
	Parse(sql string) 
}