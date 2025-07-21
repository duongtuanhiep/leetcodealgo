from enum import Enum
from abc import ABC, abstractmethod
from datetime import datetime
from collections import defaultdict

"""
Base Requirements:

Model the core classes and relationships for a messaging/chat application (like WhatsApp, Slack, or Discord).
The system should support individual users who can send and receive messages.
Users can participate in group chats with multiple participants.
Messages should have timestamps and delivery status tracking.
Include basic functionality for creating chat rooms, sending messages, and retrieving chat history.
Focus on the class structure, relationships, and core messaging functionality.

Key Considerations:
Think about the relationship between users, messages, and chat rooms
Consider how to handle both direct messages (1-on-1) and group chats
Design for message delivery and read status tracking
Keep extensibility in mind for future features
Take your time to design the class structure and relationships. Let me know when you're ready for feedback or want to discuss your approach!

- Define Functional requirement 5-7m: 
- Define class properties & method 5-7m: 


Entities:
- Chat
    has []chatmessage
    add(ChatMessage)

- ChatMessage
    Has userID, chatcontent
    userAck()
- Recipient
- DeliveryStatus

"""


class DeliveryStatus(Enum):
    STATUS_SENT = 1
    STATUS_DELIVERED = 2


class User:
    id: str
    name: str

    def get_id(self) -> str:
        return self.id


class UserChat:
    user: User
    chats: list["Chat"]


class ChatMessage:
    content: str
    status: DeliveryStatus
    timestamp: datetime
    sender: str
    ack: set
    parent: "ChatMessage"

    def update_status(self, status: DeliveryStatus):
        self.status = status

    def user_ack(self, user: User):
        self.ack.add(user.id())

    def get_ack_cnt(self) -> int:
        return len(self.ack)

    def get_parent(self) -> "ChatMessage":
        return self.parent


class MessageStatusUpdater(ABC):
    @abstractmethod
    def user_ack(self, user: User, chat: "Chat", message: ChatMessage):
        pass

    @abstractmethod
    def user_send(self, user: User, message: ChatMessage):
        pass


class MessageNotifier(ABC):
    @abstractmethod
    def notify(self, user: User, message: ChatMessage):
        pass


class MessageManager:
    updater: MessageStatusUpdater
    notifier: MessageNotifier

    def send_message(self, user: "User", chat: "Chat", message: "ChatMessage"):
        if user in chat.get_participants():
            chat.append_message(message)
            self.updater.user_send(user, message)

            for receiver in chat.get_participants():
                # notify all user even for threaded message
                if user.id() == receiver.id():
                    continue
                self.notifier.notify(receiver, message)


class Chat:
    id: str
    history: list[ChatMessage]
    participants: list[User]
    thread: defaultdict[ChatMessage, list[ChatMessage]]

    def get_history_threaded(self) -> list[ChatMessage]:
        # Perform Topo sort here to ensure ordering of messages according to timestamp / Parent child in thread
        pass

    def get_history_flatten(self):  # Return timestamp ordered message
        pass

    def append_message(self, msg: "ChatMessage"):
        self.history.append(msg)
        if msg.get_parent():
            self.thread[msg.get_parent].append(msg)

    def get_participants(self) -> list[User]:
        return self.participants

    def add_participants(self, user: "User"):
        if user not in self.participants:
            self.participants.append(user)


class ChatManager:
    chats: list[Chat]

    def create_chat(self, user: User, participants: list["User"]) -> "Chat":
        return Chat()

    def get_chat(self, user: User) -> "UserChat":
        return UserChat()

    def join_chat(self, user: User, chat: "UserChat"):
        pass
