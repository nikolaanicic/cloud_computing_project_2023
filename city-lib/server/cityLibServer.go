package server

import "rac_oblak_proj/city-lib/repositories"

type CityLibServer struct {
	rentals *repositories.RentalRepo
	books   *repositories.BookRepo
}
