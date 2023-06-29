from fastapi import FastAPI, HTTPException, Depends, Request
from fastapi.responses import JSONResponse
from fastapi_jwt_auth import AuthJWT
from fastapi_jwt_auth.exceptions import AuthJWTException

from database.database import DBHelper
import models.models as models
import logging
from config.config import EXPIRES_TIME
import uvicorn

Db = DBHelper()
app = FastAPI(docs_url='/auth/docs',
              redoc_url='/auth/redoc',
              openapi_url='/auth/openapi.json')
logging.basicConfig(filename='app.log', filemode='w', format='%(name)s - %(levelname)s - %(message)s')


@AuthJWT.load_config
def get_config():
    return models.Settings()

@app.exception_handler(AuthJWTException)
def authjwt_exception_handler(request: Request, exc: AuthJWTException):
    return JSONResponse(
        status_code=exc.status_code,
        content={"detail": exc.message} 
    )

@app.post('/auth/login')
async def login(user: models.User, Authorize: AuthJWT = Depends()):
    try:
        result = Db.get_user(username=user.username, password=user.password)
        if result['status'] != 'Success':
            # print(result)
            return JSONResponse(content={'detail': result['status']}, status_code=400)
        access_token = Authorize.create_access_token(
            subject=user.username, expires_time=int(EXPIRES_TIME)*60)
        Authorize.set_access_cookies(access_token)
        return {'data': {
                "username": user.username,
                "role": result['role']}}
    except AuthJWTException:
        logging.error("Unauthorized token is missing")
        return JSONResponse(content={'detail': 'Unauthorized token is missing'}, status_code=401)
    except Exception:
        logging.error(result['status'])
        return JSONResponse(content={'detail': result['status']}, status_code=400)


@app.delete('/auth/logout')
async def logout(Authorize: AuthJWT = Depends()):
    try:
        Authorize.jwt_required()
        Authorize.unset_jwt_cookies()
        return {"msg": "Successfully logout"}
    except AuthJWTException:
        logging.error("Unauthorized token is missing")
        raise HTTPException(
            status_code=401, detail=f'Unauthorized token is missing')
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail="error")


@app.post('/auth/addUser')
async def add_user(user: models.AdminUser, Authorize: AuthJWT = Depends()):
    try:
        Authorize.jwt_required()
        current_username = Authorize.get_jwt_subject()
        if user.role > 3 or user.role < 1:
            logging.error("Invalid role")
            raise Exception("Invalid role")
        result = Db.create_user(username=user.username, password=user.password,
                                role=user.role, current_user=current_username)

        if result['status'] != 'User is successfully added':
            logging.error(result['status'])
            return JSONResponse(content={"detail": result['status']}, status_code=400)
        return {"success": result['status']}

    except AuthJWTException:
        logging.error("Unauthorized token is missing")
        raise HTTPException(
            status_code=401, detail=f'Unauthorized token is missing')
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')


@app.put('/auth/changePassword')
async def change_password(user: models.AdminPassword, Authorize: AuthJWT = Depends()):
    try:
        Authorize.jwt_required()
        current_username = Authorize.get_jwt_subject()
        result = Db.change_password(
            username=current_username, password=user.password, new_password=user.new_password)
        if result['status'] != 'Password is successfully changed':
            logging.error(result['status'])
            return JSONResponse(content={"detail": result['status']}, status_code=400)
        return {"success": result['status']}

    except AuthJWTException:
        logging.error('Unauthorized token is missing')
        raise HTTPException(
            status_code=401, detail=f'Unauthorized token is missing')
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')


@app.get('/auth/admin/owners')
async def get_admin_owners(Authorize: AuthJWT = Depends()):
    try:
        Authorize.jwt_required()
        res = Db.get_owners(Authorize.get_jwt_subject())
        if res:
            if type(res) is dict:
                logging.error("Access denied")
                raise Exception("Access denied")
            return res
        else:
            logging.error("error")
            raise Exception("error")

    except AuthJWTException:
        logging.error('Unauthorized token is missing')
        raise HTTPException(
            status_code=401, detail=f'Unauthorized token is missing')
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')


@app.get('/auth/admin/operators')
async def get_admin_operators(Authorize: AuthJWT = Depends()):
    try:
        Authorize.jwt_required()
        res = Db.get_operators(Authorize.get_jwt_subject())
        if res:
            if type(res) is dict:
                logging.error('Access denied')
                raise Exception("Access denied")
            return res
        else:
            logging.error('error')
            raise Exception("error")

    except AuthJWTException:
        logging.error('Unauthorized token is missing')
        raise HTTPException(
            status_code=401, detail=f'Unauthorized token is missing')
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')


@app.delete('/auth/admin/owners')
async def get_admin_owners(user: models.DeleteUser, Authorize: AuthJWT = Depends()):
    try:
        Authorize.jwt_required()
        res = Db.delete_owners(
            role_username=Authorize.get_jwt_subject(), username=user.username)
        if res['status'] != "User is successfully deleted":
            logging.error(res['status'])
            raise Exception(res['status'])
        return {"success": res['status']}
    except AuthJWTException:
        logging.error('Unauthorized token is missing')
        raise HTTPException(
            status_code=401, detail=f'Unauthorized token is missing')
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')


@app.delete('/auth/admin/operators')
async def get_admin_owners(user: models.DeleteUser, Authorize: AuthJWT = Depends()):
    try:
        Authorize.jwt_required()
        res = Db.delete_operators(
            role_username=Authorize.get_jwt_subject(), username=user.username)
        if res['status'] != "User is successfully deleted":
            logging.error(res['status'])
            raise Exception(res['status'])
        return {"success": res['status']}
    except AuthJWTException:
        logging.error('Unauthorized token is missing')
        raise HTTPException(
            status_code=401, detail=f'Unauthorized token is missing')
    except Exception as e:
        logging.error(e)
        raise HTTPException(status_code=400, detail=f'{e}')


if __name__ == "__main__":
    Db.setup()
    uvicorn.run(app, host="0.0.0.0", port=5677)
