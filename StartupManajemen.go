package main

import "fmt"

const NMAX int = 10

type Startup [NMAX]string

var A Startup
var i int
var StrtUp string

func main() {
	var TM int
	i = 0

	tampilMenu(TM)
}

func tampilMenu(pilihan int) string {
	var nData, data int
	fmt.Println("--- Menu Utama ---")
	fmt.Println("1. Menambah StartUp")
	fmt.Println("2. Menambah Anggota")
	fmt.Println("3. Tampilkan StartUp")
	fmt.Println("4. Keluar")

	fmt.Println("Pilih opsi(1-4): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		return tambahStartup(StrtUp)
	case 3:
		return tampilkanStartup(&data, nData)
	}
	return "ENTER"
}
func tambahStartup(StrtUp string) string {
	var i int
	fmt.Scan(&StrtUp)
	for StrtUp != "" {
		A[i] = StrtUp
		i = i + 1
		fmt.Println("Startup ke-", i, ":", StrtUp)
		fmt.Print("Masukkan nama startup lagi, tekan enter jika selesai: ")
		fmt.Scanln(&StrtUp)
	}
	return StrtUp
}

func tampilkanStartup(A *Startup, j int) {
	for j = 0; j < i; j++ {
		fmt.Print(A[i], " ")
	}
}
