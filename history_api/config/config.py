from environs import Env

env = Env()

env.read_env()
POSTGRES = env.str('POSTGRES')
