package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
)

type IEventUsecase interface {
	CreateNewEvent(ctx context.Context, name, detail, location, adminIdString string) (*entity.Event, error)
	GetEventById(ctx context.Context, eventIdString string) (*entity.Event, error)
	DeleteEventById(ctx context.Context, eventIdString string) error
	UpdateEventById(ctx context.Context, eventIdString string, name,detail,location string) (*entity.Event, error)
	ChangeEventStatusOfId(ctx context.Context, eventIdString string, stateString string) (*entity.Event, error)
	GetUserEvents(ctx context.Context, userIdString string) ([]*entity.Event,error)
}

type EventUsecase struct {
	repo repository.IEventRepository
}

func NewEventUsecae(r repository.IEventRepository) IEventUsecase {
	return &EventUsecase{
		repo: r,
	}
}

func (uc *EventUsecase) CreateNewEvent(ctx context.Context, name, detail, location, adminIdString string) (*entity.Event, error) {
	if name == "" {
		return nil, fmt.Errorf("EventUsecase: name is empty")
	}
	adminId := entity.UserId(adminIdString)
	if adminId == "" {
		return nil, fmt.Errorf("EventUsecase: adminId parse failed")
	}
	e := &entity.Event{
		Name:     name,
		Detail:   detail,
		Location: location,
	}
	return uc.repo.CreateNewEvent(ctx, e, adminId)
}

func (uc *EventUsecase) GetEventById(ctx context.Context, eventIdString string) (*entity.Event, error) {
	// TODO: serviceとかでcast関数を用意すべき
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("EventUsecase: eventId parse failed")
	}
	return uc.repo.GetEventById(ctx, eventId)
}

func (uc *EventUsecase) DeleteEventById(ctx context.Context, eventIdString string) error {
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return fmt.Errorf("EventUsecase: eventId parse failed")
	}
	// TODO: adminuser か確認
	return uc.repo.DeleteEventById(ctx,eventId)
}

func (uc *EventUsecase) UpdateEventById(ctx context.Context, eventIdString string, name,detail,location string) (*entity.Event, error) {	
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("EventUsecase: eventId parse failed")
	}
	// TODO: adminuser か確認
	if name == "" {
		return nil, fmt.Errorf("EventUsecase: name is empty")
	}
	e := &entity.Event{
		Name:     name,
		Detail:   detail,
		Location: location,
	}
	return uc.repo.UpdateEventById(ctx,eventId,e)
}


func (uc *EventUsecase) ChangeEventStatusOfId(ctx context.Context, eventIdString string, stateString string) (*entity.Event, error) {
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("EventUsecase: eventId parse failed")
	}
	state := entity.State(stateString)
	if state == "" {
		return nil, fmt.Errorf("EventUsecase: state parse failed")
	}
	// すでにclose,cancelだった場合
	if state == entity.CLOSE_STATE || state == entity.CANCEL_STATE {
		return nil, fmt.Errorf("EventUsecase: this event is already close or cancel")
	}
	switch state {
	case entity.CLOSE_STATE:
		return uc.repo.ChangeEventStatusToCloseOfId(ctx, eventId)
	case entity.CANCEL_STATE:
		return uc.repo.ChangeEventStatusToCancelOfId(ctx, eventId)
	default:
		return nil, fmt.Errorf("EventUsecase: selected state is not matched")
	}
}

func (uc *EventUsecase) GetUserEvents(ctx context.Context, userIdString string) ([]*entity.Event, error) {
	// TODO: serviceとかでcast関数を用意すべき
	userId := entity.UserId(userIdString)
	if userId == "" {
		return nil, fmt.Errorf("EventUsecase: userId parse failed")
	}
	return uc.repo.GetUserEvents(ctx,userId)
}
