package types

import (
	"strconv"
	"strings"
)

// Recipe domain
type Recipe struct {
	OID         string   `bson:"_id"`
	ID          string   `bson:"id"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Steps       []string `bson:"steps"`
	Ingredients []string `bson:"ingredients"`
}

// GetOID function
func (r *Recipe) GetOID() string {
	return r.OID
}

// SetOID function
func (r *Recipe) SetOID(id string) {
	r.OID = id
}

// GetID function
func (r *Recipe) GetID() string {
	return r.ID
}

// SetID function
func (r *Recipe) SetID(id string) {
	r.ID = id
}

// GetName function
func (r *Recipe) GetName() string {
	return r.Name
}

// SetName function
func (r *Recipe) SetName(name string) {
	r.Name = name
}

// GetDescription function
func (r *Recipe) GetDescription() string {
	return r.Description
}

// SetDescription function
func (r *Recipe) SetDescription(description string) {
	r.Description = description
}

// GetSteps function
func (r *Recipe) GetSteps() []string {
	return r.Steps
}

// SetSteps function
func (r *Recipe) SetSteps(steps []string) {
	r.Steps = steps
}

// SetIngredients function
func (r *Recipe) SetIngredients(ingredients []string) {
	r.Ingredients = ingredients
}

// GetIngredients function
func (r *Recipe) GetIngredients() []string {
	return r.Ingredients
}

// GetObjectInfo function
func (r *Recipe) GetObjectInfo() string {
	info := []string{
		r.GetName(),
		r.GetDescription(),
	}
	resp := strings.Join(info, ": ")

	for i, step := range r.GetSteps() {
		resp += "\nStep " + strconv.Itoa(i) + ":" + step
	}

	return resp
}
