package main

import (
	"fmt"
)

// soal 1
type fruit struct {
	name        string
	warna       string
	ada_bijinya bool
	harga       int
}

// soal 2

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) LuasSegitiga() int {
	return (s.alas * s.tinggi) / 2
}
func (p persegi) LuasPersegi() int {
	return p.sisi * p.sisi
}
func (pp persegiPanjang) LuasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

// soal 3
// soal 4
type movie struct {
	title    string
	duration int
	genre    string
	year     int
}

var DataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int) {
	NewMovie := movie{title: title, duration: duration, genre: genre, year: year}
	DataFilm = append(DataFilm, NewMovie)
}
func main() {
	// soal 1
	fmt.Println("===== Soal 1 =====")
	fruits := []fruit{
		{
			name:        "Nanas",
			warna:       "Kuning",
			ada_bijinya: false,
			harga:       9000,
		},
		{
			name:        "Jeruk",
			warna:       "Oranye",
			ada_bijinya: true,
			harga:       8000,
		},
		{
			name:        "Semangka",
			warna:       "Hijau dan Merah",
			ada_bijinya: true,
			harga:       10000,
		},
		{
			name:        "Pisang",
			warna:       "Kuning",
			ada_bijinya: false,
			harga:       5000,
		},
	}

	fmt.Println(fruits)

	// soal 2
	fmt.Println("===== Soal 2 =====")
	var segitiga = segitiga{20, 10}
	fmt.Printf("luas Persegi %d \n", segitiga.LuasSegitiga())
	var persegi = persegi{20}
	fmt.Printf("luas persegi %d \n", persegi.LuasPersegi())
	var persegiPanjang = persegiPanjang{20, 10}
	fmt.Printf("luas Persegi Panjang %d \n", persegiPanjang.LuasPersegiPanjang())
	// soal 3
	fmt.Println("===== Soal 3 =====")

	// soal 4
	fmt.Println("===== Soal 4 =====")
	tambahDataFilm("LOTR", 120, "action", 1999)
	tambahDataFilm("avenger", 120, "action", 2019)
	tambahDataFilm("spiderman", 120, "action", 2004)
	tambahDataFilm("juon", 120, "horror", 2004)
	fmt.Println(DataFilm)

}
