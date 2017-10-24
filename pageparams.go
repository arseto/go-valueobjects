package valueobjects

import (
	"net/http"
	"strconv"
)

type PagingContract interface {
	GetPagingParams() PageParams
}

type PageParams struct {
	size int
	page int
}

func (p *PageParams) GetSize() int {
	return p.size
}

func (p *PageParams) GetPage() int {
	return p.page
}

func (p *PageParams) GetOffset() int {
	return (p.page - 1) * p.size
}

func (p *PageParams) GetSQLStatement() string {
	return " LIMIT " + strconv.Itoa(p.GetSize()) +
		" OFFSET " + strconv.Itoa(p.GetOffset())
}

func (p *PageParams) validate() (err error) {
	if p.page < 1 {
		err = NewValidationError(
			"Page",
			CannotLessThan,
			"000422",
			"1",
		)
		return
	}
	if p.size < 1 {
		err = NewValidationError(
			"Size",
			CannotLessThan,
			"000422",
			"1",
		)
		return
	}
	return
}

func NewPageParamFromRequestQuery(r *http.Request) (pageParam *PageParams,
	err error) {

	sizeStr := r.FormValue("size")
	if sizeStr == "" {
		sizeStr = "10"
	}

	pageStr := r.FormValue("page")
	if pageStr == "" {
		pageStr = "1"
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return
	}

	pageParam = &PageParams{
		size,
		page,
	}

	err = pageParam.validate()
	if err != nil {
		return
	}

	return
}
