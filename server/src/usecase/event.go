package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/constant"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/service"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/google/uuid"
)

type IEvent interface {
	CreateNewEvent(ctx context.Context, adminIdString, name, detail, location string) (*ent.Event, error)
	GetEventById(ctx context.Context, eventIdString string) (*ent.Event, error)
	DeleteEventById(ctx context.Context, eventIdString string) error
	UpdateEventById(ctx context.Context, eventIdString string, name, detail, location string) (*ent.Event, error)
	GetEventAdminById(ctx context.Context, eventIdString string) (*ent.User, error)
	GetEventComments(ctx context.Context, eventIdString string) ([]*ent.Comment, error)
	AddNewEventParticipant(ctx context.Context, eventIdString, name, comment string) error
	ChangeEventStatusOfId(ctx context.Context, eventIdString string, stateString string) (*ent.Event, error)
	GetEventUsers(ctx context.Context, eventIdString string) ([]*ent.User, error)
}

type Event struct {
	repo repository.IEvent
}

func NewEvent(r repository.IEvent) IEvent {
	return &Event{
		repo: r,
	}
}

func (uc *Event) CreateNewEvent(ctx context.Context, adminIdString, name, detail, location string) (*ent.Event, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	adminId, err := uuid.Parse(adminIdString)
	if err != nil {
		return nil, fmt.Errorf("adminId Parse: %w", err)
	}
	ee := &ent.Event{
		Name:     name,
		Detail:   detail,
		Location: location,
	}
	return uc.repo.CreateNewEvent(ctx, adminId, ee)
}

func (uc *Event) GetEventById(ctx context.Context, eventIdString string) (*ent.Event, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventById(ctx, eventId)
}

func (uc *Event) DeleteEventById(ctx context.Context, eventIdString string) error {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return fmt.Errorf("eventId Parse: %w", err)
	}
	// TODO: adminuser か確認
	return uc.repo.DeleteEventById(ctx, eventId)
}

func (uc *Event) UpdateEventById(ctx context.Context, eventIdString string, name, detail, location string) (*ent.Event, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	// TODO: adminuser か確認
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	ee := &ent.Event{
		Name:     name,
		Detail:   detail,
		Location: location,
	}
	return uc.repo.UpdateEventById(ctx, eventId, ee)
}

func (uc *Event) GetEventAdminById(ctx context.Context, eventIdString string) (*ent.User, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventAdminById(ctx, eventId)
}

func (uc *Event) GetEventComments(ctx context.Context, eventIdString string) ([]*ent.Comment, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventComments(ctx, eventId)
}

func (uc *Event) AddNewEventParticipant(ctx context.Context, eventIdString, name, comment string) error {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return fmt.Errorf("eventId Parse: %w", err)
	}
	if name == "" {
		return fmt.Errorf("name is empty")
	}
	eu := &ent.User{
		Name: name,
		Icon: service.GetRandomDefaultIcon(),
	}
	err = uc.repo.AddNewEventParticipant(ctx, eventId, eu, comment)
	if err != nil {
		return fmt.Errorf("AddNewEventParticipant: %w", err)
	}
	return nil
}

func (uc *Event) ChangeEventStatusOfId(ctx context.Context, eventIdString string, state string) (*ent.Event, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	if state == "" {
		return nil, fmt.Errorf("state is Empty")
	}
	// TODO:すでにclose,cancelだった場合
	// TODO:また,open以外の時はparticipantできないようにする処理もいる
	switch state {
	case constant.STATE_CLOSE:
		return uc.repo.ChangeEventStatusToCloseOfId(ctx, eventId)
	case constant.STATE_CANCEL:
		return uc.repo.ChangeEventStatusToCancelOfId(ctx, eventId)
	default:
		return nil, fmt.Errorf("received state is not matched")
	}
}

func (uc *Event) GetEventUsers(ctx context.Context, eventIdString string) ([]*ent.User, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventUsers(ctx, eventId)
}
