package main

import "fmt"

const NMAXGT = 1024

type mAnggota struct {
	namaAnggota  string
	peranAnggota string
}

type tim [NMAXGT]mAnggota

const NMAX int = 1024

type mStartup struct {
	namaStrtup  string
	bidangUsaha string
	produk      string
	pendanaan   int
	timAnggota  tim
	jmlhAnggota int
	jumStartup  int
	thnBerdiri  int
}

type Startup [NMAX - 1]mStartup

var A Startup
var i int
var StrtUp, bidangUsaha, produk, kode string
var pendanaan, jumStartup, thnBerdiri, opsi, pilih int
var found bool

func main() {
	var pilMenu int
	var logIn, pass string
	var akses bool = false
	fmt.Println("--- Selamat datang ---")
	for akses != true {
		fmt.Print("Username: ")
		fmt.Scan(&logIn)
		fmt.Print("Password: ")
		fmt.Scan(&pass)
		if (logIn == "A" && pass == "A") || (logIn == "Alfan" && pass == "AlfanTelkom") {
			akses = true
			for pilMenu != 8 {
				tampilMenu(&pilMenu)
			}
		} else {
			fmt.Println("Username atau password salah. Silahkan coba lagi!")
		}
	}
}

func tampilMenu(pilihan *int) {
	fmt.Println("--- Menu Utama ---")
	fmt.Println()
	fmt.Println("1. Menambah StartUp")
	fmt.Println()
	fmt.Println("2. Mengubah StartUp")
	fmt.Println()
	fmt.Println("3. Tampilkan StartUp")
	fmt.Println()
	fmt.Println("4. Menghapus StartUp")
	fmt.Println()
	fmt.Println("5. Menambah anggota")
	fmt.Println()
	fmt.Println("6. Tampilkan anggota")
	fmt.Println()
	fmt.Println("7. Pencarian Startup")
	fmt.Println()
	fmt.Println("8. Keluar")

	fmt.Print("Pilih opsi(1-6): ")
	fmt.Scan(pilihan)

	switch *pilihan {
	case 1:
		tambahStartup()
	case 2:
		mengubahStartup()
	case 3:
		tampilkanStartup()
	case 4:
		menghapusStartup(&opsi, &kode)
	case 5:
		tambahAnggota()
	case 6:
		tampilkanAnggota()
	case 7:
		idx := mencariStartup(&pilih)
		if idx != -1 {
			fmt.Println("\nData startup ditemukan:")
			fmt.Println("===============================================================")
			fmt.Printf("Nama Startup   : %s\n", A[idx].namaStrtup)
			fmt.Printf("Bidang Usaha   : %s\n", A[idx].bidangUsaha)
			fmt.Printf("Produk         : %s\n", A[idx].produk)
			fmt.Printf("Pendanaan      : %d\n", A[idx].pendanaan)
			fmt.Printf("Tahun Berdiri  : %d\n", A[idx].thnBerdiri)
			fmt.Println("===============================================================")
		} else {
			fmt.Println("Data Startup tidak ditemukan")
		}
	}
}
func tambahStartup() {
	fmt.Print("Masukkan Jumlah Startup yang ingin ditambahkan: ")
	fmt.Scan(&jumStartup)
	for j := 0; j < jumStartup; j++ {
		fmt.Println("=======================================================================================")
		fmt.Println("Startup ke-", j+1)
		fmt.Println("=======================================================================================")
		fmt.Print("Masukkan nama StartUp: ")
		fmt.Scan(&StrtUp)
		A[i].namaStrtup = StrtUp
		fmt.Println()
		fmt.Print("Masukkan tahun berdiri Startup: ")
		fmt.Scan(&thnBerdiri)
		A[i].thnBerdiri = thnBerdiri
		fmt.Println()
		fmt.Print("Masukkan tipe Bidang Usaha: ")
		fmt.Scan(&bidangUsaha)
		A[i].bidangUsaha = bidangUsaha
		fmt.Println()
		fmt.Print("Masukkan nama produk perusahaan: ")
		fmt.Scan(&produk)
		A[i].produk = produk
		fmt.Println()
		fmt.Print("Masukkan jumlah pendanaan: ")
		fmt.Scan(&pendanaan)
		A[i].pendanaan = pendanaan
		i++
	}
}

func mengubahStartup() {
	var newStartup, newProduk, newBidangUsaha string
	var newPendanaan, newThnBerdiri int
	var idx int
	tampilkanStartup()
	if i == 0 {
		fmt.Println("Belum ada StartUp yang dapat diubah")
	} else {
		fmt.Print("Masukkan ID StartUp yang ingin diubah: ")
		fmt.Scan(&idx)
		if idx >= 1 && idx <= i {
			fmt.Print("Masukkan nama baru untuk Startup: ")
			fmt.Scan(&newStartup)
			A[idx-1].namaStrtup = newStartup
			fmt.Println()
			fmt.Print("Masukkan tahun berdiri perusahaan: ")
			fmt.Scan(&thnBerdiri)
			A[idx-1].thnBerdiri = newThnBerdiri
			fmt.Println()
			fmt.Print("Masukkan Bidang usaha baru yang diinginkan: ")
			fmt.Scan(&newBidangUsaha)
			A[idx-1].bidangUsaha = newBidangUsaha
			fmt.Println()
			fmt.Print("Masukkan Produk baru yang ingin ditambahkan: ")
			fmt.Scan(&newProduk)
			A[idx-1].produk = newProduk
			fmt.Println()
			fmt.Print("Masukkan total Pendanaan baru: ")
			fmt.Scan(&pendanaan)
			A[idx-1].pendanaan = newPendanaan
			fmt.Println("Data StartUp berhasil diubah")
		} else {
			fmt.Println("Nomor tidak ada di daftar!")
		}
	}
}

func tampilkanStartup() {
	fmt.Printf("\nDaftar StartUp yang telah ditambahkan:\n")
	fmt.Println("=======================================================================================")
	fmt.Printf("| %-3s | %-20s | %-15s | %-15s | %-10s | %-5s |\n",
		"ID", "Nama StartUp", "Bidang Usaha", "Produk", "Pendanaan", "Tahun")
	fmt.Println("=======================================================================================")

	for j := 0; j < i; j++ {
		fmt.Printf("| %-3d | %-20s | %-15s | %-15s | %-10d | %-5d |\n",
			j+1,
			A[j].namaStrtup,
			A[j].bidangUsaha,
			A[j].produk,
			A[j].pendanaan,
			A[j].thnBerdiri)
	}
	fmt.Println("=======================================================================================")
}

func menghapusStartup(opsi *int, kode *string) {
	var idx int
	tampilkanStartup()
	fmt.Println("=======================================================================================")
	fmt.Println("1. ID")
	fmt.Println("2. Data")
	fmt.Println("Opsi yang ingin dihapus (1/2): ")
	fmt.Scan(opsi)
	switch *opsi {
	case 1:
		if i == 0 {
			fmt.Println("Tidak ada ID yang dapat dihapus")
		} else {
			fmt.Print("Masukkan ID StartUp yang ingin dihapus: ")
			fmt.Scan(&idx)

			if idx >= 1 && idx <= i {
				for j := idx - 1; j < i-1; j++ {
					A[j] = A[j+1]
				}
				A[i-1] = mStartup{}
				i--
				fmt.Println("ID berhasil dihapus!")
			} else {
				fmt.Println("Nomor tidak valid!, harap masukkan nomer dengan benar")
			}
		}
	case 2:
		fmt.Println("Opsi delete data: ")
		fmt.Println("A. Nama Startup")
		fmt.Println("B. Tipe/Bidang Usaha")
		fmt.Println("C. Produk")
		fmt.Println("D. Dana Anggaran")
		fmt.Println("E. Keterangan tahun")

		tampilkanStartup()

		fmt.Println("Masukkan ID data yang ingin dihapus (1-5): ")
		fmt.Scan(&idx)

		fmt.Println("Opsi delete data: ")
		fmt.Println("A. Nama Startup")
		fmt.Println("B. Tipe/Bidang Usaha")
		fmt.Println("C. Produk")
		fmt.Println("D. Dana Anggaran")
		fmt.Println("E. Keterangan tahun")
		fmt.Println("Masukkan data Startup yang ingin dihapus (A-E): ")
		fmt.Scan(kode)
		if idx >= 1 && idx <= i {
			switch *kode {
			case "A":
				A[idx-1].namaStrtup = ""
				fmt.Println("Nama Startup berhasil dihapus.")
			case "B":
				A[idx-1].bidangUsaha = ""
				fmt.Println("Bidang usaha berhasil dihapus.")
			case "C":
				A[idx-1].produk = ""
				fmt.Println("Produk berhasil dihapus.")
			case "D":
				A[idx-1].pendanaan = 0
				fmt.Println("Dana Anggaran berhasil dihapus.")
			case "E":
				A[idx-1].thnBerdiri = 0
				fmt.Println("Tahun berdiri berhasil dihapus.")
			default:
				fmt.Println("Pilihan tidak valid!")
			}
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
}

func tambahAnggota() {
	var idx int
	var anggotaBaru mAnggota
	if i == 0 {
		fmt.Println("StartUp belum tersedia, anggota tidak dapat ditambahkan")
	} else {
		fmt.Print("Masukkan ID StartUp yang diinginkan: ")
		fmt.Scan(&idx)
		if idx >= 1 && idx <= i {
			if A[idx-1].jmlhAnggota < NMAXGT {
				fmt.Print("Masukkan nama anggota yang diinginkan: ")
				fmt.Scan(&anggotaBaru.namaAnggota)
				fmt.Print("Masukkan peran anggota: ")
				fmt.Scan(&anggotaBaru.peranAnggota)
				A[idx-1].timAnggota[A[idx-1].jmlhAnggota] = anggotaBaru
				A[idx-1].jmlhAnggota++
				fmt.Println("Anggota berhasil ditambahkan!")
			}
		} else {
			fmt.Println("Nomor StartUp tidak tersedia!")
		}
	}
}

func tampilkanAnggota() {
	var idx int
	if i == 0 {
		fmt.Println("Startup belum tersedia, anggota tidak dapat ditambahkan.")
	} else {
		fmt.Print("Masukkan ID Startup yang ingin dilihat anggotanya: ")
		fmt.Scan(&idx)
		if idx >= 1 && idx <= i {
			if A[idx-1].jmlhAnggota == 0 {
				fmt.Println("Belum ada anggota yang tersedia")
			} else {
				fmt.Printf("\nDaftar anggota untuk Startup '%s':\n", A[idx-1].namaStrtup)
				fmt.Println("============================================================")
				fmt.Printf("| %-3s | %-20s | %-20s |\n", "No", "Nama Anggota", "Peran")
				fmt.Println("============================================================")
				for j := 0; j < A[idx-1].jmlhAnggota; j++ {
					fmt.Printf("| %-3d | %-20s | %-20s |\n",
						j+1,
						A[idx-1].timAnggota[j].namaAnggota,
						A[idx-1].timAnggota[j].peranAnggota)
				}
				fmt.Println("============================================================")
			}
		} else {
			fmt.Println("ID Startup tidak ditemukan!")
		}
	}
}

func mencariStartup(pilihan *int) int {
	var X string
	var found int = -1
	var low, mid, high, idx int
	fmt.Println("METHODS")
	fmt.Println("1. Sequential search")
	fmt.Println("2. Binary search")
	fmt.Scan(pilihan)
	switch *pilihan {
	case 1:
		fmt.Print("Masukkan nama startup atau bidang usaha yang dicari (Sequential Search): ")
		fmt.Scan(&X)
		for idx < i && found == -1 {

			if (A[idx].namaStrtup == X) || (A[idx].bidangUsaha == X) {
				found = idx
			}
			idx++
		}
		return found
	case 2:
		fmt.Print("Masukkan nama startup atau bidang usaha yang ingin dicari (Binary Search): ")
		fmt.Scan(&X)

		low = 0
		high = i - 1

		for low <= high && found == -1 {
			mid = (low + high) / 2
			if A[mid].namaStrtup == X {
				found = mid
			} else if A[mid].namaStrtup < X {
				low = mid + 1
			} else if A[mid].namaStrtup > X {
				high = mid - 1
			}
		}

	}
	return found

}
