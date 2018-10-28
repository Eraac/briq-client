package briq

type (
	Organization struct {
		ID        string `json:"id"`
		Key       string `json:"key"`
		CreatedAt string `json:"created_at"`
	}

	User struct {
		ID              string `json:"id"`
		Email           string `json:"email"`
		Image           string `json:"image"`
		OrganizationID  string `json:"organization_id"`
		DisplayName     string `json:"displayName"`
		Role            string `json:"role"`
		CreatedAt       string `json:"created_at"`
		Username        string `json:"username"`
		ExternalRef     string `json:"externalRef"`
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		ActiveBalance   int    `json:"activeBalance"`
		InactiveBalance int    `json:"inactiveBalance"`
	}

	Transaction struct {
		ID             string `json:"id"`
		From           string `json:"from"`
		To             string `json:"to"`
		Amount         int    `json:"amount"`
		App            string `json:"app"`
		Comment        string `json:"comment"`
		Reaction       string `json:"reaction"`
		ReactedAt      string `json:"reactedAt"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
		ExpiredAt      string `json:"expiredAt"`
		OrganizationID string `json:"organization_id"`
		UserFromID     string `json:"user_from_id"`
		UserToID       string `json:"user_to_id"`
	}

	TransactionInput struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Amount  int    `json:"amount"`
		Comment string `json:"comment"`
	}
)
