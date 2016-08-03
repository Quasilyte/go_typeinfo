package typeinfo

import (
	"reflect"
)

var (
	// Attributes:
	commentType = reflect.TypeOf(Comment{})

	// Markers:
	podMarkerType         = reflect.TypeOf(PodMarker{})
	pointerFreeMarkerType = reflect.TypeOf(PointerFreeMarker{})
)

func searchString(object interface{}, needle reflect.Type) string {
	rt := reflect.TypeOf(object)

	for i := 0; i < rt.NumField(); i++ {
		if rt.Field(i).Type == needle {
			if result := rt.Field(i).Tag.Get("typeinfo"); result == "" {
				return string(rt.Field(i).Tag)
			} else {
				return result
			}
		}
	}

	return ""
}

func searchMarker(object interface{}, needle reflect.Type) bool {
	rt := reflect.TypeOf(object)

	for i := 0; i < rt.NumField(); i++ {
		if rt.Field(i).Type == needle {
			return true
		}
	}

	return false
}
