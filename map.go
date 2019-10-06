package main

import (
	"encoding/xml"
	"image"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	tileset     *ebiten.Image
	tilesetOpts *ebiten.DrawImageOptions
	mapData     *mapxml
	tiles       []*ebiten.Image
)

type chunk struct {
	DataStr string `xml:",chardata"`
	X       int    `xml:"x,attr"`
	Y       int    `xml:"y,attr"`
	Width   int    `xml:"width,attr"`
	Height  int    `xml:"height,attr"`
	Data    []uint
}

type layer struct {
	XMLName xml.Name `xml:"layer"`
	Chunk   []*chunk `xml:"data>chunk"`
}

type mapxml struct {
	XMLName xml.Name `xml:"map"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Layer   []*layer `xml:"layer"`
}

type tilesetxml struct {
	XMLName xml.Name `xml:"tileset"`
}

func init() {
	var err error
	tileset, _, err = ebitenutil.NewImageFromFile("tileset.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	tilesetOpts = &ebiten.DrawImageOptions{}

	mapfile, _ := ioutil.ReadFile("map.tmx")
	mapData = &mapxml{}
	xml.Unmarshal(mapfile, mapData)

	for _, l := range mapData.Layer {
		for _, c := range l.Chunk {
			indices := strings.Split(c.DataStr, ",")
			c.Data = make([]uint, len(indices))
			for i := range indices {
				str := strings.TrimSpace(indices[i])
				index, _ := strconv.ParseUint(str, 10, 32)
				c.Data[i] = uint(index)
			}
		}
	}
	tiles = make([]*ebiten.Image, 16)
	for i := range tiles {
		tw, _ := tileset.Size()
		x := (i % tw) * 24
		y := (i / tw) * 24
		tiles[i] = tileset.SubImage(image.Rect(x, y, x+24, y+24)).(*ebiten.Image)
	}

}

func mapUpdate(screen *ebiten.Image) {

	for _, l := range mapData.Layer {
		for _, c := range l.Chunk {
			tilesetOpts.GeoM.Reset()
			tilesetOpts.GeoM.Translate(cameraX, cameraY)
			tilesetOpts.GeoM.Translate(float64((c.X)*24), float64((c.Y)*24))

			for i, d := range c.Data {
				if d > 0 {
					screen.DrawImage(tiles[d-1], tilesetOpts)
				}
				tilesetOpts.GeoM.Translate(24, 0)
				if i%c.Width == c.Width-1 {
					tilesetOpts.GeoM.Reset()
					tilesetOpts.GeoM.Translate(cameraX, cameraY)
					tilesetOpts.GeoM.Translate(float64((c.X)*24), float64((c.Y)*24))
					tilesetOpts.GeoM.Translate(0, 24*float64((i+1)/c.Width))
				}
			}
		}
	}
}
