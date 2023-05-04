package utils

// MergeIntoMap merges `src` map into `dst` map.
// `dst` - destination map, `src` - source map
// if the same key exists in both maps, the value from the second map will be used.
// if the first map is nil, a new map will be created.
// if the second map is nil, the first map will be returned.
// if both maps are nil, a new map will be created.
func MergeIntoMap(dst, src map[string]interface{}) map[string]interface{} {
	if dst == nil {
		dst = make(map[string]interface{})
	}

	for k, v := range src {
		dst[k] = v
	}

	return dst
}

// MergeIntoMapRecursively recursively merges `src` map into `dst` map.
// `dst` - destination map, `src` - source map
// if the same key exists in both maps, the value from the second map will be used.
// if the first map is nil, a new map will be created.
// if the second map is nil, the first map will be returned.
// if both maps are nil, a new map will be created.
// if the value for a given key is a map, it will be merged recursively.
func MergeIntoMapRecursively(dst, src map[string]interface{}) map[string]interface{} {
	if dst == nil {
		return src
	}
	if src == nil {
		return dst
	}

	mergeMapsRecursivelyHelper(dst, src)

	return dst
}

func mergeMapsRecursivelyHelper(dst, src map[string]interface{}) {
	for k, v := range src {
		if _, ok := dst[k]; ok {
			if dstMap, ok := dst[k].(map[string]interface{}); ok {
				if srcMap, ok := src[k].(map[string]interface{}); ok {
					mergeMapsRecursivelyHelper(dstMap, srcMap)
					continue
				}
			}
		}

		dst[k] = v
	}
}
