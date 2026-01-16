package tool

// 去重切片
func Unique[T comparable](items []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(items))

	for _, v := range items {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// RemoveElements 泛型版本，可用于任何可比较类型
func RemoveElements[T comparable](source []T, toRemove []T) []T {
	if len(toRemove) == 0 {
		return source
	}

	removeMap := make(map[T]bool, len(toRemove))
	for _, item := range toRemove {
		removeMap[item] = true
	}

	result := make([]T, 0, len(source))
	for _, item := range source {
		if !removeMap[item] {
			result = append(result, item)
		}
	}

	return result
}
