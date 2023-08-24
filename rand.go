package zzrand

import (
	"math/bits"
	"sync"
	"time"
)

func Intn(n int) int {
	v := pool.Get()
	if v == nil {
		v = &RNG{}
	}
	r := v.(*RNG)
	x := r.Intn(n)
	pool.Put(r)
	return x
}

func Int() int {
	v := pool.Get()
	if v == nil {
		v = &RNG{}
	}
	r := v.(*RNG)
	x := r.Int()
	pool.Put(r)
	return x
}

func Uint32() uint32 {
	v := pool.Get()
	if v == nil {
		v = &RNG{}
	}
	r := v.(*RNG)
	x := r.Uint32()
	pool.Put(r)
	return x
}

func Int32() int32 {
	v := pool.Get()
	if v == nil {
		v = &RNG{}
	}
	r := v.(*RNG)
	x := r.Int32()
	pool.Put(r)
	return x
}

func Uint64() uint64 {
	v := pool.Get()
	if v == nil {
		v = &RNG{}
	}
	r := v.(*RNG)
	x := r.Uint64()
	pool.Put(r)
	return x
}
func Int64() int64 {
	v := pool.Get()
	if v == nil {
		v = &RNG{}
	}
	r := v.(*RNG)
	x := r.Int64()
	pool.Put(r)
	return x
}

var pool sync.Pool

type RNG struct {
	x uint64
}

func (r *RNG) Uint64() uint64 {
	for r.x == 0 {
		r.x = randomSeedUint64()
	}
	// wyhash
	r.x += 0xa0761d6478bd642f
	hi, lo := bits.Mul64(r.x, r.x^0xe7037ed1a0b428db)
	return hi ^ lo
}

func (r *RNG) Uint32() uint32 {
	return uint32(r.Uint64())
}

func (r *RNG) Int32() int32 {
	return int32(r.Uint64() >> 33)
}

func (r *RNG) Int() int {
	return int(r.Uint64() >> 1)
}
func (r *RNG) Int64() int64 {
	return int64(r.Uint64() >> 1)
}

func (r *RNG) Intn(n int) int {
	x := r.Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return int((uint64(x) * uint64(n)) >> 32)
}

func randomSeedUint64() uint64 {
	x := time.Now().UnixNano()
	return uint64((x >> 32) ^ x)
}
