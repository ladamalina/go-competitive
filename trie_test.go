package main

import (
	"testing"
)

func TestTriePutGet(t *testing.T) {
	trie := NewTrie[int]()

	// Test basic Put and Get
	trie.Put("hello", 42)
	val, ok := trie.Get("hello")
	if !ok || val != 42 {
		t.Errorf("Get failed: want (42, true), got (%d, %v)", val, ok)
	}

	// Test Get non-existent key
	_, ok = trie.Get("world")
	if ok {
		t.Errorf("Get non-existent key should return false, got (_, true)")
	}
}

func TestTriePutGetMultiple(t *testing.T) {
	trie := NewTrie[string]()

	tests := []struct {
		key string
		val string
	}{
		{"apple", "fruit"},
		{"app", "software"},
		{"application", "software"},
		{"apply", "verb"},
		{"banana", "fruit"},
		{"band", "music"},
	}

	// Insert all
	for _, test := range tests {
		trie.Put(test.key, test.val)
	}

	// Verify all
	for _, test := range tests {
		val, ok := trie.Get(test.key)
		if !ok || val != test.val {
			t.Errorf("Get(%q) failed: want (%q, true), got (%q, %v)", test.key, test.val, val, ok)
		}
	}
}

func TestTriePutOverwrite(t *testing.T) {
	trie := NewTrie[int]()

	trie.Put("key", 1)
	val, ok := trie.Get("key")
	if !ok || val != 1 {
		t.Errorf("First Put failed: want (1, true), got (%d, %v)", val, ok)
	}

	// Overwrite
	trie.Put("key", 2)
	val, ok = trie.Get("key")
	if !ok || val != 2 {
		t.Errorf("Overwrite failed: want (2, true), got (%d, %v)", val, ok)
	}
}

func TestTrieBPutBGet(t *testing.T) {
	trie := NewTrie[float64]()

	// Test with byte slices
	trie.BPut([]byte("test"), 3.14)
	val, ok := trie.BGet([]byte("test"))
	if !ok || val != 3.14 {
		t.Errorf("BGet failed: want (3.14, true), got (%f, %v)", val, ok)
	}

	// Test non-existent byte key
	_, ok = trie.BGet([]byte("missing"))
	if ok {
		t.Errorf("BGet non-existent key should return false")
	}
}

func TestTriePrefixRelationship(t *testing.T) {
	trie := NewTrie[int]()

	// Add "cat"
	trie.Put("cat", 1)

	// "c" and "ca" should not be in trie
	_, ok := trie.Get("c")
	if ok {
		t.Errorf("Prefix 'c' should not exist")
	}

	_, ok = trie.Get("ca")
	if ok {
		t.Errorf("Prefix 'ca' should not exist")
	}

	// "cat" should exist
	val, ok := trie.Get("cat")
	if !ok || val != 1 {
		t.Errorf("'cat' should exist")
	}

	// "cath" should not exist
	_, ok = trie.Get("cath")
	if ok {
		t.Errorf("Extension 'cath' should not exist")
	}

	// Add "ca" as a separate key
	trie.Put("ca", 2)
	val, ok = trie.Get("ca")
	if !ok || val != 2 {
		t.Errorf("'ca' should now exist with value 2, got (%d, %v)", val, ok)
	}

	// Both "ca" and "cat" should still exist with original values
	val, ok = trie.Get("cat")
	if !ok || val != 1 {
		t.Errorf("'cat' should still exist with value 1")
	}
}

func TestTrieEmptyString(t *testing.T) {
	trie := NewTrie[int]()

	trie.Put("", 42)
	val, ok := trie.Get("")
	if !ok || val != 42 {
		t.Errorf("Empty string Get failed: want (42, true), got (%d, %v)", val, ok)
	}
}

func TestTrieSpecialCharacters(t *testing.T) {
	trie := NewTrie[string]()

	keys := []string{
		"hello-world",
		"hello_world",
		"hello@world",
		"hello.world",
		"hello/world",
		"hello:world",
	}

	for _, key := range keys {
		trie.Put(key, key)
	}

	for i, key := range keys {
		val, ok := trie.Get(key)
		if !ok || val != key {
			t.Errorf("Special char key %d (%q) failed: want (%q, true), got (%q, %v)", i, key, key, val, ok)
		}
	}
}

func TestTrieNumericKeys(t *testing.T) {
	trie := NewTrie[int]()

	// Numeric keys as strings
	keys := []string{"0", "1", "10", "123", "999"}

	for i, key := range keys {
		trie.Put(key, i)
	}

	for i, key := range keys {
		val, ok := trie.Get(key)
		if !ok || val != i {
			t.Errorf("Numeric key %q failed: want (%d, true), got (%d, %v)", key, i, val, ok)
		}
	}
}

func TestTrieLongKeys(t *testing.T) {
	trie := NewTrie[int]()

	// Test with very long keys
	longKey := ""
	for i := 0; i < 1000; i++ {
		longKey += "a"
	}

	trie.Put(longKey, 42)
	val, ok := trie.Get(longKey)
	if !ok || val != 42 {
		t.Errorf("Long key failed: want (42, true), got (%d, %v)", val, ok)
	}

	// Shorter prefix should not exist
	_, ok = trie.Get(longKey[:999])
	if ok {
		t.Errorf("Shorter prefix of long key should not exist")
	}
}

func TestTrieZeroValues(t *testing.T) {
	trie := NewTrie[int]()

	// Store zero value (default int value is 0)
	trie.Put("zero", 0)
	val, ok := trie.Get("zero")
	if !ok {
		t.Errorf("Zero value should be retrievable, got ok=false")
	}
	if val != 0 {
		t.Errorf("Zero value should be 0, got %d", val)
	}

	// Non-existent key should also return 0 but ok=false
	_, ok = trie.Get("nonexistent")
	if ok {
		t.Errorf("Non-existent key should return ok=false")
	}
}

func TestTrieStringValues(t *testing.T) {
	trie := NewTrie[string]()

	trie.Put("name", "Alice")
	trie.Put("age", "30")
	trie.Put("city", "NYC")

	tests := []struct {
		key   string
		want  string
		found bool
	}{
		{"name", "Alice", true},
		{"age", "30", true},
		{"city", "NYC", true},
		{"country", "", false},
	}

	for _, test := range tests {
		val, ok := trie.Get(test.key)
		if ok != test.found || val != test.want {
			t.Errorf("Get(%q) failed: want (%q, %v), got (%q, %v)", test.key, test.want, test.found, val, ok)
		}
	}
}

func TestTrieByteSliceKeys(t *testing.T) {
	trie := NewTrie[int]()

	keys := [][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}

	for i, key := range keys {
		trie.BPut(key, i)
	}

	for _, key := range keys {
		_, ok := trie.BGet(key)
		if !ok {
			t.Errorf("BGet(%q) failed: key not found", key)
		}
	}
}

func TestTrieCommonPrefixes(t *testing.T) {
	trie := NewTrie[int]()

	// Test with many keys sharing common prefixes
	keys := []struct {
		key string
		val int
	}{
		{"a", 1},
		{"ab", 2},
		{"abc", 3},
		{"abcd", 4},
		{"abcde", 5},
		{"abd", 6},
		{"ac", 7},
	}

	for _, k := range keys {
		trie.Put(k.key, k.val)
	}

	for _, k := range keys {
		val, ok := trie.Get(k.key)
		if !ok || val != k.val {
			t.Errorf("Get(%q) failed: want (%d, true), got (%d, %v)", k.key, k.val, val, ok)
		}
	}

	// Test non-existent intermediate keys
	_, ok := trie.Get("ab" + "x")
	if ok {
		t.Errorf("Non-existent key 'abx' should not be found")
	}
}

func TestTrieConcurrentNodes(t *testing.T) {
	trie := NewTrie[string]()

	// Test keys that diverge at different points
	keys := []struct {
		key string
		val string
	}{
		{"test", "value1"},
		{"team", "value2"},
		{"tea", "value3"},
		{"top", "value4"},
		{"tip", "value5"},
	}

	for _, k := range keys {
		trie.Put(k.key, k.val)
	}

	for _, k := range keys {
		val, ok := trie.Get(k.key)
		if !ok || val != k.val {
			t.Errorf("Get(%q) failed: want (%q, true), got (%q, %v)", k.key, k.val, val, ok)
		}
	}
}

func TestTrieBinaryData(t *testing.T) {
	trie := NewTrie[int]()

	// Test with various byte values (0-255)
	keys := []struct {
		key   []byte
		val   int
		descr string
	}{
		{[]byte{0}, 1, "null byte"},
		{[]byte{255}, 2, "max byte"},
		{[]byte{1, 2, 3}, 3, "sequential bytes"},
		{[]byte{255, 255, 255}, 4, "max bytes"},
		{[]byte{0, 128, 255}, 5, "mixed bytes"},
	}

	for _, k := range keys {
		trie.BPut(k.key, k.val)
	}

	for _, k := range keys {
		val, ok := trie.BGet(k.key)
		if !ok || val != k.val {
			t.Errorf("BGet(%s) failed: want (%d, true), got (%d, %v)", k.descr, k.val, val, ok)
		}
	}
}
