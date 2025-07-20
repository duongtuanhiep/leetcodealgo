from abc import ABC, abstractmethod
from enum import Enum
from datetime import datetime
from dataclasses import dataclass
from collections import defaultdict
from threading import Lock

"""
Core Features of Ticket booking

Search movies by city
Select a show (movie + time + cinema)
Pick seats from available ones
Book tickets and pay

Entities:
Movie
Show(movie/time/cinema)
Seats
Ticket
"""


class Movie:
    id: str
    name: str


class Show:
    id: str
    movie: Movie
    location: str
    start: datetime
    seats: list["Seat"]


class Seat:
    id: str
    position: str  # Some row, col number
    price: int


class TicketReservation:
    show: Show
    seat: Seat
    paid: bool


class ShowSearch(ABC):
    @abstractmethod
    def find_show(self, time: datetime, movie_name: str) -> list[Show]:
        pass


class SeatReservation(ABC):
    @abstractmethod
    def reserve_seats(self, show: "Show", seat: "Seat") -> "TicketReservation":
        pass

    @abstractmethod
    def get_available_seats(self, show: Show) -> list[Seat]:
        pass

class Payment(ABC):
    @abstractmethod
    def pay(self, reservation: "TicketReservation") -> bool:
        pass

class CardPayment(Payment):
    def pay(self, reservation) -> bool:
        print("user paid")
        return True
    

class ShowSearcher(ShowSearch):
    shows_by_name: defaultdict[str, list[Show]]

    def find_show(self, time, movie_name):
        if movie_name in self.shows_by_name:
            return [show for show in self.shows_by_name.get(movie_name) if show.start == time]

class ReservationHandler(SeatReservation):
    seats_reservation : dict["Seat", bool]
    lock: Lock

    def reserve_seats(self, show, seat):
        with self.lock:
            if seat.id in self.seats_reservation and not self.seats_reservation.get(seat.id):
                reservation = TicketReservation() # TODO: real implementation
                self.seats_reservation.get(seat.id) = True
                return reservation
        
        
    def get_available_seats(self, show) -> list[Seat]:
        with self.lock:
            return [seat for seat in show.seats if not self.seats_reservation.get(seat.id)]
    
class TicketBooker:
    shows: ShowSearch
    payment: Payment
    ticket: SeatReservation

    def search_show(self, time, movie_name):
        # Do more stuff
        return self.shows.find_show(time, movie_name)

    def book_ticket(self, show: Show, seat: Seat):
        reservation = self.ticket.reserve_seats(show, seat)
        if reservation: 
            success = self.payment.pay(reservation)
            if success: 
                return reservation
        