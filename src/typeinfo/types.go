package typeinfo

// This file declares all public types for metadata.
//
// If type has more than one attribute/marker of the same type,
// only first one is used.

/**
 * Attributes
 */

// Attributes have string value which is initialized via field tag.
// Value can be in one of the 2 formats:
// 1) Go conventional `typeinfo:"this is value"`.
// 2) Plain `this is value`. This will drive linter mad and can break things.
// The second option is there for those who do not need any other
// tags for those type markers/attributes & can tolerate linter warnings.
// The benefits are: tag content is more compact and readable,
// also access time is faster.

// Comment -- holds short type description.
type Comment struct{}

/**
 * Markers (used for predicates)
 */

// All markers are `single` and have no value.

// PodMarker -- tells if described type should be treated as POD.
type PodMarker struct{}

// PointerFreeMarker -- type has no pointer members,
// so the copy of all its bytes is a deep copy.
type PointerFreeMarker struct{}
