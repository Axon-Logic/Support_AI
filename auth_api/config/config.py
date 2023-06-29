from environs import Env

env = Env()

env.read_env()

POSTGRES = env.str('POSTGRES')
SECRET_KEY = env.str('SECRET_KEY')
EXPIRES_TIME = env.str('EXPIRES_TIME')
