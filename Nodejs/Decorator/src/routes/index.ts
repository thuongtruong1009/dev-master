import 'reflect-metadata';
import { EMethod } from '../enums/methods.enum';
import { routeDecorator } from '../decorators/route.decorator';

export const Get = routeDecorator(EMethod.get);
export const Post = routeDecorator(EMethod.post);
export const Put = routeDecorator(EMethod.put);
export const Delete = routeDecorator(EMethod.delete);
export const Patch = routeDecorator(EMethod.patch);
