import { Response, NextFunction } from 'express';
import Controller from '../decorators/controller.decorator';
import { Get } from '../routes/index';
import useDecorator from '../decorators/use.decorator';
import { RequestWithBody } from '../interfaces/default.interface';

function requireAuth(
    req: RequestWithBody,
    res: Response,
    next: NextFunction,
): void {
    if (req.session && req.session.loggedIn) {
        next();
    }
    res.status(403).send('Not permitted');
}

@Controller('')
export default class AppController {
    @Get('/')
    async getApp(req: RequestWithBody, res: Response): Promise<void> {
        if (req.session && req.session.loggedIn) {
            res.status(200).send(`
                <h1>you are logged in</h1>
                <a href='/auth/logout'>log out</a>
            `);
        } else {
            res.status(401).send(`
              <h1>you are not logged in</h1>
              <a href='/auth/login'>log in</a>
            `);
        }
    }

    @Get('/protected')
    // @useDecorator(requireAuth)
    async getProtected(req: RequestWithBody, res: Response): Promise<void> {
        res.send('Welcome to protected route, logged in user');
    }
}
