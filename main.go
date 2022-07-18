package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/drewstinnett/box-packing/packit"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"

	"github.com/gedex/bp3d"
)

var capacity int64

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Int64("capacity", capacity).Msg("Starting packit with capacity")

	p := packit.Packit{}
	_ = p
	fullSum := int64(0)
	for _, a := range flag.Args() {
		s, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to parse size")
		}
		i := SampleItem{Size: s}
		fullSum += s
		p.AddItem(&i)
	}
	minBucketC := int(math.Ceil(float64(fullSum) / float64(capacity)))
	log.Info().Int64("fullSum", fullSum).Int("min-buckets", minBucketC).Msg("Total size of all items")
	err := p.QuickPackWithMaxSize(capacity)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to pack")
	}
	boxCount := len(p.Boxes)
	for x := 0; x < boxCount; x++ {
		b := p.Boxes[x]
		fmt.Printf("Box: %v %v/%v\n", x+1, b.Used, b.Capacity)
		for _, i := range b.Items {
			fmt.Printf("  %v\n", i)
		}
	}
	if boxCount != minBucketC {
		log.Fatal().Msg("Optimal bucket count does not match actual bucket count")
	}
}

func init() {
	flag.Int64Var(&capacity, "capacity", 10, "Capacity of the box")
	flag.Parse()
}

func displayPacked(bins []*bp3d.Bin) {
	for _, b := range bins {
		fmt.Println(b)
		fmt.Println(" packed items:")
		for _, i := range b.Items {
			fmt.Println("  ", i)
		}
	}
}

type SampleItem struct {
	Size int64
}

func (s *SampleItem) GetSize() int64 {
	return s.Size
}
