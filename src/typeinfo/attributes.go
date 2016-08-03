package typeinfo

// CommentOf -- get value of `Comment` type attribute.
func CommentOf(object interface{}) string {
	return searchString(object, commentType)
}
