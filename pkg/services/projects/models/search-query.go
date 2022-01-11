package models

import (
	"strconv"

	"lastimplementation.com/internal/validate"
)

const (
	defaultPage  = 1
	defaultLimit = 20
)

type SearchQP struct {
	Query string `validate:"max=100"`
	Page  int    `validate:"min=1,max=100"`
	Limit int    `validate:"min=20"`
}

func NewSearchQP(query, page, limit string) (SearchQP, error) {
	var res SearchQP
	if page != "" {
		pageNum, err := strconv.Atoi(page)
		if err != nil {
			return res, err
		}
		res.Page = pageNum
	} else {
		res.Page = defaultPage
	}
	if limit != "" {
		limitNum, err := strconv.Atoi(limit)
		if err != nil {
			return res, err
		}
		res.Limit = limitNum
	} else {
		res.Limit = defaultLimit
	}
	res.Query = query
	if err := validate.Get().Struct(res); err != nil {
		return res, err
	}
	return res, nil
}
