
type Donation struct {
    ID       int    `json:"id"`
    Item     string `json:"item"`
    Quantity int    `json:"quantity"`
}

type FoodRequest struct {
    ID       int    `json:"id"`
    Item     string `json:"item"`
    Quantity int    `json:"quantity"`
}

var donations []Donation
var requests []FoodRequest

func saveDonation(donation Donation) {
    donations = append(donations, donation)
}

func saveRequest(request FoodRequest) {
    requests = append(requests, request)
}

func getAllDonations() []Donation {
    return donations
}

func getAllRequests() []FoodRequest {
    return requests