package main

import "fmt"

// soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return (panjang * 2) + (lebar * 2)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// soal 2
func introduce(nama string, jk string, job string, umur string) string {
	return "Pak " + nama + " adalah seorang " + job + " yang berusia " + umur + " tahun"
}

// soal3
func yourHobbies(name string, buah ...string) {
	fmt.Printf("halo nama saya %s dan buah favorit saya adalah", name)
	for _, buah := range buah {
		fmt.Printf(buah + " ")
	}
}

// soal 4

func main() {
	// soal 1

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2

	john := introduce("John", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)
	fmt.Println(john)

	// soal 3

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	yourHobbies("John", buah...)
	// soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	// tambahDataFilm("LOTR", "2 jam", "action", "1999")
	// tambahDataFilm("avenger", "2 jam", "action", "2019")
	// tambahDataFilm("spiderman", "2 jam", "action", "2004")
	// tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
