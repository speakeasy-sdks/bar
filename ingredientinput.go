// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package bar

type IngredientInput struct {
	// The name of the ingredient.
	Name string `json:"name"`
	// The type of ingredient.
	Type IngredientType `json:"type"`
	// The product code of the ingredient, only available when authenticated.
	ProductCode *string `json:"productCode,omitempty"`
}

func (o *IngredientInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *IngredientInput) GetType() IngredientType {
	if o == nil {
		return IngredientType("")
	}
	return o.Type
}

func (o *IngredientInput) GetProductCode() *string {
	if o == nil {
		return nil
	}
	return o.ProductCode
}
