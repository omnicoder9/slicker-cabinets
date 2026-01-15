package models

type Service struct {
	ID          string
	Name        string
	Category    string // "Core" or "Specialized"
	Description string
	ImageURL    string
}

var CoreServices = []Service{
	{ID: "core-kitchen", Name: "Custom Kitchen Cabinetry", Category: "Core", Description: "Transform your culinary space with bespoke kitchen cabinets tailored to your style and storage needs.", ImageURL: "https://images.unsplash.com/photo-1484154218962-a1c002085d2f?w=600&h=400&fit=crop&v=2"},
	{ID: "core-vanity", Name: "Bathroom Vanities", Category: "Core", Description: "Elegant and functional bathroom vanities that maximize space and enhance your daily routine.", ImageURL: "https://images.unsplash.com/photo-1620626011761-996317b8d101?w=600&h=400&fit=crop&v=2"},
	{ID: "core-builtin", Name: "Built-in Cabinets", Category: "Core", Description: "Seamless built-in units for living rooms, offices, and libraries, adding sophistication and utility.", ImageURL: "https://images.unsplash.com/photo-1513694203232-719a280e022f?w=600&h=400&fit=crop&v=2"},
	{ID: "core-pantry", Name: "Pantry Systems", Category: "Core", Description: "Organized pantry solutions to keep your ingredients accessible and your kitchen clutter-free.", ImageURL: "https://images.unsplash.com/photo-1609347744403-2306e8a83d41?w=600&h=400&fit=crop&v=2"},
	{ID: "core-laundry", Name: "Laundry Room Cabinetry", Category: "Core", Description: "Durable and practical cabinetry designed to make laundry day efficient and orderly.", ImageURL: "https://images.unsplash.com/photo-1610557892470-55d9e80c0bce?w=600&h=400&fit=crop&v=2"},
	{ID: "core-mudroom", Name: "Mudroom Storage Systems", Category: "Core", Description: "Keep your entryway tidy with custom mudroom storage for coats, shoes, and bags.", ImageURL: "https://images.unsplash.com/photo-1556912173-3db996e7c603?w=600&h=400&fit=crop&v=2"},
	{ID: "core-closet", Name: "Closet Cabinetry", Category: "Core", Description: "Luxurious wood closets, offering a superior alternative to wire systems for your wardrobe.", ImageURL: "https://images.unsplash.com/photo-1595514020180-27f385880d18?w=600&h=400&fit=crop&v=2"},
	{ID: "core-garage", Name: "Garage Cabinets", Category: "Core", Description: "Heavy-duty wood or hybrid garage systems to organize tools and gear effectively.", ImageURL: "https://images.unsplash.com/photo-1589939705384-5185137a7f0f?w=600&h=400&fit=crop&v=2"},
}

var SpecializedServices = []Service{
	{ID: "spec-construction", Name: "Inset & Overlay Construction", Category: "Specialized", Description: "Precision-crafted inset, overlay, and full-overlay options for a distinct look.", ImageURL: "https://images.unsplash.com/photo-1533090481720-856c6e3c1fdc?w=600&h=400&fit=crop&v=2"},
	{ID: "spec-euro", Name: "Face-frame & Frameless", Category: "Specialized", Description: "Choose between traditional face-frame or modern frameless (European) styles.", ImageURL: "https://images.unsplash.com/photo-1556228453-efd6c1ff04f6?w=600&h=400&fit=crop&v=2"},
	{ID: "spec-drawers", Name: "Custom Drawer Systems", Category: "Specialized", Description: "High-quality dovetail joints and soft-close mechanisms for smooth, lasting operation.", ImageURL: "https://images.unsplash.com/photo-1599619351208-3e6c839d6828?w=600&h=400&fit=crop&v=2"},
	{ID: "spec-panel", Name: "Appliance Paneling", Category: "Specialized", Description: "Integrate your appliances seamlessly with custom panels for dishwashers and refrigerators.", ImageURL: "https://images.unsplash.com/photo-1584622050111-993a426fbf0a?w=600&h=400&fit=crop&v=2"},
	{ID: "spec-floating", Name: "Floating Cabinets", Category: "Specialized", Description: "Modern floating cabinets that create an airy, contemporary feel in any room.", ImageURL: "https://images.unsplash.com/photo-1507089947368-19c1da9775ae?w=600&h=400&fit=crop&v=2"},
	{ID: "spec-island", Name: "Island Cabinetry", Category: "Specialized", Description: "Stunning center islands that serve as the focal point of your kitchen.", ImageURL: "https://images.unsplash.com/photo-1565538810643-b5bdb714032a?w=600&h=400&fit=crop&v=2"},
	{ID: "spec-wine", Name: "Wine Storage", Category: "Specialized", Description: "Temperature-controlled and aesthetically pleasing wine storage solutions.", ImageURL: "https://images.unsplash.com/photo-1568214238374-e924a56012e3?w=600&h=400&fit=crop&v=2"},
	{ID: "spec-display", Name: "Display Cabinets", Category: "Specialized", Description: "Showcase your treasures with glass doors and integrated lighting systems.", ImageURL: "https://images.unsplash.com/photo-1600585154526-990dced4db0d?w=600&h=400&fit=crop&v=2"},
}

func GetAllServices() []Service {
	return append(CoreServices, SpecializedServices...)
}

func GetServiceByID(id string) *Service {
	all := GetAllServices()
	for _, s := range all {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

type ContactForm struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required"`
}

type Review struct {
	Author string
	Text   string
	Rating int
}

var Reviews = []Review{
	{Author: "Jane Doe", Text: "Chris did an amazing job on our kitchen! The craftsmanship is top-notch.", Rating: 5},
	{Author: "John Smith", Text: "Slicker Cabinets transformed our garage. Highly recommended!", Rating: 5},
	{Author: "Emily Davis", Text: "Professional, timely, and beautiful work on our custom built-ins.", Rating: 5},
}
