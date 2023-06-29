from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse
from fastapi_utils.tasks import repeat_every

from database.database import DBHelper
import models.models as models
import uvicorn
import logging


Db = DBHelper()
app = FastAPI(docs_url='/history/docs',
              redoc_url='/history/redoc',
              openapi_url='/history/openapi.json')

@app.post('/history/update/{id}')
async def update(id: str):
    try:
        res=Db.update_date(id)
        return res
    except Exception as e:
        print(e)
        raise HTTPException(status_code=400, detail=f'{e}')
    
@app.post('/history/newMessage')
async def add_new_message( new_message: models.NewMessage):
    try:
        status,res=Db.new_message(new_message.id, new_message.type, new_message.message, new_message.owner,new_message.is_admin,new_message.client_name,new_message.message_id,new_message.reply,new_message.caption, new_message.user_name)
        if status['status'] != 'Success':
            return JSONResponse(content={'detail': status['status']}, status_code=400)
        return res
    except Exception as e:
        print(e)
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')

@app.post('/history/switch')
async def switch( switch: models.switch):
    try:
        res=Db.switch(id=switch.id, is_support=switch.is_support)
        if res['status'] != 'Success':
            return JSONResponse(content={'detail': res['status']}, status_code=400)
        return {switch.id:res}
    except Exception as e:
        print(e)
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')
    

@app.get('/history/switch_position/{id}')
async def switch_position( id: str):
    try:
        res=Db.switch_pos(id=id)
        return {"switch":res}
    except Exception as e:
        print(e)
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')
    

@app.on_event("startup")
@repeat_every(seconds=60*60*24)
async def auto_switch():
    try:
        Db.auto_switch()
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')


# @app.get('/chat/get_auto_complete')
# async def auto_complete():
#     try:
#         res = Db.get_auto_complete()
#         return res
#     except Exception as e:
#         logging.error(e)
#         raise HTTPException(status_code=400, detail=f'{e}')
    
# @app.put('/chat/update_auto_complete/{id}')
# async def update_auto_complete(id: int, message: models.AutoCompleteMessage):
#     try:
#         Db.update_auto_complete(id, message.message)
#         return {'success': id}
#     except Exception as e:
#         logging.error(e)
#         raise HTTPException(status_code=400, detail=f'{e}')

# @app.post('/chat/add_auto_complete')
# async def add_auto_complete(message: models.AutoCompleteMessage):
#     try:
#         Db.add_auto_complete(message.message)
#         return {'success': message}
#     except Exception as e:
#         logging.error(e)
#         raise HTTPException(status_code=400, detail=f'{e}')

# @app.delete('/chat/delete_auto_complete/{id}')
# async def delete_auto_complete(id: int):
#     try:
#         Db.delete_auto_complete(id)
#         return {'success': id}
#     except Exception as e:
#         logging.error(e)
#         raise HTTPException(status_code=400, detail=f'{e}')

@app.get('/history/getChats')
async def get_chats():
    try:
        res = Db.get_chats()
        # if res['status'] != 'Success':
        #     return JSONResponse(content={'detail': res['status']}, status_code=400)
        return res
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')

@app.get('/history/getChatsbyid/{id}')
async def get_chats_byid(id: str):
    try:
        res=Db.get_chats_by_id(id)
        return res
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')

@app.get('/history/getMessages/{id}')
async def get_messages(id: str):
    try:
        res = Db.get_messages(id)
        # if res['status'] != 'Success':
        #     return JSONResponse(content={'detail': res['status']}, status_code=400)
        return res
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')

@app.get('/history/checkStatus/{id}')
async def check_Status(id: str):
    try:
        res = Db.check_rasa(id)
        if res['status'] != 'Success':
            return JSONResponse(content={'detail': res['status']}, status_code=400)
        return res
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')

@app.get('/history/media/{id}')
async def get_media(id: str):
    try:
        res = Db.get_media(id)
        if res['status'] != 'Success':
            return JSONResponse(content={'detail': res['status']}, status_code=400)
        return res['text']
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')

if __name__ == "__main__":
    Db.setup()
    uvicorn.run(app, host="0.0.0.0", port=5679)
