from sqlalchemy import create_engine
from config.config import POSTGRES
from passlib.context import CryptContext
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


class DBHelper:

    def __init__(self):
        self.conn = create_engine(POSTGRES)

    def setup(self):
        try:
            stmt = """
                    CREATE TABLE IF NOT EXISTS admins(
                    user_name text NOT NULL,
                    password text,
                    role smallint,
                    PRIMARY KEY (user_name),
                    CONSTRAINT user_name  UNIQUE (user_name)
                    );
                """
            self.conn.execute(stmt)
            print(1)
        except Exception as e:
            print(e)
    
    def get_user(self, username, password):
        try:
            stmt = f"""Select password, role from admins where user_name='{username}'"""
            passw = self.conn.execute(stmt).fetchone()
            if passw:
                if self.verify_password(password, passw[0]):
                    return {"status": 'Success', "role": ("SuperAdmin" if passw[1] == 1 else "Owner" if passw[1] == 2 else "Operator")}
                else:
                    return {"status": 'Password is not correct'}
            else:
                return {"status": "User not found"}
        except Exception as e:
            return {"status": e}


    def create_user(self, username, password, role, current_user):
        """
                ROLE
            super admin => 1
            admin       => 2
            operator    => 3
        """
        try:
            stmt = f"""Select role from admins where user_name='{current_user}'"""
            res = self.conn.execute(stmt).fetchone()[0]
            if role <= res:
                return {"status": "Access denied"}
            stmt = f"""Select * from admins where user_name='{username}'"""
            res = self.conn.execute(stmt).fetchone()
            if not res:
                stmt = f"""Insert into admins(user_name, password, role) values('{username}','{self.get_password_hash(password)}',{role})"""
                self.conn.execute(stmt)
                return {"status": 'User is successfully added'}
            else:
                return {"status": "User is already exist"}
        except Exception as e:
            return {"status": e}
    
    def change_password(self, username, password, new_password):
        try:
            check_user_password = self.get_user(
                username=username, password=password)
            if check_user_password['status'] != 'Success':
                return {"status": "Password is wrong"}
            stmt = f"""UPDATE admins SET password = '{self.get_password_hash(new_password)}' WHERE user_name = '{username}'"""
            self.conn.execute(stmt)
            return {"status": 'Password is successfully changed'}
        except Exception as e:
            return {"status": e}

    def get_owners(self, username):
        try:
            stmt = f"""Select role from admins where user_name='{username}'"""
            res = self.conn.execute(stmt).fetchone()[0]
            if res > 1:
                return {"status": "Access denied"}
            stmt = f"""SELECT user_name, role FROM admins WHERE role >= 2"""
            res = self.conn.execute(stmt).fetchall()
            res = list(map(lambda x: dict(x), res))
            return res
        except Exception as e:
            return None

    def get_operators(self, username):
        try:
            stmt = f"""Select role from admins where user_name='{username}'"""
            res = self.conn.execute(stmt).fetchone()[0]
            if res > 2:
                return {"status": "Access denied"}
            stmt = f"""SELECT user_name, role FROM admins WHERE role = 3"""
            res = self.conn.execute(stmt).fetchall()
            res = list(map(lambda x: dict(x), res))
            return res
        except Exception as e:
            return None

    def delete_owners(self, role_username, username):
        try:
            stmt = f"""Select role from admins where user_name='{role_username}'"""
            res = self.conn.execute(stmt).fetchone()[0]
            if res > 1:
                return {"status": "Access denied"}
            stmt = f"""DELETE FROM admins WHERE user_name = '{username}'"""
            self.conn.execute(stmt)
            return {"status": "User is successfully deleted"}
        except Exception as e:
            return None

    def delete_operators(self, role_username, username):
        try:
            stmt = f"""Select role from admins where user_name='{role_username}'"""
            res = self.conn.execute(stmt).fetchone()[0]
            if res > 2:
                return {"status": "Access denied"}
            stmt = f"""DELETE FROM admins WHERE user_name = '{username}'"""
            res = self.conn.execute(stmt)
            return {"status": "User is successfully deleted"}
        except Exception as e:
            return None

    def verify_password(self, plain_password, hashed_password):
        return pwd_context.verify(plain_password, hashed_password)

    def get_password_hash(self, password):
        return pwd_context.hash(password)
