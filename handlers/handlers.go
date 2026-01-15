package handlers

import (
	"net/http"
	"slicker-cabinets/models"

	"github.com/gin-gonic/gin"
)

func RenderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"CoreServices":        models.CoreServices,
		"SpecializedServices": models.SpecializedServices,
		"Title":               "Home",
	})
}

func RenderContact(c *gin.Context) {
	serviceID := c.Query("service")
	prefillMessage := ""
	
	if serviceID != "" {
		service := models.GetServiceByID(serviceID)
		if service != nil {
			prefillMessage = "I am interested in getting a quote for " + service.Name + "."
		}
	}

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"Title":          "Contact Us",
		"PrefillMessage": prefillMessage,
	})
}

func RenderServices(c *gin.Context) {
	// Re-using index for now as services are on the main page, 
	// but this could be a dedicated page if needed. 
	// For now, let's redirect to home with a fragment or just render home.
	c.Redirect(http.StatusFound, "/#services")
}

func RenderReviews(c *gin.Context) {
	c.HTML(http.StatusOK, "reviews.html", gin.H{
		"Title":   "Reviews",
		"Reviews": models.Reviews,
	})
}

func RenderAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"Title": "About Chris",
	})
}

func RenderLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "Internal Portal",
	})
}

func HandleContactSubmit(c *gin.Context) {
	var form models.ContactForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, we would send an email here.
	// For now, we just acknowledge receipt.
	c.JSON(http.StatusOK, gin.H{
		"message": "Thank you, " + form.Name + "! Your message regarding '" + form.Message + "' has been received. We will contact you at " + form.Email + " shortly.",
	})
}

func HandleLoginSubmit(c *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var creds LoginRequest
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if creds.Username == "admin" && creds.Password == "slicker123" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": "fake-jwt-token"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
