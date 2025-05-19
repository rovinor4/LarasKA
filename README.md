

# LarasKA (Layanan Reservasi Kereta Api Sederhana)

LarasKA adalah aplikasi reservasi tiket kereta api sederhana yang dibangun pakai Golang. Dengan LarasKA, lo bisa:

- Lihat daftar stasiun dan rute kereta
- Cari rute berdasarkan stasiun asal, tujuan, tanggal, harga, dan jam keberangkatan
- Pesan tiket kereta dengan input data penumpang (nama & NIK)
- Sorting rute berdasarkan jam, harga, kapasitas, atau kode kereta

## Teknologi

- Golang
- Modul standar Go (fmt, time, bufio, dll)

## Cara Instalasi

1. Clone repo  
   `git clone https://github.com/username/LarasKA.git`
2. Masuk ke direktori proyek  
   `cd LarasKA`
3. Build aplikasi  
   `go build -o laraska`
4. Jalankan  
   `./laraska`

## Cara Pakai

1. Jalankan `./laraska`
2. Ikuti menu interaktif untuk:
   - Melihat stasiun
   - Memilih stasiun asal & tujuan
   - Memasukkan tanggal perjalanan
   - Memilih rute & jumlah penumpang
   - Input data penumpang dan generate kode tiket

## Kontribusi

PR dan issue welcome!  

## Lisensi

MIT License