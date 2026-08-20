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

// UniqueFunc 去重切片，使用自定义相等函数判断元素是否重复。
// 与 Unique 不同，它支持对不可直接比较的类型（如结构体按某个字段）去重。
func UniqueFunc[T any](items []T, equalFunc func(T, T) bool) []T {
	result := make([]T, 0, len(items))

	for _, v := range items {
		dup := false
		for _, u := range result {
			if equalFunc(v, u) {
				dup = true
				break
			}
		}

		if !dup {
			result = append(result, v)
		}
	}

	return result
}

// UniqueBy 去重切片，使用 keyFunc 提取可比较的键来判断元素是否重复。
// 相比 UniqueFunc 的 O(n²) 复杂度，本函数借助 map 实现 O(n) 去重。
func UniqueBy[T any, K comparable](items []T, keyFunc func(T) K) []T {
	seen := make(map[K]struct{}, len(items))
	result := make([]T, 0, len(items))

	for _, v := range items {
		key := keyFunc(v)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
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

	removeMap := make(map[T]struct{}, len(toRemove))
	for _, item := range toRemove {
		removeMap[item] = struct{}{}
	}

	result := make([]T, 0, len(source))
	for _, item := range source {
		if _, ok := removeMap[item]; !ok {
			result = append(result, item)
		}
	}

	return result
}
