package slice_utils

import "github.com/codeus-node/fail"

func First[T any](s []T, filter func(item T) bool) (T, fail.CustomError) {
	var res T
	for _, item := range s {
		if filter(item) {
			return item, nil
		}
	}
	return res, fail.Wrap(nil, "no item found").AsWarning()
}
