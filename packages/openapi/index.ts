import 'reflect-metadata';
import express from 'express';
import * as bodyParser from 'body-parser';
import helmet from 'helmet';
const cors = require('cors');
import router from './lib/router';


async function server() {

  const app = express();
  app.use(bodyParser.json());
  app.use(bodyParser.urlencoded({ extended: false }));
  app.use(cors());
  app.use(helmet());

  app.use(router);

  app.use((err, req, res, next) => {
    res.locals.message = err.message;
    res.locals.error = req.app.get('env') === 'development' ? err : {};

    res.status(err.status || 500);
    res.render('error');
  });

  let port = 8008;

  app.listen(port);
  console.log('🚀 Server ready at http://localhost:' + port);
}

export async function start() {
    try {
        await server();
    } catch (error) {
        console.error(error);
    }
}

start();