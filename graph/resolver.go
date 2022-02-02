package graph

import (
	_commentRepository "github.com/justjundana/event-planner/repository/comment"
	_eventRepository "github.com/justjundana/event-planner/repository/event"
	_participantRepository "github.com/justjundana/event-planner/repository/participant"
	_userRepository "github.com/justjundana/event-planner/repository/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRepository        _userRepository.UserInterface
	eventRepository       _eventRepository.EventInterface
	participantRepository _participantRepository.ParticipantInterface
	commentRepository     _commentRepository.CommentInterface
}

func NewResolver(
	ur _userRepository.UserInterface,
	er _eventRepository.EventInterface,
	pr _participantRepository.ParticipantInterface,
	cr _commentRepository.CommentInterface,
) *Resolver {
	return &Resolver{
		userRepository:        ur,
		eventRepository:       er,
		participantRepository: pr,
		commentRepository:     cr,
	}
}
