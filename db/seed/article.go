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
				Title:    "Manfaat Pola Makan Seimbang",
				Category: domain.ArticleCategoryHealth,
				Body: `
	## Pentingnya Pola Makan Seimbang
	Mengadopsi pola makan seimbang sangat penting untuk menjaga kesehatan dan kebugaran tubuh. Dengan mengonsumsi makanan yang mengandung nutrisi lengkap, kita dapat mendukung fungsi tubuh secara optimal.
	
	### Manfaat Pola Makan Seimbang
	Pola makan yang baik mampu meningkatkan sistem kekebalan tubuh, mengurangi risiko penyakit kronis, dan meningkatkan energi harian. Sebagai contoh, buah dan sayuran kaya akan vitamin dan mineral yang dapat melindungi tubuh dari infeksi. Sedangkan, protein dari daging, ikan, dan kacang-kacangan membantu memperbaiki jaringan tubuh. 
	
	Menerapkan pola makan seimbang bukan hanya tentang menghindari makanan cepat saji, tetapi juga memilih bahan makanan segar yang diolah dengan cara yang sehat. Dengan demikian, kita dapat mencapai keseimbangan nutrisi yang dibutuhkan oleh tubuh setiap harinya.
				`,
				Likes:     150,
				Date:      time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:       2,
				Title:    "Olahraga Terbaik untuk Membangun Kekuatan",
				Category: domain.ArticleCategorySport,
				Body: `
	## Panduan Dasar Membangun Kekuatan
	Membangun kekuatan fisik sangat penting untuk menjaga kebugaran dan meningkatkan kualitas hidup. Beberapa latihan yang efektif untuk membangun kekuatan meliputi squat, push-up, dan deadlift.
	
	### Jenis Olahraga dan Manfaatnya
	Squat sangat baik untuk memperkuat otot kaki dan pinggul. Push-up membantu meningkatkan kekuatan tubuh bagian atas, terutama otot dada, bahu, dan lengan. Deadlift adalah latihan yang melibatkan seluruh tubuh dan sangat efektif untuk kekuatan otot inti serta punggung.
	
	Olahraga kekuatan perlu dilakukan secara bertahap dan konsisten. Dengan menggabungkan latihan kekuatan ini ke dalam rutinitas harian, kita dapat mencapai tubuh yang lebih kuat dan lebih sehat.
				`,
				Likes:     102,
				Date:      time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:       3,
				Title:    "Resep Smoothie Sehat dan Lezat",
				Category: domain.ArticleCategoryFood,
				Body: `
	## Smoothie Sehat untuk Menyegarkan Hari Anda
	Smoothie adalah minuman yang menyegarkan dan kaya nutrisi, sangat cocok untuk memulai hari atau dinikmati di siang hari yang panas. Kombinasi buah dan sayuran dalam smoothie dapat memberikan asupan vitamin dan mineral penting bagi tubuh.
	
	### Rekomendasi Resep
	- **Smoothie Berry Blast**: Campuran stroberi, blueberry, dan pisang memberikan rasa manis alami dan kaya antioksidan.
	- **Green Detox**: Bayam, apel hijau, dan jahe, ideal untuk detoksifikasi tubuh.
	- **Tropical Twist**: Kombinasi mangga, nanas, dan air kelapa yang menyegarkan dan kaya akan vitamin C.
	
	Cara membuatnya mudah, cukup campurkan bahan-bahan segar dalam blender, lalu sajikan. Smoothie ini tidak hanya lezat tetapi juga dapat membantu menjaga kesehatan tubuh.
				`,
				Likes:     90,
				Date:      time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		if err := db.Model(&domain.Article{}).CreateInBatches(&articles, len(articles)).Error; err != nil {
			return err
		}

		return nil
	}
}
