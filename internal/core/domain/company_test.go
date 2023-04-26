package domain_test

import (
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Input   string
		IsValid bool
	}{
		{
			"",
			false,
		},
		{
			"random name",
			true,
		},
		{
			"qwe",
			true,
		},
		{
			"very long company name",
			false,
		},
	}

	for _, tc := range testCases {
		err := domain.ValidateName(tc.Input)
		if tc.IsValid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestValidateDescription(t *testing.T) {
	t.Parallel()

	longString := make([]rune, 3001)
	for i := range longString {
		longString[i] = 'a'
	}

	testCases := []struct {
		Input   string
		IsValid bool
	}{
		{
			"",
			true,
		},
		{
			"random description",
			true,
		},
		{
			string(longString[:3000]),
			true,
		},
		{
			string(longString),
			false,
		},
	}

	for _, tc := range testCases {
		err := domain.ValidateDescription(tc.Input)
		if tc.IsValid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestValidateType(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Input   string
		IsValid bool
	}{
		{
			"Corporations",
			true,
		},
		{
			"NonProfit",
			true,
		},
		{
			"Cooperative",
			true,
		},
		{
			"Sole Proprietorship",
			true,
		},
		{
			"Sole Proprietorshi",
			false,
		},
		{
			"ole Proprietorship",
			false,
		},
	}

	for _, tc := range testCases {
		err := domain.ValidateType(tc.Input)
		if tc.IsValid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Input   domain.Company
		IsValid bool
	}{
		{
			domain.Company{Name: "name", Description: "desc", Type: "NonProfit"},
			true,
		},
		{
			domain.Company{Name: "name", Description: "desc", Type: "NonaProfit"},
			false,
		},
		{
			domain.Company{Name: "", Description: "desc", Type: "NonProfit"},
			false,
		},
		{
			domain.Company{Name: "na", Description: "desc", Type: "NonProfit"},
			false,
		},
	}
	for _, tc := range testCases {
		err := tc.Input.Validate()
		if tc.IsValid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
