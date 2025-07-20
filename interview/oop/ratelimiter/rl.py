from collections import Protocol
from datetime import datetime
from threading import Lock

"""
Implement a simple rate limiter:
- Rate limit request by IP

"""


class Client:
    ip: str


class RateLimit(Protocol):
    def allow(self, client: Client) -> bool: ...


# We store last request timestamp and increase request frequency
class TimeWindowRateLimiter:
    def __init__(self, frequency, duration):
        self.frequency = frequency
        self.duration = duration
        self.requestFrequency = {}  # Key: IP, Val: (windowStart, count)

    def allow(self, client: Client) -> bool:
        windowStart, lastCount = self.requestFrequency.get(client.ip, (0, 0))
        currentTime = datetime.now().timestamp()

        if (
            windowStart < currentTime - self.duration
        ):  # if it's out of range then move it to "now"
            newCount = 1
            windowStart = currentTime
            self.requestFrequency[client.ip] = (windowStart, newCount)
        elif lastCount < self.frequency:
            newCount = lastCount + 1
            self.requestFrequency[client.ip] = (windowStart, newCount)
        else:
            return False

        return True

    def clean_up(self):
        expiredKeys = [
            k
            for k, v in self.requestFrequency.items()
            if v < datetime.now().timestamp() - self.duration
        ]

        for key in expiredKeys:
            self.requestFrequency.pop(key, None)


class TokenBucketRateLimiter:
    def __init__(self, bucket_size: int, filling_rate: int):
        self.ip_to_bucket = {}
        self.bucket_size = bucket_size
        self.fillint_rate = filling_rate  # time elapsed to fill 1
        self.lock = Lock()

    def allow(self, client: Client) -> bool:
        with self.lock:
            currentTime = datetime.now()
            token, lastRefill = self.ip_to_bucket.get(
                client.ip, (self.bucket_size, currentTime)
            )

            time_passed = currentTime - lastRefill
            new_token = min(self.bucket_size, token + time_passed // self.fillint_rate)

            if new_token:
                self.ip_to_bucket[client.ip] = (new_token, currentTime)
                return True
            else:
                self.ip_to_bucket[client.ip] = (new_token, lastRefill)
                return False
