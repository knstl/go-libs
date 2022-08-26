package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	i, err := initRedis(t)
	assert.Nil(t, err)
	assert.NotNil(t, i)
}

func TestRedis_Set(t *testing.T) {
	r, err := initRedis(t)
	assert.Nil(t, err)

	testData := struct {
		Field string
		F     float64
	}{"1", 200.1}

	err = r.Set(context.TODO(), "test", testData, time.Second)
	assert.Nil(t, err)

	var newValue struct {
		Field string
		F     float64
	}
	err = r.Get(context.TODO(), "test", &newValue)
	assert.Nil(t, err)
	assert.Equal(t, testData, newValue)

	ttl := r.client.TTL(context.TODO(), "test")
	assert.Equal(t, time.Second, ttl.Val())
}

func TestRedis_Get(t *testing.T) {
	r, err := initRedis(t)
	assert.Nil(t, err)

	type TestStruct struct {
		Field string
		F     float64
	}
	testData := TestStruct{"1", 200.1}

	err = r.Set(context.TODO(), "test", testData, time.Second)
	assert.Nil(t, err)

	var newValue TestStruct
	err = r.Get(context.TODO(), "test", &newValue)
	assert.Nil(t, err)
	assert.Equal(t, testData, newValue)

	ttl := r.client.TTL(context.TODO(), "test")
	assert.Equal(t, time.Second, ttl.Val())

	var empty interface{}
	err = r.Get(context.TODO(), "1", empty)
	assert.Equal(t, err, ErrKeyNotFound)
	assert.Nil(t, empty)
}

func TestRedis_Delete(t *testing.T) {
	r, err := initRedis(t)
	assert.Nil(t, err)
	err = r.Set(context.TODO(), "test", []byte{0, 1}, time.Second)
	assert.Nil(t, err)

	err = r.Delete(context.TODO(), "test")
	assert.Nil(t, err)

	var v []byte
	err = r.Get(context.TODO(), "test", &v)
	assert.NotNil(t, err)
	assert.Equal(t, string([]byte{}), string(v))
}

func TestRedis_IsAvailable(t *testing.T) {
	r, err := initRedis(t)
	assert.Nil(t, err)

	assert.True(t, r.IsAvailable(context.TODO()))
}

func TestRedis_Reconnect(t *testing.T) {
	r, err := initRedis(t)
	assert.Nil(t, err)

	mr, err := miniredis.Run()
	assert.NotNil(t, mr)
	assert.Nil(t, err)

	err = r.Reconnect(context.TODO(), fmt.Sprintf("redis://%s", mr.Addr()))
	assert.NoError(t, err)
}

func initRedis(t *testing.T) (*Redis, error) {
	mr, err := miniredis.Run()
	assert.NotNil(t, mr)
	assert.Nil(t, err)

	c, err := NewClient(context.TODO(), fmt.Sprintf("redis://%s", mr.Addr()))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	return c, nil
}
