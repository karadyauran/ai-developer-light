
import (
    "errors"
    "strings"
)

func validateDonation(donation Donation) error {
    if strings.TrimSpace(donation.Item) == "" {
        return errors.New("item cannot be empty")
    }
    if donation.Quantity <= 0 {
        return errors.New("quantity must be greater than zero")
    }
    return nil
}

func validateRequest(request FoodRequest) error {
    if strings.TrimSpace(request.Item) == "" {
        return errors.New("item cannot be empty")
    }
    if request.Quantity <= 0 {
        return errors.New("quantity must be greater than zero")
    }
    return nil