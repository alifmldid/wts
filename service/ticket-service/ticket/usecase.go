package ticket

import "context"

type TicketUsecase interface{
	Insert(c context.Context, payload TicketPayload) (ticket NewTicketResponse, err error)
	GetTicket(c context.Context, id string) (ticket TicketData, err error)
	DataUpdate(c context.Context, id string, payload TicketPayload) (err error)
	StatusUpdate(c context.Context, id string, status string) (err error)
}

type ticketUsecase struct{
	ticketRepository TicketRepository
}

func NewTicketUsecase(ticketRepository TicketRepository) TicketUsecase{
	return &ticketUsecase{ticketRepository}
}

func (uc *ticketUsecase) Insert(c context.Context, payload TicketPayload) (ticket NewTicketResponse, err error){
	id, err := uc.ticketRepository.Save(c, payload)

	ticket.Id = id

	return ticket, err
}

func (uc *ticketUsecase) GetTicket(c context.Context, id string) (ticket TicketData, err error){
	ticket, err = uc.ticketRepository.GetById(c, id)

	return ticket, err
}

func (uc *ticketUsecase) DataUpdate(c context.Context, id string, payload TicketPayload) (err error){
	err = uc.ticketRepository.DataUpdate(c, id, payload)

	return err
}

func (uc *ticketUsecase) StatusUpdate(c context.Context, id string, status string) (err error){
	err = uc.ticketRepository.StatusUpdate(c, id, status)

	return err
}