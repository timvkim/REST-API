package models

type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price"`
	AuthorID    int    `json:"author_id"`
}

type UpdateCourse struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price"`
}

func (u UpdateCourse) Validate() error {

	errs := ErrorFields{
		Fields: make(map[string]string),
	}

	if u.Price <= 0 {
		errs.Fields["price"] = "the price has to be a positive"
	}

	if len(u.Title) > 20 {
		errs.Fields["title"] = "the title has to be a less than 20 symbols"
	}

	if len(errs.Fields) == 0 {
		return nil
	}

	return errs

}
