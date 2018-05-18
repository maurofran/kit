package metadata_test

import (
	"testing"

	"github.com/maurofran/kit/metadata"
	. "github.com/maurofran/kit/testing"
)

func anEmptyContainer() *metadata.Container {
	return metadata.Empty()
}

func aContainer() *metadata.Container {
	return metadata.From(map[string]interface{}{
		"key1": "value1",
		"key2": 2,
	})
}

func TestEmpty(t *testing.T) {
	c := metadata.Empty()

	Assert(t, c != nil, "Metadata container bust not be nil")
	Equals(t, c.Empty(), true)
	_, ok := c.Get("aKey")
	Equals(t, ok, false)
}

func TestWith(t *testing.T) {
	c := metadata.With("aKey", "aValue")

	Assert(t, c != nil, "Metadata container must not be nil")
	Equals(t, c.Empty(), false)
	val, ok := c.Get("aKey")
	Equals(t, ok, true)
	Equals(t, val, "aValue")
}

func TestFrom(t *testing.T) {
	from := make(map[string]interface{})
	from["key1"] = "value1"
	from["key2"] = 2
	c := metadata.From(from)

	Assert(t, c != nil, "Metadata container must not be nil")
	Equals(t, c.Empty(), false)
	val, ok := c.Get("key1")
	Equals(t, ok, true)
	Equals(t, val, "value1")
	v2, ok := c.Get("key2")
	Equals(t, ok, true)
	Equals(t, v2, 2)
}

func TestGet_OnEmpty(t *testing.T) {
	_, ok := anEmptyContainer().Get("aKey")

	Equals(t, false, ok)
}

func TestGet_MissingKey(t *testing.T) {
	_, ok := aContainer().Get("missing")

	Equals(t, false, ok)
}

func TestGet_ExistingKey(t *testing.T) {
	value, ok := aContainer().Get("key1")

	Equals(t, value, "value1")
	Equals(t, ok, true)
}

func TestKeys_OnEmpty(t *testing.T) {
	keys := anEmptyContainer().Keys()

	Equals(t, 0, len(keys))
}

func TestKeys_TwoKeys(t *testing.T) {
	keys := aContainer().Keys()

	Equals(t, 2, len(keys))
}

func TestAnd(t *testing.T) {
	c := aContainer()
	m := c.And("aKey", "aValue")

	Assert(t, m != c, "should return a new container")
	value, ok := m.Get("aKey")
	Equals(t, true, ok)
	Equals(t, "aValue", value)
	Equals(t, 3, len(m.Keys()))
}

func TestAndIfNotPresent_OnEmpty(t *testing.T) {
	invoked := false
	c := anEmptyContainer()
	m := c.AndIfNotPresent("aKey", func() interface{} {
		invoked = true
		return "aValue"
	})

	Assert(t, m != c, "should return a new container")
	value, ok := m.Get("aKey")
	Equals(t, true, ok)
	Equals(t, "aValue", value)
	Equals(t, 1, len(m.Keys()))
	Equals(t, true, invoked)
}

func TestAndIfNotPresent_MissingKey(t *testing.T) {
	invoked := false
	c := aContainer()
	m := c.AndIfNotPresent("aKey", func() interface{} {
		invoked = true
		return "aValue"
	})

	Assert(t, m != c, "should return a new container")
	value, ok := m.Get("aKey")
	Equals(t, true, ok)
	Equals(t, "aValue", value)
	Equals(t, 3, len(m.Keys()))
	Equals(t, true, invoked)
}

func TestAndIfNotPresent_ExistingKey(t *testing.T) {
	invoked := false
	c := aContainer()
	m := c.AndIfNotPresent("key1", func() interface{} {
		invoked = true
		return "aValue"
	})

	Assert(t, m == c, "should return the same container")
	value, ok := m.Get("key1")
	Equals(t, true, ok)
	Equals(t, "value1", value)
	Equals(t, 2, len(m.Keys()))
	Equals(t, false, invoked)
}

func TestMergedWith_EmptyEntries(t *testing.T) {
	c := aContainer()
	m := c.MergedWith(nil)

	Assert(t, m == c, "should return the same container")
}

func TestMergedWith_OnEmpty(t *testing.T) {
	c := anEmptyContainer()
	m := c.MergedWith(map[string]interface{}{
		"key3": "value3",
		"key4": "value4",
	})

	Assert(t, c != m, "should return a new container")
	Equals(t, 2, len(m.Keys()))
}

func TestMergedWith_OnExisting(t *testing.T) {
	c := aContainer()
	m := c.MergedWith(map[string]interface{}{
		"key3": "value3",
		"key4": "value4",
	})

	Assert(t, c != m, "should return a new container")
	Equals(t, 4, len(m.Keys()))
}

func TestWithKeys_OnEmpty(t *testing.T) {
	c := anEmptyContainer()
	m := c.WithKeys("key1", "key2")

	Assert(t, c == m, "should return the same container")
}

func TestWithKeys_WithNoKeys(t *testing.T) {
	c := aContainer()
	m := c.WithKeys()

	Assert(t, c != m, "should return a new container")
	Equals(t, true, m.Empty())
}

func TestWithKeys_RetainOnlyKeys(t *testing.T) {
	c := aContainer().And("key3", "value3").And("key4", 4)
	m := c.WithKeys("key1", "key3")

	Assert(t, c != m, "should return a new container")
	Equals(t, 2, len(m.Keys()))
	value, ok := m.Get("key1")
	Equals(t, true, ok)
	Equals(t, "value1", value)
	value, ok = m.Get("key3")
	Equals(t, true, ok)
	Equals(t, "value3", value)
}

func TestWithoutKeys_OnEmpty(t *testing.T) {
	c := anEmptyContainer()
	m := c.WithoutKeys("key1", "key2")

	Assert(t, c == m, "should return the same container")
}

func TestWithoutKeys_WithNoKeys(t *testing.T) {
	c := aContainer()
	m := c.WithoutKeys()

	Assert(t, c == m, "should return the same container")
	Equals(t, 2, len(m.Keys()))
}

func TestWithoutKeys_ShouldRemoveKeys(t *testing.T) {
	c := aContainer().And("key3", "value3").And("key4", 4)
	m := c.WithoutKeys("key1", "key3")

	Assert(t, c != m, "should return a new container")
	Equals(t, 2, len(m.Keys()))
	value, ok := m.Get("key2")
	Equals(t, true, ok)
	Equals(t, 2, value)
	value, ok = m.Get("key4")
	Equals(t, true, ok)
	Equals(t, 4, value)
}
