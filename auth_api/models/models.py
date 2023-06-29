from pydantic import BaseModel
from config.config import SECRET_KEY






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


class Settings(BaseModel):
    authjwt_secret_key: str = SECRET_KEY
    authjwt_token_location: set = {"cookies"}
    authjwt_cookie_csrf_protect: bool = False
