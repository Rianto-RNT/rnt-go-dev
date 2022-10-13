//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Judul string
type Nama string

type AuditPeminjaman struct {
	checkOut time.Time
	checkIn  time.Time
}

type Anggota struct {
	nama      Nama
	semuaBuku map[Judul]AuditPeminjaman
}

type MasukkanDataBuku struct {
	jumlahKeseluruhan int
	dipinjamkan       int
}

type Perpustakaan struct {
	semuaAnggota map[Nama]Anggota
	semuaBuku    map[Judul]MasukkanDataBuku
}

func cetakAuditAnggota(anggota *Anggota) {
	for judul, audit := range anggota.semuaBuku {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[Belum dikembalikan]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(anggota.nama, ":", judul, ":", audit.checkOut.String(), "melalui", returnTime)
	}
}

func cetakAuditAuditAnggota(perpustakaan *Perpustakaan) {
	for _, anggota := range perpustakaan.semuaAnggota {
		cetakAuditAnggota(&anggota)
	}
}

func cetakBukuBukuPerpustakaan(perpustakaan *Perpustakaan) {
	fmt.Println()
	for judul, buku := range perpustakaan.semuaBuku {
		fmt.Println(judul, "/ total:", buku.jumlahKeseluruhan, "/ lended:", buku.dipinjamkan)

	}
	fmt.Println()
}

func kembalikanBuku(perpustakaan *Perpustakaan, judul Judul, anggota *Anggota) bool {
	buku, temukan := perpustakaan.semuaBuku[judul]
	if !temukan {
		fmt.Println("Buku bukan bagian dari perpustakaan")
		return false
	}

	audit, temukan := anggota.semuaBuku[judul]
	if !temukan {
		fmt.Println("Anggota tidak memeriksa buku ini")
		return true
	}

	buku.dipinjamkan -= 1
	perpustakaan.semuaBuku[judul] = buku

	audit.checkIn = time.Now()
	anggota.semuaBuku[judul] = audit
	return true
}

func checkoutBuku(perpustakaan *Perpustakaan, judul Judul, anggota *Anggota) bool {
	buku, temukan := perpustakaan.semuaBuku[judul]
	if !temukan {
		fmt.Println("Buku bukan bagian dari perpustakaan")
		return false
	}

	buku.dipinjamkan += 1
	perpustakaan.semuaBuku[judul] = buku

	anggota.semuaBuku[judul] = AuditPeminjaman{checkOut: time.Now()}

	return false
}

func main() {
	perpustakaan := Perpustakaan{
		semuaBuku:    make(map[Judul]MasukkanDataBuku),
		semuaAnggota: make(map[Nama]Anggota),
	}

	perpustakaan.semuaBuku["Webapp in RNT-Go"] = MasukkanDataBuku{
		jumlahKeseluruhan: 4,
		dipinjamkan:       0,
	}
	perpustakaan.semuaBuku["The Little RNT-Go Book"] = MasukkanDataBuku{
		jumlahKeseluruhan: 3,
		dipinjamkan:       0,
	}
	perpustakaan.semuaBuku["Let's fireup RNT-Go"] = MasukkanDataBuku{
		jumlahKeseluruhan: 2,
		dipinjamkan:       0,
	}
	perpustakaan.semuaBuku["RNT-Go Academy"] = MasukkanDataBuku{
		jumlahKeseluruhan: 1,
		dipinjamkan:       0,
	}

	perpustakaan.semuaAnggota["Ryan"] = Anggota{"Ryan", make(map[Judul]AuditPeminjaman)}
	perpustakaan.semuaAnggota["Morrison"] = Anggota{"Morrison", make(map[Judul]AuditPeminjaman)}
	perpustakaan.semuaAnggota["Lalaland"] = Anggota{"Lalaland", make(map[Judul]AuditPeminjaman)}

	fmt.Println("\nInitial:")

	cetakBukuBukuPerpustakaan(&perpustakaan)
	cetakAuditAuditAnggota(&perpustakaan)

	member := perpustakaan.semuaAnggota["Ryan"]
	checkedOut := checkoutBuku(&perpustakaan, "RNT-Go Academy", &member)
	fmt.Println("\nPeriksa buku:")
	if checkedOut {
		cetakBukuBukuPerpustakaan(&perpustakaan)
		cetakAuditAuditAnggota(&perpustakaan)
	}

	returned := kembalikanBuku(&perpustakaan, "RNT-Go Academy", &member)
	fmt.Println("\nCheck in a book:")
	if returned {
		cetakBukuBukuPerpustakaan(&perpustakaan)
		cetakAuditAuditAnggota(&perpustakaan)
	}
}
