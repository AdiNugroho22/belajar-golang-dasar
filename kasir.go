package main

import "fmt"

type Item struct {
	nama  string
	harga int
	qty   int
}

func main() {
	var items []Item
	var total int

	// Menu items
	fmt.Println("=== KASIR SEDERHANA ===")
	fmt.Println("1. Nasi Goreng - Rp 15000")
	fmt.Println("2. Mie Ayam - Rp 12000") 
	fmt.Println("3. Es Teh - Rp 5000")
	fmt.Println("4. Selesai")

	for {
		var pilihan int
		fmt.Print("Pilih menu (1-4): ")
		fmt.Scan(&pilihan)

		if pilihan == 4 {
			break
		}

		var qty int
		fmt.Print("Jumlah: ")
		fmt.Scan(&qty)

		switch pilihan {
		case 1:
			items = append(items, Item{"Nasi Goreng", 15000, qty})
		case 2:
			items = append(items, Item{"Mie Ayam", 12000, qty})
		case 3:
			items = append(items, Item{"Es Teh", 5000, qty})
		}
	}

	// Hitung total
	fmt.Println("\n=== STRUK BELANJA ===")
	for _, item := range items {
		subtotal := item.harga * item.qty
		fmt.Printf("%s x%d = Rp %d\n", item.nama, item.qty, subtotal)
		total += subtotal
	}

	fmt.Printf("\nTotal: Rp %d\n", total)

	// Pembayaran
	var bayar int
	fmt.Print("Bayar: Rp ")
	fmt.Scan(&bayar)

	kembalian := bayar - total
	if kembalian >= 0 {
		fmt.Printf("Kembalian: Rp %d\n", kembalian)
		fmt.Println("Terima kasih!")
	} else {
		fmt.Println("Uang tidak cukup!")
	}
}