package main

import (
	"fmt"
	"reflect"
	"strings"
	"typeinfo"
	"unsafe"
)

type point struct {
	typeinfo.Comment `typeinfo:"represents 2D point"`
	typeinfo.PodMarker
	typeinfo.PointerFreeMarker

	X float64
	Y float64
}

func describe(object interface{}) {
	fmt.Printf("{ %s type }\n", reflect.TypeOf(object).Name())

	if comment := typeinfo.CommentOf(object); comment != "" {
		println(comment)
	} else {
		println("-- comment is empty --")
	}

	markers := []string{}
	mapping := map[string]func(interface{}) bool{
		"POD":          typeinfo.IsPod,
		"pointer-free": typeinfo.IsPointerFree,
	}
	for marker, predicate := range mapping {
		if predicate(object) {
			markers = append(markers, marker)
		}
	}

	if len(markers) > 0 {
		fmt.Printf("markers: %s\n", strings.Join(markers, ", "))
	}
}

func main() {
	describe(point{})

	fmt.Printf("sizeof(point) = %d\n", unsafe.Sizeof(point{}))
}
