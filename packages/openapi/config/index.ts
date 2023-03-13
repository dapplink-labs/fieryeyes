require('dotenv').config()
console.log(process.env)

const configs = {
  mongoHost: process.env.MONGO_CONNECT
}

export default configs