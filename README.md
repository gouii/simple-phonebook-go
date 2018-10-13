# Simple Phonebook Go

## Kebutuhan
- glide `https://github.com/Masterminds/glide`
- sql-migrate `https://github.com/rubenv/sql-migrate`

## Cara Menjalankan
- clone atau download repository
- buka terminal dan jalankan `glide install`
- kemudian setting
    - copy dan rename `env.json.example` menjadi `env.json` dan sesuaikan settingannya dengan komputer anda
    - copy dan rename `dbconfig.yml.example` menjadi `dbconfig.yml` dan sesuaikan settingannya dengan komputer anda
- jalankan di terminal `sql-migrate up`
- kemudian ketik `go run main.go`