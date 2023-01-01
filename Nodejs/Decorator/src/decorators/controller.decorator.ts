import 'reflect-metadata';
import Router from '../routes/instance.route';
import { EMetadata } from '../enums/metadata.enum';
import { EMethod } from '../enums/methods.enum';
import bodyValidator from '../middlewares/validator.middleware';

export default function Controller(routePrefix?: string) {
    return function (target: Function) {
        const router = Router.getInstance();

        for (let key in target.prototype) {
            const routeHandler = target.prototype[key];

            const path = Reflect.getMetadata(
                EMetadata.path,
                target.prototype,
                key,
            );

            const method: EMethod = Reflect.getMetadata(
                EMetadata.method,
                target.prototype,
                key,
            );

            const middleware =
                Reflect.getMetadata(
                    EMetadata.middleware,
                    target.prototype,
                    key,
                ) || [];

            const requiredBodyProps =
                Reflect.getMetadata(
                    EMetadata.validator,
                    target.prototype,
                    key,
                ) || [];

            const validator = bodyValidator(requiredBodyProps);
            if (path) {
                router[method](
                    routePrefix + path,
                    ...middleware,
                    validator,
                    routeHandler,
                );
            }
        }
    };
}
