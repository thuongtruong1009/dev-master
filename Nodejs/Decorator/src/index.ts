import express from 'express';
import bodyParser from 'body-parser';
import cookieSession from 'cookie-session';
import Router from './routes/instance.route';

import './controllers/app.controller';
import './controllers/auth.controller';

const port = 3000;

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cookieSession({ keys: ['secretkey'] }));
app.use(Router.getInstance());

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});
