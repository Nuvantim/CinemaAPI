// booking service
package service

import (
	"fmt"
	"time"

	req "api/internal/app/request"
	"api/internal/gateway"
	rds "api/redis"
	model "booking/pkg/monorepo"
	dto "cinema/pkg/monorepo"
)

type ListBookingRow struct {
	ID          int64               `json:"id"`
	UserID      int64               `json:"user_id"`
	ShowTime    dto.ListShowTimeRow `json:"showtime"`
	BookingTime time.Time           `json:"booking_time"`
	TotalAmount float64             `json:"total_amount"`
}

type DataBookingRow struct {
	ID          int64              `json:"id"`
	UserID      int64              `json:"user_id"`
	ShowTime    dto.GetShowTimeRow `json:"showtime"`
	BookingTime time.Time          `json:"booking_time"`
	TotalAmount float64            `json:"total_amount"`
}

func ListBooking(body req.BookingPayment) ([]ListBookingRow, error) {
	// get data booking
	data_booking, err := func() ([]model.Booking, error) {
		// check data on redis
		key := fmt.Sprintf("list:booking:%d", body.UserID)
		redis_data, err := rds.GetData[[]model.Booking](key)
		if err == nil && redis_data != nil {
			return *redis_data, nil
		}

		// get data from service
		url := "/bookings"
		booking_data, err := gateway.PostBooking[any, []model.Booking](url, body)
		if err != nil {
			return []model.Booking{}, err
		}
		return booking_data, nil
	}()
	if err != nil {
		return []ListBookingRow{}, err
	}

	// get showtime data
	data_showtime, err := ListShowTime()
	if err != nil {
		return []ListBookingRow{}, err
	}

	//create map for showtime
	showtimeMap := make(map[int64]dto.ListShowTimeRow)
	for _, s := range data_showtime {
		showtimeMap[s.ID] = s
	}

	// crete response
	var data []ListBookingRow
	for _, b := range data_booking {
		showtime := showtimeMap[b.ShowtimeID]
		response := ListBookingRow{
			ID:          b.ID,
			UserID:      b.UserID,
			ShowTime:    showtime,
			BookingTime: b.BookingTime,
			TotalAmount: b.TotalAmount,
		}

		data = append(data, response)
	}

	return data, nil
}

func GetBooking(id int64) (model.Booking, error) {
	url := fmt.Sprintf("/booking/%d", id)

	data, err := gateway.GetBooking[model.Booking](url)
	if err != nil {
		return model.Booking{}, err
	}

	return data, nil
}

func CreateBooking(body model.CreateBookingParams) (DataBookingRow, error) {
	var url = "/booking/create"

	booking, err := gateway.PostBooking[model.CreateBookingParams, model.Booking](url, body)
	if err != nil {
		return DataBookingRow{}, err
	}
	showtime, err := GetShowTime(booking.ShowtimeID)
	if err != nil {
		return DataBookingRow{}, err
	}

	data := DataBookingRow{
		ID:          booking.ID,
		UserID:      booking.UserID,
		ShowTime:    showtime,
		BookingTime: booking.BookingTime,
		TotalAmount: booking.TotalAmount,
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
