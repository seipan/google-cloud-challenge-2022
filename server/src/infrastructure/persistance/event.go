package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/constant"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/comment"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/event"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

type Event struct {
	client *ent.Client
}

func NewEvent(c *ent.Client) repository.IEvent {
	return &Event{
		client: c,
	}
}

func (repo *Event) CreateNewEvent(ctx context.Context, adminId uuid.UUID, ee *ent.Event) (*ent.Event, error) {
	event, err := repo.client.Event.
		Create().
		SetName(ee.Name).
		SetDetail(ee.Detail).
		SetLocation(ee.Location).
		SetAdminID(adminId).
		SetType(string(constant.ONCE_TYPE)).
		SetState(string(constant.OPEN_STATE)).
		AddUserIDs(adminId).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.Create: %w", err)
	}
	return repo.getEventById(ctx, event.ID)
}

func (repo *Event) GetEventById(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	return repo.getEventById(ctx, eventId)
}

func (repo *Event) DeleteEventById(ctx context.Context, eventId uuid.UUID) error {
	err := repo.client.Event.
		DeleteOneID(eventId).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("Event.DeleteOneID: %w", err)
	}
	return nil
}

func (repo *Event) UpdateEventById(ctx context.Context, eventId uuid.UUID, ee *ent.Event) (*ent.Event, error) {
	event, err := repo.client.Event.
		UpdateOneID(eventId).
		SetName(ee.Name).
		SetDetail(ee.Detail).
		SetLocation(ee.Location).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (repo *Event) GetEventAdminById(ctx context.Context, eventId uuid.UUID) (*ent.User, error) {
	event, err := repo.client.Event.
		Query().
		Where(event.ID(eventId)).
		WithAdmin().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.Query: %w", err)
	}
	return event.Edges.Admin, nil
}

func (repo *Event) GetEventComments(ctx context.Context, eventId uuid.UUID) ([]*ent.Comment, error) {
	comments, err := repo.client.Comment.
		Query().
		Where(comment.HasEventWith(event.ID(eventId))).
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Comment.Query: %w", err)
	}
	return comments, nil
}

func (repo *Event) AddNewEventParticipant(ctx context.Context, eventId uuid.UUID, eu *ent.User, comment string) error {
	user, err := repo.client.User.
		Create().
		SetName(eu.Name).
		SetIcon(eu.Icon).
		AddEventIDs(eventId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("Userepo.Create: %w", err)
	}
	if comment == "" {
		return nil
	}
	_, err = repo.client.Comment.
		Create().
		SetBody(comment).
		SetEventID(eventId).
		SetUserID(user.ID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("Comment.Create: %w", err)
	}
	return nil
}

func (repo *Event) ChangeEventStatusToCloseOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	event, err := repo.client.Event.
		UpdateOneID(eventId).
		SetState(string(constant.CLOSE_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (repo *Event) ChangeEventStatusToCancelOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	event, err := repo.client.Event.
		UpdateOneID(eventId).
		SetState(string(constant.CANCEL_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (repo *Event) GetEventUsers(ctx context.Context, eventId uuid.UUID) ([]*ent.User, error) {
	users, err := repo.client.User.
		Query().
		Where(user.HasEventsWith(event.ID(eventId))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Userepo.Query: %w", err)
	}
	return users, nil
}

func (repo *Event) getEventById(ctx context.Context, eventUuid uuid.UUID) (*ent.Event, error) {
	event, err := repo.client.Event.
		Query().
		Where(event.ID(eventUuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.Query: %w", err)
	}
	return event, nil
}
