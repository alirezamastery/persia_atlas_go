package pagination

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/url"
	"strconv"
)

var (
	pageQueryParam     = "page"
	pageSizeQueryParam = "page_size"
	defaultPageSize    = 20
)

type PageNumberPagination struct {
	ctx      *gin.Context
	db       *gorm.DB
	scopes   []func(*gorm.DB) *gorm.DB
	lst      any
	count    int64
	pageSize int
	pageNum  int
	numPages int
	response *PaginatedResponse
}

type PaginatedResponse struct {
	Count      int64   `json:"count"`
	TotalPages int     `json:"total_pages"`
	Next       *string `json:"next"`
	Previous   *string `json:"previous"`
	Items      any     `json:"items"`
}

func NewPageNumberPaginator(
	ctx *gin.Context,
	db *gorm.DB,
	scopes []func(*gorm.DB) *gorm.DB,
	lst any,
) *PageNumberPagination {
	return &PageNumberPagination{
		ctx:    ctx,
		db:     db,
		scopes: scopes,
		lst:    lst,
	}
}

func (p *PageNumberPagination) GetPaginatedResponse() *PaginatedResponse {
	p.prepare()
	p.paginate()
	return &PaginatedResponse{
		Count:      p.count,
		TotalPages: p.numPages,
		Next:       p.getNextLink(),
		Previous:   p.getPreviousLink(),
		Items:      p.lst,
	}
}

func (p *PageNumberPagination) prepare() {
	p.pageSize = p.getPageSize()
	p.pageNum = p.getPageNumber()
}

func (p *PageNumberPagination) paginate() {
	p.scopes = append(p.scopes, p.getPaginationScope())

	p.performQuery()
	p.getCount()

	p.numPages = p.getNumPages()
}

func (p *PageNumberPagination) getPaginationScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.pageNum - 1) * p.pageSize
		return db.Offset(offset).Limit(p.pageSize)
	}
}

func (p *PageNumberPagination) performQuery() {
	p.db.Scopes(p.scopes...).Find(p.lst)
}

func (p *PageNumberPagination) getCount() {
	p.db.Scopes(p.scopes...).Find(p.lst).Count(&p.count)
}

func (p *PageNumberPagination) getPageNumber() int {
	page, _ := strconv.Atoi(p.ctx.Query(pageQueryParam))
	if page == 0 {
		page = 1
	}
	fmt.Println("page num:", page)
	return page
}

func (p *PageNumberPagination) getPageSize() int {
	pageSize, _ := strconv.Atoi(p.ctx.Query(pageSizeQueryParam))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = defaultPageSize
	}
	fmt.Println("getPageSize:", pageSize)
	return pageSize
}

func (p *PageNumberPagination) getNumPages() int {
	if p.count == 0 {
		return 0
	}
	return int(math.Ceil(float64(p.count) / float64(p.pageSize)))
}

func (p *PageNumberPagination) getNextLink() *string {
	if p.pageNum >= p.numPages {
		return nil
	}

	absURL := p.buildAbsoluteURL()
	nextUrl, err := url.Parse(absURL)
	if err != nil {
		fmt.Println("error in parsing url:", err.Error())
		return nil
	}
	nextPageNum := p.pageNum + 1
	v := nextUrl.Query()
	if v.Has(pageQueryParam) {
		v.Set(pageQueryParam, strconv.Itoa(nextPageNum))
	} else {
		v.Add(pageQueryParam, strconv.Itoa(nextPageNum))
	}
	nextUrl.RawQuery = v.Encode()
	nextLink := nextUrl.String()

	return &nextLink
}

func (p *PageNumberPagination) getPreviousLink() *string {
	if p.pageNum <= 1 {
		return nil
	}

	absURL := p.buildAbsoluteURL()
	nextUrl, err := url.Parse(absURL)
	if err != nil {
		fmt.Println("error in parsing url:", err.Error())
		return nil
	}

	prevPageNum := p.pageNum - 1
	v := nextUrl.Query()
	if prevPageNum == 1 {
		v.Del(pageQueryParam)
	} else {
		if v.Has(pageQueryParam) {
			v.Set(pageQueryParam, strconv.Itoa(prevPageNum))
		} else {
			v.Add(pageQueryParam, strconv.Itoa(prevPageNum))
		}
	}

	nextUrl.RawQuery = v.Encode()
	nextLink := nextUrl.String()

	return &nextLink
}

func (p *PageNumberPagination) buildAbsoluteURL() string {
	proto := p.ctx.Request.Header.Get("X-Forwarded-Proto")
	scheme := "http"
	if p.ctx.Request.TLS != nil || proto == "https" {
		scheme = "https"
	}

	host := p.ctx.Request.Host
	fullPath := p.ctx.FullPath()
	rawQuery := p.ctx.Request.URL.RawQuery

	return scheme + "://" + host + fullPath + "?" + rawQuery
}
