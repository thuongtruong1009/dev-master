import { Response } from 'express';
import Controller from '../decorators/controller.decorator';
import { Get, Post } from '../routes/index';
import useDecorator from '../decorators/use.decorator';
import { RequestWithBody } from '../interfaces/default.interface';

@Controller('/auth')
export default class AuthController {
    @Get('/login')
    async getLogin(req: RequestWithBody, res: Response): Promise<void> {
        res.status(200).send(`
      <form method="POST">
        <div>
          <label>Email</label>
          <input name="email" />
        </div>
        <div>
          <label>Password</label>
          <input name="password" type="password" />
        </div>
        <button type='submit'>Submit</button>
      </form>
    `);
    }

    @Post('/login')
    async postLogin(req: RequestWithBody, res: Response): Promise<void> {
        const { email, password } = req.body;
        if (email === 'email@gmail.com' && password === 'password') {
            req.session = { loggedIn: true };
            res.redirect('/');
        } else {
            res.status(401).send('Invalid email or password');
        }
    }

    @Get('/logout')
    async getLogout(req: RequestWithBody, res: Response): Promise<void> {
        req.session = undefined;
        res.redirect('/');
    }
}
