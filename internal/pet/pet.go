package pet

import (
	"context"
	"errors"
	"log"
	"time"
)

var (
	ErrFetchingPet    = errors.New("failed to fetch Pet by ID")
	ErrNotImplemented = errors.New("not implemented")
)

// Pet - a representatoin for the pet
// structure for for our pet adoption service
type Pet struct {
	ID           string
	Available    bool
	Name         string
	Animal       string
	Breed        string
	Location     string
	Date_Rescued time.Time
}

// PetStore - this interface defines all of the methods
// that our service needs to operate
type PetStore interface {
	GetPetFromStore(context.Context, string) (Pet, error)
}

// Service - is the struct on which all our logical
// will be built on top of
// When we instantite service and pass into repository layer it needs
// to match the store interface
type Service struct {
	PetStore PetStore
}

// NewService - returns a pointer for new service
func NewService(petStore PetStore) *Service {
	return &Service{
		PetStore: petStore,
	}
}

func (s *Service) GetPet(ctx context.Context, id string) (Pet, error) {
	log.Println("Retrieving a Pet")
	// Logger.Info("Retrieving a Pet")
	pet, err := s.PetStore.GetPetFromStore(ctx, id)
	if err != nil {
		log.Println(err)
		// Unique errors won't expose underlying imp details
		return Pet{}, ErrFetchingPet
	}
	return pet, nil
}

func (s *Service) UpdatePet(ctx context.Context, pet Pet) error {
	log.Println("Updating a Pet")
	return ErrNotImplemented
}

func (s *Service) DeletePet(ctx context.Context, id string) error {
	log.Println("Deleting a Pet")
	return ErrNotImplemented
}

func (s *Service) AddPet(ctx context.Context, pet Pet) (Pet, error) {
	log.Println("Adding a Pet")
	return Pet{}, ErrNotImplemented
}

// Benefit of this this structure

// Service - uses -> Store
// Store - implements -> Pet business logic
// Pet business logic modifies Pet business data
// Pet Model defines relevant business data

// We can have a number of different implementations for storage
// We've essentially decopuled the business logic from data layer

// We can effectively test our application in units
// We can mock
