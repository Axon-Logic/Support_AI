from sqlalchemy import create_engine
from config.config import POSTGRES
from passlib.context import CryptContext
import uuid
import base64
import glob
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


class DBHelper:

    def __init__(self):
        self.conn = create_engine(POSTGRES)

    def setup(self):
        try:
            stmt = """
                    CREATE TABLE IF NOT EXISTS client(
                    client_id text NOT NULL,
                    message_count bigint,
                    user_name text,
                    first_name text,
                    last_name text,
                    client_name text,
                    message text,
                    date timestamp,
                    type text,
                    is_support bool DEFAULT false,
                    PRIMARY KEY (client_id),
                    CONSTRAINT client_id  UNIQUE (client_id)
                    );

                    CREATE TABLE IF NOT EXISTS users(
                    user_name text NOT NULL,
                    password text,
                    role smallint,
                    PRIMARY KEY (user_name),
                    CONSTRAINT user_name  UNIQUE (user_name)
                    );

                    CREATE TABLE IF NOT EXISTS auto_complete_message(
                    message text NOT NULL,
                    id serial NOT NULL,
                    PRIMARY KEY (id),
                    CONSTRAINT id  UNIQUE (id)
                    );
                """
            self.conn.execute(stmt)
        except Exception as e:
            print(e)
            return {"status": e}

    def create_new_table(self, id):
        try:
            stmt = f"""CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
                    
                """
            self.conn.execute(stmt)
            stmt = f"""CREATE TABLE IF NOT EXISTS messages_{id}(
                    message_id text DEFAULT UUID_GENERATE_V4(),
                    client_id text NOT NULL,
                    type text,
                    date timestamp,
                    owner bool,
                    is_admin bool,
                    message text,
                    reply text,
                    caption text
                    );"""
            self.conn.execute(stmt)
        except Exception as e:
            print(e)
            return {"status": e}

    def update_date(self, id):
        try:
            stmt = f"""UPDATE client SET message_count=0 where client_id = '{id}' RETURNING client_id,message_count,user_name,
            first_name,last_name,extract(epoch from date) as date,  
            CASE
                WHEN type = 'text' THEN message
                ELSE type
                END as message"""
            res = self.conn.execute(stmt).fetchone()
            return res
        except Exception as e:
            print(e)

    def new_message(self, id, type, message, owner, is_admin, client_name, message_id, reply, caption, username):
        try:
            if not username:
                username = id
            self.new_client(id, client_name=client_name, username=username)
            if type != 'text':
                filename = ''
                myuuid = str(uuid.uuid4())
                data = base64.b64decode(message)
                if type == 'voice':
                    filename = myuuid+'.ogg'
                elif type == 'photo':
                    filename = myuuid+'.png'
                elif type == 'video':
                    filename = myuuid+'.mp4'
                with open('media/'+filename, 'wb') as f:
                    f.write(data)
                message = myuuid
            message = message.replace("'", "\\'")
            reply = self.check_var(reply)
            message_id = self.check_var(message_id)
            caption = self.check_var(caption)
            stmt = f"""insert into messages_{id}(client_id, type, message, date, owner,is_admin, reply, message_id, caption)
                values('{id}', '{type}', e'{message}', now() + interval '5 hour', {owner},'{is_admin}','{reply}','{message_id}','{caption}') RETURNING extract(epoch from date) as date, message_id"""
            stmt = stmt.replace("'None'", "DEFAULT")
            date, message_id = self.conn.execute(stmt).fetchone()
            if not owner:
                stmt = f"""Update client set type='{type}' ,date=now() + interval '5 hour',message=e'{message}',message_count=message_count+1 where client_id='{id}' RETURNING message_count, is_support """
            else:
                stmt = f"""Update client set type='{type}' ,date=now() + interval '5 hour',message=e'{message}',message_count=0,is_support='{is_admin}' where client_id='{id}' RETURNING message_count, is_support"""
            count, is_support = self.conn.execute(stmt).fetchone()
            return {"status": 'Success'}, {"Status": is_support, "MessageDate": date, "MessageCount": count, "MessageId": message_id}
        except Exception as e:
            print(e)
            return {"status": e}, {"Status": "err"}

    def check_var(self, var):
        print(var, var == "")
        if not var:
            return None
        return var

    def check_rasa(self, id):
        try:
            stmt = f"""SELECT is_admin
                FROM (
                SELECT client_id, date, is_admin,
                ROW_NUMBER() OVER (PARTITION BY client_id ORDER BY date DESC) AS row_num
                FROM messages_{id}
                ) AS subquery
                WHERE row_num = 1 and now()+ interval '5 hour'- date <interval '1 day' and is_admin='true'"""
            res = self.conn.execute(stmt).fetchone()
            if res:
                return res[0]
            else:
                return False
        except Exception as e:
            print(e)
            return {"status": e}

    def switch(self, id, is_support):
        try:
            stmt = f"""Update client set is_support='{is_support}' where client_id='{id}'"""
            self.conn.execute(stmt)
            return {"status": 'Success'}
        except Exception as e:
            print(e)
            return {"status": e}

    def switch_pos(self, id):
        try:
            stmt = f"""select is_support from client where client_id='{id}'"""
            res = self.conn.execute(stmt).fetchone()
            return res
        except Exception as e:
            print(e)
            return {"status": e}

    def auto_switch(self):
        try:
            stmt = f"""Update client set is_support='false' where now()+ interval '5 hour'- date >interval '1 day' and is_support='true'"""
            self.conn.execute(stmt)
        except Exception as e:
            print(e)

    def get_auto_complete(self):
        try:
            stmt = """select * from auto_complete_message order by id"""
            res = self.conn.execute(stmt).fetchall()
            res = list(map(lambda x: dict(x), res))
            return res
        except Exception as e:
            print(e)
            return {"status": e}

    def update_auto_complete(self, id, message):
        try:
            stmt = f"""update auto_complete_message set message = '{message}'
                where id = {id}"""
            self.conn.execute(stmt)
        except Exception as e:
            print(e)
            return {"status": e}

    def delete_auto_complete(self, id):
        try:
            stmt = f"""delete from auto_complete_message
                where id = {id}"""
            self.conn.execute(stmt)
        except Exception as e:
            print(e)
            return {"status": e}

    def add_auto_complete(self, message):
        try:
            stmt = f"""insert into auto_complete_message(message)
                values('{message}')"""
            self.conn.execute(stmt)
        except Exception as e:
            print(e)
            return {"status": e}

    def get_messages(self, id):
        try:
            stmt = f"""SELECT client_id,
                        type,
                        message,
                        extract(epoch from date) as date,
                        owner,
                        reply,
                        message_id,
                        caption 
                        from messages_{id} order by date"""
            res = self.conn.execute(stmt).fetchall()
            res = list(map(lambda x: dict(x), res))
            # remove this code after changing UI
            # for i in res:
            #     if i['type'] != 'text':
            #         i['message'] = self.get_media(i['message'])
            return res
        except Exception as e:
            print(e)
            # return {"status": e}

    def get_media(self, id):
        try:
            res = glob.glob("media/"+id + '.*')
            with open(res[0], "rb") as File:
                text = base64.b64encode(File.read())
            return {"status": "Success","text":text}
        except Exception as e:
            print(e)
            return {"status": e}

    def new_client(self, id, client_name, username):
        try:

            stmt = f"""insert into client(client_id, user_name, message_count, client_name)
                values('{id}','{username}' ,0, '{client_name}') ON CONFLICT(client_id) DO NOTHING"""
            self.conn.execute(stmt)
            self.create_new_table(id=id)
        except Exception as e:
            print(e)
            return {"status": e}

    def get_chats(self):
        try:
            stmt = """select client_id,message_count,user_name,
            first_name,last_name,date,  
            CASE
                WHEN type = 'text' THEN message
                ELSE type
                END as message,
            CASE 
                WHEN message_count>0 THEN true
                ELSE false
                END as count
              from client order by message_count desc"""
            res = self.conn.execute(stmt).fetchall()
            res = list(map(lambda x: dict(x), res))
            return res
        except Exception as e:
            print(e)
            # return {"status": e}

    def get_chats_by_id(self, id):
        try:
            stmt = f"""select client_id,message_count,user_name,
            first_name,last_name,date,
            CASE
                WHEN type = 'text' THEN message
                ELSE type
                END as message,
            CASE 
                WHEN message_count>0 THEN true
                ELSE false
                END as count 
            from client where client_id='{id}' order by message_count desc"""
            return dict(self.conn.execute(stmt).fetchone())
        except Exception as e:
            print(e)
            return {"status": e}
