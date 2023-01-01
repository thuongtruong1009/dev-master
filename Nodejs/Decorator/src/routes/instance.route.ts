import express from 'express';

export default class Router {
    private static instance: express.Router;

    public static getInstance(): express.Router {
        if (!Router.instance) {
            Router.instance = express.Router();
        }

        return Router.instance;
    }
}
