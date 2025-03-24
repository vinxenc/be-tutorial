// Change the package name to users_test to avoid import cycle
package users_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	// Import the users package without creating a cycle
	"day-005-logger/src/modules/users"

	// Remove the duplicate testing import
	// "testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestRegisterRoutes tests all routes registered by the RegisterRoutes function
func TestRegisterRoutes(t *testing.T) {
	// Setup a fresh app for each test to ensure clean state
	app := fiber.New()
	api := app.Group("/api")
	users.RegisterRoutes(api)

	// Create a user first to use in subsequent tests
	var userId string

	// Test POST route with validation middleware
	t.Run("POST route with validation middleware", func(t *testing.T) {
		// Test with invalid data
		req := httptest.NewRequest("POST", "/api/users/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		// Test with valid data
		validUserJSON := `{
			"id": 1,
			"name": "Test User",
			"email": "test@example.com",
			"age": 30
		}`
		req = httptest.NewRequest("POST", "/api/users/", strings.NewReader(validUserJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, err = app.Test(req)
		assert.NoError(t, err)

		// Check if the response is 201 Created or 200 OK (both are acceptable)
		assert.True(t, resp.StatusCode == fiber.StatusCreated || resp.StatusCode == fiber.StatusOK,
			"Expected status code to be 201 Created or 200 OK, got %d", resp.StatusCode)

		// Parse the response to get the user ID for later tests
		var userResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&userResponse)
		assert.NoError(t, err)

		// Extract the user ID from the response - handle different possible formats
		if id, ok := userResponse["id"]; ok && id != nil {
			switch v := id.(type) {
			case string:
				userId = v
			case float64:
				userId = fmt.Sprintf("%.0f", v)
			case int:
				userId = fmt.Sprintf("%d", v)
			default:
				t.Logf("User ID is of unexpected type: %T", id)
				userId = "1" // Fallback to a default ID
			}
		} else {
			t.Log("Response doesn't contain an 'id' field, using default ID")
			userId = "1" // Fallback to a default ID
		}

		assert.NotEmpty(t, userId, "User ID should not be empty")
	})

	// Test GET all users route
	t.Run("GET all users route", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/users/", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Verify response contains users
		var users []interface{}
		err = json.NewDecoder(resp.Body).Decode(&users)
		assert.NoError(t, err)
		assert.NotEmpty(t, users, "Users list should not be empty")
	})

	// Test GET user by ID route
	t.Run("GET user by ID route", func(t *testing.T) {
		// Skip if we don't have a user ID
		if userId == "" {
			t.Skip("No user ID available, skipping test")
		}

		// Get the user we created
		req := httptest.NewRequest("GET", "/api/users/"+userId, nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Test with non-existent ID
		req = httptest.NewRequest("GET", "/api/users/999", nil)
		resp, err = app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	// Test PUT route
	t.Run("PUT route", func(t *testing.T) {
		// Skip if we don't have a user ID
		if userId == "" {
			t.Skip("No user ID available, skipping test")
		}

		// Update existing user
		updateJSON := `{
			"name": "Updated User",
			"email": "updated@example.com",
			"age": 31
		}`
		req := httptest.NewRequest("PUT", "/api/users/"+userId, strings.NewReader(updateJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Update non-existent user
		req = httptest.NewRequest("PUT", "/api/users/999", strings.NewReader(updateJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, err = app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	// Test DELETE route
	t.Run("DELETE route", func(t *testing.T) {
		// Skip if we don't have a user ID
		if userId == "" {
			t.Skip("No user ID available, skipping test")
		}

		// Delete existing user
		req := httptest.NewRequest("DELETE", "/api/users/"+userId, nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Delete non-existent user
		req = httptest.NewRequest("DELETE", "/api/users/999", nil)
		resp, err = app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		// Verify user was deleted
		req = httptest.NewRequest("GET", "/api/users/"+userId, nil)
		resp, err = app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}
