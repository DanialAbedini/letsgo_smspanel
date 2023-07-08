package ports

import (
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/google/uuid"
)

type TemplateRepositoryContract interface {
	AddTemplate(temp entities.Template) error
}

type SnedSMSRepositoryContract interface {
	RequestContactList(id uuid.UUID) ([]entities.Contact, error)
	RequestNumber(id uuid.UUID) (entities.Number, error)
	RequestUser(id uuid.UUID) (entities.User, error)
}
type PhoneBookRepositoryContract interface {
	CreatePhoneBookList(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	GetPhoneBookList(phoneBookModel entities.PhoneBook) ([]entities.PhoneBook, error)
	GetPhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	UpdatePhoneBookById(phoneBookModel entities.PhoneBook) (entities.PhoneBook, error)
	DeletePhoneBookById(phoneBookModel entities.PhoneBook) error
}

type ContactRepositoryContract interface {
	CreateContact(contactModel entities.Contact) (entities.Contact, error)
	GetContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactByUsername(contactModel entities.Contact) (entities.Contact, error)
	DeleteContactByUsername(contactModel entities.Contact) error
	GetContactList(contactModel entities.Contact) ([]entities.Contact, error)
	GetContactById(contactModel entities.Contact) (entities.Contact, error)
	UpdateContactById(contactModel entities.Contact) (entities.Contact, error)
	DeleteContactById(contactModel entities.Contact) error
}
