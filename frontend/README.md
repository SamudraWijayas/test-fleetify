# Procurement Frontend

Frontend web untuk sistem procurement yang dibangun dengan HTML, CSS, dan JavaScript.

## Deskripsi

Aplikasi frontend ini menyediakan antarmuka web untuk sistem procurement, termasuk halaman login, dashboard, dan pembelian. Menggunakan Tailwind CSS untuk styling dan jQuery untuk interaksi.

## Prerequisites

- Web browser modern (Chrome)
- Backend API server berjalan (lihat README backend)

## Instalasi

1. Clone repository ini.
2. Masuk ke direktori frontend:
   ```
   cd frontend
   ```

## Setup

Tidak ada setup khusus yang diperlukan. Pastikan backend API server sudah berjalan.

## Menjalankan Aplikasi

Buka file `login.html` di web browser Anda, atau gunakan server lokal seperti:

```
python -m http.server 8000
```

Kemudian akses `http://localhost:8000/login.html`.

## Struktur File

- `login.html` - Halaman login
- `dashboard.html` - Halaman dashboard setelah login
- `purchase.html` - Halaman pembelian
- `js/api.js` - Fungsi untuk API calls

## Teknologi yang Digunakan

- HTML5
- CSS3 (Tailwind CSS)
- JavaScript (jQuery)
- SweetAlert2 (untuk notifikasi)
