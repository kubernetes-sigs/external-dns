// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

type V4PagePaginationResult[T any] struct {
	Items []T                        `json:"items"`
	JSON  v4PagePaginationResultJSON `json:"-"`
}

// v4PagePaginationResultJSON contains the JSON metadata for the struct
// [V4PagePaginationResult[T]]
type v4PagePaginationResultJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationResult[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v4PagePaginationResultJSON) RawJSON() string {
	return r.raw
}

type V4PagePaginationResultInfo struct {
	Page    int64                          `json:"page"`
	PerPage int64                          `json:"per_page"`
	JSON    v4PagePaginationResultInfoJSON `json:"-"`
}

// v4PagePaginationResultInfoJSON contains the JSON metadata for the struct
// [V4PagePaginationResultInfo]
type v4PagePaginationResultInfoJSON struct {
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v4PagePaginationResultInfoJSON) RawJSON() string {
	return r.raw
}

type V4PagePagination[T any] struct {
	Result     V4PagePaginationResult[T]  `json:"result"`
	ResultInfo V4PagePaginationResultInfo `json:"result_info"`
	JSON       v4PagePaginationJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// v4PagePaginationJSON contains the JSON metadata for the struct
// [V4PagePagination[T]]
type v4PagePaginationJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v4PagePaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *V4PagePagination[T]) GetNextPage() (res *V4PagePagination[T], err error) {
	if len(r.Result.Items) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *V4PagePagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &V4PagePagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type V4PagePaginationAutoPager[T any] struct {
	page *V4PagePagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewV4PagePaginationAutoPager[T any](page *V4PagePagination[T], err error) *V4PagePaginationAutoPager[T] {
	return &V4PagePaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *V4PagePaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Result.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *V4PagePaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *V4PagePaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *V4PagePaginationAutoPager[T]) Index() int {
	return r.run
}

type V4PagePaginationArrayResultInfo struct {
	Page    int64                               `json:"page"`
	PerPage int64                               `json:"per_page"`
	JSON    v4PagePaginationArrayResultInfoJSON `json:"-"`
}

// v4PagePaginationArrayResultInfoJSON contains the JSON metadata for the struct
// [V4PagePaginationArrayResultInfo]
type v4PagePaginationArrayResultInfoJSON struct {
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationArrayResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v4PagePaginationArrayResultInfoJSON) RawJSON() string {
	return r.raw
}

type V4PagePaginationArray[T any] struct {
	Result     []T                             `json:"result"`
	ResultInfo V4PagePaginationArrayResultInfo `json:"result_info"`
	JSON       v4PagePaginationArrayJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// v4PagePaginationArrayJSON contains the JSON metadata for the struct
// [V4PagePaginationArray[T]]
type v4PagePaginationArrayJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationArray[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v4PagePaginationArrayJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *V4PagePaginationArray[T]) GetNextPage() (res *V4PagePaginationArray[T], err error) {
	if len(r.Result) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *V4PagePaginationArray[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &V4PagePaginationArray[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type V4PagePaginationArrayAutoPager[T any] struct {
	page *V4PagePaginationArray[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewV4PagePaginationArrayAutoPager[T any](page *V4PagePaginationArray[T], err error) *V4PagePaginationArrayAutoPager[T] {
	return &V4PagePaginationArrayAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *V4PagePaginationArrayAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result) == 0 {
			return false
		}
	}
	r.cur = r.page.Result[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *V4PagePaginationArrayAutoPager[T]) Current() T {
	return r.cur
}

func (r *V4PagePaginationArrayAutoPager[T]) Err() error {
	return r.err
}

func (r *V4PagePaginationArrayAutoPager[T]) Index() int {
	return r.run
}

type CursorPaginationResultInfo struct {
	Count   int64                          `json:"count"`
	Cursor  string                         `json:"cursor"`
	PerPage int64                          `json:"per_page"`
	JSON    cursorPaginationResultInfoJSON `json:"-"`
}

// cursorPaginationResultInfoJSON contains the JSON metadata for the struct
// [CursorPaginationResultInfo]
type cursorPaginationResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorPaginationResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorPaginationResultInfoJSON) RawJSON() string {
	return r.raw
}

type CursorPagination[T any] struct {
	Result     []T                        `json:"result"`
	ResultInfo CursorPaginationResultInfo `json:"result_info"`
	JSON       cursorPaginationJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// cursorPaginationJSON contains the JSON metadata for the struct
// [CursorPagination[T]]
type cursorPaginationJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPagination[T]) GetNextPage() (res *CursorPagination[T], err error) {
	if len(r.Result) == 0 {
		return nil, nil
	}
	next := r.ResultInfo.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPaginationAutoPager[T any] struct {
	page *CursorPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewCursorPaginationAutoPager[T any](page *CursorPagination[T], err error) *CursorPaginationAutoPager[T] {
	return &CursorPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result) == 0 {
			return false
		}
	}
	r.cur = r.page.Result[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPaginationAutoPager[T]) Index() int {
	return r.run
}

type CursorLimitPaginationResultInfo struct {
	Count   int64                               `json:"count"`
	Cursor  string                              `json:"cursor"`
	PerPage int64                               `json:"per_page"`
	JSON    cursorLimitPaginationResultInfoJSON `json:"-"`
}

// cursorLimitPaginationResultInfoJSON contains the JSON metadata for the struct
// [CursorLimitPaginationResultInfo]
type cursorLimitPaginationResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorLimitPaginationResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorLimitPaginationResultInfoJSON) RawJSON() string {
	return r.raw
}

type CursorLimitPagination[T any] struct {
	Result     []T                             `json:"result"`
	ResultInfo CursorLimitPaginationResultInfo `json:"result_info"`
	JSON       cursorLimitPaginationJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// cursorLimitPaginationJSON contains the JSON metadata for the struct
// [CursorLimitPagination[T]]
type cursorLimitPaginationJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorLimitPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorLimitPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorLimitPagination[T]) GetNextPage() (res *CursorLimitPagination[T], err error) {
	if len(r.Result) == 0 {
		return nil, nil
	}
	next := r.ResultInfo.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorLimitPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorLimitPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorLimitPaginationAutoPager[T any] struct {
	page *CursorLimitPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewCursorLimitPaginationAutoPager[T any](page *CursorLimitPagination[T], err error) *CursorLimitPaginationAutoPager[T] {
	return &CursorLimitPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorLimitPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result) == 0 {
			return false
		}
	}
	r.cur = r.page.Result[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorLimitPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorLimitPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorLimitPaginationAutoPager[T]) Index() int {
	return r.run
}

type SinglePage[T any] struct {
	Result []T            `json:"result"`
	JSON   singlePageJSON `json:"-"`
	cfg    *requestconfig.RequestConfig
	res    *http.Response
}

// singlePageJSON contains the JSON metadata for the struct [SinglePage[T]]
type singlePageJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinglePage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singlePageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SinglePage[T]) GetNextPage() (res *SinglePage[T], err error) {
	if len(r.Result) == 0 {
		return nil, nil
	}
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SinglePage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SinglePage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SinglePageAutoPager[T any] struct {
	page *SinglePage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSinglePageAutoPager[T any](page *SinglePage[T], err error) *SinglePageAutoPager[T] {
	return &SinglePageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SinglePageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result) == 0 {
			return false
		}
	}
	r.cur = r.page.Result[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SinglePageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SinglePageAutoPager[T]) Err() error {
	return r.err
}

func (r *SinglePageAutoPager[T]) Index() int {
	return r.run
}
