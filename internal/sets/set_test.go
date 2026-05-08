/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sets_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/internal/sets"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		want  sets.Set[int]
	}{
		{
			name:  "empty",
			items: []int{},
			want:  sets.New[int](),
		},
		{
			name:  "single item",
			items: []int{1},
			want:  sets.New(1),
		},
		{
			name:  "multiple items with duplicates",
			items: []int{1, 2, 3, 2, 1},
			want:  sets.New(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sets.New(tt.items...)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestNewFromMapKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[int]int
		want sets.Set[int]
	}{
		{
			name: "empty map",
			m:    map[int]int{},
			want: sets.New[int](),
		},
		{
			name: "map with unique keys",
			m:    map[int]int{1: 10, 2: 20, 3: 30},
			want: sets.New(1, 2, 3),
		},
		{
			name: "map with duplicate values",
			m:    map[int]int{1: 10, 2: 20, 3: 10},
			want: sets.New(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sets.NewFromMapKeys(tt.m)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestSet_Insert(t *testing.T) {
	tests := []struct {
		name string
		s    sets.Set[int]
		item int
		want sets.Set[int]
	}{
		{
			name: "insert into empty set",
			s:    sets.New[int](),
			item: 1,
			want: sets.New(1),
		},
		{
			name: "insert existing item",
			s:    sets.New(1, 2, 3),
			item: 2,
			want: sets.New(1, 2, 3),
		},
		{
			name: "insert new item",
			s:    sets.New(1, 2, 3),
			item: 4,
			want: sets.New(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Insert(tt.item)
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestSet_Delete(t *testing.T) {
	tests := []struct {
		name string
		set  sets.Set[int]
		item int
		want sets.Set[int]
	}{
		{
			name: "delete from empty set",
			set:  sets.New[int](),
			item: 1,
			want: sets.New[int](),
		},
		{
			name: "delete existing item",
			set:  sets.New(1, 2, 3),
			item: 2,
			want: sets.New(1, 3),
		},
		{
			name: "delete non-existing item",
			set:  sets.New(1, 2, 3),
			item: 4,
			want: sets.New(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Delete(tt.item)
			assert.Equal(t, tt.want, tt.set)
		})
	}
}

func TestSet_Has(t *testing.T) {
	tests := []struct {
		name string
		set  sets.Set[int]
		item int
		want bool
	}{
		{
			name: "item in set",
			set:  sets.New(1, 2, 3),
			item: 2,
			want: true,
		},
		{
			name: "item not in set",
			set:  sets.New(1, 2, 3),
			item: 4,
			want: false,
		},
		{
			name: "empty set",
			set:  sets.New[int](),
			item: 1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.Has(tt.item))
		})
	}
}

func TestSet_List(t *testing.T) {
	tests := []struct {
		name string
		set  sets.Set[int]
		want []int
	}{
		{
			name: "empty set",
			set:  sets.New[int](),
			want: []int{},
		},
		{
			name: "set with unique items",
			set:  sets.New(1, 2, 3),
			want: []int{1, 2, 3},
		},
		{
			name: "set with duplicates (should not affect output)",
			set:  sets.New(1, 2, 3, 2, 1),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, tt.set.List())
		})
	}
}

func TestSorted(t *testing.T) {
	tests := []struct {
		name string
		s    sets.Set[int]
		want []int
	}{
		{
			name: "empty set",
			s:    sets.New[int](),
			want: []int{},
		},
		{
			name: "set with unique items",
			s:    sets.New(3, 1, 2),
			want: []int{1, 2, 3},
		},
		{
			name: "set with duplicates",
			s:    sets.New(3, 1, 2, 2, 1),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sets.Sorted(tt.s)
			assert.Equal(t, tt.want, s)
		})
	}
}
