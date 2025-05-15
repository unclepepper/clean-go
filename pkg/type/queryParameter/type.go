package queryParameter

import (
	"github.com/unclepepper/clean-go/pkg/type/pagination"
	"github.com/unclepepper/clean-go/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}
