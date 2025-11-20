// booking seat service
package service

import (
	"fmt"

	"github.com/google/uuid"

	"api/internal/gateway"
	model "booking/pkg/monorepo"
	dto "cinema/pkg/monorepo"
)

type ListBookingSeatRow struct {
	ID        uuid.UUID       `json:"id"`
	BookingID int64           `json:"booking_id"`
	Seat      dto.ListSeatRow `json:"seat"`
	PricePaid float64         `json:"price_paid"`
}

type DataBookingSeatRow struct {
	ID        uuid.UUID      `json:"id"`
	BookingID int64          `json:"booking_id"`
	Seat      dto.GetSeatRow `json:"seat"`
	PricePaid float64        `json:"price_paid"`
}

func ListBookingSeat(body any) ([]ListBookingSeatRow, error) {
	url := "/booking/seats"

	booking_seat, err := gateway.PostBooking[any, []model.BookingSeat](url, body)
	if err != nil {
		return []ListBookingSeatRow{}, err
	}

	data_seat, err := ListSeat()
	if err != nil {
		return []ListBookingSeatRow{}, err
	}

	seatMap := make(map[int64]dto.ListSeatRow)
	for _, s := range data_seat {
		seatMap[s.ID] = s
	}

	var data []ListBookingSeatRow
	for _, b := range booking_seat {
		seat := seatMap[b.SeatID]
		response := ListBookingSeatRow{
			ID:        b.ID,
			BookingID: b.BookingID,
			Seat:      seat,
			PricePaid: b.PricePaid,
		}
		data = append(data, response)
	}

	return data, nil
}

func CreateBookingSeat(body model.CreateBookingSeatParams) (DataBookingSeatRow, error) {
	url := "/booking/seat/create"

	booking_seat, err := gateway.PostBooking[model.CreateBookingSeatParams, model.BookingSeat](url, body)
	if err != nil {
		return DataBookingSeatRow{}, err
	}

	seat, err := GetSeat(booking_seat.SeatID)
	if err != nil {
		return DataBookingSeatRow{}, err
	}

	data := DataBookingSeatRow{
		ID:        booking_seat.ID,
		BookingID: booking_seat.BookingID,
		Seat:      seat,
		PricePaid: booking_seat.PricePaid,
	}

	return data, nil
}

func DeleteBookingSeat(id string) error {
	url := fmt.Sprintf("/booking/seat/delete/%s", id)

	if err := gateway.DeleteBooking(url); err != nil {
		return err
	}
	return nil
}
