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
				Title:    "Understanding Diabetes: Types and Causes",
				Category: domain.ArticleCategoryMedical,
				Body: `
## Understanding Diabetes: Types and Causes  

Diabetes is a chronic health condition that affects the way your body converts food into energy. This occurs when the body either does not produce enough insulin or cannot effectively use the insulin it produces.  

### Types of Diabetes  

1. **Type 1 Diabetes**  
  Type 1 diabetes is an autoimmune condition where the body attacks insulin-producing cells in the pancreas. It typically develops in childhood or adolescence and requires insulin therapy for management.  

2. **Type 2 Diabetes**  
  Type 2 diabetes is more common and usually develops in adulthood. It occurs when the body becomes resistant to insulin or when the pancreas does not produce enough insulin.  

3. **Gestational Diabetes**  
  Gestational diabetes occurs during pregnancy and can pose risks to both mother and baby. While it often resolves after childbirth, it increases the risk of developing type 2 diabetes later in life.  

### Causes and Risk Factors  

- **Genetic Factors**: A family history of diabetes can increase your risk.  
- **Lifestyle Choices**: Poor diet, physical inactivity, and obesity are major contributors.  
- **Age**: The risk of type 2 diabetes increases with age, especially after 45.  
- **Ethnicity**: Some ethnic groups are more prone to diabetes, such as African Americans, Hispanics, and Asians.  

Early diagnosis and lifestyle modifications are essential to managing diabetes and reducing complications.  
				`,
				Date: time.Now().AddDate(0, 0, -1),
			},
			{
				ID:       2,
				Author:   "Gede Indra Adi Brata",
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/nutrition_diabetes.jpg",
				Title:    " The Role of Nutrition in Diabetes Management",
				Category: domain.ArticleCategorySport,
				Body: `
## The Role of Nutrition in Diabetes Management  

Managing diabetes effectively requires a balanced and well-planned diet. Nutrition plays a crucial role in regulating blood sugar levels and preventing complications.  

### Key Principles of a Diabetic Diet  

1. **Focus on Low-Glycemic Index Foods**  
  Foods with a low glycemic index, such as whole grains, legumes, and non-starchy vegetables, help maintain stable blood sugar levels.  

2. **Include Healthy Fats**  
  Incorporate sources of healthy fats like avocados, nuts, seeds, and olive oil to support overall health without spiking blood sugar.  

3. **Opt for Lean Proteins**  
  Lean proteins such as chicken, fish, tofu, and legumes help control hunger and stabilize glucose levels.  

4. **Monitor Carbohydrate Intake**  
  Carbohydrates have the most significant impact on blood sugar. Choose complex carbs and monitor portion sizes.  

### Example Meal Plan  

- **Breakfast**: Oatmeal with fresh berries and a sprinkle of chia seeds.  
- **Lunch**: Grilled chicken salad with mixed greens, cherry tomatoes, and olive oil dressing.  
- **Dinner**: Steamed fish with quinoa and roasted vegetables.  

Balanced nutrition, combined with regular exercise, forms the foundation of effective diabetes management.  
				`,
				Date: time.Now().AddDate(0, 0, -2),
			},
			{
				ID:       3,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/exercise_diabetes.jpg",
				Author:   "Gede Indra Adi Brata",
				Title:    "Exercise and Diabetes: What You Need to Know",
				Category: domain.ArticleCategorySport,
				Body: `
## Exercise and Diabetes: What You Need to Know  

Regular physical activity is essential for managing diabetes and improving overall health. Exercise helps regulate blood sugar, enhance insulin sensitivity, and reduce the risk of complications.  

### Benefits of Exercise for Diabetics  

1. **Improved Insulin Sensitivity**  
  Exercise makes your cells more responsive to insulin, lowering blood sugar levels.  

2. **Weight Management**  
  Regular workouts help maintain a healthy weight, a critical factor in managing type 2 diabetes.  

3. **Cardiovascular Health**  
  Physical activity reduces the risk of heart disease, which is a common complication of diabetes.  

### Recommended Exercises  

- **Aerobic Exercise**: Walking, cycling, or swimming for at least 150 minutes a week.  
- **Strength Training**: Lifting weights or resistance band exercises 2-3 times a week.  
- **Flexibility and Balance**: Yoga or Tai Chi to improve balance and reduce stress.  

### Tips for Safe Exercise  

- Monitor blood sugar before and after exercise.  
- Stay hydrated and carry a quick source of sugar in case of hypoglycemia.  
- Wear comfortable, supportive shoes to prevent foot injuries.  

Incorporating exercise into your routine can significantly enhance your quality of life as a diabetic.   
				`,
				Date: time.Now().AddDate(0, 0, -3),
			},
			{
				ID:       4,
				Author:   "Gede Indra Adi Brata",
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/sympton_diabetes.jpg",
				Title:    "Recognizing Early Signs of Diabetes",
				Category: domain.ArticleCategoryMedical,
				Body: `
## Recognizing Early Signs of Diabetes  

Early detection of diabetes can prevent severe complications and allow for timely management. Understanding the warning signs is crucial for seeking prompt medical attention.  

### Common Symptoms of Diabetes  

1. **Increased Thirst and Frequent Urination**  
  High blood sugar levels lead to excessive thirst and frequent urination as the body tries to flush out excess glucose.  

2. **Fatigue**  
  When glucose cannot enter cells for energy, it results in persistent fatigue.  

3. **Unexplained Weight Loss**  
  Sudden weight loss can occur due to the body breaking down fat and muscle for energy.  

4. **Blurred Vision**  
  High blood sugar levels can cause swelling in the lens of the eye, leading to blurred vision.  

5. **Slow Healing Wounds**  
  Diabetes impairs circulation and affects the immune system, causing slow healing of cuts and wounds.  

### When to Seek Help  

If you experience any of these symptoms, consult a healthcare professional immediately. Early intervention can significantly reduce the risk of complications.  
				`,
				Date: time.Now().AddDate(0, 0, -7),
			},
			{
				ID:       5,
				Author:   "Gede Indra Adi Brata",
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/myth_diabetes.jpg",
				Title:    "Recognizing Early Signs of Diabetes",
				Category: domain.ArticleCategoryKnowledge,
				Body: `
## Debunking Common Myths About Diabetes  

Misconceptions about diabetes can lead to fear and misinformation. Let's debunk some of the most common myths.  

### Myth 1: Diabetes is Caused by Eating Too Much Sugar  
**Truth**: While excessive sugar consumption can contribute to weight gain, which increases the risk of type 2 diabetes, it is not the sole cause. Genetics and lifestyle also play a significant role.  

### Myth 2: Only Overweight People Get Diabetes  
**Truth**: While obesity is a risk factor for type 2 diabetes, anyone, regardless of weight, can develop the condition.  

### Myth 3: People with Diabetes Cannot Eat Carbohydrates  
**Truth**: Diabetics can enjoy carbohydrates in moderation, focusing on complex carbs and portion control.  

### Myth 4: Insulin is a Sign of Failing Health  
**Truth**: Insulin therapy is a vital treatment for some people with diabetes and helps manage blood sugar levels effectively.  

### Myth 5: Diabetes is Not a Serious Disease  
**Truth**: If left unmanaged, diabetes can lead to severe complications like heart disease, kidney failure, and nerve damage.  

Understanding the facts about diabetes empowers individuals to make informed decisions about their health.  
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
