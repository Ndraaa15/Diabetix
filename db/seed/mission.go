package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func MissionSeeder() Seeder {
	return func(db *gorm.DB) error {
		missions := []domain.Mission{
			{
				ID:     1,
				Image:  "https://example.com/images/yoga.png",
				Exp:    150,
				Calory: 200,
				Title:  "Yoga Pagi untuk Keseimbangan Tubuh",
				Body: `
	**Yoga pagi** dapat membantu menyeimbangkan tubuh dan pikiran sebelum memulai hari. Dengan melakukan yoga di pagi hari, kita dapat meningkatkan kelenturan tubuh, menenangkan pikiran, serta mengurangi tingkat stres. Gerakan-gerakan sederhana dalam yoga memungkinkan tubuh untuk bergerak secara alami dan membantu mengurangi kekakuan otot.
	
	Meluangkan waktu beberapa menit setiap pagi untuk yoga bisa menjadi rutinitas yang menenangkan dan memberi energi. Gunakan pernapasan yang dalam untuk menyeimbangkan emosi, dan rasakan tubuh menjadi lebih ringan setelah melakukannya.

	Benefit:
	- **Meningkatkan fleksibilitas tubuh.**
	- **Mengurangi stres dan meningkatkan fokus.**
	- **Meningkatkan kesehatan mental dan fisik.**
				`,
				Category:  "Hard",
				Duration:  30,
				CreatedAt: time.Now(),
				UpdateAt:  time.Now(),
			},
			{
				ID:     2,
				Image:  "https://example.com/images/smoothie.png",
				Exp:    100,
				Calory: 120,
				Title:  "Membuat Smoothie Sehat",
				Body: `
	Smoothie adalah cara yang praktis dan lezat untuk mendapatkan berbagai nutrisi dari buah dan sayur. Dalam misi ini, Anda akan belajar membuat smoothie yang kaya serat dan vitamin. Cobalah kombinasi buah seperti pisang, bayam, dan alpukat untuk mendapatkan tekstur yang lembut dan rasa yang menyegarkan.
	
	Smoothie ini bisa menjadi sarapan atau camilan yang menyehatkan. Dengan menambahkan biji chia atau oat, Anda bisa meningkatkan kadar serat, yang baik untuk pencernaan.
	
	Benefit:
	- **Memberikan asupan serat dan vitamin yang tinggi.**
	- **Menjaga kesehatan pencernaan.**
	- **Memberi energi yang tahan lama.**
	
	`,
				Category:  "Easy",
				Duration:  10,
				CreatedAt: time.Now(),
				UpdateAt:  time.Now(),
			},
			{
				ID:     3,
				Image:  "https://example.com/images/meditasi.png",
				Exp:    200,
				Calory: 50,
				Title:  "Meditasi untuk Relaksasi",
				Body: `
	Meditasi adalah teknik yang efektif untuk menenangkan pikiran dan mengurangi stres. Dalam misi ini, Anda akan dilatih untuk melakukan meditasi selama 15 menit setiap hari. Duduk dengan nyaman, fokus pada pernapasan, dan biarkan pikiran Anda beristirahat.
	
	Meditasi rutin dapat membantu mengendalikan emosi, meningkatkan fokus, dan membawa ketenangan dalam kehidupan sehari-hari. Luangkan waktu sejenak untuk diri sendiri dan nikmati momen-momen kedamaian dalam kesibukan sehari-hari.
	
	Benefit:
	- **Mengurangi stres dan kecemasan.**
	- **Meningkatkan fokus dan konsentrasi.**
	- **Membantu meningkatkan kualitas tidur.**

	`,
				Category:  "Medium",
				Duration:  15,
				CreatedAt: time.Now(),
				UpdateAt:  time.Now(),
			},
			{
				ID:     4,
				Image:  "https://example.com/images/jogging.png",
				Exp:    250,
				Calory: 300,
				Title:  "Jogging Pagi di Taman",
				Body: `
	Jogging adalah cara yang menyenangkan dan mudah untuk menjaga kebugaran. Misi ini mengajak Anda untuk melakukan jogging ringan di pagi hari selama 30 menit. Berlari di taman yang sejuk tidak hanya menyehatkan fisik tetapi juga menyegarkan pikiran.
	
	Jogging secara rutin dapat meningkatkan kesehatan jantung, membantu menjaga berat badan, dan memperkuat otot kaki. Ini juga memberikan kesempatan untuk menikmati alam dan udara segar, yang baik untuk kesehatan mental.

	Benfit: 
	- **Meningkatkan kesehatan jantung dan stamina.**
	- **Membantu menurunkan berat badan.**
	- **Memberikan energi positif untuk hari yang lebih baik.**
				`,
				Category:  "Hard",
				Duration:  30,
				CreatedAt: time.Now(),
				UpdateAt:  time.Now(),
			},
			{
				ID:     5,
				Image:  "https://example.com/images/buku.png",
				Exp:    120,
				Calory: 50,
				Title:  "Membaca Buku Inspiratif",
				Body: `
	Membaca buku inspiratif dapat membuka wawasan dan memberikan motivasi baru. Dalam misi ini, Anda diminta membaca setidaknya 10 halaman dari buku pilihan Anda yang membahas tema pengembangan diri atau motivasi.
	
	Kegiatan ini bisa menjadi waktu yang baik untuk merenung dan mendapatkan pemahaman baru yang dapat diterapkan dalam kehidupan sehari-hari. Menjadikan membaca sebagai kebiasaan juga dapat meningkatkan pengetahuan dan memperluas perspektif.
				

	Benefit:
	- **Meningkatkan pengetahuan dan wawasan.**
	- **Memotivasi diri untuk mencapai tujuan.**
	- **Meningkatkan konsentrasi dan fokus.**	
	`,
				Category:  "Medium",
				Duration:  20,
				CreatedAt: time.Now(),
				UpdateAt:  time.Now(),
			},
		}

		if err := db.Model(&domain.Mission{}).CreateInBatches(&missions, len(missions)).Error; err != nil {
			return err
		}

		return nil
	}
}
