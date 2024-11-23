package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func ArticleSeeder() Seeder {
	return func(db *gorm.DB) error {
		articles := []domain.Article{
			{
				ID:       1,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/type_diabetes.jpeg",
				Author:   "Gede Indra Adi Brata",
				Title:    "Memahami Diabetes: Jenis dan Penyebabnya",
				Category: domain.ArticleCategoryMedical,
				Body: `
		## Memahami Diabetes: Jenis dan Penyebabnya  
		
		Diabetes adalah kondisi kesehatan kronis yang memengaruhi cara tubuh mengubah makanan menjadi energi. Hal ini terjadi ketika tubuh tidak memproduksi cukup insulin atau tidak dapat menggunakan insulin dengan efektif.  
		
		### Jenis Diabetes  
		
		1. **Diabetes Tipe 1**  
			Diabetes tipe 1 adalah kondisi autoimun di mana tubuh menyerang sel-sel penghasil insulin di pankreas. Biasanya berkembang pada masa kanak-kanak atau remaja dan memerlukan terapi insulin untuk pengelolaannya.  
		
		2. **Diabetes Tipe 2**  
			Diabetes tipe 2 lebih umum dan biasanya berkembang pada usia dewasa. Terjadi ketika tubuh menjadi resisten terhadap insulin atau ketika pankreas tidak menghasilkan cukup insulin.  
		
		3. **Diabetes Gestasional**  
			Diabetes gestasional terjadi selama kehamilan dan dapat menimbulkan risiko bagi ibu dan bayi. Meskipun sering sembuh setelah melahirkan, kondisi ini meningkatkan risiko diabetes tipe 2 di kemudian hari.  
		
		### Penyebab dan Faktor Risiko  
		
		- **Faktor Genetik**: Riwayat keluarga dengan diabetes dapat meningkatkan risiko.  
		- **Gaya Hidup**: Pola makan buruk, kurang aktivitas fisik, dan obesitas adalah faktor utama.  
		- **Usia**: Risiko diabetes tipe 2 meningkat seiring bertambahnya usia, terutama setelah 45 tahun.  
		- **Etnisitas**: Beberapa kelompok etnis, seperti Afrika-Amerika, Hispanik, dan Asia, lebih rentan terhadap diabetes.  
		
		Diagnosis dini dan perubahan gaya hidup sangat penting untuk mengelola diabetes dan mengurangi komplikasi.  
				`,
				Date: time.Now().AddDate(0, 0, -1),
			},
			{
				ID:       2,
				Author:   "Gede Indra Adi Brata",
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/nutrition_diabetes.jpg",
				Title:    "Pentingnya Nutrisi dalam Pengelolaan Diabetes",
				Category: domain.ArticleCategorySport,
				Body: `
		## Pentingnya Nutrisi dalam Pengelolaan Diabetes  
		
		Mengelola diabetes secara efektif memerlukan pola makan yang seimbang dan terencana. Nutrisi memainkan peran penting dalam mengatur kadar gula darah dan mencegah komplikasi.  
		
		### Prinsip Utama Pola Makan untuk Diabetes  
		
		1. **Fokus pada Makanan dengan Indeks Glikemik Rendah**  
			Makanan seperti biji-bijian utuh, kacang-kacangan, dan sayuran non-tepung membantu menjaga kadar gula darah tetap stabil.  
		
		2. **Konsumsi Lemak Sehat**  
			Sertakan lemak sehat seperti alpukat, kacang-kacangan, biji-bijian, dan minyak zaitun untuk mendukung kesehatan secara keseluruhan tanpa meningkatkan gula darah.  
		
		3. **Pilih Protein Tanpa Lemak**  
			Protein tanpa lemak seperti ayam, ikan, tahu, dan kacang-kacangan membantu mengontrol rasa lapar dan menstabilkan kadar glukosa.  
		
		4. **Pantau Asupan Karbohidrat**  
			Karbohidrat memiliki dampak terbesar pada gula darah. Pilih karbohidrat kompleks dan perhatikan ukuran porsinya.  
		
		### Contoh Rencana Makan  
		
		- **Sarapan**: Oatmeal dengan buah beri segar dan taburan biji chia.  
		- **Makan Siang**: Salad ayam panggang dengan sayuran hijau, tomat ceri, dan dressing minyak zaitun.  
		- **Makan Malam**: Ikan kukus dengan quinoa dan sayuran panggang.  
		
		Nutrisi yang seimbang, dikombinasikan dengan olahraga teratur, membentuk dasar pengelolaan diabetes yang efektif.  
				`,
				Date: time.Now().AddDate(0, 0, -2),
			},
			{
				ID:       3,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/exercise_diabetes.jpg",
				Author:   "Gede Indra Adi Brata",
				Title:    "Olahraga dan Diabetes: Apa yang Harus Anda Ketahui",
				Category: domain.ArticleCategorySport,
				Body: `
		## Olahraga dan Diabetes: Apa yang Harus Anda Ketahui  
		
		Aktivitas fisik teratur sangat penting untuk mengelola diabetes dan meningkatkan kesehatan secara keseluruhan. Olahraga membantu mengatur gula darah, meningkatkan sensitivitas insulin, dan mengurangi risiko komplikasi.  
		
		### Manfaat Olahraga bagi Penderita Diabetes  
		
		1. **Peningkatan Sensitivitas Insulin**  
			Olahraga membuat sel-sel tubuh lebih responsif terhadap insulin, sehingga menurunkan kadar gula darah.  
		
		2. **Pengelolaan Berat Badan**  
			Latihan rutin membantu menjaga berat badan yang sehat, faktor penting dalam mengelola diabetes tipe 2.  
		
		3. **Kesehatan Kardiovaskular**  
			Aktivitas fisik mengurangi risiko penyakit jantung, komplikasi umum dari diabetes.  
		
		### Jenis Olahraga yang Dianjurkan  
		
		- **Aerobik**: Jalan kaki, bersepeda, atau berenang setidaknya 150 menit per minggu.  
		- **Latihan Kekuatan**: Angkat beban atau latihan dengan resistance band 2-3 kali seminggu.  
		- **Fleksibilitas dan Keseimbangan**: Yoga atau Tai Chi untuk meningkatkan keseimbangan dan mengurangi stres.  
		
		### Tips Berolahraga dengan Aman  
		
		- Pantau kadar gula darah sebelum dan sesudah berolahraga.  
		- Tetap terhidrasi dan bawa sumber gula cepat untuk mencegah hipoglikemia.  
		- Kenakan sepatu yang nyaman dan mendukung untuk mencegah cedera kaki.  
		
		Mengintegrasikan olahraga ke dalam rutinitas harian Anda dapat secara signifikan meningkatkan kualitas hidup sebagai penderita diabetes.  
				`,
				Date: time.Now().AddDate(0, 0, -3),
			},
			{
				ID:       4,
				Author:   "Gede Indra Adi Brata",
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/sympton_diabetes.jpg",
				Title:    "Mengenali Gejala Awal Diabetes",
				Category: domain.ArticleCategoryMedical,
				Body: `
		## Mengenali Gejala Awal Diabetes  
		
		Deteksi dini diabetes dapat mencegah komplikasi serius dan memungkinkan pengelolaan yang tepat waktu. Memahami tanda-tanda peringatan sangat penting untuk mencari perhatian medis segera.  
		
		### Gejala Umum Diabetes  
		
		1. **Haus Berlebih dan Sering Buang Air Kecil**  
			Kadar gula darah tinggi menyebabkan haus berlebih dan sering buang air kecil karena tubuh mencoba mengeluarkan glukosa berlebih.  
		
		2. **Kelelahan**  
			Ketika glukosa tidak dapat masuk ke sel untuk energi, hal ini menyebabkan kelelahan yang terus-menerus.  
		
		3. **Penurunan Berat Badan yang Tidak Bisa Dijelaskan**  
			Penurunan berat badan tiba-tiba dapat terjadi karena tubuh memecah lemak dan otot untuk energi.  
		
		4. **Penglihatan Kabur**  
			Kadar gula darah tinggi dapat menyebabkan pembengkakan pada lensa mata, yang menyebabkan penglihatan kabur.  
		
		5. **Luka yang Sulit Sembuh**  
			Diabetes mengganggu sirkulasi dan memengaruhi sistem kekebalan tubuh, menyebabkan penyembuhan luka yang lambat.  
		
		### Kapan Harus Mencari Bantuan  
		
		Jika Anda mengalami salah satu gejala ini, konsultasikan dengan profesional kesehatan segera. Intervensi dini dapat secara signifikan mengurangi risiko komplikasi.  
				`,
				Date: time.Now().AddDate(0, 0, -7),
			},
			{
				ID:       5,
				Author:   "Gede Indra Adi Brata",
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/myth_diabetes.jpg",
				Title:    "Mengungkap Mitos Umum Tentang Diabetes",
				Category: domain.ArticleCategoryKnowledge,
				Body: `
		## Mengungkap Mitos Umum Tentang Diabetes  
		
		Kesalahpahaman tentang diabetes dapat menyebabkan ketakutan dan informasi yang salah. Mari kita ungkap beberapa mitos paling umum.  
		
		### Mitos 1: Diabetes Disebabkan oleh Makan Terlalu Banyak Gula  
		**Fakta**: Meskipun konsumsi gula berlebih dapat menyebabkan kenaikan berat badan, yang meningkatkan risiko diabetes tipe 2, itu bukan satu-satunya penyebab. Genetika dan gaya hidup juga memainkan peran penting.  
		
		### Mitos 2: Hanya Orang Gemuk yang Bisa Terkena Diabetes  
		**Fakta**: Meskipun obesitas meningkatkan risiko, orang dengan berat badan normal juga dapat terkena diabetes tipe 1 atau tipe 2.  
		
		### Mitos 3: Diabetes Selalu Diturunkan  
		**Fakta**: Riwayat keluarga memang meningkatkan risiko, tetapi gaya hidup dan faktor lingkungan juga memengaruhi.  
		
		### Mitos 4: Penderita Diabetes Tidak Boleh Makan Karbohidrat  
		**Fakta**: Penderita diabetes dapat makan karbohidrat, tetapi harus fokus pada sumber karbohidrat yang sehat dan memantau ukuran porsi.  
		
		### Mitos 5: Saya Tidak Memiliki Gejala, Jadi Saya Tidak Punya Diabetes  
		**Fakta**: Diabetes tipe 2 sering berkembang perlahan dengan gejala minimal atau tanpa gejala. Skrining rutin penting, terutama bagi mereka yang memiliki faktor risiko.  
		
		Memahami fakta sebenarnya tentang diabetes dapat memberdayakan individu untuk membuat keputusan yang tepat tentang kesehatan mereka.  
				`,
				Date: time.Now().AddDate(0, 0, -5),
			},
		}

		if err := db.Model(&domain.Article{}).CreateInBatches(&articles, len(articles)).Error; err != nil {
			return err
		}

		return nil
	}
}
