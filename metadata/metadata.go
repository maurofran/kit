package metadata

// Supplier is the interface for supplying a value
type Supplier func() interface{}

// Container is the metadata containe object.
type Container struct {
	data map[string]interface{}
}

// Get will retrieve the value of supplied key.
func (c *Container) Get(key string) (interface{}, bool) {
	if c.Empty() {
		return nil, false
	}
	val, ok := c.data[key]
	return val, ok
}

// Keys will retrieve a slice of all the keys in container.
func (c *Container) Keys() []string {
	keys := make([]string, 0)
	if c.Empty() {
		return keys
	}
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}

// Empty will check if container is empty
func (c *Container) Empty() bool {
	return len(c.data) == 0
}

// And will return a new metadata instance with supplied key and value.
func (c *Container) And(key string, value interface{}) *Container {
	res := From(c.data)
	res.data[key] = value
	return res
}

// AndIfNotPresent will return a new metadata instance with supplied key and value obtained from
// supplier, invoked only if key is not present.
func (c *Container) AndIfNotPresent(key string, value Supplier) *Container {
	if c.Empty() {
		return c.And(key, value())
	}
	if _, ok := c.data[key]; !ok {
		return c.And(key, value())
	}
	return c
}

// MergedWith will return a new metadata object with supplied entries.
func (c *Container) MergedWith(entries map[string]interface{}) *Container {
	if len(entries) == 0 {
		return c
	}
	if c.Empty() {
		return From(entries)
	}
	res := From(c.data)
	for key, value := range entries {
		res.data[key] = value
	}
	return res
}

// WithoutKeys will retrieve a new metadata object without the supplied keys.
func (c *Container) WithoutKeys(keys ...string) *Container {
	if len(keys) == 0 {
		return c
	}
	if c.Empty() {
		return c
	}
	res := From(c.data)
	for _, key := range keys {
		delete(res.data, key)
	}
	return res
}

// WithKeys will retrieve a new metadata object with only supplied keys.
func (c *Container) WithKeys(keys ...string) *Container {
	if len(keys) == 0 {
		return Empty()
	}
	if c.Empty() {
		return c
	}
	res := &Container{data: make(map[string]interface{})}
	for _, key := range keys {
		res.data[key] = c.data[key]
	}
	return res
}

// Empty will return an empty metadata object.
func Empty() *Container {
	return new(Container)
}

// From will initialize a metadata object with supplied data.
func From(source map[string]interface{}) *Container {
	c := &Container{data: make(map[string]interface{})}
	for key, value := range source {
		c.data[key] = value
	}
	return c
}

// With will create a new metadata object with supplied key and value.
func With(key string, value interface{}) *Container {
	c := &Container{data: make(map[string]interface{})}
	c.data[key] = value
	return c
}
