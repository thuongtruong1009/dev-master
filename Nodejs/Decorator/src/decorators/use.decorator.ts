import { RequestHandler } from 'express';
import 'reflect-metadata';
import { EMetadata } from '../enums/metadata.enum';

export default function useDecorator(middleware: RequestHandler) {
    return function (target: any, key: string, desc: PropertyDecorator) {
        const middllewares =
            Reflect.getMetadata(EMetadata.middleware, target, key) || [];

        middllewares.push(middleware);

        Reflect.defineMetadata(EMetadata.middleware, middleware, target, key);
    };
}
