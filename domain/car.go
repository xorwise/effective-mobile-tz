package domain

type Car struct {
	RegNum          string `db:"reg_num" json:"reg_num"`
	Mark            string `db:"mark" json:"mark"`
	Model           string `db:"model" json:"model"`
	Year            int    `db:"year" json:"year"`
	OwnerName       string `db:"owner_name" json:"owner_name"`
	OwnerSurname    string `db:"owner_surname" json:"owner_surname"`
	OwnerPatronymic string `db:"owner_patronymic" json:"owner_patronymic"`
}

type UpdateCarRequest struct {
	Mark            string `json:"mark"`
	Model           string `json:"model"`
	Year            int    `json:"year"`
	OwnerName       string `json:"owner_name"`
	OwnerSurname    string `json:"owner_surname"`
	OwnerPatronymic string `json:"owner_patronymic"`
}

type CreateCarRequest struct {
	RegNums []string `json:"regNums"`
}

type ExternalResponse struct {
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
	} `json:"owner"`
}
