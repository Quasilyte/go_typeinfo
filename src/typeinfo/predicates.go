package typeinfo

// IsPod -- checks if `object` type has `PodMarker`.
// Uses: `typeinfo.PodMarker`.
func IsPod(object interface{}) bool {
	return searchMarker(object, podMarkerType)
}

// IsPointerFree -- checks if `object` type has `PointerFreeMarker`.
func IsPointerFree(object interface{}) bool {
	return searchMarker(object, pointerFreeMarkerType)
}
