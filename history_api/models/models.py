from pydantic import BaseModel


class NewMessage(BaseModel):
    id: str
    type: str
    message: str
    is_admin: bool = True
    owner: bool = True
    client_name: str = "telegram"
    message_id: str = None
    reply: str = None
    caption: str = None
    user_name: str = None


class switch(BaseModel):
    id: str
    is_support: bool


class AutoCompleteMessage(BaseModel):
    message: str


class User(BaseModel):
    username: str
    password: str


class AdminUser(BaseModel):
    username: str
    password: str
    role: int


class AdminPassword(BaseModel):
    password: str
    new_password: str


class DeleteUser(BaseModel):
    username: str
