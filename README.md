### What is this?

An example of how to define type bound information (type meta data).
2 most useful typeinfo classes are shown: predicates and attributes.

### How?

We utilize a property of `struct{}` type.
It almost always has 0 size and it always has zero size
when it is a leading member of the struct.
For example, if we have a struct with first N members with
`struct{}` type, memory overhead will be zero.
And still, they have a type associated with them which can be
examined with reflection and is usable inside type assertions or
type switches. It is also possible to attach field tag for those members.
This way, we can add some data for out types.

### What for?

This is a kind of poor mans class annotations for Go.

1) Generic algorithms which can use some type information to use
more optimized paths.
2) To tie type description to code more closely.
This gives us an ability to generate, for example, documentation.

### Example

```Go
// 1) define a struct which uses 
type point struct {
	typeinfo.Comment `typeinfo:"represents 2D point"`
	typeinfo.PodMarker
	typeinfo.PointerFreeMarker

	X float64
	Y float64
}

// sizeof(point) == sizeof(float64) + sizeof(float64)

// 2) use provided functions to extract info:
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
```

### Attributes

Attributes have string value which is initialized via field tag.
Value can be in one of the 2 formats:
1) Go conventional `typeinfo:"this is value"`.
2) Plain `this is value`. This will drive linter mad and can break things.
The second option is there for those who do not need any other
tags for those type markers/attributes & can tolerate linter warnings.
The benefits are: tag content is more compact and readable,
also access time is faster.

### Predicates

Predicates use type markers.
If a type has marker, then the predicate returns true.
All markers are `single` and have no value.