import { RequestHandler } from 'express';
import 'reflect-metadata';
import { EMetadata } from '../enums/metadata.enum';

interface RouteHandlerDescriptor extends PropertyDescriptor {
    value?: RequestHandler;
}

export function routeDecorator(method: string) {
    return function get(path: string) {
        return function (
            target: any,
            key: string,
            desc: RouteHandlerDescriptor,
        ) {
            Reflect.defineMetadata(EMetadata.path, path, target, key);
            Reflect.defineMetadata(EMetadata.method, method, target, key);
        };
    };
}
