import { Request, Response, NextFunction, RequestHandler } from 'express';

export default function bodyValidator(keys: string[]): RequestHandler {
    return function (req: Request, res: Response, next: NextFunction) {
        if (!req.body) {
            return res.status(422).send('Invalid request');
        }
        for (let key in keys) {
            if (!req.body[key]) {
                return res.status(422).send('Invalid request');
            }
        }

        next();
    };
}
