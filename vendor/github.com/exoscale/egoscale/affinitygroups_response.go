// code generated; DO NOT EDIT.

package egoscale

import "fmt"

<<<<<<< HEAD
<<<<<<< HEAD
// Response returns the struct to unmarshal.
func (ListAffinityGroups) Response() interface{} {
	return new(ListAffinityGroupsResponse)
}

// ListRequest returns itself.
func (ls *ListAffinityGroups) ListRequest() (ListCommand, error) {
	if ls == nil {
		return nil, fmt.Errorf("%T cannot be nil", ls)
	}
	return ls, nil
}

// SetPage sets the current page.
func (ls *ListAffinityGroups) SetPage(page int) {
	ls.Page = page
}

// SetPageSize sets the page size.
func (ls *ListAffinityGroups) SetPageSize(pageSize int) {
	ls.PageSize = pageSize
}

// Each triggers the callback for each, valid answer or any non 404 issue.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Response returns the struct to unmarshal
||||||| parent of 5ce8c7613 (update vendored files)
// Response returns the struct to unmarshal
=======
// Response returns the struct to unmarshal.
>>>>>>> 5ce8c7613 (update vendored files)
func (ListAffinityGroups) Response() interface{} {
	return new(ListAffinityGroupsResponse)
}

// ListRequest returns itself.
func (ls *ListAffinityGroups) ListRequest() (ListCommand, error) {
	if ls == nil {
		return nil, fmt.Errorf("%T cannot be nil", ls)
	}
	return ls, nil
}

// SetPage sets the current page.
func (ls *ListAffinityGroups) SetPage(page int) {
	ls.Page = page
}

// SetPageSize sets the page size.
func (ls *ListAffinityGroups) SetPageSize(pageSize int) {
	ls.PageSize = pageSize
}

<<<<<<< HEAD
// Each triggers the callback for each, valid answer or any non 404 issue
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// Each triggers the callback for each, valid answer or any non 404 issue
=======
// Each triggers the callback for each, valid answer or any non 404 issue.
>>>>>>> 5ce8c7613 (update vendored files)
func (ListAffinityGroups) Each(resp interface{}, callback IterateItemFunc) {
	items, ok := resp.(*ListAffinityGroupsResponse)
	if !ok {
		callback(nil, fmt.Errorf("wrong type, ListAffinityGroupsResponse was expected, got %T", resp))
		return
	}

	for i := range items.AffinityGroup {
		if !callback(&items.AffinityGroup[i], nil) {
			break
		}
	}
}
