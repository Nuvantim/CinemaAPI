// booking service
package service

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"api/internal/gateway"
	model "booking/pkgs/monorepo"
	dto "cinema/pkgs/monorepo"
)

type ListBookingRow struct {
	ID          int64               `json:"id"`
	UserID      int64               `json:"user_id"`
	ShowtimeID  int64               `json:"showtime_id"`
	ShowTime    dto.ListShowTimeRow `json:"showtime"`
	BookingTime pgtype.Timestamp    `json:"booking_time"`
	TotalAmount pgtype.Float8       `json:"total_amount"`
}

func ListBooking(body any) ([]ListBookingRow, error) {
	url := "/bookings"

	data_booking, err := gateway.PostBooking[any, []model.Booking](url, body)
	if err != nil {
		return []ListBookingRow{}, err
	}

	data_showtime, err := ListShowTime()
	if err != nil {
		return []ListBookingRow{}, err
	}

	showtimeMap := make(map[int64]dto.ListShowTimeRow)
	for _, s := range data_showtime {
		showtimeMap[s.ID] = s
	}

	var data []ListBookingRow
	for _, b := range data_booking {
		showtime := showtimeMap[b.ShowtimeID]
		response := ListBookingRow{
			ID:          b.ID,
			UserID:      b.UserID,
			ShowtimeID:  b.ShowtimeID,
			ShowTime:    showtime,
			TotalAmount: b.TotalAmount,
		}

		data = append(data, response)
	}

	return data, nil
}

func CreateBooking(body model.CreateBookingParams) ([]model.Booking, error) {
	var url = "/booking/create"

	data, err := gateway.PostBooking[model.CreateBookingParams, []model.Booking](url, body)
	if err != nil {
		return []model.Booking{}, err
	}

	return data, nil
}

func DeleteBooking(id int64) error {
	var url = fmt.Sprintf("/booking/delete/%d", id)

	if err := gateway.DeleteBooking(url); err != nil {
		return err
	}
	return nil
}
