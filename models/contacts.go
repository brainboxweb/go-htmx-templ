package models

type Contacts []Contact

func (cc Contacts) Add(c Contact) {
	cc = append(cc, c)
}

type Contact struct {
	Name  string
	Email string
}

func NewContact(name, email string) Contact {
	return Contact{name, email}
}

type Data struct {
	Contacts Contacts
}

func NewData() Data {
	return Data{
		Contacts: []Contact{
			NewContact("John", "john@john.com"),
			NewContact("Clara", "clara@clara.com"),
		},
	}
}

func (d *Data) HasEmail(email string) bool {
	for _, val := range d.Contacts {
		if val.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func NewPage() Page { // Start here. The starting values are the problem!!!
	return Page{
		Data: NewData(),
		Form: NewFormData(),
	}
}
