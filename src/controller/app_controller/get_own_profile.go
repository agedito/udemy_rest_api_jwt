package app_controller

import (
	"fmt"
	"net/http"
)

// FEATURE Implement this
func (c *AppController) GetOwnProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected")
	email, err := c.getEmailFromTokenRequest(w, r)
	fmt.Println(email, err)
}
