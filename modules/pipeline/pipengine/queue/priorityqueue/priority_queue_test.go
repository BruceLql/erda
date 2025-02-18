// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package priorityqueue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue_Get(t *testing.T) {
	pq := NewPriorityQueue()

	get := pq.Get("k1")
	assert.Nil(t, get, "no BaseItem now")

	pq.Add(NewItem("k1", 0, time.Time{}))
	get = pq.Get("k1")
	assert.NotNil(t, get, "only k1 now")
	assert.Equal(t, "k1", get.Key(), "only k1 now")
}

func TestPriorityQueue_Peek(t *testing.T) {
	pq := NewPriorityQueue()

	now := time.Now()

	pq.Add(NewItem("k1", 1, time.Now()))
	peeked := pq.Peek()
	assert.NotNil(t, peeked, "only k1 now")
	assert.Equal(t, peeked.Key(), "k1", "only k1 now")

	pq.Add(NewItem("k1", 1, now))
	pq.Add(NewItem("k2", 1, now.Add(time.Second)))
	peeked = pq.Peek()
	assert.NotNil(t, peeked, "k1 and k2")
	assert.Equal(t, peeked.Key(), "k1", "k1,k2 have same priority, but k1 is earlier")

	pq.Add(NewItem("k3", 1, now.Add(-time.Second)))
	pq.Add(NewItem("k3", 1, now.Add(-time.Second)))
	peeked = pq.Peek()
	assert.NotNil(t, peeked, "k1, k2, k3")
	assert.Equal(t, peeked.Key(), "k3", "k1,k2,k3 have same priority, but k3 is the earliest")

	pq.Add(NewItem("k4", 2, now.Add(time.Hour)))
	peeked = pq.Peek()
	assert.NotNil(t, peeked, "k1, k2, k3, k4")
	assert.Equal(t, peeked.Key(), "k4", "k4's priority is highest")

	// priority: k4 > k3 > k1 > k2
	assert.Equal(t, "k4", pq.data.items[0].Key())
	assert.Equal(t, "k3", pq.data.items[1].Key())
	assert.Equal(t, "k1", pq.data.items[2].Key())
	assert.Equal(t, "k2", pq.data.items[3].Key())
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq := NewPriorityQueue()
	popped := pq.Pop()
	assert.Nil(t, popped, "no BaseItem now")

	now := time.Now()

	pq.Add(NewItem("k1", 1, now))
	popped = pq.Pop()
	assert.NotNil(t, popped)
	assert.Equal(t, popped.Key(), "k1", "only k1 now")

	pq.Add(NewItem("k2", 1, now))
	pq.Add(NewItem("k3", 2, now))
	popped = pq.Pop()
	assert.NotNil(t, popped)
	assert.Equal(t, popped.Key(), "k3", "k3's priority is higher than k2")
	popped = pq.Pop()
	assert.NotNil(t, popped)
	assert.Equal(t, popped.Key(), "k2", "only k2 now")
	popped = pq.Pop()
	assert.Nil(t, popped, "no BaseItem now, all popped out")
	popped = pq.Pop()
	assert.Nil(t, popped, "no BaseItem now, all popped out")

	pq.Add(NewItem("k4", 1, now))
	pq.Add(NewItem("k5", 1, now.Add(-time.Second)))
	popped = pq.Pop()
	assert.NotNil(t, popped)
	assert.Equal(t, popped.Key(), "k5", "k5 is earlier than p4")
}

func TestPriorityQueue_Add(t *testing.T) {
	pq := NewPriorityQueue()
	assert.Equal(t, 0, len(pq.data.items), "no BaseItem now")

	now := time.Now()
	pq.Add(NewItem("k1", 1, now))
	assert.Equal(t, 1, len(pq.data.items), "only k1 now")
	get := pq.Get("k1")
	assert.Equal(t, "k1", get.Key(), "only k1 now")
	assert.Equal(t, int64(1), get.Priority())

	pq.Add(NewItem("k1", 2, now))
	assert.Equal(t, 1, len(pq.data.itemByKey), "still only k1 now")
	get = pq.Get("k1")
	assert.Equal(t, "k1", get.Key(), "still only k1 now")
	assert.Equal(t, int64(2), get.Priority(), "k1's priority updated to 2")
}

type obj struct {
	name      string
	createdAt time.Time
	index     int
	priority  int64
}

func (o obj) Key() string             { return o.name }
func (o obj) Priority() int64         { return 1 }
func (o obj) SetPriority(i int64)     { o.priority = i }
func (o obj) CreationTime() time.Time { return o.createdAt }
func (o obj) Index() int              { return o.index }
func (o obj) SetIndex(i int)          { o.index = i }

func TestPriorityQueue_Add2(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Add(obj{name: "k1", createdAt: time.Now(), priority: 1})
	pq.Add(obj{name: "k2", createdAt: time.Now(), priority: 1})
	popped := pq.Pop()
	assert.NotNil(t, popped)
	assert.Equal(t, "k1", popped.Key())
}

func TestPriorityQueue_Remove(t *testing.T) {
	pq := NewPriorityQueue()

	removed := pq.Remove("k1")
	assert.Nil(t, removed, "no BaseItem now")

	pq.Add(NewItem("k1", 0, time.Time{}))
	removed = pq.Remove("k2")
	assert.Nil(t, removed, "k2 not exist")
	removed = pq.Remove("k1")
	assert.NotNil(t, removed, "k1 exist and removed")
	removed = pq.Remove("k1")
	assert.Nil(t, removed, "k1 already been removed")
}
