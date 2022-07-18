package packit

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

type SampleItem struct {
	Size int64
}

func (s *SampleItem) GetSize() int64 {
	return s.Size
}

func TestNewPackit(t *testing.T) {
	p := Packit{}
	require.NotNil(t, p)

	i := SampleItem{Size: 1}
	err := p.AddItem(&i)
	require.NoError(t, err)
}

func TestQuickPack(t *testing.T) {
	p := Packit{}
	sizes := []int64{1, 2, 3}
	for _, size := range sizes {
		i := SampleItem{Size: size}
		err := p.AddItem(&i)
		require.NoError(t, err)
	}
	err := p.QuickPackWithMaxSize(3)
	require.NoError(t, err)

	for _, b := range p.Boxes {
		log.Info().Msgf("Box: %+v", b)
	}
	require.Equal(t, &Box{Capacity: 3, Used: 3, Items: []BoxItem{&SampleItem{Size: 3}}}, p.Boxes[0])
	require.Equal(t, &Box{Capacity: 3, Used: 3, Items: []BoxItem{
		&SampleItem{Size: 2},
		&SampleItem{Size: 1},
	}}, p.Boxes[1])
}

func TestQuickPackOdd(t *testing.T) {
	p := Packit{}
	sizes := []int64{1, 2, 3, 4}
	for _, size := range sizes {
		i := SampleItem{Size: size}
		err := p.AddItem(&i)
		require.NoError(t, err)
	}
	err := p.QuickPackWithMaxSize(4)
	require.NoError(t, err)

	for _, b := range p.Boxes {
		log.Info().Msgf("Box: %+v", b)
	}
	require.Equal(t, 3, len(p.Boxes))
}

func TestNoMax(t *testing.T) {
	p := Packit{}
	sizes := []int64{1, 2}
	for _, size := range sizes {
		i := SampleItem{Size: size}
		err := p.AddItem(&i)
		require.NoError(t, err)
	}
	err := p.QuickPackWithMaxSize(0)
	require.NoError(t, err)
	require.Equal(t, &Box{Capacity: 0, Used: 3, Items: []BoxItem{
		&SampleItem{Size: 2},
		&SampleItem{Size: 1},
	}}, p.Boxes[0])
}

func TestOverMax(t *testing.T) {
	p := Packit{}
	sizes := []int64{1, 2}
	for _, size := range sizes {
		i := SampleItem{Size: size}
		err := p.AddItem(&i)
		require.NoError(t, err)
	}
	err := p.QuickPackWithMaxSize(1)
	require.Error(t, err)
	require.EqualError(t, err, "index containes at least one file that is too large")
}

func TestSlowPack(t *testing.T) {
	p := Packit{}
	sizes := []int64{4, 5, 5, 5, 4, 3, 5}
	for _, size := range sizes {
		i := SampleItem{Size: size}
		err := p.AddItem(&i)
		require.NoError(t, err)
	}
	cap := int64(10)
	err := p.SlowPackWithMaxSize(cap)
	require.NoError(t, err)

	for _, b := range p.Boxes {
		log.Info().Msgf("Box: %+v", b)
	}
	require.Equal(t, 4, len(p.Boxes))
}
