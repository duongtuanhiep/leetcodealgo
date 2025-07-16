import bisect
import collections
import sortedcontainers

"""
Basic Room Booking

The system must allow users to book meeting rooms. This includes:

- Reserve a room for a specific date and time
- Check room availability before booking
- Cancel existing reservations

Entities
MeetingRoom
Calendar
Reservation
"""


class MeetingRoom:
    def __init__(self, id: int):
        self.id = id
        self.scheduled_meetings = sortedcontainers.SortedList()

    def assign_meeting(self, start, end) -> bool:
        meeting = (start, end)
        meeting_idx = bisect.bisect_left(
            self.scheduled_meetings, start, key=lambda x: x[0]
        )
        n = len(self.scheduled_meetings)
        if (
            meeting_idx == n
            and self.scheduled_meetings[-1][1] <= start
            or 0 < meeting_idx < n
            and self.scheduled_meetings[meeting_idx - 1][1] <= start
            and end <= self.scheduled_meetings[meeting_idx][0]
            or meeting_idx == 0
            and end <= self.scheduled_meetings[meeting_idx][0]
        ):
            self.scheduled_meetings.append(meeting)
            return True
        return False

    # Get all timeslot in between current meetings up to a time
    def get_all_available_timeslot(self, end_time: int) -> list[tuple[int, int]]:
        available = []
        free_start, free_end = 0, 0
        for i, (start, end) in enumerate(self.scheduled_meetings):
            if start <= free_start:
                free_start = max(free_start, end)
                # TODO: Perform cancellation

    def cancel_meeting(self, meeting_reservation: "Reservation") -> bool:
        meeting = (meeting_reservation.start, meeting_reservation.end)
        if meeting in self.scheduled_meetings:
            self.scheduled_meetings.discard(meeting)
            return True
        return False


class Reservation:
    def __init__(self, room: MeetingRoom, start: int, end: int):
        self.room = room
        self.start = start
        self.end = end


class BookingSystem:
    def __init__(self, meeting_rooms: list[MeetingRoom]):
        self.rooms = meeting_rooms

    def reserve_meeting(self, start, end) -> Reservation:
        for room in self.rooms:
            if room.assign_meeting(start, end):
                print(f"Reserved room {room.id} for meeting!")
                return Reservation(room, start, end)
        raise Exception("No meeting rooms available!")

    def get_all_availability(self, end_time: int):
        for room in self.rooms:
            print(f"Room {room.id} availability")
