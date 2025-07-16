from abc import ABC, abstractmethod
from enum import Enum
from datetime import datetime
from dataclasses import dataclass
from collections import defaultdict

"""
Problem Statement
Design the core classes and relationships for a system that allows customers to leave reviews for products on Amazon.

Base Requirements:
Customers can leave reviews for products they have purchased.
Each review includes a rating (e.g., 1-5 stars), a text comment, and a timestamp.
Customers can edit or delete their own reviews.
Reviews are associated with specific products and customers.
The system should support displaying all reviews for a product, sorted by recency or rating.
Reviews can be upvoted or marked as helpful by other customers.
The system should prevent a customer from reviewing the same product more than once.

Twist 1: Review Moderation and Seller Responses
Requirements:
Review Moderation System: 
- Add support for content moderation where reviews can be flagged by users as inappropriate, spam, or fake. 
- Flagged reviews should go through a review process before being visible to other customers.

Seller Response Feature: 
- Allow product sellers/vendors to respond to customer reviews. 
- Sellers can only respond once per review, and their responses should be clearly marked as coming from the seller.

Review Status Management: 
- Reviews can have different states (pending, approved, rejected, flagged, hidden) and only approved reviews should be visible in normal product browsing.
"""


class CustomerRating(Enum):
    RATING_REAL_BAD = 1
    RATING_BAD = 2
    RATING_MEH = 3
    RATING_GOOD = 4
    RATING_REAL_GOOD = 5


class ReviewSortKey(Enum):
    SORTKEY_CREATED_AT = "created_at"
    SORTKEY_RATING = "rating"

class FlagReason(Enum):
    FLAG_REASON_SPAM = "spam"
    FLAG_REASON_INAPPROPRIATE = "inappropriate"
    FLAG_REASON_FAKE = "fake"


@dataclass
class Product:
    id: str
    name: str
    avg_rating: float
    review_cnt: int
    description: str


@dataclass
class User:
    id: str

@dataclass
class SellerResponse:
    id: str
    review_id: str
    content: str


@dataclass
class Review:
    id: str
    user_id: str
    product_id: str
    rating: CustomerRating
    description: str
    created_at: datetime
    modified_at: datetime
    upvotes: set  # set(user_id)
    helpfuls: set  # set(user_id)
    visible: bool
    seller_resp: SellerResponse

class ReviewImpression(ABC):
    @abstractmethod
    def upvote(self, user: User, review: Review):
        pass

    @abstractmethod
    def helpful(self, user: User, review: Review):
        pass


class ReviewCrud(ABC):
    @abstractmethod
    def get_by_id(self, id) -> Review:
        pass

    @abstractmethod
    def create_review(self, review: Review):
        pass

    @abstractmethod
    def modify_review(
        self, user: User, review: Review, description: str, rating: CustomerRating
    ):
        pass

    @abstractmethod
    def delete_review(self, user: User, review: Review):
        pass


class ProductReview(ABC):
    @abstractmethod
    def add_product_reviews(self, product, reviews):
        pass

    @abstractmethod
    def get_reviews_by_product(
        self, product_id: str, sortKey: ReviewSortKey = ReviewSortKey.SORTKEY_CREATED_AT
    ) -> list[Review]:
        pass


class ReviewFlag(ABC):
    @abstractmethod
    def flag_review(self, review: Review):
        pass

class ReviewFlagger(ABC):
    reviews_provider: ReviewCrud

    def flag_review(self, review: Review):
        review.


class ReviewImpressionStore(ReviewImpression):
    reviews_provider: ReviewCrud

    def upvote(self, user: User, review: Review):
        if user.id not in self.reviews_provider.get_by_id(review.id).upvotes:
            self.reviews_provider.get_by_id(review.id).upvotes.add(user.id)

    def helpful(self, user: User, review: Review):
        if user.id not in self.reviews_provider.get_by_id(review.id).helpfuls:
            self.reviews_provider.get_by_id(review.id).helpfuls.add(user.id)


class ProductReviewProvider(ProductReview):
    product_to_reviews = defaultdict[str, list[Review]]

    def get_reviews_by_product(
        self, product_id: str, sortKey: ReviewSortKey = ReviewSortKey.SORTKEY_CREATED_AT
    ) -> list[Review]:

        items = self.product_to_reviews.get(product_id)

        if sortKey == ReviewSortKey.SORTKEY_CREATED_AT:
            items = items.sort(key=lambda x: x.created_at, reverse=True)
        elif sortKey == ReviewSortKey.SORTKEY_RATING:
            items = items.sort(key=lambda x: x.rating, reverse=True)
        return items

    def add_product_review(self, product: Product, review: Review):
        self.product_to_reviews.get(product.id).append(review)
        product.avg_rating = (
            product.avg_rating * product.review_cnt + review.rating
        ) / (product.review_cnt + 1)
        product.review_cnt += 1


class ReviewStore(ReviewCrud):
    reviews: dict[str, Review]  # reviewID - review

    def get_by_id(self, id: str) -> Review:
        return self.reviews.get(id, None)

    def create_review(self, review: Review):
        # TODO: gen Id internally
        self.reviews[review.id] = Review

    def modify_review(
        self, user: User, review: Review, description: str, rating: CustomerRating
    ):
        self.reviews[review.id].description = description
        self.reviews[review.id].rating = rating

    def delete_review(self, user: User, review: Review):
        if (
            self.reviews.get(review.id)
            and self.reviews.get(review.id).user_id == user.id
        ):
            self.reviews.pop(review.id)
