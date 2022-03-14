package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

const (
	EntityCity         = "City"
	EntityCustomer     = "Customer"
	EntityProduct      = "Product"
	EntityCategory     = "Category"
	EntityInvoiceOrder = "Invoice Order"
	EntityOrderDetail  = "Order Detail"
	Entity             = "entity"
)

type Customer struct {
	DID     string    `json:"did"`
	XID     uuid.UUID `json:"xid"`
	Entity  string    `json:"entity"`
	Name    string    `json:"name" faker:"name"`
	Address City      `json:"address"`
}
type City struct {
	DID    string    `json:"did"`
	XID    uuid.UUID `json:"xid"`
	Entity string    `json:"entity"`
	Name   string    `json:"name"`
}

type Product struct {
	DID                  string          `json:"did"`
	XID                  uuid.UUID       `json:"xid"`
	Entity               string          `json:"entity"`
	Name                 string          `json:"name" faker:"name"`
	Price                decimal.Decimal `json:"price"`
	CommissionPercentage int             `json:"commission_percentage"`
	CommissionAmount     decimal.Decimal `json:"commission_amount"`
	AddressOrigin        City            `json:"address_origin"`
}

type Category struct {
	DID    string    `json:"did"`
	XID    uuid.UUID `json:"xid"`
	Entity string    `json:"entity"`
	Name   string    `json:"name" faker:"name"`
}

type MutationResult struct {
	Data struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Queries interface{} `json:"queries"`
		Uids    struct {
			Entity string `json:"entity"`
		} `json:"uids"`
	} `json:"data"`
}

var (
	CityMap     map[string]City     // key A1 - A34
	CustomerMap map[string]Customer // key Customer C1 - C1000000
	ProductMap  map[string]Product  // key Product P1 - P15000
	CategoryMap map[string]Category // key Product G1 - G15000

	DgraphHost = "http://localhost:8080"
)

func init() {
	checkpoint := time.Now()
	log.Printf("Generate City ")
	CityMap = GenerateCityMap()
	GenerateRDFCity(CityMap)
	log.Printf("Time Spent %s \n", time.Since(checkpoint))

	checkpoint = time.Now()
	log.Printf("Generate Category ")
	CategoryMap = GenerateCategoryMap()
	GenerateRDFCategory(CategoryMap)
	log.Printf("Time Spent %s \n", time.Since(checkpoint))

	checkpoint = time.Now()
	log.Printf("Generate Customer ")
	CustomerMap = GenerateCustomerMap(10000)
	GenerateRDFCustomer(CustomerMap)
	log.Printf("Time Spent %s \n", time.Since(checkpoint))

	checkpoint = time.Now()
	log.Printf("Generate Product ")
	ProductMap = GenerateProductMap(1000)
	GenerateRDFProduct(ProductMap)
	log.Printf("Time Spent %s \n", time.Since(checkpoint))

	checkpoint = time.Now()
	log.Printf("Generate Invoice ")
	GenerateRDFInvoice()
	log.Printf("Time Spent %s \n", time.Since(checkpoint))
}

func main() {
}

func GenerateCityMap() (newCityMap map[string]City) {
	provinceNames := []string{
		"Banda Aceh",
		"Medan",
		"Palembang",
		"Padang",
		"Bengkulu",
		"Pekanbaru",
		"Tanjung Pinang",
		"Jambi",
		"Bandar Lampung",
		"Pangkal Pinang",
		"Pontianak",
		"Samarinda",
		"Banjarmasin",
		"Palangkaraya",
		"Tanjung Selor",
		"Serang",
		"Jakarta",
		"Bandung",
		"Semarang",
		"Yogyakarta",
		"Surabaya",
		"Denpasar",
		"Kupang",
		"Mataram",
		"Gorontalo",
		"Mamuju",
		"Palu",
		"Manado",
		"Kendari",
		"Makassar",
		"Ternate",
		"Ambon",
		"Manokwari",
		"Jayapura",
	}

	newCityMap = make(map[string]City)
	for i := range provinceNames {
		newCityMap[fmt.Sprintf("A%d", i+1)] = NewCity(provinceNames[i])
	}

	return
}

func NewCity(name string) (newAddress City) {
	newAddress.XID, _ = uuid.NewV4()
	newAddress.Name = name
	newAddress.Entity = EntityCity
	return
}

func GenerateCategoryMap() (newCategoryMap map[string]Category) {
	categories := []string{
		"Pulsa dan Tagihan",
		"Prabayar",
		"Pulsa Seluler",
		"Token PLN",
		"Paket Data",
		"Paket Telefon & SMS",
		"Voucher Internet",
		"E - Money",
		"Voucher Game",
		"Voucher Digital",
		"Voucher Pulsa selular",
		"Pascabayar",
		"Tagihan PLN",
		"Rumah Tangga",
		"Kamar Mandi",
		"Gayung",
		"Cermin Kamar Mandi",
		"Dispenser Odol",
		"Gantungan Handuk",
		"Keset Anti Slip",
		"Rak Toilet",
		"Tempat Sikat Gigi",
		"Handuk Mandi",
		"Tempat Sabun",
		"Kamar Mandi Lainnya",
		"Kamar Tidur",
		"Bantal",
		"Kasur",
		"Matras",
		"Selimut",
		"Sprei dan Bed Cover",
		"Kamar Tidur Lainnya",
		"Ruang Tamu & Keluarga",
		"Karpet & Tikar",
		"Bantal Sofa",
		"Cover Sofa",
		"Gorden",
		"Sarung Bantal Sofa",
		"Ruang Tamu & Keluarga Lainnya",
		"Dekorasi",
		"Cover Kursi",
		"Hiasan Dinding",
		"Jam Meja",
		"Keset",
		"Lilin",
		"Lilin Aroma Terapi",
		"Lukisan",
		"Stiker Kaca",
		"Tanaman Artifical",
		"Taplak Meja",
		"Vas Bunga",
		"Wall Sticker",
		"Dekorasi Lainnya",
		"Furniture",
		"Cermin Badan",
		"Lemari Pakaian",
		"Meja Makan",
		"Meja Rias",
		"Meja Tamu",
		"Meja TV",
		"Pengaman Furniture",
		"Rak",
		"Sofa",
		"Furniture Lainnya",
		"Kursi",
		"Alat Kebersihan",
		"Alat-Alat Pel",
		"Asbak",
		"Ember & Baskom",
		"Kain Lap",
		"Kantong Sampah",
		"Kemoceng",
		"Alat Kebersihan Lainnya",
		"Pengki",
		"Sapu",
		"Sapu Lidi",
		"Sarung Tangan Karet",
		"Selang Air",
		"Sikat",
		"Tempat Sampah",
		"Kebutuhan Rumah",
		"Baterai",
		"Gembok",
		"Humidifier",
		"Payung",
		"Penahan Pintu",
		"Kebutuhan Rumah Lainnya",
		"Laundry",
		"Cover Mesin Cuci",
		"Gantungan Baju",
		"Jaring Pakaian Mesin Cuci",
		"Jemuran Baju",
		"Jepit Jemuran",
		"Laundry Bag",
		"Papan Cuci Baju",
		"Roll Pembersih Pakaian",
		"Tempat Penyimpanan",
		"Botol",
		"Keranjang",
		"Kotak",
		"Laci",
		"Tempat Penyimpanan Lainnya",
		"Stand Hanger",
		"Storage Box Multifungsi",
		"Tempat Pakaian",
		"Tempat Perhiasan & Aksesoris",
		"Tempat Sepatu & Sandal",
		"Tempat Tas",
		"Tempat Tissue",
		"Taman",
		"Pot",
		"Tanaman",
		"Media Tanam",
		"Pupuk",
		"Hiasan Taman",
		"Dapur",
		"Aksesoris Dapur",
		"Alat Pemotong Serbaguna",
		"Capit Makanan",
		"Celemek",
		"Chopper",
		"Grinder",
		"Gunting Dapur",
		"Korek Kompor",
		"Parutan",
		"Peeler",
		"Pelindung Tangan",
		"Pengasah Pisau",
		"Pisau Dapur",
		"Pisau Set",
		"Talenan",
		"Bekal",
		"Botol Minum",
		"Cetakan Bento",
		"Kotak Makan",
		"Lunch Box Set",
		"Partisi Bento",
		"Rantang",
		"Tas Bekal",
		"Termos Air",
		"Penyimpanan Makanan",
		"Aluminium Foil",
		"Box Telur",
		"Cooler Box",
		"Food Display",
		"Food Warmer",
		"Ice - Rice Bucket",
		"Plastik Klip",
		"Sealer Makanan",
		"Tempat Buah & Sayur",
		"Tempat Bumbu",
		"Tempat Saos & Kecap",
		"Toples Makanan",
		"Peralatan Baking",
		"Cetakan Kue",
		"Kocokan Telur",
		"Kuas Kue",
		"Pisau Kue",
		"Tatakan Kue",
		"Peralatan Dapur",
		"Dispenser Air",
		"Pompa Galon",
		"Rak Dapur",
		"Rak Piring",
		"Regulator & Penghemat Gas",
		"Sarung Galon",
		"Sarung Kulkas",
		"Timbangan Dapur",
		"Peralatan Makan & Minum",
		"Cangkir",
		"Centong Nasi",
		"Gelas & Mug",
		"Mangkok Makan",
		"Nampan",
		"Peralatan Makan Set",
		"Peralatan Minum Set",
		"Piring & Mangkok Saji",
		"Piring Makan",
		"Pitcher Minuman",
		"Sedotan",
		"Sendok & Garpu Dessert",
		"Sendok & Garpu Makan",
		"Sendok Bebek",
		"Sendok Sayur & Kuah",
		"Sumpit Makan",
		"Tatakan Gelas & Piring",
		"Tempat Sendok & Garpu",
		"Tudung Saji",
		"Tutup Gelas & Piring",
		"Peralatan Masak",
		"Food Processor",
		"Cetakan Es, Puding, Coklat",
		"Cobek",
		"Deep Fryer",
		"Gelas Takar",
		"Gilingan Daging",
		"Griller",
		"Kompor",
		"Panci",
		"Presto",
		"Saringan Masak",
		"Sendok Takar",
		"Spatula & Sutil",
		"Steamer",
		"Teko & Pemanas Air",
		"Wajan",
		"Perlengkapan Cuci Piring",
		"Dish Dryer",
		"Saringan Bak Cuci Piring",
		"Sikat Cuci Botol",
		"Sponge Cuci Piring",
		"Fashion Muslim",
		"Aksesoris Muslim",
		"Bros Hijab",
		"Headpiece Hijab",
		"Kaos Kaki Wudhu",
		"Klip Turki",
		"Peniti Hijab",
		"Atasan Muslim Wanita",
		"Blouse Muslim Wanita",
		"Manset Muslim Wanita",
		"Setelan Syari Wanita",
		"Tunik Muslim",
		"Baju Renang Muslim",
		"Pakaian Renang Muslim",
		"Bawahan Muslim Wanita",
		"Celana Muslim",
		"Legging Wudhu",
		"Palazzo",
		"Rok Muslim",
		"Dress Muslim Wanita",
		"Dress Abaya",
		"Gamis Wanita",
		"Jumpsuit Muslim",
		"Kaftan",
		"Jilbab",
		"Cadar",
		"Ciput",
		"Jilbab Instan",
		"Jilbab Segi Empat",
		"Jilbab Olahraga",
		"Jilbab Khimar",
		"Jilbab Pashmina",
		"Jilbab Turban",
		"Outerwear Muslim Wanita",
		"Cape Muslim",
		"Cardigan Muslim",
		"Coat Muslim",
		"Vest Muslim",
		"Outer Wanita Muslim",
		"Muslim Pria",
		"Baju Koko Pria",
		"Baju Koko Set Pria",
		"Celana Sirwal",
		"Pakaian Gamis Pria",
		"Kain",
		"Kafan",
		"Fashion Dewasa Muslim",
		"Seragam Group Wanita",
		"Seragam Couple",
		"Seragam Keluarga Sarimbit",
		"Al-Quran & Buku Islami",
		"Hard Copy",
		"Al-Quran",
		"Buku Islam",
		"e-Book",
		"Al-Quran",
		"Fashion Anak & Bayi",
		"Fashion Bayi",
		"Pakaian Bayi",
		"Aksesoris Bayi",
		"Fashion Anak Laki-laki",
		"Atasan Anak Laki-Laki",
		"Celana Anak Laki-Laki",
		"Tas Anak Laki-Laki",
		"Sepatu dan Sandal Anak Laki-Laki",
		"Aksesoris Anak Laki-Laki",
		"Setelan Set Anak Laki-Laki",
		"Baju Tidur Anak Laki-Laki",
		"Fashion Anak Perempuan",
		"Bawahan Anak Perempuan",
		"Tas Anak Perempuan",
		"Sepatu Anak Perempuan",
		"Aksesoris Anak Perempuan",
		"Setelan Set Anak Perempuan",
		"Baju Tidur Anak Perempuan",
		"Baju Anak Perempuan",
		"Seragam Sekolah",
		"Atasan Seragam",
		"Bawahan Seragam",
		"Aksesoris Seragam Sekolah",
		"Pakaian Muslim Anak",
		"Hijab Anak",
		"Baju Koko Anak",
		"Busana Muslim Family Set",
		"Busana Muslim Set Anak",
		"Pakaian Gamis Anak",
		"Rok Muslim Anak",
		"Fashion Dewasa",
		"Fashion Pria",
		"Kaos Dan Kemeja Pria",
		"Jaket dan Sweater Pria",
		"Celana Pria",
		"Tas Pria",
		"Sepatu Pria",
		"Aksesoris Pria",
		"Pakaian Dalam Pria",
		"Fashion Wanita",
		"Atasan Wanita",
		"Outer Wanita",
		"Bawahan Wanita",
		"Tas Wanita",
		"Sepatu Wanita",
		"Aksesoris Wanita",
		"Kain",
		"Baju Tidur Wanita",
		"Pakaian Dalam Wanita",
		"Fashion Ibu Hamil",
		"Atasan Bumil",
		"Bawahan Bumil",
		"Seragam",
		"Seragam Group Pria",
		"Makanan",
		"Makanan Segar",
		"Beras",
		"Buah",
		"Sayur",
		"Umbi",
		"Daging",
		"Unggas",
		"Telur",
		"Ikan & Hasil Laut",
		"Bumbu Dapur",
		"Penyedap Makanan",
		"Bumbu masak instan",
		"Rempah-rempah",
		"Saus",
		"Paket Sembako",
		"Minyak Goreng",
		"Gula, Garam & Merica",
		"Makanan Siap Saji",
		"Makanan Kaleng",
		"Makanan Cup",
		"Makanan Olahan Jadi",
		"Makanan Ringan",
		"Cokelat",
		"Permen",
		"Snack",
		"Selai",
		"Kacang & Keripik",
		"Kue & Cake",
		"Kue Bolu",
		"Roti Gandum",
		"Kue Kering",
		"Sembako",
		"Makanan Hewan",
		"Pakan Ternak",
		"Bahan Kue",
		"Bahan Puding & Agar - Agar",
		"Baking Powder",
		"Baking Soda",
		"Coklat Bubuk",
		"Coklat Masak",
		"Perisa Makanan",
		"Pewarna Makanan",
		"Ragi",
		"Topping & Penghias Kue",
		"Tepung",
		"Makanan Beku",
		"Bakso & Daging Olahan Lainnya",
		"Camilan Beku",
		"Dessert",
		"Kentang Beku",
		"Nugget",
		"Sosis",
		"Mie & Pasta",
		"Mie Instant",
		"Produk Olahan Susu",
		"Keju",
		"Krim",
		"Mentega & Butter",
		"Susu Kental Manis",
		"Yogurt",
		"Minuman",
		"Minuman Cair",
		"Air Zam - zam",
		"Air Zam-Zam",
		"Teh",
		"Kopi",
		"Susu",
		"Soft Drink",
		"Sirup",
		"Madu",
		"Air Mineral",
		"Minuman Kesehatan",
		"Jus",
		"Minuman Bubuk",
		"Teh",
		"Kopi",
		"Susu",
		"Buah & aneka rasa",
		"Minuman Tradisional",
		"Kesehatan",
		"Obat-obatan",
		"Obat Herbal",
		"Obat Medis",
		"Suplemen & Nutrisi",
		"Pelangsing",
		"Penambah Berat Badan",
		"Lainnya",
		"Peralatan Medis",
		"Masker",
		"Sarung Tangan",
		"Alat Pelindung Diri",
		"Alkohol Medis",
		"Hand Sanitizer",
		"Kesehatan Wanita",
		"Suplemen Kewanitaan",
		"Obat Keputihan",
		"Aromatherapy",
		"Essential Oil",
		"Perlengkapan Kebersihan",
		"Deterjen Laundry",
		"Karbol",
		"Pembersih Toilet",
		"Pengharum Ruangan",
		"Pewangi Pelembut Pakaian",
		"Sabun Cuci Piring",
		"Tissue",
		"Perlengkapan Medis",
		"Termometer",
		"Tulang Otot & Sendi",
		"Minyak Pijat",
		"Vitamin & Multivitamin",
		"Sistem Kekebalan Tubuh",
		"Suplemen Vitamin Rambut",
		"Vitamin & Nutrisi",
		"Vitamin Anak",
		"Vitamin C",
		"Vitamin D",
		"Peralatan Ibadah",
		"Wanita",
		"Mukena Dewasa",
		"Anak Perempuan",
		"Mukena Anak",
		"Pria",
		"Sarung Dewasa",
		"Peci Dewasa",
		"Sorban",
		"Anak Laki-laki",
		"Sarung Anak",
		"Peci Anak",
		"Peralatan Ibadah Umum",
		"Sajadah Anak",
		"Sajadah",
		"Tasbih",
		"Rompi Sholat",
		"Perlengkapan Haji & Umroh",
		"Pakaian Ihram Pria",
		"Buku",
		"Hard Copy",
		"Teknologi & Sains",
		"Bisnis",
		"Masakan",
		"Buku Anak",
		"Novel",
		"e-Book",
		"Teknologi & Sains",
		"Bisnis",
		"Masakan",
		"Kecantikan",
		"Aksesoris Rambut",
		"Bando Bandana",
		"Ikat Rambut",
		"Jepitan Rambut",
		"Mahkota & Headpiece",
		"Brush Applicator",
		"Beauty Sponge",
		"Make Up Brush",
		"Make Up Brush Set",
		"Pembersih Brush Make Up",
		"Eyebrow Kit",
		"Pensil Alis",
		"Eyebrow Mascara",
		"Hand & Nail Art",
		"Henna",
		"Kuteks Halal",
		"Lip Color & Lip Care",
		"Lip Balm & Oil",
		"Lip Cream",
		"Lipgloss",
		"Lip Scrub",
		"Lipstik",
		"Lip Tint & Lip Stain",
		"Make up Mata",
		"Eye Liner",
		"Eye Shadow",
		"Mascara",
		"Peralatan Make Up",
		"Cermin Make Up",
		"Laci & Tempat Make Up",
		"Pinset Komedo",
		"Tas Kosmetik",
		"Make Up Wajah",
		"BB Cream",
		"Bedak Wajah",
		"Blush On",
		"CC Cream",
		"Concealer & Color Corrector",
		"Cushion",
		"Face Primer",
		"Foundation",
		"Setting Spray",
		"Masker Kecantikan",
		"Masker Bibir",
		"Masker Wajah",
		"Pembersih Make Up",
		"Kapas Wajah",
		"Make Up Remover Balm",
		"Make Up Remover Oil",
		"Micellar Water",
		"Pembersih Mata Bibir",
		"Perawatan Wajah",
		"Cleanser Wajah",
		"Face Mist",
		"Krim Mata",
		"Krim Wajah",
		"Minyak Wajah",
		"Paket Perawatan Wajah",
		"Scrub Wajah",
		"Serum Wajah & Mata",
		"Skincare Tools",
		"Sunblock Wajah",
		"Toner Wajah",
		"Penghilang Bekas Jerawat",
		"Styling Rambut Wanita",
		"Hair Dryer",
		"Sisir Rambut",
		"Stationery & Craft",
		"Stationery",
		"Kalkulator & Kamus Elektronik",
		"Kalkulator",
		"Kalkulator Ilmiah",
		"Kamus Elektronik",
		"Rumah Tangga",
		"Otomotif",
		"Motor",
		"Aksesoris Motor",
		"Helm",
		"Bike Tag",
		"Aksesoris Pengendara Motor",
		"Mobil",
		"Hiasan Mobil",
		"Pengharum Mobil",
		"Interior Mobil",
		"Perawatan Mobil",
		"Elektronik",
		"Kamera",
		"Aksesoris Kamera",
		"Tas Kamera",
		"Handphone",
		"Aksesoris Handphone",
		"Casing Handphone",
		"Android",
		"Jam",
		"Jam Digital",
		"Audio",
		"Speaker",
		"Elektronik Rumah Tangga",
		"Elektronik Dapur",
		"Setrika",
		"Vacuum Cleaner",
		"Lampu",
		"Bohlam",
		"Lampu Darurat",
		"Travel",
		"Perjalanan Wisata",
		"Perjalanan Ibadah",
		"Tiket & Perjalanan",
		"Haji & Umroh",
		"Donasi",
		"Zakat",
		"Zakat",
		"Infaq/sodaqah",
		"Infaq/sodaqah",
		"Wakaf",
		"Wakaf",
		"Qurban",
		"Qurban Hidup",
		"Qurban Kemasan",
		"Voucher",
		"Makanan & Minuman",
		"Travel",
		"Olahraga",
		"Olahraga Darat",
		"Panahan",
		"Sepeda",
		"Olahraga Air",
		"Pakaian Olahraga",
		"Pakaian Olahraga Wanita",
		"Pakaian Olahraga Pria",
		"Pakaian Olahraga Anak",
		"Sepatu Olahraga",
		"Sepatu Olahraga Wanita",
		"Sepatu Olahraga Pria",
		"Gym & Fitness",
		"Alat Fitness",
		"Hiking & Camping",
		"Peralatan Hiking & Camping",
		"Aksesoris Olahraga",
		"Aksesoris Olahraga Lainnya",
		"Member",
		"Online Course",
		"Personal Development",
		"Parenting & Relationship",
		"Pelajar (SMA)",
		"Mahasiswa",
		"Agama Islam",
		"Bahasa Arab",
		"Business",
		"Finance",
		"Entrepreneurship",
		"Communication",
		"Management",
		"Sales",
		"Strategy",
		"Voucher Diskon",
		"Keanggotaan",
		"Premium",
		"Personal Care",
		"Perawatan Gigi dan Mulut",
		"Pasta Gigi",
		"Sikat Gigi",
		"Perawatan Kuku",
		"Gunting Kuku",
		"Perawatan Kuku Lainnya",
		"Perawatan Kulit",
		"Body Butter",
		"Body Lotion",
		"Body Oil",
		"Body Scrub",
		"Deodorant",
		"Pemutih Tubuh & Ketiak",
		"Penghilang Bekas Luka",
		"Stretchmark Cream",
		"Sunblock",
		"Perawatan Rambut",
		"Conditioner",
		"Hair Tonic",
		"Masker Rambut",
		"Produk Styling Rambut",
		"Shampoo",
		"Vitamin & Serum Rambut",
		"Perawatan Tubuh",
		"Sabun Mandi",
		"Hair Wax & Pomade",
		"Produk Kewanitaan",
		"Pembalut",
		"Perawatan Tubuh Wanita",
		"Sabun Kewanitaan",
		"Perawatan Mata",
		"Cairan Pembersih Sofltens",
		"Softlens",
		"Perawatan Kaki & Tangan",
		"Foot Mask",
		"Foot Scrub",
		"Foot Spray",
		"Hand Cream",
		"Sabun Cuci Tangan",
		"Parfume",
		"Parfume Anak",
		"Parfume Pria",
		"Parfume Wanita",
		"Ibu & Bayi",
		"Kamar Bayi",
		"Boks & Matras Tidur Bayi",
		"Matras & Sprei",
		"Keamanan Bayi",
		"Kelambu",
		"Kesehatan Bayi",
		"Perawatan Kulit Bayi",
		"Mainan",
		"Mainan Bayi & Anak",
		"Mainan Boneka",
		"Mainan Edukatif",
		"Mainan Olahraga & Outdoor",
		"Mainan Peran",
		"Mainan Robot",
		"Perlengkapan Ibu Hamil",
		"Bantal Ibu Hamil",
		"Penyangga Perut",
		"Perlengkapan Makan Bayi",
		"Celemek Bayi",
		"Dot Bayi",
		"Kursi Makan Bayi",
		"Perlengkapan Botol Susu",
		"Perlengkapan Menyusui",
		"Perlengkapan Mandi Bayi",
		"Alat & Aksesoris Mandi",
		"Alat Perawatan Bayi",
		"Bak Mandi & Dudukan",
		"Jas Mandi, Handuk, & Lap Bayi",
		"Perlengkapan Travelling Bayi",
		"Aksesoris Dudukan Mobil & Motor",
		"Gendongan Bayi",
		"Tas Perlengkapan Bayi",
		"Popok & Pispot",
		"Popok Sekali Pakai",
		"Susu Formula & Makanan Bayi",
		"Makanan Bayi",
		"Stationery & Craft",
		"Kerajinan Tangan",
		"Sulam",
		"Pernak Pernik dan Hadiah",
		"Alat Tulis",
		"Correction (Tip-Ex)",
		"Textliner",
		"Jangka",
		"Paket Alat Tulis",
		"Papan Tulis & Tempel",
		"Penghapus",
		"Pensil",
		"Pulpen",
		"Rautan",
		"Papan Jalan",
		"Spidol Papan Tulis",
		"Spidol Permanen",
		"Tempat Pensil",
		"Tinta",
		"Buku Tulis",
		"Agenda & Planner",
		"Buku Keuangan",
		"Buku Tulis Sekolah",
		"Notebook & Notepad",
		"Document Organizer",
		"Binder",
		"Box File",
		"Kalender",
		"Kotak Kartu Nama",
		"Lemari File - Filling Cabinet",
		"Map",
		"Pembatas Buku",
		"Rak Kertas",
		"Stationery Stand",
		"Kertas",
		"Kertas Folio",
		"Kertas HVS",
		"Kertas Thermal",
		"Sticky Notes",
		"Tambahan",
		"Ongkir Khusus",
		"Ongkir Sembako",
		"Online Course",
	}
	newCategoryMap = make(map[string]Category)
	for i := range categories {
		newCategoryMap[fmt.Sprintf("G%d", i+1)] = NewCategory(categories[i])
	}
	return
}

func NewCategory(name string) (newCategory Category) {
	newCategory.XID, _ = uuid.NewV4()
	newCategory.Name = name
	newCategory.Entity = EntityCategory
	return
}

func GenerateProductMap(numOfProduct int) (newProductMap map[string]Product) {
	newProductMap = make(map[string]Product)

	for i := 0; i < numOfProduct; i++ {
		newProductMap[fmt.Sprintf("P%d", i+1)] = NewProduct("Product " + fmt.Sprintf("P%d", i+1))
	}

	return
}

func NewProduct(name string) (newProduct Product) {
	newProduct.XID, _ = uuid.NewV4()
	newProduct.Name = name
	newProduct.Price = decimal.NewFromInt(Random(10000, 150000, 2500))
	newProduct.CommissionPercentage = int(Random(5, 10, 1))
	newProduct.CommissionAmount = decimal.NewFromInt(newProduct.Price.IntPart() * int64(newProduct.CommissionPercentage) / 100)
	newProduct.Entity = EntityProduct
	return
}

func Random(min, max, multiplier int) int64 {
	if multiplier <= 0 {
		multiplier = 1
	} else if multiplier > max {
		multiplier = max
	}

	if min/multiplier != 0 {
		min = min / multiplier
	}

	if max/multiplier != 0 {
		max = max / multiplier
	}

	rand.Seed(time.Now().UnixNano())
	return int64((rand.Intn(max-min+1) + min) * multiplier)
}

func GenerateCustomerMap(numOfCustomer int) (newCustomerMap map[string]Customer) {
	newCustomerMap = make(map[string]Customer)

	for i := 0; i < numOfCustomer; i++ {
		var fakeName string

		newCustomerMap[fmt.Sprintf("C%d", i+1)] = NewCustomer(fakeName)
	}

	return
}

func NewCustomer(name string) (newCustomer Customer) {
	newCustomer.XID, _ = uuid.NewV4()
	newCustomer.Name = faker.Name()
	newCustomer.Entity = EntityCustomer
	return
}

func GenerateRDFCity(existingCityMap map[string]City) {
	for key, city := range existingCityMap {
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "name", city.Name))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "xid", city.XID))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "entity", city.Entity))
	}
}

func GenerateRDFCategory(existingCategoryMap map[string]Category) {
	for key, category := range existingCategoryMap {
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "name", category.Name))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "xid", category.XID))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "entity", category.Entity))
	}
}

func GenerateRDFCustomer(existingCustomerMap map[string]Customer) {
	for key, customer := range existingCustomerMap {
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "name", customer.Name))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "xid", customer.XID))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "entity", customer.Entity))

		RandomCityKey := fmt.Sprintf("A%d", Random(1, len(CityMap), 1))
		WriteToFile(fmt.Sprintf(`<%s> <%s> <%s> .`, key, "destination", RandomCityKey))
	}
}

func GenerateRDFProduct(existingProductMap map[string]Product) {
	for key, product := range existingProductMap {
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "name", product.Name))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "xid", product.XID))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "entity", product.Entity))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "price", product.Price.String()))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "commission_amount", product.CommissionAmount.String()))
		WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, key, "commission_percentage", fmt.Sprint(product.CommissionPercentage)))

		RandomCategoryKey := fmt.Sprintf("G%d", Random(1, len(CategoryMap), 1))
		WriteToFile(fmt.Sprintf(`<%s> <%s> <%s> .`, key, "category", RandomCategoryKey))

		RandomCityKey := fmt.Sprintf("A%d", Random(1, len(CityMap), 1))
		WriteToFile(fmt.Sprintf(`<%s> <%s> <%s> .`, key, "origin", RandomCityKey))
	}
}

func GenerateRDFInvoice() {
	invoiceCount := 1

	// first purchase
	for customerKey := range CustomerMap {
		// repeat maker
		repeatSeed := invoiceCount % 100
		switch true {
		case repeatSeed <= 65:
			purchaseRepeat := Random(1, 1, 1)
			for i := 0; i < int(purchaseRepeat); i++ {
				purchaseAmount := Random(1, 2, 1)
				SeedPurchase(purchaseAmount, invoiceCount, customerKey)
				invoiceCount++
			}
		case repeatSeed > 65 && repeatSeed <= 95:
			purchaseRepeat := Random(1, 3, 1)
			for i := 0; i < int(purchaseRepeat); i++ {
				purchaseAmount := Random(1, 3, 1)
				SeedPurchase(purchaseAmount, invoiceCount, customerKey)
				invoiceCount++
			}
		case repeatSeed > 95 && repeatSeed <= 99:
			purchaseRepeat := Random(2, 8, 1)
			for i := 0; i < int(purchaseRepeat); i++ {
				purchaseAmount := Random(1, 5, 1)
				SeedPurchase(purchaseAmount, invoiceCount, customerKey)
				invoiceCount++
			}
		case repeatSeed == 100:
			purchaseRepeat := Random(5, 12, 1)
			for i := 0; i < int(purchaseRepeat); i++ {
				purchaseAmount := Random(1, 7, 1)
				SeedPurchase(purchaseAmount, invoiceCount, customerKey)
				invoiceCount++
			}
		}
	}
}

func SeedPurchase(purchaseAmount int64, invoiceCount int, customerKey string) {
	invoiceKey := fmt.Sprintf("IV%d", invoiceCount)
	WriteToFile(fmt.Sprintf(`<%s> <%s> <%s> .`, customerKey, "order", invoiceKey))
	invoiceCount++

	invoiceUUID, _ := uuid.NewV4()
	orderDetailUUID, _ := uuid.NewV4()
	WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, invoiceKey, "xid", invoiceUUID))

	itemKey := fmt.Sprintf("IT%d", invoiceCount)
	WriteToFile(fmt.Sprintf(`<%s> <%s> <%s> .`, invoiceKey, "order_detail", itemKey))

	purchaseProduct := Random(1, len(ProductMap), 1)
	randomDay := Random(1, 28, 1)
	var randomDate string
	if randomDay <= 9 {
		randomDate = "0" + fmt.Sprint(randomDay)
	} else {
		randomDate = fmt.Sprint(randomDay)
	}
	purchaseDate := fmt.Sprintf("2022-02-%sT15:00:00+00:00", randomDate)
	WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, invoiceKey, "purchase_date", purchaseDate))
	WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, invoiceKey, "entity", EntityInvoiceOrder))
	WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, itemKey, "xid", orderDetailUUID))
	WriteToFile(fmt.Sprintf(`<%s> <%s> "%s" .`, itemKey, "entity", EntityOrderDetail))
	WriteToFile(fmt.Sprintf(`<%s> <%s> "%d" .`, itemKey, "order_amount", purchaseAmount))
	WriteToFile(fmt.Sprintf(`<%s> <%s> <P%d> .`, itemKey, "order_product", purchaseProduct))
}

func WriteToFile(text string) {
	f, err := os.OpenFile("dataset.rdf",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text + "\n"); err != nil {
		log.Println(err)
	}
}
