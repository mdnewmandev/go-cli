package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapEntries(t *testing.T) {
	cases := []struct {
		name         string
		reapInterval time.Duration
		waitTime     time.Duration
		shouldReap   bool
	}{
		{
			name:         "Entry should remain when checked before reap interval",
			reapInterval: 100 * time.Millisecond,
			waitTime:     50 * time.Millisecond,
			shouldReap:   false,
		},
		{
			name:         "Entry should be reaped when checked after reap interval",
			reapInterval: 100 * time.Millisecond,
			waitTime:     150 * time.Millisecond,
			shouldReap:   true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(c.reapInterval)
			cache.Add("test-key", []byte("test-value"))
			
			// Force a reap after waiting
			time.Sleep(c.waitTime)
			cache.reapEntries(c.reapInterval)
			
			_, ok := cache.Get("test-key")
			if c.shouldReap && ok {
				t.Error("expected entry to be reaped but it still exists")
			}
			if !c.shouldReap && !ok {
				t.Error("expected entry to exist but it was reaped")
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}